// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zproto_sync.proto

package mtproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 同步类型
type SyncType int32

const (
	SyncType_SYNC_TYPE_UNKNOWN            SyncType = 0
	SyncType_SYNC_TYPE_USER               SyncType = 1
	SyncType_SYNC_TYPE_USER_NOTME         SyncType = 2
	SyncType_SYNC_TYPE_AUTH_KEY           SyncType = 3
	SyncType_SYNC_TYPE_AUTH_KEY_USER      SyncType = 4
	SyncType_SYNC_TYPE_AUTH_KEY_USERNOTME SyncType = 5
)

var SyncType_name = map[int32]string{
	0: "SYNC_TYPE_UNKNOWN",
	1: "SYNC_TYPE_USER",
	2: "SYNC_TYPE_USER_NOTME",
	3: "SYNC_TYPE_AUTH_KEY",
	4: "SYNC_TYPE_AUTH_KEY_USER",
	5: "SYNC_TYPE_AUTH_KEY_USERNOTME",
}
var SyncType_value = map[string]int32{
	"SYNC_TYPE_UNKNOWN":            0,
	"SYNC_TYPE_USER":               1,
	"SYNC_TYPE_USER_NOTME":         2,
	"SYNC_TYPE_AUTH_KEY":           3,
	"SYNC_TYPE_AUTH_KEY_USER":      4,
	"SYNC_TYPE_AUTH_KEY_USERNOTME": 5,
}

func (x SyncType) String() string {
	return proto.EnumName(SyncType_name, int32(x))
}
func (SyncType) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

type MessageDataEmpty struct {
}

func (m *MessageDataEmpty) Reset()                    { *m = MessageDataEmpty{} }
func (m *MessageDataEmpty) String() string            { return proto.CompactTextString(m) }
func (*MessageDataEmpty) ProtoMessage()               {}
func (*MessageDataEmpty) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

type MessageData struct {
	MessageType int32 `protobuf:"varint,1,opt,name=message_type,json=messageType" json:"message_type,omitempty"`
}

func (m *MessageData) Reset()                    { *m = MessageData{} }
func (m *MessageData) String() string            { return proto.CompactTextString(m) }
func (*MessageData) ProtoMessage()               {}
func (*MessageData) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *MessageData) GetMessageType() int32 {
	if m != nil {
		return m.MessageType
	}
	return 0
}

type VoidRsp struct {
}

func (m *VoidRsp) Reset()                    { *m = VoidRsp{} }
func (m *VoidRsp) String() string            { return proto.CompactTextString(m) }
func (*VoidRsp) ProtoMessage()               {}
func (*VoidRsp) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

type PushClientID struct {
	AuthKeyId       int64 `protobuf:"varint,1,opt,name=auth_key_id,json=authKeyId" json:"auth_key_id,omitempty"`
	SessionId       int64 `protobuf:"varint,2,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	NetlibSessionId int64 `protobuf:"varint,3,opt,name=netlib_session_id,json=netlibSessionId" json:"netlib_session_id,omitempty"`
}

func (m *PushClientID) Reset()                    { *m = PushClientID{} }
func (m *PushClientID) String() string            { return proto.CompactTextString(m) }
func (*PushClientID) ProtoMessage()               {}
func (*PushClientID) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

func (m *PushClientID) GetAuthKeyId() int64 {
	if m != nil {
		return m.AuthKeyId
	}
	return 0
}

func (m *PushClientID) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *PushClientID) GetNetlibSessionId() int64 {
	if m != nil {
		return m.NetlibSessionId
	}
	return 0
}

// PushMessage state
type ClientUpdatesState struct {
	Pts      int32 `protobuf:"varint,1,opt,name=pts" json:"pts,omitempty"`
	PtsCount int32 `protobuf:"varint,2,opt,name=pts_count,json=ptsCount" json:"pts_count,omitempty"`
	Qts      int32 `protobuf:"varint,3,opt,name=qts" json:"qts,omitempty"`
	QtsCount int32 `protobuf:"varint,4,opt,name=qts_count,json=qtsCount" json:"qts_count,omitempty"`
	Seq      int32 `protobuf:"varint,5,opt,name=seq" json:"seq,omitempty"`
	SeqStart int32 `protobuf:"varint,6,opt,name=seq_start,json=seqStart" json:"seq_start,omitempty"`
	Date     int32 `protobuf:"varint,7,opt,name=date" json:"date,omitempty"`
}

func (m *ClientUpdatesState) Reset()                    { *m = ClientUpdatesState{} }
func (m *ClientUpdatesState) String() string            { return proto.CompactTextString(m) }
func (*ClientUpdatesState) ProtoMessage()               {}
func (*ClientUpdatesState) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

func (m *ClientUpdatesState) GetPts() int32 {
	if m != nil {
		return m.Pts
	}
	return 0
}

func (m *ClientUpdatesState) GetPtsCount() int32 {
	if m != nil {
		return m.PtsCount
	}
	return 0
}

func (m *ClientUpdatesState) GetQts() int32 {
	if m != nil {
		return m.Qts
	}
	return 0
}

func (m *ClientUpdatesState) GetQtsCount() int32 {
	if m != nil {
		return m.QtsCount
	}
	return 0
}

func (m *ClientUpdatesState) GetSeq() int32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *ClientUpdatesState) GetSeqStart() int32 {
	if m != nil {
		return m.SeqStart
	}
	return 0
}

func (m *ClientUpdatesState) GetDate() int32 {
	if m != nil {
		return m.Date
	}
	return 0
}

// /////////////////////////////////////////////////////////////////////
// SERVER_AUTH_REQ
type ConnectToSessionServerReq struct {
}

func (m *ConnectToSessionServerReq) Reset()                    { *m = ConnectToSessionServerReq{} }
func (m *ConnectToSessionServerReq) String() string            { return proto.CompactTextString(m) }
func (*ConnectToSessionServerReq) ProtoMessage()               {}
func (*ConnectToSessionServerReq) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{5} }

type SessionServerConnectedRsp struct {
	ServerId   int32  `protobuf:"varint,1,opt,name=server_id,json=serverId" json:"server_id,omitempty"`
	ServerName string `protobuf:"bytes,2,opt,name=server_name,json=serverName" json:"server_name,omitempty"`
}

func (m *SessionServerConnectedRsp) Reset()                    { *m = SessionServerConnectedRsp{} }
func (m *SessionServerConnectedRsp) String() string            { return proto.CompactTextString(m) }
func (*SessionServerConnectedRsp) ProtoMessage()               {}
func (*SessionServerConnectedRsp) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{6} }

func (m *SessionServerConnectedRsp) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *SessionServerConnectedRsp) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

// PushUpdatesData --> VoidRsp
type PushUpdatesData struct {
	ClientId      *PushClientID       `protobuf:"bytes,1,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	State         *ClientUpdatesState `protobuf:"bytes,2,opt,name=state" json:"state,omitempty"`
	RawDataHeader uint32              `protobuf:"varint,3,opt,name=raw_data_header,json=rawDataHeader" json:"raw_data_header,omitempty"`
	RawData       []byte              `protobuf:"bytes,4,opt,name=raw_data,json=rawData,proto3" json:"raw_data,omitempty"`
}

