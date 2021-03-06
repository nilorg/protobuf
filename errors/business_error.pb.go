// Code generated by protoc-gen-go. DO NOT EDIT.
// source: errors/business_error.proto

package errors

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// BusinessError 业务错误
type BusinessError struct {
	// 错误代码
	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 错误消息
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BusinessError) Reset()         { *m = BusinessError{} }
func (m *BusinessError) String() string { return proto.CompactTextString(m) }
func (*BusinessError) ProtoMessage()    {}
func (*BusinessError) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e369708ce400e60, []int{0}
}

func (m *BusinessError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BusinessError.Unmarshal(m, b)
}
func (m *BusinessError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BusinessError.Marshal(b, m, deterministic)
}
func (m *BusinessError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BusinessError.Merge(m, src)
}
func (m *BusinessError) XXX_Size() int {
	return xxx_messageInfo_BusinessError.Size(m)
}
func (m *BusinessError) XXX_DiscardUnknown() {
	xxx_messageInfo_BusinessError.DiscardUnknown(m)
}

var xxx_messageInfo_BusinessError proto.InternalMessageInfo

func (m *BusinessError) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *BusinessError) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*BusinessError)(nil), "nilorg.protobuf.errors.BusinessError")
}

func init() { proto.RegisterFile("errors/business_error.proto", fileDescriptor_2e369708ce400e60) }

var fileDescriptor_2e369708ce400e60 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0x2d, 0x2a, 0xca,
	0x2f, 0x2a, 0xd6, 0x4f, 0x2a, 0x2d, 0xce, 0xcc, 0x4b, 0x2d, 0x2e, 0x8e, 0x07, 0xf3, 0xf5, 0x0a,
	0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xc4, 0xf2, 0x32, 0x73, 0xf2, 0x8b, 0xd2, 0x21, 0xbc, 0xa4, 0xd2,
	0x34, 0x3d, 0x88, 0x62, 0x25, 0x53, 0x2e, 0x5e, 0x27, 0xa8, 0x7a, 0x57, 0x90, 0x88, 0x90, 0x10,
	0x17, 0x4b, 0x72, 0x7e, 0x4a, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x6f, 0x10, 0x98, 0x2d, 0x24,
	0xc0, 0xc5, 0x9c, 0x5b, 0x9c, 0x2e, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x62, 0x3a, 0x65,
	0x72, 0x49, 0x25, 0xe7, 0xe7, 0xea, 0x61, 0x37, 0xd4, 0x09, 0xd5, 0xc8, 0x00, 0xc6, 0x28, 0xc5,
	0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x88, 0x1e, 0x7d, 0x98, 0x1e,
	0x7d, 0x88, 0x9e, 0x55, 0x4c, 0x62, 0x7e, 0x10, 0xc3, 0x02, 0x60, 0x86, 0x81, 0x35, 0x17, 0x27,
	0xb1, 0x81, 0x55, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x76, 0xc0, 0x69, 0xdf, 0x00,
	0x00, 0x00,
}
