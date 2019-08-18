// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: job/job.proto

package job

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/nilorg/protobuf/errors"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for TaskService service

type TaskService interface {
	// Execute 执行
	Execute(ctx context.Context, in *ExecuteRequest, opts ...client.CallOption) (*ExecuteResponse, error)
	// ExecuteAsync 执行异步
	ExecuteAsync(ctx context.Context, in *ExecuteRequest, opts ...client.CallOption) (*ExecuteResponse, error)
}

type taskService struct {
	c    client.Client
	name string
}

func NewTaskService(name string, c client.Client) TaskService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "nilorg.protobuf.job"
	}
	return &taskService{
		c:    c,
		name: name,
	}
}

func (c *taskService) Execute(ctx context.Context, in *ExecuteRequest, opts ...client.CallOption) (*ExecuteResponse, error) {
	req := c.c.NewRequest(c.name, "TaskService.Execute", in)
	out := new(ExecuteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskService) ExecuteAsync(ctx context.Context, in *ExecuteRequest, opts ...client.CallOption) (*ExecuteResponse, error) {
	req := c.c.NewRequest(c.name, "TaskService.ExecuteAsync", in)
	out := new(ExecuteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskService service

type TaskServiceHandler interface {
	// Execute 执行
	Execute(context.Context, *ExecuteRequest, *ExecuteResponse) error
	// ExecuteAsync 执行异步
	ExecuteAsync(context.Context, *ExecuteRequest, *ExecuteResponse) error
}

func RegisterTaskServiceHandler(s server.Server, hdlr TaskServiceHandler, opts ...server.HandlerOption) error {
	type taskService interface {
		Execute(ctx context.Context, in *ExecuteRequest, out *ExecuteResponse) error
		ExecuteAsync(ctx context.Context, in *ExecuteRequest, out *ExecuteResponse) error
	}
	type TaskService struct {
		taskService
	}
	h := &taskServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&TaskService{h}, opts...))
}

type taskServiceHandler struct {
	TaskServiceHandler
}

func (h *taskServiceHandler) Execute(ctx context.Context, in *ExecuteRequest, out *ExecuteResponse) error {
	return h.TaskServiceHandler.Execute(ctx, in, out)
}

func (h *taskServiceHandler) ExecuteAsync(ctx context.Context, in *ExecuteRequest, out *ExecuteResponse) error {
	return h.TaskServiceHandler.ExecuteAsync(ctx, in, out)
}