func (m *PushUpdatesData) Reset()                    { *m = PushUpdatesData{} }
func (m *PushUpdatesData) String() string            { return proto.CompactTextString(m) }
func (*PushUpdatesData) ProtoMessage()               {}
func (*PushUpdatesData) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{7} }

func (m *PushUpdatesData) GetClientId() *PushClientID {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *PushUpdatesData) GetState() *ClientUpdatesState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *PushUpdatesData) GetRawDataHeader() uint32 {
	if m != nil {
		return m.RawDataHeader
	}
	return 0
}

func (m *PushUpdatesData) GetRawData() []byte {
	if m != nil {
		return m.RawData
	}
	return nil
}

type SyncShortMessageRequest struct {
	ClientId     *PushClientID `protobuf:"bytes,1,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	PushType     SyncType      `protobuf:"varint,2,opt,name=push_type,json=pushType,enum=mtproto.SyncType" json:"push_type,omitempty"`
	PushtoUserId int32         `protobuf:"varint,3,opt,name=pushto_user_id,json=pushtoUserId" json:"pushto_user_id,omitempty"`
	// int32 peer_type = 4;
	PeerId   int32                 `protobuf:"varint,4,opt,name=peer_id,json=peerId" json:"peer_id,omitempty"`
	PushData *TLUpdateShortMessage `protobuf:"bytes,5,opt,name=push_data,json=pushData" json:"push_data,omitempty"`
}

func (m *SyncShortMessageRequest) Reset()                    { *m = SyncShortMessageRequest{} }
func (m *SyncShortMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*SyncShortMessageRequest) ProtoMessage()               {}
func (*SyncShortMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{8} }

func (m *SyncShortMessageRequest) GetClientId() *PushClientID {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *SyncShortMessageRequest) GetPushType() SyncType {
	if m != nil {
		return m.PushType
	}
	return SyncType_SYNC_TYPE_UNKNOWN
}

func (m *SyncShortMessageRequest) GetPushtoUserId() int32 {
	if m != nil {
		return m.PushtoUserId
	}
	return 0
}

func (m *SyncShortMessageRequest) GetPeerId() int32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *SyncShortMessageRequest) GetPushData() *TLUpdateShortMessage {
	if m != nil {
		return m.PushData
	}
	return nil
}

// UpdateShortMessageRequest --> DeliveryRsp
type UpdateShortMessageRequest struct {
	PushType     SyncType `protobuf:"varint,1,opt,name=push_type,json=pushType,enum=mtproto.SyncType" json:"push_type,omitempty"`
	PushtoUserId int32    `protobuf:"varint,2,opt,name=pushto_user_id,json=pushtoUserId" json:"pushto_user_id,omitempty"`
	// int32 peer_type = 4;
	PeerId   int32                 `protobuf:"varint,3,opt,name=peer_id,json=peerId" json:"peer_id,omitempty"`
	PushData *TLUpdateShortMessage `protobuf:"bytes,4,opt,name=push_data,json=pushData" json:"push_data,omitempty"`
}

func (m *UpdateShortMessageRequest) Reset()                    { *m = UpdateShortMessageRequest{} }
func (m *UpdateShortMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateShortMessageRequest) ProtoMessage()               {}
func (*UpdateShortMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{9} }

func (m *UpdateShortMessageRequest) GetPushType() SyncType {
	if m != nil {
		return m.PushType
	}
	return SyncType_SYNC_TYPE_UNKNOWN
}

func (m *UpdateShortMessageRequest) GetPushtoUserId() int32 {
	if m != nil {
		return m.PushtoUserId
	}
	return 0
}

func (m *UpdateShortMessageRequest) GetPeerId() int32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *UpdateShortMessageRequest) GetPushData() *TLUpdateShortMessage {
	if m != nil {
		return m.PushData
	}
	return nil
}

// UpdatShortChatMessageRequest --> DeliveryRsp
type SyncShortChatMessageRequest struct {
	ClientId     *PushClientID             `protobuf:"bytes,1,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	PushType     SyncType                  `protobuf:"varint,2,opt,name=push_type,json=pushType,enum=mtproto.SyncType" json:"push_type,omitempty"`
	PushtoUserId int32                     `protobuf:"varint,3,opt,name=pushto_user_id,json=pushtoUserId" json:"pushto_user_id,omitempty"`
	PeerChatId   int32                     `protobuf:"varint,4,opt,name=peer_chat_id,json=peerChatId" json:"peer_chat_id,omitempty"`
	PushData     *TLUpdateShortChatMessage `protobuf:"bytes,5,opt,name=push_data,json=pushData" json:"push_data,omitempty"`
}

