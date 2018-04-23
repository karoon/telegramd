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
	"github.com/nebulaim/telegramd/biz/core/user"
	"github.com/nebulaim/telegramd/biz/core/message"
	"github.com/nebulaim/telegramd/biz/core/chat"
	update2 "github.com/nebulaim/telegramd/biz/core/update"
)

// messages.getPeerDialogs#2d9776b9 peers:Vector<InputPeer> = messages.PeerDialogs;
func (s *MessagesServiceImpl) MessagesGetPeerDialogs(ctx context.Context, request *mtproto.TLMessagesGetPeerDialogs) (*mtproto.Messages_PeerDialogs, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.getPeerDialogs#2d9776b9 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): peers:Vector<InputPeer>
	dialogs := user.GetPinnedDialogs(md.UserId)
	peerDialogs := mtproto.NewTLMessagesPeerDialogs()

	messageIdList := []int32{}
	userIdList := []int32{md.UserId}
	chatIdList := []int32{}

	for _, dialog2 := range dialogs {
		// dialog.Peer
		dialog := dialog2.To_Dialog()
		messageIdList = append(messageIdList, dialog.GetTopMessage())
		peer := base.FromPeer(dialog.GetPeer())
		// TODO(@benqi): 先假设只有PEER_USER
		if peer.PeerType == base.PEER_USER {
			userIdList = append(userIdList, peer.PeerId)
		} else if peer.PeerType == base.PEER_SELF {
			userIdList = append(userIdList, md.UserId)
		} else if peer.PeerType == base.PEER_CHAT {
			chatIdList = append(chatIdList, peer.PeerId)
		}
		peerDialogs.Data2.Dialogs = append(peerDialogs.Data2.Dialogs, dialog.To_Dialog())
	}

	glog.Infof("messageIdList - %v", messageIdList)
	if len(messageIdList) > 0 {
		peerDialogs.SetMessages(message.GetMessagesByPeerAndMessageIdList2(md.UserId, messageIdList))
	}

	users := user.GetUsersBySelfAndIDList(md.UserId, userIdList)
	peerDialogs.SetUsers(users)
	//for _, user := range users {
	//	if user.GetId() == md.UserId {
	//		user.SetSelf(true)
	//	} else {
	//		user.SetSelf(false)
	//	}
	//	user.SetContact(true)
	//	user.SetMutualContact(true)
	//	peerDialogs.Data2.Users = append(peerDialogs.Data2.Users, user.To_User())
	//}

	if len(chatIdList) > 0 {
		peerDialogs.Data2.Chats = chat.GetChatListBySelfAndIDList(md.UserId, chatIdList)
	}

	state := update2.GetUpdatesState(md.AuthId, md.UserId)
	peerDialogs.SetState(state.To_Updates_State())

	glog.Infof("messages.getPeerDialogs#2d9776b9 - reply: %s", logger.JsonDebugData(peerDialogs))
	return peerDialogs.To_Messages_PeerDialogs(), nil
}
