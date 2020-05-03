// Code generated by protoc-gen-go. DO NOT EDIT.
// source: job/job.proto

package job

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	errors "github.com/nilorg/protobuf/errors"
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

// ExecuteRequest 执行请求
type ExecuteRequest struct {
	// 任务ID
	TaskId               string   `protobuf:"bytes,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	Body                 *any.Any `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteRequest) Reset()         { *m = ExecuteRequest{} }
func (m *ExecuteRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteRequest) ProtoMessage()    {}
func (*ExecuteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f97a06ee6e4937, []int{0}
}

func (m *ExecuteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteRequest.Unmarshal(m, b)
}
func (m *ExecuteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteRequest.Marshal(b, m, deterministic)
}
func (m *ExecuteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteRequest.Merge(m, src)
}
func (m *ExecuteRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteRequest.Size(m)
}
func (m *ExecuteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteRequest proto.InternalMessageInfo

func (m *ExecuteRequest) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *ExecuteRequest) GetBody() *any.Any {
	if m != nil {
		return m.Body
	}
	return nil
}

// ExecuteResponse 执行请求响应
type ExecuteResponse struct {
	Body                 *any.Any              `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	Err                  *errors.BusinessError `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ExecuteResponse) Reset()         { *m = ExecuteResponse{} }
func (m *ExecuteResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteResponse) ProtoMessage()    {}
func (*ExecuteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c3f97a06ee6e4937, []int{1}
}

func (m *ExecuteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteResponse.Unmarshal(m, b)
}
func (m *ExecuteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteResponse.Marshal(b, m, deterministic)
}
func (m *ExecuteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteResponse.Merge(m, src)
}
func (m *ExecuteResponse) XXX_Size() int {
	return xxx_messageInfo_ExecuteResponse.Size(m)
}
func (m *ExecuteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteResponse proto.InternalMessageInfo

func (m *ExecuteResponse) GetBody() *any.Any {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *ExecuteResponse) GetErr() *errors.BusinessError {
	if m != nil {
		return m.Err
	}
	return nil
}

func init() {
	proto.RegisterType((*ExecuteRequest)(nil), "nilorg.protobuf.job.ExecuteRequest")
	proto.RegisterType((*ExecuteResponse)(nil), "nilorg.protobuf.job.ExecuteResponse")
}

func init() { proto.RegisterFile("job/job.proto", fileDescriptor_c3f97a06ee6e4937) }

var fileDescriptor_c3f97a06ee6e4937 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0xe9, 0x94, 0x89, 0x99, 0x7f, 0x20, 0x13, 0x99, 0x7d, 0x90, 0x31, 0x15, 0x0a, 0x42,
	0x0a, 0x13, 0xf4, 0x79, 0x83, 0x3d, 0xe8, 0x83, 0x48, 0xdd, 0x8b, 0xbe, 0x48, 0xd3, 0x5d, 0x6b,
	0xbb, 0x99, 0x3b, 0x73, 0x53, 0xb1, 0x5f, 0xc9, 0x0f, 0xe1, 0x67, 0x93, 0x35, 0xd9, 0x86, 0x43,
	0xf4, 0x65, 0x8f, 0xe7, 0xe6, 0x77, 0x4f, 0xce, 0x49, 0xd8, 0x6e, 0x8e, 0x32, 0xcc, 0x51, 0x8a,
	0xa9, 0x46, 0x83, 0xbc, 0xa9, 0xb2, 0x09, 0xea, 0xd4, 0x2a, 0x59, 0x3c, 0x8b, 0x1c, 0xa5, 0x7f,
	0x94, 0x22, 0xa6, 0x13, 0x08, 0xe7, 0xc3, 0x30, 0x56, 0xa5, 0x25, 0xfc, 0x73, 0xcb, 0x2f, 0x8f,
	0x40, 0x6b, 0xd4, 0x14, 0xca, 0x82, 0x32, 0x05, 0x44, 0x4f, 0x95, 0xb6, 0x70, 0x27, 0x62, 0x7b,
	0x83, 0x0f, 0x48, 0x0a, 0x03, 0x11, 0xbc, 0x15, 0x40, 0x86, 0x1f, 0xb2, 0xba, 0x89, 0x69, 0x7c,
	0x3d, 0x6a, 0x79, 0x6d, 0x2f, 0xd8, 0x8e, 0x9c, 0xe2, 0x01, 0xdb, 0x94, 0x38, 0x2a, 0x5b, 0xb5,
	0xb6, 0x17, 0x34, 0xba, 0x07, 0xc2, 0x06, 0x58, 0xa6, 0xea, 0xa9, 0x32, 0xaa, 0x88, 0x8e, 0x61,
	0xfb, 0x0b, 0x4f, 0x9a, 0xa2, 0x22, 0x58, 0x2c, 0x7b, 0xff, 0x2d, 0xf3, 0x2b, 0xb6, 0x01, 0x5a,
	0xbb, 0x5b, 0xce, 0xc4, 0x6a, 0x77, 0xdb, 0x45, 0xf4, 0x5d, 0x97, 0xc1, 0x4c, 0x46, 0xb3, 0x8d,
	0xee, 0x97, 0xc7, 0x1a, 0xc3, 0x98, 0xc6, 0xf7, 0xa0, 0xdf, 0xb3, 0x04, 0xf8, 0x90, 0x6d, 0xb9,
	0x14, 0xfc, 0x44, 0xfc, 0xf2, 0x84, 0xe2, 0x67, 0x6f, 0xff, 0xf4, 0x6f, 0xc8, 0x15, 0x79, 0x60,
	0x3b, 0x6e, 0xd4, 0xa3, 0x52, 0x25, 0x6b, 0xb4, 0xee, 0x5f, 0x3e, 0x1e, 0xa7, 0x99, 0x79, 0x29,
	0xa4, 0x48, 0xf0, 0x35, 0x5c, 0xfd, 0xc4, 0x1c, 0xe5, 0x67, 0xad, 0x79, 0x6b, 0x7d, 0xee, 0xe6,
	0x3e, 0x37, 0x28, 0x65, 0xbd, 0x62, 0x2e, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x39, 0x88,
	0x61, 0x37, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TaskServiceClient is the client API for TaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskServiceClient interface {
	// Execute 执行
	Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error)
	// ExecuteAsync 执行异步
	ExecuteAsync(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error)
}