func (m *SyncShortChatMessageRequest) Reset()                    { *m = SyncShortChatMessageRequest{} }
func (m *SyncShortChatMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*SyncShortChatMessageRequest) ProtoMessage()               {}
func (*SyncShortChatMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{10} }

func (m *SyncShortChatMessageRequest) GetClientId() *PushClientID {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *SyncShortChatMessageRequest) GetPushType() SyncType {
	if m != nil {
		return m.PushType
	}
	return SyncType_SYNC_TYPE_UNKNOWN
}

func (m *SyncShortChatMessageRequest) GetPushtoUserId() int32 {
	if m != nil {
		return m.PushtoUserId
	}
	return 0
}

func (m *SyncShortChatMessageRequest) GetPeerChatId() int32 {
	if m != nil {
		return m.PeerChatId
	}
	return 0
}

func (m *SyncShortChatMessageRequest) GetPushData() *TLUpdateShortChatMessage {
	if m != nil {
		return m.PushData
	}
	return nil
}

type UpdateShortChatMessageRequest struct {
	PushType     SyncType                  `protobuf:"varint,1,opt,name=push_type,json=pushType,enum=mtproto.SyncType" json:"push_type,omitempty"`
	PushtoUserId int32                     `protobuf:"varint,2,opt,name=pushto_user_id,json=pushtoUserId" json:"pushto_user_id,omitempty"`
	PeerChatId   int32                     `protobuf:"varint,3,opt,name=peer_chat_id,json=peerChatId" json:"peer_chat_id,omitempty"`
	PushData     *TLUpdateShortChatMessage `protobuf:"bytes,4,opt,name=push_data,json=pushData" json:"push_data,omitempty"`
}

func (m *UpdateShortChatMessageRequest) Reset()                    { *m = UpdateShortChatMessageRequest{} }
func (m *UpdateShortChatMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateShortChatMessageRequest) ProtoMessage()               {}
func (*UpdateShortChatMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{11} }

func (m *UpdateShortChatMessageRequest) GetPushType() SyncType {
	if m != nil {
		return m.PushType
	}
	return SyncType_SYNC_TYPE_UNKNOWN
}

func (m *UpdateShortChatMessageRequest) GetPushtoUserId() int32 {
	if m != nil {
		return m.PushtoUserId
	}
	return 0
}

func (m *UpdateShortChatMessageRequest) GetPeerChatId() int32 {
	if m != nil {
		return m.PeerChatId
	}
	return 0
}

func (m *UpdateShortChatMessageRequest) GetPushData() *TLUpdateShortChatMessage {
	if m != nil {
		return m.PushData
	}
	return nil
}

