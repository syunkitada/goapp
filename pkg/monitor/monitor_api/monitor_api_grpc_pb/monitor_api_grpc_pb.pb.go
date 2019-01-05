// Code generated by protoc-gen-go. DO NOT EDIT.
// source: monitor_api_grpc_pb.proto

package monitor_api_grpc_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_api_grpc_pb_d28824081f9542c0, []int{0}
}
func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequest.Unmarshal(m, b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
}
func (dst *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(dst, src)
}
func (m *StatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatusRequest.Size(m)
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

type StatusReply struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusReply) Reset()         { *m = StatusReply{} }
func (m *StatusReply) String() string { return proto.CompactTextString(m) }
func (*StatusReply) ProtoMessage()    {}
func (*StatusReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_api_grpc_pb_d28824081f9542c0, []int{1}
}
func (m *StatusReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusReply.Unmarshal(m, b)
}
func (m *StatusReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusReply.Marshal(b, m, deterministic)
}
func (dst *StatusReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusReply.Merge(dst, src)
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

//
// Node
//
type GetNodeRequest struct {
	TraceId              string   `protobuf:"bytes,1,opt,name=trace_id,json=traceId" json:"trace_id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	RoleName             string   `protobuf:"bytes,3,opt,name=role_name,json=roleName" json:"role_name,omitempty"`
	ProjectName          string   `protobuf:"bytes,4,opt,name=project_name,json=projectName" json:"project_name,omitempty"`
	ProjectRoleName      string   `protobuf:"bytes,5,opt,name=project_role_name,json=projectRoleName" json:"project_role_name,omitempty"`
	Target               string   `protobuf:"bytes,6,opt,name=target" json:"target,omitempty"`
	Cluster              string   `protobuf:"bytes,7,opt,name=cluster" json:"cluster,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNodeRequest) Reset()         { *m = GetNodeRequest{} }
func (m *GetNodeRequest) String() string { return proto.CompactTextString(m) }
func (*GetNodeRequest) ProtoMessage()    {}
func (*GetNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_api_grpc_pb_d28824081f9542c0, []int{2}
}
func (m *GetNodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNodeRequest.Unmarshal(m, b)
}
func (m *GetNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNodeRequest.Marshal(b, m, deterministic)
}
func (dst *GetNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNodeRequest.Merge(dst, src)
}
func (m *GetNodeRequest) XXX_Size() int {
	return xxx_messageInfo_GetNodeRequest.Size(m)
}
func (m *GetNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNodeRequest proto.InternalMessageInfo

func (m *GetNodeRequest) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *GetNodeRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *GetNodeRequest) GetRoleName() string {
	if m != nil {
		return m.RoleName
	}
	return ""
}

func (m *GetNodeRequest) GetProjectName() string {
	if m != nil {
		return m.ProjectName
	}
	return ""
}

func (m *GetNodeRequest) GetProjectRoleName() string {
	if m != nil {
		return m.ProjectRoleName
	}
	return ""
}

func (m *GetNodeRequest) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *GetNodeRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

type GetNodeReply struct {
	StatusCode           int64    `protobuf:"varint,1,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
	Nodes                []*Node  `protobuf:"bytes,3,rep,name=nodes" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNodeReply) Reset()         { *m = GetNodeReply{} }
func (m *GetNodeReply) String() string { return proto.CompactTextString(m) }
func (*GetNodeReply) ProtoMessage()    {}
func (*GetNodeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_api_grpc_pb_d28824081f9542c0, []int{3}
}
func (m *GetNodeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNodeReply.Unmarshal(m, b)
}
func (m *GetNodeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNodeReply.Marshal(b, m, deterministic)
}
func (dst *GetNodeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNodeReply.Merge(dst, src)
}
func (m *GetNodeReply) XXX_Size() int {
	return xxx_messageInfo_GetNodeReply.Size(m)
}
func (m *GetNodeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNodeReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetNodeReply proto.InternalMessageInfo

func (m *GetNodeReply) GetStatusCode() int64 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *GetNodeReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func (m *GetNodeReply) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type UpdateNodeRequest struct {
	TraceId              string   `protobuf:"bytes,1,opt,name=trace_id,json=traceId" json:"trace_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Kind                 string   `protobuf:"bytes,3,opt,name=kind" json:"kind,omitempty"`
	Role                 string   `protobuf:"bytes,4,opt,name=role" json:"role,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
	StatusReason         string   `protobuf:"bytes,6,opt,name=status_reason,json=statusReason" json:"status_reason,omitempty"`
	State                string   `protobuf:"bytes,7,opt,name=state" json:"state,omitempty"`
	StateReason          string   `protobuf:"bytes,8,opt,name=state_reason,json=stateReason" json:"state_reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateNodeRequest) Reset()         { *m = UpdateNodeRequest{} }
func (m *UpdateNodeRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateNodeRequest) ProtoMessage()    {}
func (*UpdateNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_api_grpc_pb_d28824081f9542c0, []int{4}
}
func (m *UpdateNodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateNodeRequest.Unmarshal(m, b)
}
func (m *UpdateNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateNodeRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateNodeRequest.Merge(dst, src)
}
func (m *UpdateNodeRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateNodeRequest.Size(m)
}
func (m *UpdateNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateNodeRequest proto.InternalMessageInfo

func (m *UpdateNodeRequest) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *UpdateNodeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateNodeRequest) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *UpdateNodeRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *UpdateNodeRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *UpdateNodeRequest) GetStatusReason() string {
	if m != nil {
		return m.StatusReason
	}
	return ""
}

func (m *UpdateNodeRequest) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *UpdateNodeRequest) GetStateReason() string {
	if m != nil {
		return m.StateReason
	}
	return ""
}

type UpdateNodeReply struct {
	StatusCode           int64    `protobuf:"varint,1,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateNodeReply) Reset()         { *m = UpdateNodeReply{} }
func (m *UpdateNodeReply) String() string { return proto.CompactTextString(m) }
func (*UpdateNodeReply) ProtoMessage()    {}
func (*UpdateNodeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_api_grpc_pb_d28824081f9542c0, []int{5}
}
func (m *UpdateNodeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateNodeReply.Unmarshal(m, b)
}
func (m *UpdateNodeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateNodeReply.Marshal(b, m, deterministic)
}
func (dst *UpdateNodeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateNodeReply.Merge(dst, src)
}
func (m *UpdateNodeReply) XXX_Size() int {
	return xxx_messageInfo_UpdateNodeReply.Size(m)
}
func (m *UpdateNodeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateNodeReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateNodeReply proto.InternalMessageInfo

func (m *UpdateNodeReply) GetStatusCode() int64 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *UpdateNodeReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type Node struct {
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	Cluster              string               `protobuf:"bytes,3,opt,name=cluster" json:"cluster,omitempty"`
	Name                 string               `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Kind                 string               `protobuf:"bytes,5,opt,name=kind" json:"kind,omitempty"`
	Role                 string               `protobuf:"bytes,6,opt,name=role" json:"role,omitempty"`
	Status               string               `protobuf:"bytes,7,opt,name=status" json:"status,omitempty"`
	StatusReason         string               `protobuf:"bytes,8,opt,name=status_reason,json=statusReason" json:"status_reason,omitempty"`
	State                string               `protobuf:"bytes,9,opt,name=state" json:"state,omitempty"`
	StateReason          string               `protobuf:"bytes,10,opt,name=state_reason,json=stateReason" json:"state_reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_api_grpc_pb_d28824081f9542c0, []int{6}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Node) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Node) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Node) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Node) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Node) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *Node) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Node) GetStatusReason() string {
	if m != nil {
		return m.StatusReason
	}
	return ""
}

func (m *Node) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Node) GetStateReason() string {
	if m != nil {
		return m.StateReason
	}
	return ""
}

// Report
type ReportRequest struct {
	Index                string   `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
	Logs                 []*Log   `protobuf:"bytes,2,rep,name=logs" json:"logs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportRequest) Reset()         { *m = ReportRequest{} }
func (m *ReportRequest) String() string { return proto.CompactTextString(m) }
func (*ReportRequest) ProtoMessage()    {}
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_api_grpc_pb_d28824081f9542c0, []int{7}
}
func (m *ReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportRequest.Unmarshal(m, b)
}
func (m *ReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportRequest.Marshal(b, m, deterministic)
}
func (dst *ReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportRequest.Merge(dst, src)
}
func (m *ReportRequest) XXX_Size() int {
	return xxx_messageInfo_ReportRequest.Size(m)
}
func (m *ReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReportRequest proto.InternalMessageInfo

func (m *ReportRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *ReportRequest) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

type Log struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Time                 string            `protobuf:"bytes,2,opt,name=time" json:"time,omitempty"`
	Log                  map[string]string `protobuf:"bytes,3,rep,name=log" json:"log,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_api_grpc_pb_d28824081f9542c0, []int{8}
}
func (m *Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Log.Unmarshal(m, b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Log.Marshal(b, m, deterministic)
}
func (dst *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(dst, src)
}
func (m *Log) XXX_Size() int {
	return xxx_messageInfo_Log.Size(m)
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Log) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func (m *Log) GetLog() map[string]string {
	if m != nil {
		return m.Log
	}
	return nil
}

type ReportReply struct {
	StatusCode           int64    `protobuf:"varint,1,opt,name=status_code,json=statusCode" json:"status_code,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportReply) Reset()         { *m = ReportReply{} }
func (m *ReportReply) String() string { return proto.CompactTextString(m) }
func (*ReportReply) ProtoMessage()    {}
func (*ReportReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitor_api_grpc_pb_d28824081f9542c0, []int{9}
}
func (m *ReportReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportReply.Unmarshal(m, b)
}
func (m *ReportReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportReply.Marshal(b, m, deterministic)
}
func (dst *ReportReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportReply.Merge(dst, src)
}
func (m *ReportReply) XXX_Size() int {
	return xxx_messageInfo_ReportReply.Size(m)
}
func (m *ReportReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportReply.DiscardUnknown(m)
}

var xxx_messageInfo_ReportReply proto.InternalMessageInfo

func (m *ReportReply) GetStatusCode() int64 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *ReportReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*StatusRequest)(nil), "monitor_api_grpc_pb.StatusRequest")
	proto.RegisterType((*StatusReply)(nil), "monitor_api_grpc_pb.StatusReply")
	proto.RegisterType((*GetNodeRequest)(nil), "monitor_api_grpc_pb.GetNodeRequest")
	proto.RegisterType((*GetNodeReply)(nil), "monitor_api_grpc_pb.GetNodeReply")
	proto.RegisterType((*UpdateNodeRequest)(nil), "monitor_api_grpc_pb.UpdateNodeRequest")
	proto.RegisterType((*UpdateNodeReply)(nil), "monitor_api_grpc_pb.UpdateNodeReply")
	proto.RegisterType((*Node)(nil), "monitor_api_grpc_pb.Node")
	proto.RegisterType((*ReportRequest)(nil), "monitor_api_grpc_pb.ReportRequest")
	proto.RegisterType((*Log)(nil), "monitor_api_grpc_pb.Log")
	proto.RegisterMapType((map[string]string)(nil), "monitor_api_grpc_pb.Log.LogEntry")
	proto.RegisterType((*ReportReply)(nil), "monitor_api_grpc_pb.ReportReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MonitorApiClient is the client API for MonitorApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MonitorApiClient interface {
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error)
	GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeReply, error)
	UpdateNode(ctx context.Context, in *UpdateNodeRequest, opts ...grpc.CallOption) (*UpdateNodeReply, error)
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportReply, error)
}

