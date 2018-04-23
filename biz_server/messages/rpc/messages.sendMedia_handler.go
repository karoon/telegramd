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
	"time"
	"github.com/nebulaim/telegramd/biz_server/sync_client"
	photo2 "github.com/nebulaim/telegramd/biz/core/photo"
	message2 "github.com/nebulaim/telegramd/biz/core/message"
	"github.com/nebulaim/telegramd/biz/core/user"
	"github.com/nebulaim/telegramd/biz/nbfs_client"
)

func makeMediaByInputMedia(authKeyId int64, media *mtproto.InputMedia) *mtproto.MessageMedia {
	var (
		now = int32(time.Now().Unix())
		// photoModel = model.GetPhotoModel()
		// uuid = helper.NextSnowflakeId()
	)

	switch media.GetConstructor() {
	case mtproto.TLConstructor_CRC32_inputMediaUploadedPhoto:
		uploadedPhoto := media.To_InputMediaUploadedPhoto()
		file := uploadedPhoto.GetFile()

		result, err := nbfs_client.UploadPhotoFile(authKeyId, file)
		// , file.GetData2().GetId(), file.GetData2().GetParts(), file.GetData2().GetName(), file.GetData2().GetMd5Checksum())
		if err != nil {
			glog.Errorf("UploadPhoto error: %v, by %s", err, logger.JsonDebugData(media))
		}

		// fileData := mediaData.GetFile().GetData2()
		photo := &mtproto.TLPhoto{ Data2: &mtproto.Photo_Data{
			Id:          base.NextSnowflakeId(),
			HasStickers: len(uploadedPhoto.GetStickers()) > 0,
			AccessHash:  photo2.GetFileAccessHash(file.GetData2().GetId(), file.GetData2().GetParts()),
			Date:        now,
			Sizes:       result.SizeList,
		}}

		messageMedia := &mtproto.TLMessageMediaPhoto{Data2: &mtproto.MessageMedia_Data{
			Photo_1: photo.To_Photo(),
			Caption: uploadedPhoto.GetCaption(),
			TtlSeconds: uploadedPhoto.GetTtlSeconds(),
		}}
		return messageMedia.To_MessageMedia()
	case mtproto.TLConstructor_CRC32_inputMediaUploadedDocument:
		// inputMediaUploadedDocument#e39621fd flags:# file:InputFile thumb:flags.2?InputFile mime_type:string attributes:Vector<DocumentAttribute> caption:string stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
		uploadedDocument := media.To_InputMediaUploadedDocument()
		messageMedia, _ := nbfs_client.UploadedDocumentMedia(authKeyId, uploadedDocument)
		return messageMedia.To_MessageMedia()
		// id:InputDocument caption:string ttl_seconds:flags.0?int
	}

	return mtproto.NewTLMessageMediaEmpty().To_MessageMedia()
}

func makeOutboxMessageBySendMedia(authKeyId int64, fromId int32, peer *base.PeerUtil, request *mtproto.TLMessagesSendMedia) *mtproto.TLMessage {
	return &mtproto.TLMessage{ Data2: &mtproto.Message_Data{
		Out:          true,
		Silent:       request.GetSilent(),
		FromId:       fromId,
		ToId:         peer.ToPeer(),
		ReplyToMsgId: request.GetReplyToMsgId(),
		Media: 		  makeMediaByInputMedia(authKeyId, request.GetMedia()),
		ReplyMarkup: request.GetReplyMarkup(),
		Date:        int32(time.Now().Unix()),
	}}
}

func makeUpdateNewMessageUpdates(selfUserId int32, message *mtproto.Message) *mtproto.TLUpdates {
	userIdList, _, _ := message2.PickAllIDListByMessages([]*mtproto.Message{message})
	userList := user.GetUsersBySelfAndIDList(selfUserId, userIdList)
	updateNew := &mtproto.TLUpdateNewMessage{Data2: &mtproto.Update_Data{
		Message_1: message,
	}}
	return &mtproto.TLUpdates{Data2: &mtproto.Updates_Data{
		Updates: []*mtproto.Update{updateNew.To_Update()},
		Users:   userList,
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}}

}

