// Code generated by protoc-gen-go. DO NOT EDIT.
// source: resource_cluster_controller_grpc_pb.proto

package resource_cluster_controller_grpc_pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56171b6f41e7559c, []int{0}
}

func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequest.Unmarshal(m, b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
}
func (m *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(m, src)
}
func (m *StatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatusRequest.Size(m)
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

type StatusReply struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusReply) Reset()         { *m = StatusReply{} }
func (m *StatusReply) String() string { return proto.CompactTextString(m) }
func (*StatusReply) ProtoMessage()    {}
func (*StatusReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_56171b6f41e7559c, []int{1}
}

func (m *StatusReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusReply.Unmarshal(m, b)
}
func (m *StatusReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusReply.Marshal(b, m, deterministic)
}
func (m *StatusReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusReply.Merge(m, src)
}
func (m *StatusReply) XXX_Size() int {
	return xxx_messageInfo_StatusReply.Size(m)
}
func (m *StatusReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusReply.DiscardUnknown(m)
}

var xxx_messageInfo_StatusReply proto.InternalMessageInfo

func (m *StatusReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *StatusReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*StatusRequest)(nil), "resource_cluster_controller_grpc_pb.StatusRequest")
	proto.RegisterType((*StatusReply)(nil), "resource_cluster_controller_grpc_pb.StatusReply")
}

func init() {
	proto.RegisterFile("resource_cluster_controller_grpc_pb.proto", fileDescriptor_56171b6f41e7559c)
}

var fileDescriptor_56171b6f41e7559c = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2c, 0x4a, 0x2d, 0xce,
	0x2f, 0x2d, 0x4a, 0x4e, 0x8d, 0x4f, 0xce, 0x29, 0x2d, 0x2e, 0x49, 0x2d, 0x8a, 0x4f, 0xce, 0xcf,
	0x2b, 0x29, 0xca, 0xcf, 0xc9, 0x49, 0x2d, 0x8a, 0x4f, 0x2f, 0x2a, 0x48, 0x8e, 0x2f, 0x48, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x26, 0x42, 0xa9, 0x12, 0x3f, 0x17, 0x6f, 0x70, 0x49,
	0x62, 0x49, 0x69, 0x71, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x92, 0x21, 0x17, 0x37, 0x4c,
	0xa0, 0x20, 0xa7, 0x52, 0x48, 0x80, 0x8b, 0x39, 0xb7, 0x38, 0x5d, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x33, 0x08, 0xc4, 0x04, 0x89, 0xa4, 0x16, 0x15, 0x49, 0x30, 0x41, 0x44, 0x52, 0x8b, 0x8a, 0x8c,
	0x7a, 0x19, 0xb9, 0x24, 0x83, 0xa0, 0x76, 0x39, 0x43, 0xac, 0x72, 0x86, 0xdb, 0x24, 0x54, 0xc0,
	0xc5, 0x06, 0x31, 0x50, 0xc8, 0x48, 0x8f, 0x18, 0xc7, 0xa3, 0x38, 0x47, 0xca, 0x80, 0x24, 0x3d,
	0x05, 0x39, 0x95, 0x4a, 0x0c, 0x4e, 0x06, 0x5c, 0xd2, 0x99, 0xf9, 0x7a, 0x20, 0x39, 0xbd, 0xd4,
	0x8a, 0xc4, 0xdc, 0x82, 0x9c, 0xd4, 0x62, 0xbd, 0xa2, 0xfc, 0xd2, 0x92, 0xd4, 0xf4, 0xd2, 0xcc,
	0x94, 0x54, 0x27, 0xfe, 0x20, 0x10, 0xdb, 0x1d, 0xc4, 0x0e, 0x00, 0x05, 0x54, 0x00, 0x63, 0x12,
	0x1b, 0x38, 0xc4, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xcb, 0xd4, 0xe9, 0x5e, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResourceClusterControllerClient is the client API for ResourceClusterController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourceClusterControllerClient interface {
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error)
}

type resourceClusterControllerClient struct {
	cc *grpc.ClientConn
}

func NewResourceClusterControllerClient(cc *grpc.ClientConn) ResourceClusterControllerClient {
	return &resourceClusterControllerClient{cc}
}

func (c *resourceClusterControllerClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/resource_cluster_controller_grpc_pb.ResourceClusterController/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceClusterControllerServer is the server API for ResourceClusterController service.
type ResourceClusterControllerServer interface {
	Status(context.Context, *StatusRequest) (*StatusReply, error)
}

func RegisterResourceClusterControllerServer(s *grpc.Server, srv ResourceClusterControllerServer) {
	s.RegisterService(&_ResourceClusterController_serviceDesc, srv)
}

func _ResourceClusterController_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceClusterControllerServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resource_cluster_controller_grpc_pb.ResourceClusterController/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceClusterControllerServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceClusterController_serviceDesc = grpc.ServiceDesc{
	ServiceName: "resource_cluster_controller_grpc_pb.ResourceClusterController",
	HandlerType: (*ResourceClusterControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _ResourceClusterController_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resource_cluster_controller_grpc_pb.proto",
}
