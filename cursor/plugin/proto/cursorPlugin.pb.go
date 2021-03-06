// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugin/proto/cursorPlugin.proto

package proto

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

type ExecuteTaskResponse_Status int32

const (
	ExecuteTaskResponse_UNKNOWN ExecuteTaskResponse_Status = 0
	ExecuteTaskResponse_FAILED  ExecuteTaskResponse_Status = 1
	ExecuteTaskResponse_SUCCESS ExecuteTaskResponse_Status = 2
	ExecuteTaskResponse_RUNNING ExecuteTaskResponse_Status = 3
	ExecuteTaskResponse_WAITING ExecuteTaskResponse_Status = 4
)

var ExecuteTaskResponse_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "FAILED",
	2: "SUCCESS",
	3: "RUNNING",
	4: "WAITING",
}

var ExecuteTaskResponse_Status_value = map[string]int32{
	"UNKNOWN": 0,
	"FAILED":  1,
	"SUCCESS": 2,
	"RUNNING": 3,
	"WAITING": 4,
}

func (x ExecuteTaskResponse_Status) String() string {
	return proto.EnumName(ExecuteTaskResponse_Status_name, int32(x))
}

func (ExecuteTaskResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f0f3a40d65296300, []int{4, 0}
}

// Task is the smallest unit of work, it accomplishes a small part of an entire
// pipeline. Many tasks together make up a single pipeline.
type TaskInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Children             []string `protobuf:"bytes,3,rep,name=children,proto3" json:"children,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskInfo) Reset()         { *m = TaskInfo{} }
func (m *TaskInfo) String() string { return proto.CompactTextString(m) }
func (*TaskInfo) ProtoMessage()    {}
func (*TaskInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0f3a40d65296300, []int{0}
}

func (m *TaskInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskInfo.Unmarshal(m, b)
}
func (m *TaskInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskInfo.Marshal(b, m, deterministic)
}
func (m *TaskInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskInfo.Merge(m, src)
}
func (m *TaskInfo) XXX_Size() int {
	return xxx_messageInfo_TaskInfo.Size(m)
}
func (m *TaskInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TaskInfo proto.InternalMessageInfo

func (m *TaskInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaskInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TaskInfo) GetChildren() []string {
	if m != nil {
		return m.Children
	}
	return nil
}

type GetPipelineInfoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPipelineInfoRequest) Reset()         { *m = GetPipelineInfoRequest{} }
func (m *GetPipelineInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetPipelineInfoRequest) ProtoMessage()    {}
func (*GetPipelineInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0f3a40d65296300, []int{1}
}

func (m *GetPipelineInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPipelineInfoRequest.Unmarshal(m, b)
}
func (m *GetPipelineInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPipelineInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetPipelineInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPipelineInfoRequest.Merge(m, src)
}
func (m *GetPipelineInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetPipelineInfoRequest.Size(m)
}
func (m *GetPipelineInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPipelineInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPipelineInfoRequest proto.InternalMessageInfo

type GetPipelineInfoResponse struct {
	RootTask             string               `protobuf:"bytes,1,opt,name=root_task,json=rootTask,proto3" json:"root_task,omitempty"`
	Tasks                map[string]*TaskInfo `protobuf:"bytes,2,rep,name=tasks,proto3" json:"tasks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetPipelineInfoResponse) Reset()         { *m = GetPipelineInfoResponse{} }
func (m *GetPipelineInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetPipelineInfoResponse) ProtoMessage()    {}
func (*GetPipelineInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0f3a40d65296300, []int{2}
}