type monitorApiClient struct {
	cc *grpc.ClientConn
}

func NewMonitorApiClient(cc *grpc.ClientConn) MonitorApiClient {
	return &monitorApiClient{cc}
}

func (c *monitorApiClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/monitor_api_grpc_pb.MonitorApi/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorApiClient) GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeReply, error) {
	out := new(GetNodeReply)
	err := c.cc.Invoke(ctx, "/monitor_api_grpc_pb.MonitorApi/GetNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorApiClient) UpdateNode(ctx context.Context, in *UpdateNodeRequest, opts ...grpc.CallOption) (*UpdateNodeReply, error) {
	out := new(UpdateNodeReply)
	err := c.cc.Invoke(ctx, "/monitor_api_grpc_pb.MonitorApi/UpdateNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorApiClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportReply, error) {
	out := new(ReportReply)
	err := c.cc.Invoke(ctx, "/monitor_api_grpc_pb.MonitorApi/Report", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MonitorApiServer is the server API for MonitorApi service.
type MonitorApiServer interface {
	Status(context.Context, *StatusRequest) (*StatusReply, error)
	GetNode(context.Context, *GetNodeRequest) (*GetNodeReply, error)
	UpdateNode(context.Context, *UpdateNodeRequest) (*UpdateNodeReply, error)
	Report(context.Context, *ReportRequest) (*ReportReply, error)
}

func RegisterMonitorApiServer(s *grpc.Server, srv MonitorApiServer) {
	s.RegisterService(&_MonitorApi_serviceDesc, srv)
}

func _MonitorApi_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorApiServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor_api_grpc_pb.MonitorApi/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorApiServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorApi_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorApiServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor_api_grpc_pb.MonitorApi/GetNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorApiServer).GetNode(ctx, req.(*GetNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorApi_UpdateNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorApiServer).UpdateNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor_api_grpc_pb.MonitorApi/UpdateNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorApiServer).UpdateNode(ctx, req.(*UpdateNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorApi_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorApiServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitor_api_grpc_pb.MonitorApi/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorApiServer).Report(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MonitorApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "monitor_api_grpc_pb.MonitorApi",
	HandlerType: (*MonitorApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _MonitorApi_Status_Handler,
		},
		{
			MethodName: "GetNode",
			Handler:    _MonitorApi_GetNode_Handler,
		},
		{
			MethodName: "UpdateNode",
			Handler:    _MonitorApi_UpdateNode_Handler,
		},
		{
			MethodName: "Report",
			Handler:    _MonitorApi_Report_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitor_api_grpc_pb.proto",
}

func init() {
	proto.RegisterFile("monitor_api_grpc_pb.proto", fileDescriptor_monitor_api_grpc_pb_d28824081f9542c0)
}

var fileDescriptor_monitor_api_grpc_pb_d28824081f9542c0 = []byte{
	// 709 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x6d, 0x4b, 0xdc, 0x5a,
	0x10, 0x36, 0x9b, 0x7d, 0x9d, 0xd5, 0xbb, 0xd7, 0x73, 0xe5, 0x12, 0xd7, 0x0f, 0x6a, 0xbc, 0x5c,
	0xa4, 0x94, 0x58, 0x14, 0x4a, 0xdb, 0x4f, 0xd5, 0xb6, 0x48, 0xc1, 0xca, 0x12, 0xdb, 0x6f, 0x85,
	0x10, 0x37, 0xd3, 0x90, 0x9a, 0xe4, 0xa4, 0x27, 0x27, 0xc5, 0xfd, 0x2b, 0xfd, 0x45, 0xfd, 0x19,
	0x85, 0x42, 0x7f, 0x47, 0x99, 0x73, 0x4e, 0xd4, 0xc5, 0xb8, 0x16, 0x3f, 0x2c, 0xcc, 0xcb, 0x33,
	0xb3, 0x99, 0x67, 0x9e, 0x33, 0xb0, 0x9e, 0xf1, 0x3c, 0x91, 0x5c, 0x04, 0x61, 0x91, 0x04, 0xb1,
	0x28, 0xa6, 0x41, 0x71, 0xee, 0x15, 0x82, 0x4b, 0xce, 0xfe, 0x69, 0x48, 0x8d, 0x37, 0x63, 0xce,
	0xe3, 0x14, 0xf7, 0x14, 0xe4, 0xbc, 0xfa, 0xb4, 0x27, 0x93, 0x0c, 0x4b, 0x19, 0x66, 0x85, 0xae,
	0x72, 0x47, 0xb0, 0x72, 0x26, 0x43, 0x59, 0x95, 0x3e, 0x7e, 0xa9, 0xb0, 0x94, 0xee, 0x26, 0x0c,
	0xeb, 0x40, 0x91, 0xce, 0xd8, 0xdf, 0x60, 0x67, 0x65, 0xec, 0x58, 0x5b, 0xd6, 0xee, 0xc0, 0x27,
	0xd3, 0xfd, 0x65, 0xc1, 0x5f, 0xc7, 0x28, 0x4f, 0x79, 0x84, 0xa6, 0x86, 0xad, 0x43, 0x5f, 0x8a,
	0x70, 0x8a, 0x41, 0x12, 0x19, 0x64, 0x4f, 0xf9, 0x6f, 0x23, 0xb6, 0x01, 0x83, 0xaa, 0x44, 0x11,
	0xe4, 0x61, 0x86, 0x4e, 0x4b, 0xe5, 0xfa, 0x14, 0x38, 0x0d, 0x33, 0xa4, 0xa4, 0xe0, 0x29, 0xea,
	0xa4, 0xad, 0x93, 0x14, 0x50, 0xc9, 0x6d, 0x58, 0x2e, 0x04, 0xff, 0x8c, 0x53, 0xa9, 0xf3, 0x6d,
	0x95, 0x1f, 0x9a, 0x98, 0x82, 0x3c, 0x82, 0xd5, 0x1a, 0x72, 0xdd, 0xa7, 0xa3, 0x70, 0x23, 0x93,
	0xf0, 0xeb, 0x76, 0xff, 0x42, 0x57, 0x86, 0x22, 0x46, 0xe9, 0x74, 0x15, 0xc0, 0x78, 0xcc, 0x81,
	0xde, 0x34, 0xad, 0x4a, 0x89, 0xc2, 0xe9, 0xe9, 0x4f, 0x37, 0xae, 0x2b, 0x60, 0xf9, 0x6a, 0x4e,
	0xa2, 0x62, 0x13, 0x86, 0xa5, 0x62, 0x26, 0x98, 0xf2, 0x08, 0xd5, 0xa0, 0xb6, 0x0f, 0x3a, 0xf4,
	0x8a, 0x47, 0x48, 0x5c, 0xa1, 0x10, 0x66, 0x4a, 0x32, 0xd9, 0x1e, 0x74, 0x72, 0x1e, 0x61, 0xe9,
	0xd8, 0x5b, 0xf6, 0xee, 0x70, 0x7f, 0xdd, 0x6b, 0x5a, 0x9f, 0xfa, 0x07, 0x8d, 0x73, 0x7f, 0x58,
	0xb0, 0xfa, 0xa1, 0x88, 0x42, 0x89, 0x7f, 0xc8, 0x2f, 0x83, 0xf6, 0x0d, 0x6a, 0x95, 0x4d, 0xb1,
	0x8b, 0x24, 0x8f, 0x0c, 0xa3, 0xca, 0xa6, 0x18, 0x51, 0x64, 0x58, 0x54, 0x36, 0x51, 0xa2, 0xbf,
	0xde, 0x70, 0x66, 0x3c, 0xb6, 0x03, 0x2b, 0x66, 0x50, 0x81, 0x61, 0xc9, 0x73, 0xc3, 0xd8, 0x72,
	0x69, 0x74, 0x41, 0x31, 0xb6, 0x06, 0x1d, 0xf2, 0xd1, 0xb0, 0xa6, 0x1d, 0x5a, 0x9a, 0x32, 0xea,
	0xca, 0xbe, 0x5e, 0x9a, 0x8a, 0xe9, 0x42, 0xf7, 0x35, 0x8c, 0x6e, 0x4e, 0xf8, 0x30, 0x66, 0xdd,
	0xef, 0x2d, 0x68, 0x53, 0x03, 0xf6, 0x1c, 0xa0, 0x52, 0xed, 0xa2, 0x20, 0x94, 0xaa, 0x74, 0xb8,
	0x3f, 0xf6, 0xb4, 0xec, 0xbd, 0x5a, 0xf6, 0xde, 0xfb, 0x5a, 0xf6, 0xfe, 0xc0, 0xa0, 0x0f, 0x25,
	0x95, 0x4e, 0x05, 0xd6, 0xa5, 0xad, 0xfb, 0x4b, 0x0d, 0xfa, 0x70, 0x4e, 0x35, 0xf6, 0x9c, 0x6a,
	0xae, 0x16, 0xd2, 0x6e, 0x58, 0x48, 0xa7, 0x61, 0x21, 0xdd, 0xc6, 0x85, 0xf4, 0x16, 0x2f, 0xa4,
	0xbf, 0x68, 0x21, 0x83, 0x45, 0x0b, 0x81, 0xdb, 0x0b, 0x39, 0x83, 0x15, 0x1f, 0x0b, 0x2e, 0x64,
	0x2d, 0xb7, 0x35, 0xe8, 0x24, 0x79, 0x84, 0x97, 0x46, 0x6b, 0xda, 0x61, 0x8f, 0xa1, 0x9d, 0xf2,
	0xb8, 0x74, 0x5a, 0x4a, 0xca, 0x4e, 0xa3, 0x94, 0x4f, 0x78, 0xec, 0x2b, 0x94, 0xfb, 0xcd, 0x02,
	0xfb, 0x84, 0xc7, 0x57, 0x74, 0x58, 0xf3, 0x74, 0xd0, 0x19, 0xaa, 0x35, 0x4b, 0x36, 0x3b, 0x00,
	0x3b, 0xe5, 0xb1, 0x79, 0x27, 0xdb, 0x77, 0x35, 0xa7, 0xdf, 0x9b, 0x5c, 0x8a, 0x99, 0x4f, 0xe8,
	0xf1, 0x53, 0xe8, 0xd7, 0x01, 0x92, 0xc8, 0x05, 0xce, 0xea, 0x43, 0x75, 0x81, 0x33, 0x1a, 0xe3,
	0x6b, 0x98, 0x56, 0xf5, 0xff, 0x68, 0xe7, 0x45, 0xeb, 0x99, 0xe5, 0xbe, 0x84, 0x61, 0x3d, 0xf1,
	0xc3, 0xe4, 0xb7, 0xff, 0xb3, 0x05, 0xf0, 0x4e, 0x7f, 0xe3, 0x61, 0x91, 0xb0, 0x09, 0x74, 0xf5,
	0xd1, 0x64, 0x6e, 0xe3, 0xa7, 0xcf, 0x9d, 0xd8, 0xf1, 0xd6, 0x42, 0x4c, 0x91, 0xce, 0xdc, 0x25,
	0x76, 0x06, 0x3d, 0x73, 0x7c, 0xd8, 0x4e, 0x23, 0x7c, 0xfe, 0x04, 0x8f, 0xb7, 0x17, 0x83, 0x74,
	0xd3, 0x8f, 0x00, 0xd7, 0x4f, 0x8f, 0xfd, 0xdf, 0x58, 0x72, 0xeb, 0xfa, 0x8c, 0xff, 0xbb, 0x17,
	0xa7, 0xbb, 0x4f, 0xa0, 0xab, 0x59, 0xbd, 0x83, 0x84, 0x39, 0x91, 0xdd, 0x41, 0xc2, 0x8d, 0xb5,
	0xb8, 0x4b, 0x47, 0x4f, 0x60, 0x23, 0xe1, 0x1e, 0xe5, 0x3c, 0xbc, 0x0c, 0xb3, 0x22, 0xc5, 0xd2,
	0x13, 0xbc, 0x92, 0x18, 0x57, 0x49, 0x84, 0x47, 0x23, 0x9f, 0xec, 0x63, 0xb2, 0x27, 0xf4, 0x5a,
	0x27, 0xd6, 0x79, 0x57, 0x3d, 0xdb, 0x83, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xca, 0x86, 0x5b,
	0x3f, 0x28, 0x07, 0x00, 0x00,
}
