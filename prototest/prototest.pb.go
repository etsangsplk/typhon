// Code generated by protoc-gen-go. DO NOT EDIT.
// source: prototest.proto

package prototest

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

type Greeting struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Priority             int32    `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_d68ecc031ae9d98e, []int{0}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Greeting) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func init() {
	proto.RegisterType((*Greeting)(nil), "Greeting")
}

func init() { proto.RegisterFile("prototest.proto", fileDescriptor_d68ecc031ae9d98e) }

var fileDescriptor_d68ecc031ae9d98e = []byte{
	// 101 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0x49, 0x2d, 0x2e, 0xd1, 0x03, 0xb3, 0x94, 0x1c, 0xb8, 0x38, 0xdc, 0x8b, 0x52, 0x53,
	0x4b, 0x32, 0xf3, 0xd2, 0x85, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x25,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x21, 0x29, 0x2e, 0x8e, 0x82, 0xa2, 0xcc, 0xfc,
	0xa2, 0xcc, 0x92, 0x4a, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x38, 0xdf, 0x89, 0x3b, 0x8a,
	0x13, 0x6e, 0x68, 0x12, 0x1b, 0x98, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x88, 0xd0, 0x6a,
	0xc9, 0x68, 0x00, 0x00, 0x00,
}
