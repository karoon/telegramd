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
	message2 "github.com/nebulaim/telegramd/biz/core/message"
	"github.com/nebulaim/telegramd/biz/core/user"
	"github.com/nebulaim/telegramd/biz_server/sync_client"
	"time"
)

func makeOutboxMessageBySendMessage(fromId int32, peer *base.PeerUtil, request *mtproto.TLMessagesSendMessage) *mtproto.TLMessage {
	var (
		out = true
	)

	if peer.PeerType == base.PEER_USER && peer.PeerId == fromId {
		out = false
	}

	return &mtproto.TLMessage{ Data2: &mtproto.Message_Data{
		Out:          out,
		Silent:       request.GetSilent(),
		FromId:       fromId,
		ToId:         peer.ToPeer(),
		ReplyToMsgId: request.GetReplyToMsgId(),
		Message:      request.GetMessage(),
		Media: &mtproto.MessageMedia{
			Constructor: mtproto.TLConstructor_CRC32_messageMediaEmpty,
			Data2:       &mtproto.MessageMedia_Data{},
		},
		ReplyMarkup: request.GetReplyMarkup(),
		Entities:    request.GetEntities(),
		Date:        int32(time.Now().Unix()),
	}}
}

// 流程：
//  1. 入库: 1）存消息数据，2）分别存到发件箱和收件箱里
//  2. 离线推送
//  3. 在线推送
// messages.sendMessage#fa88																																																						427a flags:# no_webpage:flags.1?true silent:flags.5?true background:flags.6?true clear_draft:flags.7?true peer:InputPeer reply_to_msg_id:flags.0?int message:string random_id:long reply_markup:flags.2?ReplyMarkup entities:flags.3?Vector<MessageEntity> = Updates;
func (s *MessagesServiceImpl) MessagesSendMessage(ctx context.Context, request *mtproto.TLMessagesSendMessage) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.sendMessage#fa88427a - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): check request data invalid
	if request.GetPeer().GetConstructor() ==  mtproto.TLConstructor_CRC32_inputPeerEmpty {
		// retuurn
	}

	// TODO(@benqi): ???
	// request.NoWebpage
	// request.Background
	// request.ClearDraft

	// peer
	var (
		peer *base.PeerUtil
		sentMessage *mtproto.TLUpdateShortSentMessage
		err error
	)

	if request.GetPeer().GetConstructor() == mtproto.TLConstructor_CRC32_inputPeerEmpty {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error("messages.sendMessage#fa88427a - invalid peer", err)
		return nil, err
	}

	// TODO(@benqi): check user or channels's access_hash
	if request.GetPeer().GetConstructor() == mtproto.TLConstructor_CRC32_inputPeerSelf {
		peer = &base.PeerUtil{
			PeerType: base.PEER_USER,
			PeerId:   md.UserId,
		}
	} else {
		peer = base.FromInputPeer(request.GetPeer())
	}

	outboxMessage := makeOutboxMessageBySendMessage(md.UserId, peer, request)
	messageOutbox := message2.CreateMessageOutboxByNew(md.UserId, peer, request.GetRandomId(), outboxMessage.To_Message(), func(messageId int32) {
		// 更新会话信息
		user.CreateOrUpdateByOutbox(md.UserId, peer.PeerType, peer.PeerId, messageId, outboxMessage.GetMentioned(), request.GetClearDraft())
	})

	shortMessage := message2.MessageToUpdateShortMessage(outboxMessage.To_Message())
	state, err := sync_client.GetSyncClient().SyncUpdatesData(md.AuthId, md.SessionId, md.UserId, shortMessage.To_Updates())
	if err != nil {
		glog.Error(err)
		return nil, err
	}

	// 返回给客户端
	sentMessage = message2.MessageToUpdateShortSentMessage(outboxMessage.To_Message())
	sentMessage.SetPts(state.Pts)
	sentMessage.SetPtsCount(state.PtsCount)

	// 收件箱
	if request.GetPeer().GetConstructor() != mtproto.TLConstructor_CRC32_inputPeerSelf {
		inBoxes, _ := messageOutbox.InsertMessageToInbox(md.UserId, peer, func(inBoxUserId, messageId int32) {
			switch peer.PeerType {
			case base.PEER_USER:
				user.CreateOrUpdateByInbox(inBoxUserId, peer.PeerType, md.UserId, messageId, outboxMessage.GetMentioned())
			case base.PEER_CHAT, base.PEER_CHANNEL:
				user.CreateOrUpdateByInbox(inBoxUserId, peer.PeerType, peer.PeerId, messageId, outboxMessage.GetMentioned())
			}
		})

		for i := 0; i < len(inBoxes); i++ {
			switch peer.PeerType {
			case base.PEER_USER:
				shortMessage := message2.MessageToUpdateShortMessage(inBoxes[i].Message)
				sync_client.GetSyncClient().PushToUserUpdatesData(inBoxes[i].UserId, shortMessage.To_Updates())
			case base.PEER_CHAT:
				shortMessage := message2.MessageToUpdateShortChatMessage(inBoxes[i].Message)
				sync_client.GetSyncClient().PushToUserUpdatesData(inBoxes[i].UserId, shortMessage.To_Updates())
			case base.PEER_CHANNEL:
			default:
			}
		}
	}

	glog.Infof("messages.sendMessage#fa88427a - reply: %s", logger.JsonDebugData(sentMessage))
	return sentMessage.To_Updates(), nil
}
