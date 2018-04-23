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

package rpc

import (
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/baselib/logger"
	"github.com/nebulaim/telegramd/baselib/grpc_util"
	"github.com/nebulaim/telegramd/mtproto"
	"golang.org/x/net/context"
	"github.com/nebulaim/telegramd/biz/base"
	"github.com/nebulaim/telegramd/biz/core/account"
)

// account.getNotifySettings#12b3ad31 peer:InputNotifyPeer = PeerNotifySettings;
func (s *AccountServiceImpl) AccountGetNotifySettings(ctx context.Context, request *mtproto.TLAccountGetNotifySettings) (*mtproto.PeerNotifySettings, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("account.getNotifySettings#12b3ad31 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// peer := request.GetPeer()
	if request.GetPeer().GetConstructor() != mtproto.TLConstructor_CRC32_inputNotifyPeer {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error(err, ": peer only is inputNotifyPeer")
		return nil, err
	}

	peer := base.FromInputNotifyPeer(request.GetPeer())
	reply := account.GetNotifySettings(md.UserId, peer)

	glog.Infof("account.getNotifySettings#12b3ad31 - reply: %s", logger.JsonDebugData(reply))
	return reply, nil
}
