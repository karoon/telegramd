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
	SyncType_SYNC_TYPE_UNKNOWN    SyncType = 0
	SyncType_SYNC_TYPE_USER       SyncType = 1
	SyncType_SYNC_TYPE_USER_NOTME SyncType = 2
	SyncType_SYNC_TYPE_USER_ME    SyncType = 3
	SyncType_SYNC_TYPE_RPC_RESULT SyncType = 4
)

var SyncType_name = map[int32]string{
	0: "SYNC_TYPE_UNKNOWN",
	1: "SYNC_TYPE_USER",
	2: "SYNC_TYPE_USER_NOTME",
	3: "SYNC_TYPE_USER_ME",
	4: "SYNC_TYPE_RPC_RESULT",
}
var SyncType_value = map[string]int32{
	"SYNC_TYPE_UNKNOWN":    0,
	"SYNC_TYPE_USER":       1,
	"SYNC_TYPE_USER_NOTME": 2,
	"SYNC_TYPE_USER_ME":    3,
	"SYNC_TYPE_RPC_RESULT": 4,
}

func (x SyncType) String() string {
	return proto.EnumName(SyncType_name, int32(x))
}
func (SyncType) EnumDescriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

type VoidRsp struct {
}

func (m *VoidRsp) Reset()                    { *m = VoidRsp{} }
func (m *VoidRsp) String() string            { return proto.CompactTextString(m) }
func (*VoidRsp) ProtoMessage()               {}
func (*VoidRsp) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

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
func (*ClientUpdatesState) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

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
func (*ConnectToSessionServerReq) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

type SessionServerConnectedRsp struct {
	ServerId   int32  `protobuf:"varint,1,opt,name=server_id,json=serverId" json:"server_id,omitempty"`
	ServerName string `protobuf:"bytes,2,opt,name=server_name,json=serverName" json:"server_name,omitempty"`
}