type taskServiceClient struct {
	cc *grpc.ClientConn
}

func NewTaskServiceClient(cc *grpc.ClientConn) TaskServiceClient {
	return &taskServiceClient{cc}
}

func (c *taskServiceClient) Execute(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error) {
	out := new(ExecuteResponse)
	err := c.cc.Invoke(ctx, "/nilorg.protobuf.job.TaskService/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskServiceClient) ExecuteAsync(ctx context.Context, in *ExecuteRequest, opts ...grpc.CallOption) (*ExecuteResponse, error) {
	out := new(ExecuteResponse)
	err := c.cc.Invoke(ctx, "/nilorg.protobuf.job.TaskService/ExecuteAsync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskServiceServer is the server API for TaskService service.
type TaskServiceServer interface {
	// Execute 执行
	Execute(context.Context, *ExecuteRequest) (*ExecuteResponse, error)
	// ExecuteAsync 执行异步
	ExecuteAsync(context.Context, *ExecuteRequest) (*ExecuteResponse, error)
}

func RegisterTaskServiceServer(s *grpc.Server, srv TaskServiceServer) {
	s.RegisterService(&_TaskService_serviceDesc, srv)
}

func _TaskService_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nilorg.protobuf.job.TaskService/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).Execute(ctx, req.(*ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskService_ExecuteAsync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServiceServer).ExecuteAsync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nilorg.protobuf.job.TaskService/ExecuteAsync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServiceServer).ExecuteAsync(ctx, req.(*ExecuteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nilorg.protobuf.job.TaskService",
	HandlerType: (*TaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _TaskService_Execute_Handler,
		},
		{
			MethodName: "ExecuteAsync",
			Handler:    _TaskService_ExecuteAsync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job/job.proto",
}