func (m *GetPipelineInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPipelineInfoResponse.Unmarshal(m, b)
}
func (m *GetPipelineInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPipelineInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetPipelineInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPipelineInfoResponse.Merge(m, src)
}
func (m *GetPipelineInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetPipelineInfoResponse.Size(m)
}
func (m *GetPipelineInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPipelineInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPipelineInfoResponse proto.InternalMessageInfo

func (m *GetPipelineInfoResponse) GetRootTask() string {
	if m != nil {
		return m.RootTask
	}
	return ""
}

func (m *GetPipelineInfoResponse) GetTasks() map[string]*TaskInfo {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type ExecuteTaskRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteTaskRequest) Reset()         { *m = ExecuteTaskRequest{} }
func (m *ExecuteTaskRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteTaskRequest) ProtoMessage()    {}
func (*ExecuteTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0f3a40d65296300, []int{3}
}

func (m *ExecuteTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteTaskRequest.Unmarshal(m, b)
}
func (m *ExecuteTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteTaskRequest.Marshal(b, m, deterministic)
}
func (m *ExecuteTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteTaskRequest.Merge(m, src)
}
func (m *ExecuteTaskRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteTaskRequest.Size(m)
}
func (m *ExecuteTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteTaskRequest proto.InternalMessageInfo

func (m *ExecuteTaskRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ExecuteTaskResponse struct {
	Status               ExecuteTaskResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=proto.ExecuteTaskResponse_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ExecuteTaskResponse) Reset()         { *m = ExecuteTaskResponse{} }
func (m *ExecuteTaskResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteTaskResponse) ProtoMessage()    {}
func (*ExecuteTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0f3a40d65296300, []int{4}
}

func (m *ExecuteTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteTaskResponse.Unmarshal(m, b)
}
func (m *ExecuteTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteTaskResponse.Marshal(b, m, deterministic)
}
func (m *ExecuteTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteTaskResponse.Merge(m, src)
}
func (m *ExecuteTaskResponse) XXX_Size() int {
	return xxx_messageInfo_ExecuteTaskResponse.Size(m)
}
func (m *ExecuteTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteTaskResponse proto.InternalMessageInfo

func (m *ExecuteTaskResponse) GetStatus() ExecuteTaskResponse_Status {
	if m != nil {
		return m.Status
	}
	return ExecuteTaskResponse_UNKNOWN
}

func init() {
	proto.RegisterEnum("proto.ExecuteTaskResponse_Status", ExecuteTaskResponse_Status_name, ExecuteTaskResponse_Status_value)
	proto.RegisterType((*TaskInfo)(nil), "proto.TaskInfo")
	proto.RegisterType((*GetPipelineInfoRequest)(nil), "proto.GetPipelineInfoRequest")
	proto.RegisterType((*GetPipelineInfoResponse)(nil), "proto.GetPipelineInfoResponse")
	proto.RegisterMapType((map[string]*TaskInfo)(nil), "proto.GetPipelineInfoResponse.TasksEntry")
	proto.RegisterType((*ExecuteTaskRequest)(nil), "proto.ExecuteTaskRequest")
	proto.RegisterType((*ExecuteTaskResponse)(nil), "proto.ExecuteTaskResponse")
}

func init() { proto.RegisterFile("plugin/proto/cursorPlugin.proto", fileDescriptor_f0f3a40d65296300) }

var fileDescriptor_f0f3a40d65296300 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xd1, 0x8e, 0x94, 0x30,
	0x14, 0x15, 0xd8, 0x19, 0x67, 0x2e, 0x66, 0x97, 0x5c, 0x13, 0x45, 0x8c, 0x8a, 0x44, 0x93, 0xf1,
	0x85, 0x4d, 0xf0, 0x45, 0x7d, 0x31, 0x9b, 0x59, 0x5c, 0x89, 0x06, 0x27, 0xb0, 0x93, 0x7d, 0x31,
	0x31, 0x08, 0x55, 0x9b, 0xc1, 0x16, 0x69, 0x31, 0xee, 0xaf, 0xf8, 0x01, 0x7e, 0x88, 0x5f, 0x66,
	0x5a, 0xaa, 0x8e, 0xba, 0xe3, 0x13, 0xbd, 0xe7, 0xdc, 0x7b, 0x4f, 0xcf, 0xa1, 0x70, 0xa7, 0x6b,
	0x87, 0xf7, 0x94, 0x1d, 0x76, 0x3d, 0x97, 0xfc, 0xb0, 0x1e, 0x7a, 0xc1, 0xfb, 0x95, 0x86, 0x62,
	0x0d, 0xe1, 0x44, 0x7f, 0xa2, 0xd7, 0x30, 0x3b, 0xad, 0xc4, 0x26, 0x63, 0xef, 0x38, 0x22, 0xec,
	0xb1, 0xea, 0x23, 0xf1, 0xad, 0xd0, 0x5a, 0xcc, 0x0b, 0x7d, 0xc6, 0x10, 0xdc, 0x86, 0x88, 0xba,
	0xa7, 0x9d, 0xa4, 0x9c, 0xf9, 0xb6, 0xa6, 0xb6, 0x21, 0x0c, 0x60, 0x56, 0x7f, 0xa0, 0x6d, 0xd3,
	0x13, 0xe6, 0x3b, 0xa1, 0xb3, 0x98, 0x17, 0xbf, 0xea, 0xc8, 0x87, 0x6b, 0x27, 0x44, 0xae, 0x68,
	0x47, 0x5a, 0xca, 0x88, 0x12, 0x29, 0xc8, 0xa7, 0x81, 0x08, 0x19, 0x7d, 0xb7, 0xe0, 0xfa, 0x3f,
	0x94, 0xe8, 0x38, 0x13, 0x04, 0x6f, 0xc2, 0xbc, 0xe7, 0x5c, 0xbe, 0x91, 0x95, 0xd8, 0x98, 0xcb,
	0xcc, 0x14, 0xa0, 0x2e, 0x8a, 0x4f, 0x61, 0xa2, 0x70, 0xe1, 0xdb, 0xa1, 0xb3, 0x70, 0x93, 0x07,
	0xa3, 0x9d, 0x78, 0xc7, 0xae, 0x58, 0xcd, 0x88, 0x94, 0xc9, 0xfe, 0xbc, 0x18, 0xe7, 0x82, 0x0c,
	0xe0, 0x37, 0x88, 0x1e, 0x38, 0x1b, 0x72, 0x6e, 0x54, 0xd4, 0x11, 0xef, 0xc3, 0xe4, 0x73, 0xd5,
	0x0e, 0x44, 0x7b, 0x75, 0x93, 0x03, 0x23, 0xf0, 0x33, 0xa5, 0x62, 0x64, 0x9f, 0xd8, 0x8f, 0xac,
	0xe8, 0x1e, 0x60, 0xfa, 0x85, 0xd4, 0x83, 0x24, 0x8a, 0x35, 0xd6, 0x70, 0x1f, 0x6c, 0xda, 0x98,
	0x8d, 0x36, 0x6d, 0xa2, 0xaf, 0x16, 0x5c, 0xfd, 0xa3, 0xcd, 0xd8, 0x7c, 0x0c, 0x53, 0x21, 0x2b,
	0x39, 0x08, 0xdd, 0xbb, 0x9f, 0xdc, 0x35, 0x4a, 0x17, 0xf4, 0xc6, 0xa5, 0x6e, 0x2c, 0xcc, 0x40,
	0xf4, 0x1c, 0xa6, 0x23, 0x82, 0x2e, 0x5c, 0x5e, 0xe7, 0x2f, 0xf2, 0x57, 0x67, 0xb9, 0x77, 0x09,
	0x01, 0xa6, 0xcf, 0x8e, 0xb2, 0x97, 0xe9, 0xb1, 0x67, 0x29, 0xa2, 0x5c, 0x2f, 0x97, 0x69, 0x59,
	0x7a, 0xb6, 0x2a, 0x8a, 0x75, 0x9e, 0x67, 0xf9, 0x89, 0xe7, 0xa8, 0xe2, 0xec, 0x28, 0x3b, 0x55,
	0xc5, 0x5e, 0xf2, 0xcd, 0x82, 0x2b, 0xcb, 0xad, 0xd7, 0x81, 0x2b, 0x38, 0xf8, 0x2b, 0x4b, 0xbc,
	0xb5, 0x2b, 0x63, 0xed, 0x37, 0xb8, 0xfd, 0xff, 0x5f, 0x80, 0xc7, 0xe0, 0x6e, 0x59, 0xc2, 0x1b,
	0x17, 0xd9, 0x1c, 0x37, 0x05, 0xbb, 0x13, 0x78, 0x3b, 0xd5, 0xd4, 0xc3, 0x1f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x73, 0xeb, 0x02, 0x2f, 0xd9, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CursorPluginClient is the client API for CursorPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CursorPluginClient interface {
	// GetPipelineInfo is called upon compiling the pipeline so that cursor knowns
	// the structure of the tasks.
	// It returns a task info map that describes what kind of tasks exist in the
	// pipeline and a default task used as the root task
	GetPipelineInfo(ctx context.Context, in *GetPipelineInfoRequest, opts ...grpc.CallOption) (*GetPipelineInfoResponse, error)
	ExecuteTask(ctx context.Context, in *ExecuteTaskRequest, opts ...grpc.CallOption) (*ExecuteTaskResponse, error)
}

type cursorPluginClient struct {
	cc *grpc.ClientConn
}

func NewCursorPluginClient(cc *grpc.ClientConn) CursorPluginClient {
	return &cursorPluginClient{cc}
}

func (c *cursorPluginClient) GetPipelineInfo(ctx context.Context, in *GetPipelineInfoRequest, opts ...grpc.CallOption) (*GetPipelineInfoResponse, error) {
	out := new(GetPipelineInfoResponse)
	err := c.cc.Invoke(ctx, "/proto.CursorPlugin/GetPipelineInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cursorPluginClient) ExecuteTask(ctx context.Context, in *ExecuteTaskRequest, opts ...grpc.CallOption) (*ExecuteTaskResponse, error) {
	out := new(ExecuteTaskResponse)
	err := c.cc.Invoke(ctx, "/proto.CursorPlugin/ExecuteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CursorPluginServer is the server API for CursorPlugin service.
type CursorPluginServer interface {
	// GetPipelineInfo is called upon compiling the pipeline so that cursor knowns
	// the structure of the tasks.
	// It returns a task info map that describes what kind of tasks exist in the
	// pipeline and a default task used as the root task
	GetPipelineInfo(context.Context, *GetPipelineInfoRequest) (*GetPipelineInfoResponse, error)
	ExecuteTask(context.Context, *ExecuteTaskRequest) (*ExecuteTaskResponse, error)
}

func RegisterCursorPluginServer(s *grpc.Server, srv CursorPluginServer) {
	s.RegisterService(&_CursorPlugin_serviceDesc, srv)
}

func _CursorPlugin_GetPipelineInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPipelineInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CursorPluginServer).GetPipelineInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CursorPlugin/GetPipelineInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CursorPluginServer).GetPipelineInfo(ctx, req.(*GetPipelineInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CursorPlugin_ExecuteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CursorPluginServer).ExecuteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CursorPlugin/ExecuteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CursorPluginServer).ExecuteTask(ctx, req.(*ExecuteTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CursorPlugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CursorPlugin",
	HandlerType: (*CursorPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPipelineInfo",
			Handler:    _CursorPlugin_GetPipelineInfo_Handler,
		},
		{
			MethodName: "ExecuteTask",
			Handler:    _CursorPlugin_ExecuteTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugin/proto/cursorPlugin.proto",
}
