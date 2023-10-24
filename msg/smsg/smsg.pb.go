// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg/smsg/smsg.proto

/*
Package smsg is a generated protocol buffer package.

It is generated from these files:
	msg/smsg/smsg.proto

It has these top-level messages:
	Server2AllSession
	GaCeReqLogin
	GaCeRespLogin
	CeGaBindGameServer
	CeGamReqJoinGame
	CeGamRespJoinGame
	GamCeNoticeGameStart
	GamCeNoticeGameEnd
	CeGameUserDisconnect
	CeGameUserReconnect
	GaCeUserDisconnect
	CeGaCloseSession
	AdReqMetrics
	AdRespMetrics
*/
package smsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GaCeRespLogin_Error int32

const (
	GaCeRespLogin_Invalid GaCeRespLogin_Error = 0
)

var GaCeRespLogin_Error_name = map[int32]string{
	0: "Invalid",
}
var GaCeRespLogin_Error_value = map[string]int32{
	"Invalid": 0,
}

func (x GaCeRespLogin_Error) String() string {
	return proto.EnumName(GaCeRespLogin_Error_name, int32(x))
}
func (GaCeRespLogin_Error) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type CeGamRespJoinGame_Error int32

const (
	CeGamRespJoinGame_Invalid      CeGamRespJoinGame_Error = 0
	CeGamRespJoinGame_GameNotExist CeGamRespJoinGame_Error = 1
)

var CeGamRespJoinGame_Error_name = map[int32]string{
	0: "Invalid",
	1: "GameNotExist",
}
var CeGamRespJoinGame_Error_value = map[string]int32{
	"Invalid":      0,
	"GameNotExist": 1,
}

func (x CeGamRespJoinGame_Error) String() string {
	return proto.EnumName(CeGamRespJoinGame_Error_name, int32(x))
}
func (CeGamRespJoinGame_Error) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type AdRespMetrics_MetricsType int32

const (
	AdRespMetrics_Invalid     AdRespMetrics_MetricsType = 0
	AdRespMetrics_OnlineCount AdRespMetrics_MetricsType = 1
)

var AdRespMetrics_MetricsType_name = map[int32]string{
	0: "Invalid",
	1: "OnlineCount",
}
var AdRespMetrics_MetricsType_value = map[string]int32{
	"Invalid":     0,
	"OnlineCount": 1,
}

func (x AdRespMetrics_MetricsType) String() string {
	return proto.EnumName(AdRespMetrics_MetricsType_name, int32(x))
}
func (AdRespMetrics_MetricsType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{13, 0}
}

type Server2AllSession struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Server2AllSession) Reset()                    { *m = Server2AllSession{} }
func (m *Server2AllSession) String() string            { return proto.CompactTextString(m) }
func (*Server2AllSession) ProtoMessage()               {}
func (*Server2AllSession) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Server2AllSession) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type GaCeReqLogin struct {
	Sesid int32  `protobuf:"varint,1,opt,name=sesid" json:"sesid,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *GaCeReqLogin) Reset()                    { *m = GaCeReqLogin{} }
func (m *GaCeReqLogin) String() string            { return proto.CompactTextString(m) }
func (*GaCeReqLogin) ProtoMessage()               {}
func (*GaCeReqLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GaCeReqLogin) GetSesid() int32 {
	if m != nil {
		return m.Sesid
	}
	return 0
}

func (m *GaCeReqLogin) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GaCeRespLogin struct {
	UserID uint64              `protobuf:"varint,1,opt,name=userID" json:"userID,omitempty"`
	Err    GaCeRespLogin_Error `protobuf:"varint,2,opt,name=err,enum=smsg.GaCeRespLogin_Error" json:"err,omitempty"`
	Token  string              `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	InGame bool                `protobuf:"varint,4,opt,name=inGame" json:"inGame,omitempty"`
}

func (m *GaCeRespLogin) Reset()                    { *m = GaCeRespLogin{} }
func (m *GaCeRespLogin) String() string            { return proto.CompactTextString(m) }
func (*GaCeRespLogin) ProtoMessage()               {}
func (*GaCeRespLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GaCeRespLogin) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *GaCeRespLogin) GetErr() GaCeRespLogin_Error {
	if m != nil {
		return m.Err
	}
	return GaCeRespLogin_Invalid
}

func (m *GaCeRespLogin) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *GaCeRespLogin) GetInGame() bool {
	if m != nil {
		return m.InGame
	}
	return false
}

type CeGaBindGameServer struct {
	Sesid        int32 `protobuf:"varint,1,opt,name=sesid" json:"sesid,omitempty"`
	Gameserverid int32 `protobuf:"varint,2,opt,name=gameserverid" json:"gameserverid,omitempty"`
}