func (m *SessionServerConnectedRsp) Reset()                    { *m = SessionServerConnectedRsp{} }
func (m *SessionServerConnectedRsp) String() string            { return proto.CompactTextString(m) }
func (*SessionServerConnectedRsp) ProtoMessage()               {}
func (*SessionServerConnectedRsp) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{3} }

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
	AuthKeyId   int64               `protobuf:"varint,1,opt,name=auth_key_id,json=authKeyId" json:"auth_key_id,omitempty"`
	SessionId   int64               `protobuf:"varint,2,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	State       *ClientUpdatesState `protobuf:"bytes,3,opt,name=state" json:"state,omitempty"`
	UpdatesData []byte              `protobuf:"bytes,5,opt,name=updates_data,json=updatesData,proto3" json:"updates_data,omitempty"`
}

func (m *PushUpdatesData) Reset()                    { *m = PushUpdatesData{} }
func (m *PushUpdatesData) String() string            { return proto.CompactTextString(m) }
func (*PushUpdatesData) ProtoMessage()               {}
func (*PushUpdatesData) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{4} }

func (m *PushUpdatesData) GetAuthKeyId() int64 {
	if m != nil {
		return m.AuthKeyId
	}
	return 0
}

func (m *PushUpdatesData) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *PushUpdatesData) GetState() *ClientUpdatesState {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *PushUpdatesData) GetUpdatesData() []byte {
	if m != nil {
		return m.UpdatesData
	}
	return nil
}

// Updates
// messages_affectedHistory
// messages_affectedMessages
type RpcResultData struct {
	// int32 rpc_result_type = 1;
	Updates          *Updates                    `protobuf:"bytes,2,opt,name=updates" json:"updates,omitempty"`
	AffectedHistory  *TLMessagesAffectedHistory  `protobuf:"bytes,3,opt,name=affected_history,json=affectedHistory" json:"affected_history,omitempty"`
	AffectedMessages *TLMessagesAffectedMessages `protobuf:"bytes,4,opt,name=affected_messages,json=affectedMessages" json:"affected_messages,omitempty"`
}

func (m *RpcResultData) Reset()                    { *m = RpcResultData{} }
func (m *RpcResultData) String() string            { return proto.CompactTextString(m) }
func (*RpcResultData) ProtoMessage()               {}
func (*RpcResultData) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{5} }

func (m *RpcResultData) GetUpdates() *Updates {
	if m != nil {
		return m.Updates
	}
	return nil
}

func (m *RpcResultData) GetAffectedHistory() *TLMessagesAffectedHistory {
	if m != nil {
		return m.AffectedHistory
	}
	return nil
}

func (m *RpcResultData) GetAffectedMessages() *TLMessagesAffectedMessages {
	if m != nil {
		return m.AffectedMessages
	}
	return nil
}

// /////////////////////////////////////////////////////////////////////
// RPC
type UpdatesRequest struct {
	PushType    SyncType       `protobuf:"varint,1,opt,name=push_type,json=pushType,enum=mtproto.SyncType" json:"push_type,omitempty"`
	Layer       int32          `protobuf:"varint,2,opt,name=layer" json:"layer,omitempty"`
	ServerId    int32          `protobuf:"varint,3,opt,name=server_id,json=serverId" json:"server_id,omitempty"`
	AuthKeyId   int64          `protobuf:"varint,4,opt,name=auth_key_id,json=authKeyId" json:"auth_key_id,omitempty"`
	SessionId   int64          `protobuf:"varint,5,opt,name=session_id,json=sessionId" json:"session_id,omitempty"`
	PushUserId  int32          `protobuf:"varint,6,opt,name=push_user_id,json=pushUserId" json:"push_user_id,omitempty"`
	ClientMsgId int64          `protobuf:"varint,7,opt,name=client_msg_id,json=clientMsgId" json:"client_msg_id,omitempty"`
	Updates     *Updates       `protobuf:"bytes,8,opt,name=updates" json:"updates,omitempty"`
	RpcResult   *RpcResultData `protobuf:"bytes,9,opt,name=rpc_result,json=rpcResult" json:"rpc_result,omitempty"`
}

func (m *UpdatesRequest) Reset()                    { *m = UpdatesRequest{} }
func (m *UpdatesRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdatesRequest) ProtoMessage()               {}
func (*UpdatesRequest) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{6} }

func (m *UpdatesRequest) GetPushType() SyncType {
	if m != nil {
		return m.PushType
	}
	return SyncType_SYNC_TYPE_UNKNOWN
}

func (m *UpdatesRequest) GetLayer() int32 {
	if m != nil {
		return m.Layer
	}
	return 0
}

func (m *UpdatesRequest) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *UpdatesRequest) GetAuthKeyId() int64 {
	if m != nil {
		return m.AuthKeyId
	}
	return 0
}

func (m *UpdatesRequest) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *UpdatesRequest) GetPushUserId() int32 {
	if m != nil {
		return m.PushUserId
	}
	return 0
}

func (m *UpdatesRequest) GetClientMsgId() int64 {
	if m != nil {
		return m.ClientMsgId
	}
	return 0
}

func (m *UpdatesRequest) GetUpdates() *Updates {
	if m != nil {
		return m.Updates
	}
	return nil
}

func (m *UpdatesRequest) GetRpcResult() *RpcResultData {
	if m != nil {
		return m.RpcResult
	}
	return nil
}

func init() {
	proto.RegisterType((*VoidRsp)(nil), "mtproto.VoidRsp")
	proto.RegisterType((*ClientUpdatesState)(nil), "mtproto.ClientUpdatesState")
	proto.RegisterType((*ConnectToSessionServerReq)(nil), "mtproto.ConnectToSessionServerReq")
	proto.RegisterType((*SessionServerConnectedRsp)(nil), "mtproto.SessionServerConnectedRsp")
	proto.RegisterType((*PushUpdatesData)(nil), "mtproto.PushUpdatesData")
	proto.RegisterType((*RpcResultData)(nil), "mtproto.RpcResultData")
	proto.RegisterType((*UpdatesRequest)(nil), "mtproto.UpdatesRequest")
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
	SyncUpdatesData(ctx context.Context, in *UpdatesRequest, opts ...grpc.CallOption) (*ClientUpdatesState, error)
	PushUpdatesData(ctx context.Context, in *UpdatesRequest, opts ...grpc.CallOption) (*VoidRsp, error)
}

type rPCSyncClient struct {
	cc *grpc.ClientConn
}

func NewRPCSyncClient(cc *grpc.ClientConn) RPCSyncClient {
	return &rPCSyncClient{cc}
}

func (c *rPCSyncClient) SyncUpdatesData(ctx context.Context, in *UpdatesRequest, opts ...grpc.CallOption) (*ClientUpdatesState, error) {
	out := new(ClientUpdatesState)
	err := grpc.Invoke(ctx, "/mtproto.RPCSync/SyncUpdatesData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCSyncClient) PushUpdatesData(ctx context.Context, in *UpdatesRequest, opts ...grpc.CallOption) (*VoidRsp, error) {
	out := new(VoidRsp)
	err := grpc.Invoke(ctx, "/mtproto.RPCSync/PushUpdatesData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RPCSync service

type RPCSyncServer interface {
	SyncUpdatesData(context.Context, *UpdatesRequest) (*ClientUpdatesState, error)
	PushUpdatesData(context.Context, *UpdatesRequest) (*VoidRsp, error)
}

func RegisterRPCSyncServer(s *grpc.Server, srv RPCSyncServer) {
	s.RegisterService(&_RPCSync_serviceDesc, srv)
}

func _RPCSync_SyncUpdatesData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCSyncServer).SyncUpdatesData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mtproto.RPCSync/SyncUpdatesData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCSyncServer).SyncUpdatesData(ctx, req.(*UpdatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPCSync_PushUpdatesData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatesRequest)
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
		return srv.(RPCSyncServer).PushUpdatesData(ctx, req.(*UpdatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPCSync_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mtproto.RPCSync",
	HandlerType: (*RPCSyncServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncUpdatesData",
			Handler:    _RPCSync_SyncUpdatesData_Handler,
		},
		{
			MethodName: "PushUpdatesData",
			Handler:    _RPCSync_PushUpdatesData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zproto_sync.proto",
}

func init() { proto.RegisterFile("zproto_sync.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 697 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0x49, 0x3f, 0xd6, 0xf6, 0x74, 0x1f, 0xa9, 0x19, 0x90, 0x6d, 0x02, 0x46, 0xc4, 0x24,
	0xb4, 0x8b, 0x48, 0x14, 0x71, 0xc7, 0xd5, 0x4a, 0xc5, 0xaa, 0xad, 0x5d, 0xe5, 0xb4, 0xa0, 0x71,
	0x63, 0x79, 0x89, 0xb7, 0x56, 0x34, 0x9f, 0x76, 0x90, 0xc2, 0x15, 0x8f, 0xc0, 0x53, 0xf0, 0x06,
	0x3c, 0x0d, 0x37, 0x3c, 0x0a, 0xb2, 0x9d, 0x76, 0x6d, 0x27, 0xc6, 0x55, 0xec, 0xbf, 0x7f, 0x3e,
	0xfe, 0xfb, 0x9c, 0xe3, 0x40, 0xeb, 0x5b, 0x9c, 0x46, 0x22, 0x22, 0x3c, 0x0f, 0x3d, 0x47, 0x0d,
	0x51, 0x2d, 0x10, 0x6a, 0xb0, 0xbf, 0xcb, 0xbd, 0x09, 0x0b, 0xa8, 0x23, 0x66, 0xce, 0xed, 0xb2,
	0xdd, 0x80, 0xda, 0xc7, 0x68, 0xea, 0x63, 0x1e, 0xdb, 0xbf, 0x0c, 0x40, 0x9d, 0xd9, 0x94, 0x85,
	0x62, 0x1c, 0xfb, 0x54, 0x30, 0xee, 0x0a, 0x2a, 0x18, 0x32, 0xa1, 0x1c, 0x0b, 0x6e, 0x19, 0x87,
	0xc6, 0xab, 0x2a, 0x96, 0x43, 0x74, 0x00, 0x8d, 0x58, 0x70, 0xe2, 0x45, 0x59, 0x28, 0xac, 0x92,
	0xd2, 0xeb, 0xb1, 0xe0, 0x1d, 0x39, 0x97, 0x78, 0x22, 0xb8, 0x55, 0xd6, 0x78, 0xa2, 0xf1, 0x64,
	0x81, 0x57, 0x34, 0x9e, 0x2c, 0xe1, 0x9c, 0x25, 0x56, 0x55, 0xe3, 0x9c, 0x25, 0x12, 0xe7, 0x2c,
	0x21, 0x5c, 0xd0, 0x54, 0x58, 0x1b, 0x1a, 0xe7, 0x2c, 0x71, 0xe5, 0x1c, 0x21, 0xa8, 0x48, 0x6b,
	0x56, 0x4d, 0xe9, 0x6a, 0x6c, 0x1f, 0xc0, 0x5e, 0x27, 0x0a, 0x43, 0xe6, 0x89, 0x51, 0xe4, 0x32,
	0xce, 0xa7, 0x51, 0xe8, 0xb2, 0xf4, 0x2b, 0x4b, 0x31, 0x4b, 0xec, 0x4b, 0xd8, 0x5b, 0xd1, 0x0a,
	0x92, 0xc9, 0x1b, 0xeb, 0xa3, 0xa4, 0x4a, 0xa6, 0x7e, 0x71, 0xc1, 0xba, 0x16, 0x7a, 0x3e, 0x7a,
	0x0e, 0xcd, 0x62, 0x31, 0xa4, 0x01, 0x53, 0xf7, 0x6c, 0x60, 0xd0, 0xd2, 0x80, 0x06, 0xcc, 0xfe,
	0x69, 0xc0, 0xce, 0x30, 0xe3, 0x93, 0x22, 0x5b, 0xef, 0xa9, 0xa0, 0xe8, 0x19, 0x34, 0x69, 0x26,
	0x26, 0xe4, 0x0b, 0xcb, 0xe7, 0x31, 0xcb, 0xb8, 0x21, 0xa5, 0x33, 0x96, 0xf7, 0x7c, 0xf4, 0x14,
	0x80, 0x6b, 0x3b, 0x72, 0xb9, 0xa4, 0x97, 0x0b, 0xa5, 0xe7, 0xa3, 0xd7, 0x50, 0xe5, 0x32, 0xe9,
	0x2a, 0x7d, 0xcd, 0xf6, 0x81, 0x53, 0x14, 0xcf, 0xb9, 0x5b, 0x17, 0xac, 0x49, 0xf4, 0x02, 0x36,
	0x33, 0x2d, 0x13, 0x9f, 0x0a, 0xaa, 0x32, 0xb9, 0x89, 0x9b, 0xd9, 0xad, 0x29, 0xfb, 0x8f, 0x01,
	0x5b, 0x38, 0xf6, 0x30, 0xe3, 0xd9, 0x4c, 0x28, 0x9b, 0xc7, 0x50, 0x2b, 0x00, 0xe5, 0xa1, 0xd9,
	0x36, 0x17, 0x27, 0x15, 0x67, 0xe0, 0x39, 0x80, 0x2e, 0xc0, 0xa4, 0xd7, 0xd7, 0x2a, 0x67, 0x64,
	0x32, 0xe5, 0x22, 0x4a, 0xf3, 0xc2, 0xde, 0xcb, 0xc5, 0xa6, 0xd1, 0x39, 0x09, 0x18, 0xe7, 0xf4,
	0x86, 0x71, 0x32, 0x87, 0x4f, 0x35, 0x8b, 0x77, 0xd6, 0x04, 0x84, 0xa1, 0xb5, 0x08, 0x38, 0xdf,
	0xa5, 0xfa, 0xa2, 0xd9, 0x3e, 0xba, 0x37, 0x62, 0xbf, 0x10, 0xb0, 0xb9, 0xae, 0xd8, 0xbf, 0x4b,
	0xb0, 0x3d, 0x77, 0xce, 0x92, 0x8c, 0x71, 0x81, 0x1c, 0x68, 0xc4, 0x19, 0x9f, 0x10, 0x91, 0xc7,
	0x4c, 0x15, 0x62, 0xbb, 0xdd, 0x5a, 0x84, 0x77, 0xf3, 0xd0, 0x1b, 0xe5, 0x31, 0xc3, 0x75, 0xc9,
	0xc8, 0x11, 0xda, 0x85, 0xea, 0x8c, 0xe6, 0x2c, 0x2d, 0x3a, 0x5a, 0x4f, 0x56, 0x5b, 0xa4, 0xbc,
	0xd6, 0x22, 0x6b, 0xd5, 0xae, 0xdc, 0x5f, 0xed, 0xea, 0x7a, 0xb5, 0x0f, 0x61, 0x53, 0x39, 0xcc,
	0xb8, 0x0e, 0xaf, 0x9b, 0x1d, 0xa4, 0x36, 0xe6, 0xea, 0x00, 0x1b, 0xb6, 0x3c, 0x55, 0x79, 0x12,
	0xf0, 0x1b, 0x89, 0xd4, 0x54, 0x8c, 0xa6, 0x16, 0xfb, 0xfc, 0xa6, 0xe7, 0x2f, 0xd7, 0xb2, 0xfe,
	0xbf, 0x5a, 0xbe, 0x05, 0x48, 0x63, 0x8f, 0xa4, 0xaa, 0x13, 0xac, 0x86, 0xc2, 0x1f, 0x2f, 0xf0,
	0x95, 0x1e, 0xc1, 0x8d, 0x74, 0x3e, 0x3d, 0xfe, 0x6e, 0x40, 0x7d, 0x9e, 0x31, 0xf4, 0x08, 0x5a,
	0xee, 0xe5, 0xa0, 0x43, 0x46, 0x97, 0xc3, 0x2e, 0x19, 0x0f, 0xce, 0x06, 0x17, 0x9f, 0x06, 0xe6,
	0x03, 0x84, 0x60, 0x7b, 0x49, 0x76, 0xbb, 0xd8, 0x34, 0x90, 0x05, 0xbb, 0xab, 0x1a, 0x19, 0x5c,
	0x8c, 0xfa, 0x5d, 0xb3, 0xb4, 0x16, 0x44, 0xae, 0xf4, 0xbb, 0x66, 0x79, 0x75, 0x03, 0x1e, 0x76,
	0x08, 0xee, 0xba, 0xe3, 0xf3, 0x91, 0x59, 0x69, 0xff, 0x30, 0xa0, 0x86, 0x87, 0x1d, 0xe9, 0x02,
	0x7d, 0x80, 0x1d, 0xf9, 0x5d, 0x7e, 0x77, 0x4f, 0xee, 0xdc, 0x59, 0x77, 0xc1, 0xfe, 0x7d, 0x4f,
	0x08, 0xbd, 0xbb, 0xfb, 0x80, 0xff, 0x19, 0xe8, 0x36, 0xab, 0xc5, 0xff, 0xf2, 0xe4, 0x08, 0x1e,
	0x7a, 0x51, 0xe0, 0x84, 0xec, 0x2a, 0x9b, 0xd1, 0x69, 0xe0, 0xe8, 0x7f, 0xef, 0x09, 0x7c, 0x1e,
	0xca, 0xaf, 0x74, 0x78, 0x5a, 0x1a, 0x1a, 0x57, 0x1b, 0x4a, 0x7e, 0xf3, 0x37, 0x00, 0x00, 0xff,
	0xff, 0x16, 0x1d, 0xe0, 0x8a, 0x9c, 0x05, 0x00, 0x00,
}
