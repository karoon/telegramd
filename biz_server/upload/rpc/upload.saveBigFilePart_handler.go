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
	"github.com/nebulaim/telegramd/biz/dal/dataobject"
	"github.com/nebulaim/telegramd/biz/dal/dao"
)

// upload.saveBigFilePart#de7b673d file_id:long file_part:int file_total_parts:int bytes:bytes = Bool;
func (s *UploadServiceImpl) UploadSaveBigFilePart(ctx context.Context, request *mtproto.TLUploadSaveBigFilePart) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("UploadSaveBigFilePart - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	filePartsDO := &dataobject.FilePartsDO{
	    CreatorUserId: md.UserId,
	    FileId: request.FileId,
	    FilePart: request.FilePart,
	    IsBigFile: 1,
	    FileTotalParts: request.FileTotalParts,
	    Bytes: request.Bytes,
	}
	dao.GetFilePartsDAO(dao.DB_MASTER).Insert(filePartsDO)

	glog.Infof("UploadSaveBigFilePart - reply: {true}")
	return mtproto.ToBool(true), nil
}
