/*
 *  Copyright (c) 2017, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package grpc_util

import (
	"net"
	"time"
	"google.golang.org/grpc"
	"github.com/nebulaim/telegramd/baselib/grpc_util/middleware/recovery2"
	"github.com/nebulaim/telegramd/baselib/grpc_util/service_discovery"
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/baselib/grpc_util/service_discovery/etcd3"
	"github.com/coreos/etcd/clientv3"
	"os/signal"
	"syscall"
	"os"
)

type RPCServer struct {
	addr string
	registry *etcd3.EtcdReigistry
	s        *grpc.Server
}

func NewRpcServer(addr string, discovery *service_discovery.ServiceDiscoveryServerConfig) *RPCServer {
	etcdConfg := clientv3.Config{
		Endpoints: discovery.EtcdAddrs,
	}

	registry, err := etcd3.NewRegistry(
		etcd3.Option{
			EtcdConfig:  etcdConfg,
			RegistryDir: "/nebulaim",
			ServiceName: discovery.ServiceName,
			NodeID:      discovery.NodeID,
			NData: etcd3.NodeData{
				Addr: discovery.RPCAddr,
				//Metadata: map[string]string{"weight": "1"},
			},
			Ttl: 1000 * time.Second,
		})
	if err != nil {
		glog.Fatal(err)
		// return nil
	}

	s := grpc_recovery2.NewRecoveryServer2(BizUnaryRecoveryHandler, BizUnaryRecoveryHandler2, BizStreamRecoveryHandler)
	rs := &RPCServer{
		addr:     addr,
		registry: registry,
		s:        s,
	}
	return rs
}

// type func RegisterRPCServerHandler(s *grpc.Server)
type RegisterRPCServerFunc func(s *grpc.Server)

func (s *RPCServer) Serve(regFunc RegisterRPCServerFunc) {
	// defer s.GracefulStop()
	listener, err := net.Listen("tcp", s.addr)

	if err != nil {
		glog.Error("failed to listen: %v", err)
		return
	}
	glog.Infof("rpc listening on:%s", s.addr)

	if regFunc != nil {
		regFunc(s.s)
	}

	defer s.s.GracefulStop()
	go s.registry.Register()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s2 := <-ch
		glog.Infof("exit...")
		s.registry.Deregister()
		if i, ok := s2.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}

	}()

	if err := s.s.Serve(listener); err != nil {
		glog.Fatalf("failed to serve: %s", err)
	}
}

func (s *RPCServer) Stop() {
	s.s.GracefulStop()
}