func (m *CeGaBindGameServer) Reset()                    { *m = CeGaBindGameServer{} }
func (m *CeGaBindGameServer) String() string            { return proto.CompactTextString(m) }
func (*CeGaBindGameServer) ProtoMessage()               {}
func (*CeGaBindGameServer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CeGaBindGameServer) GetSesid() int32 {
	if m != nil {
		return m.Sesid
	}
	return 0
}

func (m *CeGaBindGameServer) GetGameserverid() int32 {
	if m != nil {
		return m.Gameserverid
	}
	return 0
}

type CeGamReqJoinGame struct {
	Userid       uint64 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Sesid        int32  `protobuf:"varint,2,opt,name=sesid" json:"sesid,omitempty"`
	Nickname     string `protobuf:"bytes,3,opt,name=nickname" json:"nickname,omitempty"`
	GateServerid int32  `protobuf:"varint,4,opt,name=gateServerid" json:"gateServerid,omitempty"`
}

func (m *CeGamReqJoinGame) Reset()                    { *m = CeGamReqJoinGame{} }
func (m *CeGamReqJoinGame) String() string            { return proto.CompactTextString(m) }
func (*CeGamReqJoinGame) ProtoMessage()               {}
func (*CeGamReqJoinGame) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CeGamReqJoinGame) GetUserid() uint64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *CeGamReqJoinGame) GetSesid() int32 {
	if m != nil {
		return m.Sesid
	}
	return 0
}

func (m *CeGamReqJoinGame) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *CeGamReqJoinGame) GetGateServerid() int32 {
	if m != nil {
		return m.GateServerid
	}
	return 0
}

type CeGamRespJoinGame struct {
	Err    CeGamRespJoinGame_Error `protobuf:"varint,1,opt,name=err,enum=smsg.CeGamRespJoinGame_Error" json:"err,omitempty"`
	Gameid int64                   `protobuf:"varint,2,opt,name=gameid" json:"gameid,omitempty"`
}

func (m *CeGamRespJoinGame) Reset()                    { *m = CeGamRespJoinGame{} }
func (m *CeGamRespJoinGame) String() string            { return proto.CompactTextString(m) }
func (*CeGamRespJoinGame) ProtoMessage()               {}
func (*CeGamRespJoinGame) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CeGamRespJoinGame) GetErr() CeGamRespJoinGame_Error {
	if m != nil {
		return m.Err
	}
	return CeGamRespJoinGame_Invalid
}

func (m *CeGamRespJoinGame) GetGameid() int64 {
	if m != nil {
		return m.Gameid
	}
	return 0
}

type GamCeNoticeGameStart struct {
	Gameid int64 `protobuf:"varint,1,opt,name=gameid" json:"gameid,omitempty"`
}

func (m *GamCeNoticeGameStart) Reset()                    { *m = GamCeNoticeGameStart{} }
func (m *GamCeNoticeGameStart) String() string            { return proto.CompactTextString(m) }
func (*GamCeNoticeGameStart) ProtoMessage()               {}
func (*GamCeNoticeGameStart) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GamCeNoticeGameStart) GetGameid() int64 {
	if m != nil {
		return m.Gameid
	}
	return 0
}

type GamCeNoticeGameEnd struct {
	Gameid int64 `protobuf:"varint,1,opt,name=gameid" json:"gameid,omitempty"`
}

func (m *GamCeNoticeGameEnd) Reset()                    { *m = GamCeNoticeGameEnd{} }
func (m *GamCeNoticeGameEnd) String() string            { return proto.CompactTextString(m) }
func (*GamCeNoticeGameEnd) ProtoMessage()               {}
func (*GamCeNoticeGameEnd) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GamCeNoticeGameEnd) GetGameid() int64 {
	if m != nil {
		return m.Gameid
	}
	return 0
}

type CeGameUserDisconnect struct {
	Userid uint64 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
}

func (m *CeGameUserDisconnect) Reset()                    { *m = CeGameUserDisconnect{} }
func (m *CeGameUserDisconnect) String() string            { return proto.CompactTextString(m) }
func (*CeGameUserDisconnect) ProtoMessage()               {}
func (*CeGameUserDisconnect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CeGameUserDisconnect) GetUserid() uint64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