// TODO(@benqi): check error
func SetupUpdatesState(state *mtproto.ClientUpdatesState, updates *mtproto.TLUpdates) *mtproto.TLUpdates {
	pts := state.GetPts() - state.GetPtsCount() + 1

	for _, update := range updates.GetUpdates() {
		switch update.GetConstructor() {
		case mtproto.TLConstructor_CRC32_updateNewMessage,
			mtproto.TLConstructor_CRC32_updateDeleteMessages,
			mtproto.TLConstructor_CRC32_updateReadHistoryOutbox,
			mtproto.TLConstructor_CRC32_updateReadHistoryInbox,
			mtproto.TLConstructor_CRC32_updateWebPage,
			mtproto.TLConstructor_CRC32_updateReadMessagesContents,
			mtproto.TLConstructor_CRC32_updateEditMessage:

			//if pts >= state.GetPtsCount() {
			//	return false
			//}
			//
			update.Data2.Pts = pts
			update.Data2.PtsCount = 1
			pts += 1
		}
	}

	return updates
	// return pts == state.GetPtsCount()
}

// messages.sendMedia#c8f16791 flags:# silent:flags.5?true background:flags.6?true clear_draft:flags.7?true peer:InputPeer reply_to_msg_id:flags.0?int media:InputMedia random_id:long reply_markup:flags.2?ReplyMarkup = Updates;
func (s *MessagesServiceImpl) MessagesSendMedia(ctx context.Context, request *mtproto.TLMessagesSendMedia) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("messages.sendMedia#c8f16791 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): ???
	// request.NoWebpage
	// request.Background

	// peer
	var (
		peer *base.PeerUtil
		err error
	)

	if request.GetPeer().GetConstructor() == mtproto.TLConstructor_CRC32_inputPeerEmpty {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
		glog.Error("messages.sendMedia#c8f16791 - invalid peer", err)
		return nil, err
	}
	// TODO(@benqi): check user or channels's access_hash

	// peer = helper.FromInputPeer2(md.UserId, request.GetPeer())
	if request.GetPeer().GetConstructor() == mtproto.TLConstructor_CRC32_inputPeerSelf {
		peer = &base.PeerUtil{
			PeerType: base.PEER_USER,
			PeerId:   md.UserId,
		}
	} else {
		peer = base.FromInputPeer(request.GetPeer())
	}

	/////////////////////////////////////////////////////////////////////////////////////
	// 发件箱
	// sendMessageToOutbox
	outboxMessage := makeOutboxMessageBySendMedia(md.AuthId, md.UserId, peer, request)
	messageOutbox := message2.CreateMessageOutboxByNew(md.UserId, peer, request.GetRandomId(), outboxMessage.To_Message(), func(messageId int32) {
		// 更新会话信息
		user.CreateOrUpdateByOutbox(md.UserId, peer.PeerType, peer.PeerId, messageId, outboxMessage.GetMentioned(), request.GetClearDraft())
	})

	syncUpdates := makeUpdateNewMessageUpdates(md.UserId, messageOutbox.Message)
	state, err := sync_client.GetSyncClient().SyncUpdatesData(md.AuthId, md.SessionId, md.UserId, syncUpdates.To_Updates())
	if err != nil {
		return nil, err
	}

	reply := SetupUpdatesState(state, syncUpdates)
	updateMessageID := &mtproto.TLUpdateMessageID{Data2: &mtproto.Update_Data{
		Id_4:     messageOutbox.MessageId,
		RandomId: request.GetRandomId(),
	}}
	updateList := []*mtproto.Update{updateMessageID.To_Update()}
	updateList = append(updateList, reply.GetUpdates()...)
	reply.SetUpdates(updateList)

	/////////////////////////////////////////////////////////////////////////////////////
	// 收件箱
	if request.GetPeer().GetConstructor() != mtproto.TLConstructor_CRC32_inputPeerSelf {
		inBoxes, _ := messageOutbox.InsertMessageToInbox(md.UserId, peer, func(inBoxUserId, messageId int32) {
			// 更新会话信息
			switch peer.PeerType {
			case base.PEER_USER:
				user.CreateOrUpdateByInbox(inBoxUserId, peer.PeerType, md.UserId, messageId, outboxMessage.GetMentioned())
			case base.PEER_CHAT, base.PEER_CHANNEL:
				user.CreateOrUpdateByInbox(inBoxUserId, peer.PeerType, peer.PeerId, messageId, outboxMessage.GetMentioned())
			}
		})

		for i := 0; i < len(inBoxes); i++ {
			syncUpdates = makeUpdateNewMessageUpdates(inBoxes[i].UserId, inBoxes[i].Message)
			sync_client.GetSyncClient().PushToUserUpdatesData(inBoxes[i].UserId, syncUpdates.To_Updates())
		}
	}

	glog.Infof("messages.sendMedia#c8f16791 - reply: %s", logger.JsonDebugData(reply))
	return reply.To_Updates(), nil
}
