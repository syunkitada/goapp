// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authproxy/authproxy_grpc_pb/authproxy_grpc_pb.proto

package authproxy_grpc_pb // import "github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"

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

type TraceContext struct {
	TraceId              string   `protobuf:"bytes,1,opt,name=TraceId" json:"TraceId,omitempty"`
	App                  string   `protobuf:"bytes,2,opt,name=App" json:"App,omitempty"`
	Host                 string   `protobuf:"bytes,3,opt,name=Host" json:"Host,omitempty"`
	ActionName           string   `protobuf:"bytes,4,opt,name=ActionName" json:"ActionName,omitempty"`
	UserName             string   `protobuf:"bytes,5,opt,name=UserName" json:"UserName,omitempty"`
	RoleName             string   `protobuf:"bytes,6,opt,name=RoleName" json:"RoleName,omitempty"`
	ProjectName          string   `protobuf:"bytes,7,opt,name=ProjectName" json:"ProjectName,omitempty"`
	ProjectRoleName      string   `protobuf:"bytes,8,opt,name=ProjectRoleName" json:"ProjectRoleName,omitempty"`
	StatusCode           int64    `protobuf:"varint,9,opt,name=StatusCode" json:"StatusCode,omitempty"`
	Err                  string   `protobuf:"bytes,10,opt,name=Err" json:"Err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TraceContext) Reset()         { *m = TraceContext{} }
func (m *TraceContext) String() string { return proto.CompactTextString(m) }
func (*TraceContext) ProtoMessage()    {}
func (*TraceContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_authproxy_grpc_pb_d281370763a33e61, []int{0}
}
func (m *TraceContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TraceContext.Unmarshal(m, b)
}
func (m *TraceContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TraceContext.Marshal(b, m, deterministic)
}
func (dst *TraceContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TraceContext.Merge(dst, src)
}
func (m *TraceContext) XXX_Size() int {
	return xxx_messageInfo_TraceContext.Size(m)
}
func (m *TraceContext) XXX_DiscardUnknown() {
	xxx_messageInfo_TraceContext.DiscardUnknown(m)
}

var xxx_messageInfo_TraceContext proto.InternalMessageInfo

func (m *TraceContext) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *TraceContext) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *TraceContext) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *TraceContext) GetActionName() string {
	if m != nil {
		return m.ActionName
	}
	return ""
}

func (m *TraceContext) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *TraceContext) GetRoleName() string {
	if m != nil {
		return m.RoleName
	}
	return ""
}

func (m *TraceContext) GetProjectName() string {
	if m != nil {
		return m.ProjectName
	}
	return ""
}

func (m *TraceContext) GetProjectRoleName() string {
	if m != nil {
		return m.ProjectRoleName
	}
	return ""
}

func (m *TraceContext) GetStatusCode() int64 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *TraceContext) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*TraceContext)(nil), "authproxy_grpc_pb.TraceContext")
}

func init() {
	proto.RegisterFile("authproxy/authproxy_grpc_pb/authproxy_grpc_pb.proto", fileDescriptor_authproxy_grpc_pb_d281370763a33e61)
}

var fileDescriptor_authproxy_grpc_pb_d281370763a33e61 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x5b, 0xfb, 0x31, 0x0a, 0xd5, 0x3d, 0x2d, 0x0a, 0x52, 0x3c, 0xf5, 0x94, 0x1c,
	0x7a, 0xf4, 0xd4, 0x96, 0xa2, 0x5e, 0xa4, 0x44, 0xbd, 0x78, 0x29, 0xdb, 0x64, 0x49, 0x63, 0x9b,
	0xce, 0xb2, 0x99, 0x85, 0xf4, 0xff, 0xf8, 0x43, 0x65, 0xa7, 0x36, 0x04, 0x0b, 0xde, 0xde, 0xf7,
	0x79, 0x42, 0xd8, 0x99, 0x81, 0x89, 0x72, 0xb4, 0x31, 0x16, 0xab, 0x43, 0x54, 0xa7, 0x55, 0x66,
	0x4d, 0xb2, 0x32, 0xeb, 0x73, 0x12, 0x1a, 0x8b, 0x84, 0xe2, 0xe6, 0x4c, 0x3c, 0x7c, 0xb7, 0xe0,
	0xea, 0xdd, 0xaa, 0x44, 0xcf, 0x71, 0x4f, 0xba, 0x22, 0x21, 0xa1, 0xc7, 0xfd, 0x25, 0x95, 0xc1,
	0x28, 0x18, 0x0f, 0xe2, 0x53, 0x15, 0xd7, 0xd0, 0x9e, 0x1a, 0x23, 0x5b, 0x4c, 0x7d, 0x14, 0x02,
	0x3a, 0xcf, 0x58, 0x92, 0x6c, 0x33, 0xe2, 0x2c, 0xee, 0x01, 0xa6, 0x09, 0xe5, 0xb8, 0x7f, 0x55,
	0x85, 0x96, 0x1d, 0x36, 0x0d, 0x22, 0x6e, 0xa1, 0xff, 0x51, 0x6a, 0xcb, 0xf6, 0x82, 0x6d, 0xdd,
	0xbd, 0x8b, 0x71, 0xa7, 0xd9, 0x75, 0x8f, 0xee, 0xd4, 0xc5, 0x08, 0x2e, 0x97, 0x16, 0xbf, 0x74,
	0x42, 0xac, 0x7b, 0xac, 0x9b, 0x48, 0x8c, 0x61, 0xf8, 0x5b, 0xeb, 0x9f, 0xf4, 0xf9, 0xab, 0xbf,
	0xd8, 0xbf, 0xf1, 0x8d, 0x14, 0xb9, 0x72, 0x8e, 0xa9, 0x96, 0x83, 0x51, 0x30, 0x6e, 0xc7, 0x0d,
	0xe2, 0x27, 0x5d, 0x58, 0x2b, 0xe1, 0x38, 0xe9, 0xc2, 0xda, 0x59, 0x01, 0x77, 0x39, 0x86, 0x7e,
	0x69, 0xa1, 0xae, 0x54, 0x61, 0x76, 0xba, 0x0c, 0x2d, 0x3a, 0xd2, 0x99, 0xcb, 0x53, 0x3d, 0x1b,
	0xc6, 0x3e, 0x3f, 0xf9, 0xbc, 0xf4, 0x9b, 0x5e, 0x06, 0x9f, 0x8f, 0x59, 0x4e, 0x1b, 0xb7, 0x0e,
	0x13, 0x2c, 0xa2, 0xf2, 0xe0, 0xf6, 0xdb, 0x9c, 0x54, 0xaa, 0xa2, 0x0c, 0x95, 0x31, 0x91, 0xd9,
	0x66, 0xd1, 0x3f, 0x07, 0x5c, 0x77, 0xf9, 0x5e, 0x93, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89,
	0xa7, 0x2a, 0x5c, 0xe6, 0x01, 0x00, 0x00,
}