type CeGameUserReconnect struct {
	Userid    uint64 `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	GateID    int32  `protobuf:"varint,2,opt,name=gateID" json:"gateID,omitempty"`
	SessionID int32  `protobuf:"varint,3,opt,name=sessionID" json:"sessionID,omitempty"`
}

func (m *CeGameUserReconnect) Reset()                    { *m = CeGameUserReconnect{} }
func (m *CeGameUserReconnect) String() string            { return proto.CompactTextString(m) }
func (*CeGameUserReconnect) ProtoMessage()               {}
func (*CeGameUserReconnect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CeGameUserReconnect) GetUserid() uint64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *CeGameUserReconnect) GetGateID() int32 {
	if m != nil {
		return m.GateID
	}
	return 0
}

func (m *CeGameUserReconnect) GetSessionID() int32 {
	if m != nil {
		return m.SessionID
	}
	return 0
}

type GaCeUserDisconnect struct {
	SessionID int32 `protobuf:"varint,1,opt,name=sessionID" json:"sessionID,omitempty"`
}

func (m *GaCeUserDisconnect) Reset()                    { *m = GaCeUserDisconnect{} }
func (m *GaCeUserDisconnect) String() string            { return proto.CompactTextString(m) }
func (*GaCeUserDisconnect) ProtoMessage()               {}
func (*GaCeUserDisconnect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *GaCeUserDisconnect) GetSessionID() int32 {
	if m != nil {
		return m.SessionID
	}
	return 0
}

type CeGaCloseSession struct {
	SessionID int32 `protobuf:"varint,1,opt,name=sessionID" json:"sessionID,omitempty"`
}

func (m *CeGaCloseSession) Reset()                    { *m = CeGaCloseSession{} }
func (m *CeGaCloseSession) String() string            { return proto.CompactTextString(m) }
func (*CeGaCloseSession) ProtoMessage()               {}
func (*CeGaCloseSession) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *CeGaCloseSession) GetSessionID() int32 {
	if m != nil {
		return m.SessionID
	}
	return 0
}

type AdReqMetrics struct {
	ReqTime int64 `protobuf:"varint,3,opt,name=reqTime" json:"reqTime,omitempty"`
}

func (m *AdReqMetrics) Reset()                    { *m = AdReqMetrics{} }
func (m *AdReqMetrics) String() string            { return proto.CompactTextString(m) }
func (*AdReqMetrics) ProtoMessage()               {}
func (*AdReqMetrics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *AdReqMetrics) GetReqTime() int64 {
	if m != nil {
		return m.ReqTime
	}
	return 0
}

type AdRespMetrics struct {
	Metrics []*AdRespMetrics_Metrics `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty"`
	ReqTime int64                    `protobuf:"varint,2,opt,name=reqTime" json:"reqTime,omitempty"`
}

func (m *AdRespMetrics) Reset()                    { *m = AdRespMetrics{} }
func (m *AdRespMetrics) String() string            { return proto.CompactTextString(m) }
func (*AdRespMetrics) ProtoMessage()               {}
func (*AdRespMetrics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *AdRespMetrics) GetMetrics() []*AdRespMetrics_Metrics {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *AdRespMetrics) GetReqTime() int64 {
	if m != nil {
		return m.ReqTime
	}
	return 0
}

