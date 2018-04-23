/*
 *  Copyright (c) 2018, https://github.com/nebulaim
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
	"github.com/nebulaim/telegramd/nbfs/biz/core/file"
)

// upload.saveBigFilePart#de7b673d file_id:long file_part:int file_total_parts:int bytes:bytes = Bool;
func (s *UploadServiceImpl) UploadSaveBigFilePart(ctx context.Context, request *mtproto.TLUploadSaveBigFilePart) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("upload.saveFilePart#b304a621 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Check file_total_parts >= bigFileSize/kMaxFileSize, Check len(bytes) <= kMaxFilePartSize, Check file_part <= file_total_parts

	filePartLogic, err := file.MakeFilePartData(md.AuthId, request.FileId, request.FilePart == 0, true)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	err = filePartLogic.SaveFilePart(request.FilePart, request.Bytes)
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	glog.Infof("upload.saveBigFilePart#de7b673d - reply: {true}")
	return mtproto.ToBool(true), nil

}