// UpdatShortChatMessageRequest --> DeliveryRsp
type SyncUpdateMessageRequest struct {
	ClientId     *PushClientID `protobuf:"bytes,1,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	PushType     SyncType      `protobuf:"varint,2,opt,name=push_type,json=pushType,enum=mtproto.SyncType" json:"push_type,omitempty"`
	PushtoUserId int32         `protobuf:"varint,3,opt,name=pushto_user_id,json=pushtoUserId" json:"pushto_user_id,omitempty"`
	PeerType     int32         `protobuf:"varint,4,opt,name=peer_type,json=peerType" json:"peer_type,omitempty"`
	PeerId       int32         `protobuf:"varint,5,opt,name=peer_id,json=peerId" json:"peer_id,omitempty"`
	PushData     *Update       `protobuf:"bytes,6,opt,name=push_data,json=pushData" json:"push_data,omitempty"`
}

func (m *SyncUpdateMessageRequest) Reset()                    { *m = SyncUpdateMessageRequest{} }
func (m *SyncUpdateMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*SyncUpdateMessageRequest) ProtoMessage()               {}
func (*SyncUpdateMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{12} }

func (m *SyncUpdateMessageRequest) GetClientId() *PushClientID {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *SyncUpdateMessageRequest) GetPushType() SyncType {
	if m != nil {
		return m.PushType
	}
	return SyncType_SYNC_TYPE_UNKNOWN
}

func (m *SyncUpdateMessageRequest) GetPushtoUserId() int32 {
	if m != nil {
		return m.PushtoUserId
	}
	return 0
}

func (m *SyncUpdateMessageRequest) GetPeerType() int32 {
	if m != nil {
		return m.PeerType
	}
	return 0
}

func (m *SyncUpdateMessageRequest) GetPeerId() int32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *SyncUpdateMessageRequest) GetPushData() *Update {
	if m != nil {
		return m.PushData
	}
	return nil
}

type PushUpdateMessageRequest struct {
	PushType     SyncType `protobuf:"varint,1,opt,name=push_type,json=pushType,enum=mtproto.SyncType" json:"push_type,omitempty"`
	PushtoUserId int32    `protobuf:"varint,2,opt,name=pushto_user_id,json=pushtoUserId" json:"pushto_user_id,omitempty"`
	PeerType     int32    `protobuf:"varint,3,opt,name=peer_type,json=peerType" json:"peer_type,omitempty"`
	PeerId       int32    `protobuf:"varint,4,opt,name=peer_id,json=peerId" json:"peer_id,omitempty"`
	PushData     *Update  `protobuf:"bytes,5,opt,name=push_data,json=pushData" json:"push_data,omitempty"`
}

func (m *PushUpdateMessageRequest) Reset()                    { *m = PushUpdateMessageRequest{} }
func (m *PushUpdateMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*PushUpdateMessageRequest) ProtoMessage()               {}
func (*PushUpdateMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{13} }

func (m *PushUpdateMessageRequest) GetPushType() SyncType {
	if m != nil {
		return m.PushType
	}
	return SyncType_SYNC_TYPE_UNKNOWN
}

func (m *PushUpdateMessageRequest) GetPushtoUserId() int32 {
	if m != nil {
		return m.PushtoUserId
	}
	return 0
}

func (m *PushUpdateMessageRequest) GetPeerType() int32 {
	if m != nil {
		return m.PeerType
	}
	return 0
}

func (m *PushUpdateMessageRequest) GetPeerId() int32 {
	if m != nil {
		return m.PeerId
	}
	return 0
}

func (m *PushUpdateMessageRequest) GetPushData() *Update {
	if m != nil {
		return m.PushData
	}
	return nil
}

type PushUpdatesRequest struct {
	PushType     SyncType   `protobuf:"varint,1,opt,name=push_type,json=pushType,enum=mtproto.SyncType" json:"push_type,omitempty"`
	PushtoUserId int32      `protobuf:"varint,2,opt,name=pushto_user_id,json=pushtoUserId" json:"pushto_user_id,omitempty"`
	PushData     *TLUpdates `protobuf:"bytes,5,opt,name=push_data,json=pushData" json:"push_data,omitempty"`
}

func (m *PushUpdatesRequest) Reset()                    { *m = PushUpdatesRequest{} }
func (m *PushUpdatesRequest) String() string            { return proto.CompactTextString(m) }
func (*PushUpdatesRequest) ProtoMessage()               {}
func (*PushUpdatesRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{14} }

func (m *PushUpdatesRequest) GetPushType() SyncType {
	if m != nil {
		return m.PushType
	}
	return SyncType_SYNC_TYPE_UNKNOWN
}

func (m *PushUpdatesRequest) GetPushtoUserId() int32 {
	if m != nil {
		return m.PushtoUserId
	}
	return 0
}

func (m *PushUpdatesRequest) GetPushData() *TLUpdates {
	if m != nil {
		return m.PushData
	}
	return nil
}

type PushUpdateShortRequest struct {
	PushType     SyncType       `protobuf:"varint,1,opt,name=push_type,json=pushType,enum=mtproto.SyncType" json:"push_type,omitempty"`
	PushtoUserId int32          `protobuf:"varint,2,opt,name=pushto_user_id,json=pushtoUserId" json:"pushto_user_id,omitempty"`
	PushData     *TLUpdateShort `protobuf:"bytes,5,opt,name=push_data,json=pushData" json:"push_data,omitempty"`
}

func (m *PushUpdateShortRequest) Reset()                    { *m = PushUpdateShortRequest{} }
func (m *PushUpdateShortRequest) String() string            { return proto.CompactTextString(m) }
func (*PushUpdateShortRequest) ProtoMessage()               {}
func (*PushUpdateShortRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{15} }

func (m *PushUpdateShortRequest) GetPushType() SyncType {
	if m != nil {
		return m.PushType
	}
	return SyncType_SYNC_TYPE_UNKNOWN
}

func (m *PushUpdateShortRequest) GetPushtoUserId() int32 {
	if m != nil {
		return m.PushtoUserId
	}
	return 0
}

func (m *PushUpdateShortRequest) GetPushData() *TLUpdateShort {
	if m != nil {
		return m.PushData
	}
	return nil
}

func init() {
	proto.RegisterType((*MessageDataEmpty)(nil), "mtproto.MessageDataEmpty")
	proto.RegisterType((*MessageData)(nil), "mtproto.MessageData")
	proto.RegisterType((*VoidRsp)(nil), "mtproto.VoidRsp")
	proto.RegisterType((*PushClientID)(nil), "mtproto.PushClientID")
	proto.RegisterType((*ClientUpdatesState)(nil), "mtproto.ClientUpdatesState")
	proto.RegisterType((*ConnectToSessionServerReq)(nil), "mtproto.ConnectToSessionServerReq")
	proto.RegisterType((*SessionServerConnectedRsp)(nil), "mtproto.SessionServerConnectedRsp")
	proto.RegisterType((*PushUpdatesData)(nil), "mtproto.PushUpdatesData")
	proto.RegisterType((*SyncShortMessageRequest)(nil), "mtproto.SyncShortMessageRequest")
	proto.RegisterType((*UpdateShortMessageRequest)(nil), "mtproto.UpdateShortMessageRequest")
	proto.RegisterType((*SyncShortChatMessageRequest)(nil), "mtproto.SyncShortChatMessageRequest")
	proto.RegisterType((*UpdateShortChatMessageRequest)(nil), "mtproto.UpdateShortChatMessageRequest")
	proto.RegisterType((*SyncUpdateMessageRequest)(nil), "mtproto.SyncUpdateMessageRequest")
	proto.RegisterType((*PushUpdateMessageRequest)(nil), "mtproto.PushUpdateMessageRequest")
	proto.RegisterType((*PushUpdatesRequest)(nil), "mtproto.PushUpdatesRequest")
	proto.RegisterType((*PushUpdateShortRequest)(nil), "mtproto.PushUpdateShortRequest")
	proto.RegisterEnum("mtproto.SyncType", SyncType_name, SyncType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RPCSync service

type RPCSyncClient interface {
	SyncUpdateShortMessage(ctx context.Context, in *SyncShortMessageRequest, opts ...grpc.CallOption) (*ClientUpdatesState, error)
	PushUpdateShortMessage(ctx context.Context, in *UpdateShortMessageRequest, opts ...grpc.CallOption) (*VoidRsp, error)
	SyncUpdateShortChatMessage(ctx context.Context, in *SyncShortChatMessageRequest, opts ...grpc.CallOption) (*ClientUpdatesState, error)
	PushUpdateShortChatMessage(ctx context.Context, in *UpdateShortChatMessageRequest, opts ...grpc.CallOption) (*VoidRsp, error)
	// rpc DeliveryUpdates2(UpdatesRequest) returns (DeliveryRsp);
	SyncUpdateMessageData(ctx context.Context, in *SyncUpdateMessageRequest, opts ...grpc.CallOption) (*ClientUpdatesState, error)
	PushUpdateMessageData(ctx context.Context, in *PushUpdateMessageRequest, opts ...grpc.CallOption) (*VoidRsp, error)
	PushUpdatesData(ctx context.Context, in *PushUpdatesRequest, opts ...grpc.CallOption) (*VoidRsp, error)
	PushUpdateShortData(ctx context.Context, in *PushUpdateShortRequest, opts ...grpc.CallOption) (*VoidRsp, error)
}

type rPCSyncClient struct {
	cc *grpc.ClientConn
}

func NewRPCSyncClient(cc *grpc.ClientConn) RPCSyncClient {
	return &rPCSyncClient{cc}
}

func (c *rPCSyncClient) SyncUpdateShortMessage(ctx context.Context, in *SyncShortMessageRequest, opts ...grpc.CallOption) (*ClientUpdatesState, error) {
	out := new(ClientUpdatesState)
	err := grpc.Invoke(ctx, "/mtproto.RPCSync/SyncUpdateShortMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSyncClient) PushUpdateShortMessage(ctx context.Context, in *UpdateShortMessageRequest, opts ...grpc.CallOption) (*VoidRsp, error) {
	out := new(VoidRsp)
	err := grpc.Invoke(ctx, "/mtproto.RPCSync/PushUpdateShortMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSyncClient) SyncUpdateShortChatMessage(ctx context.Context, in *SyncShortChatMessageRequest, opts ...grpc.CallOption) (*ClientUpdatesState, error) {
	out := new(ClientUpdatesState)
	err := grpc.Invoke(ctx, "/mtproto.RPCSync/SyncUpdateShortChatMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSyncClient) PushUpdateShortChatMessage(ctx context.Context, in *UpdateShortChatMessageRequest, opts ...grpc.CallOption) (*VoidRsp, error) {
	out := new(VoidRsp)
	err := grpc.Invoke(ctx, "/mtproto.RPCSync/PushUpdateShortChatMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSyncClient) SyncUpdateMessageData(ctx context.Context, in *SyncUpdateMessageRequest, opts ...grpc.CallOption) (*ClientUpdatesState, error) {
	out := new(ClientUpdatesState)
	err := grpc.Invoke(ctx, "/mtproto.RPCSync/SyncUpdateMessageData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSyncClient) PushUpdateMessageData(ctx context.Context, in *PushUpdateMessageRequest, opts ...grpc.CallOption) (*VoidRsp, error) {
	out := new(VoidRsp)
	err := grpc.Invoke(ctx, "/mtproto.RPCSync/PushUpdateMessageData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSyncClient) PushUpdatesData(ctx context.Context, in *PushUpdatesRequest, opts ...grpc.CallOption) (*VoidRsp, error) {
	out := new(VoidRsp)
	err := grpc.Invoke(ctx, "/mtproto.RPCSync/PushUpdatesData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSyncClient) PushUpdateShortData(ctx context.Context, in *PushUpdateShortRequest, opts ...grpc.CallOption) (*VoidRsp, error) {
	out := new(VoidRsp)
	err := grpc.Invoke(ctx, "/mtproto.RPCSync/PushUpdateShortData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RPCSync service

type RPCSyncServer interface {
	SyncUpdateShortMessage(context.Context, *SyncShortMessageRequest) (*ClientUpdatesState, error)
	PushUpdateShortMessage(context.Context, *UpdateShortMessageRequest) (*VoidRsp, error)
	SyncUpdateShortChatMessage(context.Context, *SyncShortChatMessageRequest) (*ClientUpdatesState, error)
	PushUpdateShortChatMessage(context.Context, *UpdateShortChatMessageRequest) (*VoidRsp, error)
	// rpc DeliveryUpdates2(UpdatesRequest) returns (DeliveryRsp);
	SyncUpdateMessageData(context.Context, *SyncUpdateMessageRequest) (*ClientUpdatesState, error)
	PushUpdateMessageData(context.Context, *PushUpdateMessageRequest) (*VoidRsp, error)
	PushUpdatesData(context.Context, *PushUpdatesRequest) (*VoidRsp, error)
	PushUpdateShortData(context.Context, *PushUpdateShortRequest) (*VoidRsp, error)
}

func RegisterRPCSyncServer(s *grpc.Server, srv RPCSyncServer) {
	s.RegisterService(&_RPCSync_serviceDesc, srv)
}

func _RPCSync_SyncUpdateShortMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncShortMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).SyncUpdateShortMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCSync/SyncUpdateShortMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).SyncUpdateShortMessage(ctx, req.(*SyncShortMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSync_PushUpdateShortMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateShortMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).PushUpdateShortMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCSync/PushUpdateShortMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).PushUpdateShortMessage(ctx, req.(*UpdateShortMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSync_SyncUpdateShortChatMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncShortChatMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).SyncUpdateShortChatMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCSync/SyncUpdateShortChatMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).SyncUpdateShortChatMessage(ctx, req.(*SyncShortChatMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSync_PushUpdateShortChatMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateShortChatMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).PushUpdateShortChatMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCSync/PushUpdateShortChatMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).PushUpdateShortChatMessage(ctx, req.(*UpdateShortChatMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSync_SyncUpdateMessageData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncUpdateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).SyncUpdateMessageData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCSync/SyncUpdateMessageData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).SyncUpdateMessageData(ctx, req.(*SyncUpdateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSync_PushUpdateMessageData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushUpdateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).PushUpdateMessageData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCSync/PushUpdateMessageData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).PushUpdateMessageData(ctx, req.(*PushUpdateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSync_PushUpdatesData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushUpdatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).PushUpdatesData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCSync/PushUpdatesData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).PushUpdatesData(ctx, req.(*PushUpdatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSync_PushUpdateShortData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushUpdateShortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).PushUpdateShortData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCSync/PushUpdateShortData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).PushUpdateShortData(ctx, req.(*PushUpdateShortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCSync_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mtproto.RPCSync",
	HandlerType: (*RPCSyncServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncUpdateShortMessage",
			Handler:    _RPCSync_SyncUpdateShortMessage_Handler,
		},
		{
			MethodName: "PushUpdateShortMessage",
			Handler:    _RPCSync_PushUpdateShortMessage_Handler,
		},
		{
			MethodName: "SyncUpdateShortChatMessage",
			Handler:    _RPCSync_SyncUpdateShortChatMessage_Handler,
		},
		{
			MethodName: "PushUpdateShortChatMessage",
			Handler:    _RPCSync_PushUpdateShortChatMessage_Handler,
		},
		{
			MethodName: "SyncUpdateMessageData",
			Handler:    _RPCSync_SyncUpdateMessageData_Handler,
		},
		{
			MethodName: "PushUpdateMessageData",
			Handler:    _RPCSync_PushUpdateMessageData_Handler,
		},
		{
			MethodName: "PushUpdatesData",
			Handler:    _RPCSync_PushUpdatesData_Handler,
		},
		{
			MethodName: "PushUpdateShortData",
			Handler:    _RPCSync_PushUpdateShortData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zproto_sync.proto",
}

func init() { proto.RegisterFile("zproto_sync.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 993 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0xdd, 0x6e, 0x1a, 0x47,
	0x14, 0xee, 0x02, 0x6b, 0xe0, 0x40, 0x6c, 0x3c, 0x8e, 0x6d, 0x6c, 0x9a, 0xc4, 0x5d, 0xa5, 0x51,
	0x15, 0x55, 0x28, 0xa5, 0xbd, 0xeb, 0x45, 0xdb, 0x10, 0x24, 0x23, 0x27, 0x04, 0x2d, 0x10, 0xcb,
	0xbd, 0x19, 0x8d, 0xd9, 0x51, 0x41, 0x35, 0xfb, 0xc3, 0x0c, 0x8d, 0xe8, 0x13, 0x54, 0x7d, 0x86,
	0xaa, 0x57, 0x55, 0xdf, 0xa0, 0x57, 0x7d, 0x80, 0xbe, 0x40, 0x2f, 0xfa, 0x2a, 0xbd, 0xab, 0xce,
	0xcc, 0x02, 0x0b, 0xbb, 0x4b, 0x55, 0x4b, 0xae, 0x72, 0xe5, 0xd9, 0x6f, 0x3e, 0xce, 0xf9, 0xbe,
	0x73, 0x66, 0xce, 0xae, 0x61, 0xff, 0x7b, 0x7f, 0xea, 0x49, 0x8f, 0x8a, 0xb9, 0x3b, 0xac, 0xab,
	0x25, 0xc9, 0x4f, 0xa4, 0x5a, 0x9c, 0xde, 0x17, 0xc3, 0x11, 0x9f, 0xb0, 0xba, 0xbc, 0xa9, 0xaf,
	0xb6, 0x2d, 0x02, 0x95, 0x57, 0x5c, 0x08, 0xf6, 0x0d, 0x7f, 0xc1, 0x24, 0x6b, 0x4d, 0x7c, 0x39,
	0xb7, 0x9e, 0x41, 0x29, 0x82, 0x91, 0x0f, 0xa0, 0x3c, 0xd1, 0x8f, 0x54, 0xce, 0x7d, 0x5e, 0x35,
	0xce, 0x8c, 0x8f, 0x4c, 0xbb, 0x14, 0x62, 0xfd, 0xb9, 0xcf, 0xad, 0x22, 0xe4, 0xdf, 0x78, 0x63,
	0xc7, 0x16, 0xbe, 0x35, 0x87, 0x72, 0x77, 0x26, 0x46, 0xcd, 0x9b, 0x31, 0x77, 0x65, 0xfb, 0x05,
	0x79, 0x08, 0x25, 0x36, 0x93, 0x23, 0xfa, 0x2d, 0x9f, 0xd3, 0xb1, 0xa3, 0x7e, 0x9c, 0xb5, 0x8b,
	0x08, 0x5d, 0xf0, 0x79, 0xdb, 0x21, 0x0f, 0x00, 0x04, 0x17, 0x62, 0xec, 0xb9, 0xb8, 0x9d, 0xd1,
	0xdb, 0x21, 0xd2, 0x76, 0xc8, 0x53, 0xd8, 0x77, 0xb9, 0xbc, 0x19, 0x5f, 0xd3, 0x08, 0x2b, 0xab,
	0x58, 0x7b, 0x7a, 0xa3, 0xb7, 0xe0, 0x5a, 0xbf, 0x19, 0x40, 0x74, 0xde, 0x81, 0xef, 0x30, 0xc9,
	0x45, 0x4f, 0x32, 0xc9, 0x49, 0x05, 0xb2, 0xbe, 0x14, 0xa1, 0x6c, 0x5c, 0x92, 0x1a, 0x14, 0x7d,
	0x29, 0xe8, 0xd0, 0x9b, 0xb9, 0x52, 0xa5, 0x34, 0xed, 0x82, 0x2f, 0x45, 0x13, 0x9f, 0x91, 0x1e,
	0x48, 0xa1, 0x72, 0x98, 0x36, 0x2e, 0x91, 0x1e, 0x2c, 0xe9, 0x39, 0x4d, 0x0f, 0x22, 0x74, 0xc1,
	0x83, 0xaa, 0xa9, 0xe9, 0x82, 0x07, 0x48, 0x17, 0x3c, 0xa0, 0x42, 0xb2, 0xa9, 0xac, 0xee, 0x68,
	0xba, 0xe0, 0x41, 0x0f, 0x9f, 0x09, 0x81, 0x1c, 0x4a, 0xab, 0xe6, 0x15, 0xae, 0xd6, 0x56, 0x0d,
	0x4e, 0x9a, 0x9e, 0xeb, 0xf2, 0xa1, 0xec, 0x7b, 0xa1, 0x9b, 0x1e, 0x9f, 0x7e, 0xc7, 0xa7, 0x36,
	0x0f, 0xac, 0x2b, 0x38, 0x59, 0xc3, 0x42, 0x26, 0xc7, 0x62, 0xeb, 0x54, 0x88, 0x2e, 0x4a, 0xab,
	0x52, 0x21, 0xd0, 0x76, 0xc8, 0x23, 0x28, 0x85, 0x9b, 0x2e, 0x9b, 0x70, 0xe5, 0xb3, 0x68, 0x83,
	0x86, 0x3a, 0x6c, 0xc2, 0xad, 0xdf, 0x0d, 0xd8, 0xc3, 0x5e, 0x85, 0xd5, 0x52, 0xcd, 0x6e, 0x40,
	0x71, 0xa8, 0x4a, 0xb8, 0x88, 0x58, 0x6a, 0x1c, 0xd6, 0xc3, 0x23, 0x54, 0x8f, 0x36, 0xd6, 0x2e,
	0x68, 0x5e, 0xdb, 0x21, 0x9f, 0x80, 0x29, 0xb0, 0xd2, 0x2a, 0x45, 0xa9, 0x51, 0x5b, 0xf2, 0xe3,
	0xcd, 0xb0, 0x35, 0x93, 0x3c, 0x81, 0xbd, 0x29, 0x7b, 0x4b, 0x1d, 0x26, 0x19, 0x1d, 0x71, 0xe6,
	0xf0, 0xa9, 0x2a, 0xf8, 0x3d, 0xfb, 0xde, 0x94, 0xbd, 0x45, 0x21, 0xe7, 0x0a, 0x24, 0x27, 0x50,
	0x58, 0xf0, 0x54, 0xe5, 0xcb, 0x76, 0x3e, 0x24, 0x58, 0x7f, 0x1b, 0x70, 0xdc, 0x9b, 0xbb, 0xc3,
	0xde, 0xc8, 0x9b, 0xca, 0xf0, 0xbc, 0xda, 0x3c, 0x98, 0x71, 0x21, 0x6f, 0xe5, 0xa2, 0x0e, 0x45,
	0x7f, 0x26, 0x46, 0xfa, 0x8c, 0xa3, 0x93, 0xdd, 0xc6, 0xfe, 0xf2, 0x37, 0x98, 0x08, 0x4f, 0xba,
	0x5d, 0x40, 0x0e, 0xae, 0xc8, 0x63, 0xd8, 0xc5, 0xb5, 0xf4, 0xe8, 0x4c, 0xe8, 0x06, 0xe8, 0x23,
	0x53, 0xd6, 0xe8, 0x40, 0xa8, 0x26, 0x1c, 0x43, 0xde, 0xe7, 0x7a, 0x5b, 0x9f, 0x9c, 0x1d, 0x7c,
	0x6c, 0x3b, 0xe4, 0xf3, 0x30, 0x9d, 0xb2, 0x66, 0x2a, 0x89, 0x0f, 0x97, 0xe9, 0xfa, 0x2f, 0xe9,
	0x4c, 0x55, 0x6d, 0xcd, 0x9c, 0xca, 0xad, 0xbc, 0xff, 0x61, 0xc0, 0xc9, 0x20, 0x4e, 0x08, 0xdd,
	0xaf, 0x39, 0x31, 0x6e, 0xe3, 0x24, 0xb3, 0xdd, 0x49, 0x36, 0xdd, 0x49, 0xee, 0x3f, 0x3a, 0xf9,
	0x31, 0x03, 0xb5, 0x65, 0x17, 0x9b, 0x23, 0xf6, 0xee, 0x76, 0xf2, 0x0c, 0xca, 0xca, 0xff, 0x70,
	0xc4, 0xe4, 0xaa, 0x9d, 0x80, 0x18, 0xea, 0x6e, 0x3b, 0xe4, 0x8b, 0x78, 0x4b, 0xad, 0x94, 0x42,
	0x44, 0x9d, 0xae, 0x8a, 0xf1, 0x97, 0x01, 0x0f, 0x06, 0xc9, 0xa4, 0x3b, 0x6d, 0xed, 0xa6, 0xb5,
	0xec, 0x76, 0x6b, 0xb9, 0x5b, 0x58, 0xfb, 0x21, 0x03, 0x55, 0xd4, 0xa7, 0xed, 0xbd, 0xb3, 0x4d,
	0xc6, 0x37, 0x03, 0x56, 0x42, 0x45, 0x0d, 0x47, 0x3d, 0x02, 0x2a, 0x44, 0xe4, 0x06, 0x98, 0x6b,
	0x37, 0xe0, 0xe3, 0x68, 0x75, 0x76, 0x94, 0xfe, 0xbd, 0xa5, 0x16, 0xed, 0x38, 0x52, 0x8a, 0x3f,
	0x0d, 0xa8, 0xae, 0xc6, 0xee, 0xff, 0xd2, 0xe0, 0x35, 0x5b, 0xd9, 0x74, 0x5b, 0xb9, 0x74, 0x5b,
	0xe6, 0xbf, 0xd9, 0xfa, 0xc9, 0x00, 0x12, 0x79, 0x9b, 0xdc, 0xad, 0xa1, 0x67, 0x71, 0x69, 0x07,
	0xf1, 0xf3, 0x28, 0x22, 0xf2, 0x7e, 0x31, 0xe0, 0x68, 0x25, 0x4f, 0x9d, 0xd4, 0xbb, 0x95, 0xf8,
	0x59, 0x5c, 0xe2, 0x71, 0xca, 0x95, 0x59, 0xc9, 0x7c, 0xfa, 0xab, 0x01, 0x85, 0x45, 0x4a, 0x72,
	0x08, 0xfb, 0xbd, 0xab, 0x4e, 0x93, 0xf6, 0xaf, 0xba, 0x2d, 0x3a, 0xe8, 0x5c, 0x74, 0x5e, 0x5f,
	0x76, 0x2a, 0xef, 0x11, 0x02, 0xbb, 0x11, 0xb8, 0xd7, 0xb2, 0x2b, 0x06, 0xa9, 0xc2, 0xfd, 0x75,
	0x8c, 0x76, 0x5e, 0xf7, 0x5f, 0xb5, 0x2a, 0x19, 0x72, 0x04, 0x64, 0xb5, 0xf3, 0xd5, 0xa0, 0x7f,
	0x4e, 0x2f, 0x5a, 0x57, 0x95, 0x2c, 0xa9, 0xc1, 0x71, 0x1c, 0xd7, 0xe1, 0x72, 0xe4, 0x0c, 0xde,
	0x4f, 0xd9, 0xd4, 0x61, 0xcd, 0xc6, 0xcf, 0x26, 0xe4, 0xed, 0x6e, 0x13, 0xb5, 0x92, 0x4b, 0x38,
	0x5a, 0xdd, 0xed, 0xe8, 0xa0, 0x27, 0x67, 0x6b, 0x75, 0x4c, 0x78, 0x59, 0x9d, 0x6e, 0xfb, 0x5a,
	0x20, 0x9d, 0x58, 0xcf, 0x16, 0x81, 0xad, 0x8d, 0x83, 0x98, 0x14, 0xba, 0xb2, 0xe4, 0x84, 0x1f,
	0xa7, 0x84, 0xc2, 0xe9, 0x86, 0xd0, 0xc8, 0xb4, 0x22, 0x8f, 0xe3, 0x62, 0xe3, 0x23, 0x78, 0xbb,
	0xe0, 0x37, 0x70, 0xba, 0x21, 0x38, 0x9a, 0xe0, 0x49, 0x92, 0xe8, 0x84, 0x14, 0x71, 0xe1, 0x97,
	0x70, 0x18, 0x9b, 0x9e, 0xfa, 0xe3, 0x7c, 0x4d, 0x73, 0xd2, 0x48, 0xd9, 0x2e, 0xf8, 0x25, 0x1c,
	0xc6, 0x66, 0xd1, 0x46, 0xe0, 0xb4, 0x59, 0x95, 0x20, 0xf3, 0xcb, 0xf8, 0x07, 0x65, 0x2d, 0x21,
	0x8e, 0x48, 0x8f, 0x70, 0x0e, 0x07, 0x1b, 0x05, 0x54, 0x51, 0x1e, 0x25, 0x44, 0x89, 0xde, 0xe1,
	0x78, 0xa4, 0xe7, 0x1f, 0xc2, 0xc1, 0xd0, 0x9b, 0xd4, 0x5d, 0x7e, 0x3d, 0xbb, 0x61, 0xe3, 0x49,
	0x5d, 0xff, 0x6b, 0xf4, 0x1c, 0xbe, 0xee, 0xe2, 0x5f, 0xac, 0xd6, 0x79, 0xa6, 0x6b, 0x5c, 0xef,
	0x28, 0xf8, 0xd3, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x58, 0xed, 0x98, 0xa9, 0x3b, 0x0d, 0x00,
	0x00,
}