type AdRespMetrics_Metrics struct {
	Key   AdRespMetrics_MetricsType `protobuf:"varint,1,opt,name=key,enum=smsg.AdRespMetrics_MetricsType" json:"key,omitempty"`
	Value int32                     `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
}

func (m *AdRespMetrics_Metrics) Reset()                    { *m = AdRespMetrics_Metrics{} }
func (m *AdRespMetrics_Metrics) String() string            { return proto.CompactTextString(m) }
func (*AdRespMetrics_Metrics) ProtoMessage()               {}
func (*AdRespMetrics_Metrics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13, 0} }

func (m *AdRespMetrics_Metrics) GetKey() AdRespMetrics_MetricsType {
	if m != nil {
		return m.Key
	}
	return AdRespMetrics_Invalid
}

func (m *AdRespMetrics_Metrics) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*Server2AllSession)(nil), "smsg.Server2AllSession")
	proto.RegisterType((*GaCeReqLogin)(nil), "smsg.GaCeReqLogin")
	proto.RegisterType((*GaCeRespLogin)(nil), "smsg.GaCeRespLogin")
	proto.RegisterType((*CeGaBindGameServer)(nil), "smsg.CeGaBindGameServer")
	proto.RegisterType((*CeGamReqJoinGame)(nil), "smsg.CeGamReqJoinGame")
	proto.RegisterType((*CeGamRespJoinGame)(nil), "smsg.CeGamRespJoinGame")
	proto.RegisterType((*GamCeNoticeGameStart)(nil), "smsg.GamCeNoticeGameStart")
	proto.RegisterType((*GamCeNoticeGameEnd)(nil), "smsg.GamCeNoticeGameEnd")
	proto.RegisterType((*CeGameUserDisconnect)(nil), "smsg.CeGameUserDisconnect")
	proto.RegisterType((*CeGameUserReconnect)(nil), "smsg.CeGameUserReconnect")
	proto.RegisterType((*GaCeUserDisconnect)(nil), "smsg.GaCeUserDisconnect")
	proto.RegisterType((*CeGaCloseSession)(nil), "smsg.CeGaCloseSession")
	proto.RegisterType((*AdReqMetrics)(nil), "smsg.AdReqMetrics")
	proto.RegisterType((*AdRespMetrics)(nil), "smsg.AdRespMetrics")
	proto.RegisterType((*AdRespMetrics_Metrics)(nil), "smsg.AdRespMetrics.Metrics")
	proto.RegisterEnum("smsg.GaCeRespLogin_Error", GaCeRespLogin_Error_name, GaCeRespLogin_Error_value)
	proto.RegisterEnum("smsg.CeGamRespJoinGame_Error", CeGamRespJoinGame_Error_name, CeGamRespJoinGame_Error_value)
	proto.RegisterEnum("smsg.AdRespMetrics_MetricsType", AdRespMetrics_MetricsType_name, AdRespMetrics_MetricsType_value)
}

func init() { proto.RegisterFile("msg/smsg/smsg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 573 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x66, 0xeb, 0xa4, 0x69, 0x27, 0x29, 0xa4, 0x6e, 0x54, 0x85, 0x02, 0x22, 0xda, 0x03, 0x44,
	0x2a, 0x72, 0x21, 0x88, 0x0b, 0xb7, 0xe2, 0x54, 0x51, 0x10, 0x14, 0x69, 0x5b, 0x1e, 0xc0, 0xd8,
	0xa3, 0x6a, 0x55, 0x7b, 0x9d, 0xec, 0x6e, 0x2b, 0x7a, 0xe3, 0xc0, 0x43, 0xf0, 0x7a, 0xbc, 0x09,
	0xda, 0x1f, 0x37, 0x76, 0x21, 0xe5, 0x92, 0xec, 0x8c, 0xbe, 0x6f, 0xe6, 0x9b, 0x6f, 0x67, 0x0d,
	0x7b, 0x85, 0xba, 0x38, 0x52, 0xd5, 0x4f, 0xb4, 0x90, 0xa5, 0x2e, 0xc3, 0x96, 0x39, 0xd3, 0x97,
	0xb0, 0x7b, 0x86, 0xf2, 0x1a, 0xe5, 0xe4, 0x38, 0xcf, 0xcf, 0x50, 0x29, 0x5e, 0x8a, 0x30, 0x84,
	0x56, 0x96, 0xe8, 0x64, 0x48, 0x46, 0x64, 0xdc, 0x63, 0xf6, 0x4c, 0xdf, 0x43, 0x6f, 0x96, 0xc4,
	0xc8, 0x70, 0xf9, 0xa9, 0xbc, 0xe0, 0x22, 0x1c, 0x40, 0x5b, 0xa1, 0xe2, 0x99, 0x05, 0xb5, 0x99,
	0x0b, 0x4c, 0x56, 0x97, 0x97, 0x28, 0x86, 0x1b, 0x23, 0x32, 0xde, 0x66, 0x2e, 0xa0, 0xbf, 0x08,
	0xec, 0x38, 0xb2, 0x5a, 0x38, 0xf6, 0x3e, 0x6c, 0x5e, 0x29, 0x94, 0xf3, 0xa9, 0xa5, 0xb7, 0x98,
	0x8f, 0xc2, 0x43, 0x08, 0x50, 0x4a, 0xcb, 0x7e, 0x38, 0x79, 0x1c, 0x59, 0xb9, 0x0d, 0x66, 0x74,
	0x22, 0x65, 0x29, 0x99, 0x41, 0xad, 0x9a, 0x05, 0xb5, 0x66, 0xa6, 0x34, 0x17, 0xb3, 0xa4, 0xc0,
	0x61, 0x6b, 0x44, 0xc6, 0x5b, 0xcc, 0x47, 0x74, 0x00, 0x6d, 0xcb, 0x0d, 0xbb, 0xd0, 0x99, 0x8b,
	0xeb, 0x24, 0xe7, 0x59, 0xff, 0x01, 0x3d, 0x85, 0x30, 0xc6, 0x59, 0xf2, 0x81, 0x8b, 0xcc, 0xa0,
	0x9c, 0x17, 0x6b, 0x86, 0xa3, 0xd0, 0xbb, 0x48, 0x0a, 0x54, 0x16, 0xc3, 0x33, 0xab, 0xb2, 0xcd,
	0x1a, 0x39, 0xfa, 0x83, 0x40, 0xdf, 0x14, 0x2c, 0x18, 0x2e, 0x3f, 0x96, 0xae, 0x75, 0x35, 0xad,
	0xaf, 0xe7, 0xa7, 0x75, 0x6e, 0xb9, 0x36, 0x1b, 0xf5, 0x36, 0x07, 0xb0, 0x25, 0x78, 0x7a, 0x29,
	0xcc, 0x08, 0x6e, 0xb2, 0xdb, 0xd8, 0x49, 0xd0, 0x5e, 0x26, 0xcf, 0xec, 0x88, 0x56, 0xc2, 0x2a,
	0x47, 0x7f, 0x12, 0xd8, 0xf5, 0x12, 0xd4, 0xe2, 0x56, 0xc3, 0x91, 0x73, 0x96, 0x58, 0x67, 0x9f,
	0x39, 0x67, 0xff, 0x42, 0xd5, 0xdd, 0xdd, 0x87, 0x4d, 0x33, 0x99, 0x57, 0x17, 0x30, 0x1f, 0xd1,
	0x17, 0xff, 0xf2, 0x31, 0xec, 0x9b, 0xf5, 0x28, 0xf0, 0xb4, 0xd4, 0x27, 0xdf, 0xb9, 0xd2, 0x7d,
	0x42, 0x23, 0x18, 0xcc, 0x92, 0x22, 0x36, 0x29, 0x9e, 0xa2, 0x35, 0x57, 0x27, 0x52, 0xd7, 0xea,
	0x92, 0x46, 0xdd, 0x57, 0x10, 0xde, 0xc1, 0x9f, 0x88, 0x6c, 0x2d, 0x3a, 0x82, 0x81, 0x55, 0x8f,
	0x5f, 0x15, 0xca, 0x29, 0x57, 0x69, 0x29, 0x04, 0xa6, 0x7a, 0x9d, 0xd5, 0x34, 0x85, 0xbd, 0x15,
	0x9e, 0xe1, 0x7f, 0xe0, 0xae, 0xad, 0xc6, 0xf9, 0xd4, 0x5f, 0x8d, 0x8f, 0xc2, 0xa7, 0xb0, 0xad,
	0xdc, 0x23, 0x99, 0x4f, 0xed, 0xe5, 0xb4, 0xd9, 0x2a, 0x41, 0x27, 0x66, 0x84, 0xf8, 0xae, 0xa4,
	0x06, 0x87, 0xdc, 0xe5, 0xbc, 0x76, 0xfb, 0x12, 0xe7, 0xa5, 0xc2, 0xea, 0xfd, 0xdd, 0xcf, 0x18,
	0x43, 0xef, 0x38, 0x63, 0xb8, 0xfc, 0x8c, 0x5a, 0xf2, 0x54, 0x85, 0x43, 0xe8, 0x48, 0x5c, 0x9e,
	0x73, 0xbf, 0x2e, 0x01, 0xab, 0x42, 0xfa, 0x9b, 0xc0, 0x8e, 0x81, 0xaa, 0x45, 0x85, 0x7d, 0x07,
	0x9d, 0xc2, 0x1d, 0x87, 0x64, 0x14, 0x8c, 0xbb, 0x93, 0x27, 0x6e, 0x13, 0x1a, 0xa8, 0xc8, 0xff,
	0xb3, 0x0a, 0x5b, 0x6f, 0xb1, 0xd1, 0x68, 0x71, 0xc0, 0xa0, 0x53, 0xd5, 0x7e, 0x03, 0xc1, 0x25,
	0xde, 0xf8, 0x0d, 0x7b, 0x7e, 0x4f, 0xdd, 0xf3, 0x9b, 0x05, 0x32, 0x83, 0x35, 0x0f, 0xe0, 0x3a,
	0xc9, 0xaf, 0xb0, 0x7a, 0x00, 0x36, 0xa0, 0x87, 0xd0, 0xad, 0x21, 0x9b, 0x7b, 0xf6, 0x08, 0xba,
	0x5f, 0x44, 0xce, 0x05, 0xc6, 0xe5, 0x95, 0xd0, 0x7d, 0xf2, 0x6d, 0xd3, 0x7e, 0xcd, 0xde, 0xfe,
	0x09, 0x00, 0x00, 0xff, 0xff, 0x70, 0x85, 0x65, 0x2d, 0xe4, 0x04, 0x00, 0x00,
}
