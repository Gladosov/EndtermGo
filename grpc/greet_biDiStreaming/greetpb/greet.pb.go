// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greet.proto

package greetpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{0}
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

func (m *Greeting) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Greeting) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type GreetRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{1}
}

func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetResponse) Reset()         { *m = GreetResponse{} }
func (m *GreetResponse) String() string { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()    {}
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{2}
}

func (m *GreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetResponse.Unmarshal(m, b)
}
func (m *GreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetResponse.Marshal(b, m, deterministic)
}
func (m *GreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResponse.Merge(m, src)
}
func (m *GreetResponse) XXX_Size() int {
	return xxx_messageInfo_GreetResponse.Size(m)
}
func (m *GreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResponse proto.InternalMessageInfo

func (m *GreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type GreetManytimesRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetManytimesRequest) Reset()         { *m = GreetManytimesRequest{} }
func (m *GreetManytimesRequest) String() string { return proto.CompactTextString(m) }
func (*GreetManytimesRequest) ProtoMessage()    {}
func (*GreetManytimesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{3}
}

func (m *GreetManytimesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetManytimesRequest.Unmarshal(m, b)
}
func (m *GreetManytimesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetManytimesRequest.Marshal(b, m, deterministic)
}
func (m *GreetManytimesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetManytimesRequest.Merge(m, src)
}
func (m *GreetManytimesRequest) XXX_Size() int {
	return xxx_messageInfo_GreetManytimesRequest.Size(m)
}
func (m *GreetManytimesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetManytimesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetManytimesRequest proto.InternalMessageInfo

func (m *GreetManytimesRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetManytimesResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetManytimesResponse) Reset()         { *m = GreetManytimesResponse{} }
func (m *GreetManytimesResponse) String() string { return proto.CompactTextString(m) }
func (*GreetManytimesResponse) ProtoMessage()    {}
func (*GreetManytimesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{4}
}

func (m *GreetManytimesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetManytimesResponse.Unmarshal(m, b)
}
func (m *GreetManytimesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetManytimesResponse.Marshal(b, m, deterministic)
}
func (m *GreetManytimesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetManytimesResponse.Merge(m, src)
}
func (m *GreetManytimesResponse) XXX_Size() int {
	return xxx_messageInfo_GreetManytimesResponse.Size(m)
}
func (m *GreetManytimesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetManytimesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetManytimesResponse proto.InternalMessageInfo

func (m *GreetManytimesResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type LongGreetRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LongGreetRequest) Reset()         { *m = LongGreetRequest{} }
func (m *LongGreetRequest) String() string { return proto.CompactTextString(m) }
func (*LongGreetRequest) ProtoMessage()    {}
func (*LongGreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{5}
}

func (m *LongGreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongGreetRequest.Unmarshal(m, b)
}
func (m *LongGreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongGreetRequest.Marshal(b, m, deterministic)
}
func (m *LongGreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongGreetRequest.Merge(m, src)
}
func (m *LongGreetRequest) XXX_Size() int {
	return xxx_messageInfo_LongGreetRequest.Size(m)
}
func (m *LongGreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LongGreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LongGreetRequest proto.InternalMessageInfo

func (m *LongGreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type LongGreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LongGreetResponse) Reset()         { *m = LongGreetResponse{} }
func (m *LongGreetResponse) String() string { return proto.CompactTextString(m) }
func (*LongGreetResponse) ProtoMessage()    {}
func (*LongGreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{6}
}

func (m *LongGreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongGreetResponse.Unmarshal(m, b)
}
func (m *LongGreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongGreetResponse.Marshal(b, m, deterministic)
}
func (m *LongGreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongGreetResponse.Merge(m, src)
}
func (m *LongGreetResponse) XXX_Size() int {
	return xxx_messageInfo_LongGreetResponse.Size(m)
}
func (m *LongGreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LongGreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LongGreetResponse proto.InternalMessageInfo

func (m *LongGreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type GreetEveryoneRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetEveryoneRequest) Reset()         { *m = GreetEveryoneRequest{} }
func (m *GreetEveryoneRequest) String() string { return proto.CompactTextString(m) }
func (*GreetEveryoneRequest) ProtoMessage()    {}
func (*GreetEveryoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{7}
}

func (m *GreetEveryoneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetEveryoneRequest.Unmarshal(m, b)
}
func (m *GreetEveryoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetEveryoneRequest.Marshal(b, m, deterministic)
}
func (m *GreetEveryoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetEveryoneRequest.Merge(m, src)
}
func (m *GreetEveryoneRequest) XXX_Size() int {
	return xxx_messageInfo_GreetEveryoneRequest.Size(m)
}
func (m *GreetEveryoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetEveryoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetEveryoneRequest proto.InternalMessageInfo

func (m *GreetEveryoneRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetEveryoneResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetEveryoneResponse) Reset()         { *m = GreetEveryoneResponse{} }
func (m *GreetEveryoneResponse) String() string { return proto.CompactTextString(m) }
func (*GreetEveryoneResponse) ProtoMessage()    {}
func (*GreetEveryoneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{8}
}

func (m *GreetEveryoneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetEveryoneResponse.Unmarshal(m, b)
}
func (m *GreetEveryoneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetEveryoneResponse.Marshal(b, m, deterministic)
}
func (m *GreetEveryoneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetEveryoneResponse.Merge(m, src)
}
func (m *GreetEveryoneResponse) XXX_Size() int {
	return xxx_messageInfo_GreetEveryoneResponse.Size(m)
}
func (m *GreetEveryoneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetEveryoneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetEveryoneResponse proto.InternalMessageInfo

func (m *GreetEveryoneResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type GreetWithDeadlineRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetWithDeadlineRequest) Reset()         { *m = GreetWithDeadlineRequest{} }
func (m *GreetWithDeadlineRequest) String() string { return proto.CompactTextString(m) }
func (*GreetWithDeadlineRequest) ProtoMessage()    {}
func (*GreetWithDeadlineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{9}
}

func (m *GreetWithDeadlineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetWithDeadlineRequest.Unmarshal(m, b)
}
func (m *GreetWithDeadlineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetWithDeadlineRequest.Marshal(b, m, deterministic)
}
func (m *GreetWithDeadlineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetWithDeadlineRequest.Merge(m, src)
}
func (m *GreetWithDeadlineRequest) XXX_Size() int {
	return xxx_messageInfo_GreetWithDeadlineRequest.Size(m)
}
func (m *GreetWithDeadlineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetWithDeadlineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetWithDeadlineRequest proto.InternalMessageInfo

func (m *GreetWithDeadlineRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetWithDeadlineResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetWithDeadlineResponse) Reset()         { *m = GreetWithDeadlineResponse{} }
func (m *GreetWithDeadlineResponse) String() string { return proto.CompactTextString(m) }
func (*GreetWithDeadlineResponse) ProtoMessage()    {}
func (*GreetWithDeadlineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_32c0044392f32579, []int{10}
}

func (m *GreetWithDeadlineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetWithDeadlineResponse.Unmarshal(m, b)
}
func (m *GreetWithDeadlineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetWithDeadlineResponse.Marshal(b, m, deterministic)
}
func (m *GreetWithDeadlineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetWithDeadlineResponse.Merge(m, src)
}
func (m *GreetWithDeadlineResponse) XXX_Size() int {
	return xxx_messageInfo_GreetWithDeadlineResponse.Size(m)
}
func (m *GreetWithDeadlineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetWithDeadlineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetWithDeadlineResponse proto.InternalMessageInfo

func (m *GreetWithDeadlineResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "greet.Greeting")
	proto.RegisterType((*GreetRequest)(nil), "greet.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "greet.GreetResponse")
	proto.RegisterType((*GreetManytimesRequest)(nil), "greet.GreetManytimesRequest")
	proto.RegisterType((*GreetManytimesResponse)(nil), "greet.GreetManytimesResponse")
	proto.RegisterType((*LongGreetRequest)(nil), "greet.LongGreetRequest")
	proto.RegisterType((*LongGreetResponse)(nil), "greet.LongGreetResponse")
	proto.RegisterType((*GreetEveryoneRequest)(nil), "greet.GreetEveryoneRequest")
	proto.RegisterType((*GreetEveryoneResponse)(nil), "greet.GreetEveryoneResponse")
	proto.RegisterType((*GreetWithDeadlineRequest)(nil), "greet.GreetWithDeadlineRequest")
	proto.RegisterType((*GreetWithDeadlineResponse)(nil), "greet.GreetWithDeadlineResponse")
}

func init() { proto.RegisterFile("greet.proto", fileDescriptor_32c0044392f32579) }

var fileDescriptor_32c0044392f32579 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x4f, 0x4f, 0xc2, 0x40,
	0x10, 0xc5, 0xad, 0x06, 0xa4, 0x83, 0xff, 0x18, 0x11, 0xb1, 0x40, 0x24, 0xbd, 0x48, 0x42, 0x82,
	0x04, 0xbc, 0x79, 0x30, 0x41, 0x94, 0x8b, 0x1a, 0x45, 0x13, 0x8d, 0x17, 0x53, 0x74, 0xac, 0x4d,
	0xa0, 0xc5, 0x76, 0x21, 0xe1, 0x13, 0xf8, 0xb5, 0x4d, 0x97, 0x5d, 0xa8, 0x05, 0x6c, 0xd2, 0x1b,
	0xbb, 0xef, 0xed, 0xef, 0x2d, 0xfb, 0x26, 0x85, 0xb4, 0xe9, 0x12, 0xb1, 0xda, 0xd0, 0x75, 0x98,
	0x83, 0x09, 0xbe, 0xd0, 0xaf, 0x21, 0xd5, 0xf1, 0x7f, 0x58, 0xb6, 0x89, 0x25, 0x80, 0x4f, 0xcb,
	0xf5, 0xd8, 0x9b, 0x6d, 0x0c, 0x28, 0xaf, 0x94, 0x95, 0x8a, 0xda, 0x55, 0xf9, 0xce, 0x9d, 0x31,
	0x20, 0x2c, 0x80, 0xda, 0x37, 0xa4, 0xba, 0xce, 0xd5, 0x94, 0xbf, 0xe1, 0x8b, 0xfa, 0x39, 0x6c,
	0x71, 0x4e, 0x97, 0xbe, 0x47, 0xe4, 0x31, 0xac, 0x42, 0xca, 0x14, 0x5c, 0x4e, 0x4a, 0x37, 0x76,
	0x6b, 0xd3, 0x78, 0x19, 0xd7, 0x9d, 0x19, 0xf4, 0x13, 0xd8, 0x16, 0x87, 0xbd, 0xa1, 0x63, 0x7b,
	0x84, 0x39, 0x48, 0xba, 0xe4, 0x8d, 0xfa, 0x4c, 0xdc, 0x42, 0xac, 0xf4, 0x36, 0x1c, 0x70, 0xe3,
	0xad, 0x61, 0x4f, 0x98, 0x35, 0x20, 0x2f, 0x56, 0x5c, 0x1d, 0x72, 0x61, 0x4a, 0x44, 0xee, 0x05,
	0xec, 0xdd, 0x38, 0xb6, 0x19, 0xff, 0x1f, 0x56, 0x21, 0x13, 0x00, 0x44, 0xa4, 0x5d, 0x42, 0x96,
	0x1b, 0xaf, 0xc6, 0xe4, 0x4e, 0x1c, 0x9b, 0x62, 0x25, 0x9e, 0x8a, 0xa7, 0x9a, 0x43, 0x22, 0x52,
	0x3b, 0x90, 0xe7, 0x07, 0x9e, 0x2d, 0xf6, 0xd5, 0x26, 0xe3, 0xa3, 0x6f, 0xc5, 0x4c, 0x6e, 0xc2,
	0xd1, 0x12, 0xd0, 0xff, 0xe9, 0x8d, 0x9f, 0x0d, 0x31, 0x40, 0x8f, 0xe4, 0x8e, 0xad, 0x77, 0xc2,
	0x33, 0x48, 0xf0, 0x35, 0xee, 0x07, 0x93, 0xc4, 0x85, 0xb4, 0xec, 0xdf, 0xcd, 0x29, 0x5c, 0x5f,
	0xc3, 0x07, 0xd8, 0x99, 0x55, 0xfb, 0xe4, 0x57, 0x8b, 0xc5, 0xa0, 0x33, 0x3c, 0x37, 0x5a, 0x69,
	0x85, 0x2a, 0x81, 0x75, 0x05, 0x5b, 0xa0, 0xce, 0xaa, 0xc3, 0x43, 0xe1, 0x0f, 0x4f, 0x83, 0x96,
	0x5f, 0x14, 0x24, 0xa3, 0xa2, 0xe0, 0xbd, 0x18, 0x70, 0x59, 0x06, 0x16, 0x82, 0xb9, 0xa1, 0x9e,
	0xb5, 0xe2, 0x72, 0x71, 0xce, 0xab, 0x2b, 0xf8, 0x02, 0x99, 0x85, 0x47, 0xc6, 0xe3, 0xe0, 0xc1,
	0x25, 0x3d, 0x6a, 0xe5, 0xd5, 0x06, 0x49, 0x6f, 0xa9, 0xaf, 0x9b, 0xdc, 0x34, 0xec, 0xf5, 0x92,
	0xfc, 0x53, 0xd1, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x87, 0x50, 0xf6, 0x69, 0x39, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	// unary
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
	//server streaming
	GreetManyTimes(ctx context.Context, in *GreetManytimesRequest, opts ...grpc.CallOption) (GreetService_GreetManyTimesClient, error)
	// Client Streaming
	LongGreet(ctx context.Context, opts ...grpc.CallOption) (GreetService_LongGreetClient, error)
	//BiDi Streaming
	GreetEveryone(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetEveryoneClient, error)
	GreetWithDeadline(ctx context.Context, in *GreetWithDeadlineRequest, opts ...grpc.CallOption) (*GreetWithDeadlineResponse, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) GreetManyTimes(ctx context.Context, in *GreetManytimesRequest, opts ...grpc.CallOption) (GreetService_GreetManyTimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[0], "/greet.GreetService/GreetManyTimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetManyTimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_GreetManyTimesClient interface {
	Recv() (*GreetManytimesResponse, error)
	grpc.ClientStream
}

type greetServiceGreetManyTimesClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetManyTimesClient) Recv() (*GreetManytimesResponse, error) {
	m := new(GreetManytimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) LongGreet(ctx context.Context, opts ...grpc.CallOption) (GreetService_LongGreetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[1], "/greet.GreetService/LongGreet", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceLongGreetClient{stream}
	return x, nil
}

type GreetService_LongGreetClient interface {
	Send(*LongGreetRequest) error
	CloseAndRecv() (*LongGreetResponse, error)
	grpc.ClientStream
}

type greetServiceLongGreetClient struct {
	grpc.ClientStream
}

func (x *greetServiceLongGreetClient) Send(m *LongGreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceLongGreetClient) CloseAndRecv() (*LongGreetResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(LongGreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) GreetEveryone(ctx context.Context, opts ...grpc.CallOption) (GreetService_GreetEveryoneClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GreetService_serviceDesc.Streams[2], "/greet.GreetService/GreetEveryone", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetEveryoneClient{stream}
	return x, nil
}

type GreetService_GreetEveryoneClient interface {
	Send(*GreetEveryoneRequest) error
	Recv() (*GreetEveryoneResponse, error)
	grpc.ClientStream
}

type greetServiceGreetEveryoneClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetEveryoneClient) Send(m *GreetEveryoneRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceGreetEveryoneClient) Recv() (*GreetEveryoneResponse, error) {
	m := new(GreetEveryoneResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) GreetWithDeadline(ctx context.Context, in *GreetWithDeadlineRequest, opts ...grpc.CallOption) (*GreetWithDeadlineResponse, error) {
	out := new(GreetWithDeadlineResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/GreetWithDeadline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	// unary
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
	//server streaming
	GreetManyTimes(*GreetManytimesRequest, GreetService_GreetManyTimesServer) error
	// Client Streaming
	LongGreet(GreetService_LongGreetServer) error
	//BiDi Streaming
	GreetEveryone(GreetService_GreetEveryoneServer) error
	GreetWithDeadline(context.Context, *GreetWithDeadlineRequest) (*GreetWithDeadlineResponse, error)
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) Greet(ctx context.Context, req *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (*UnimplementedGreetServiceServer) GreetManyTimes(req *GreetManytimesRequest, srv GreetService_GreetManyTimesServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetManyTimes not implemented")
}
func (*UnimplementedGreetServiceServer) LongGreet(srv GreetService_LongGreetServer) error {
	return status.Errorf(codes.Unimplemented, "method LongGreet not implemented")
}
func (*UnimplementedGreetServiceServer) GreetEveryone(srv GreetService_GreetEveryoneServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetEveryone not implemented")
}
func (*UnimplementedGreetServiceServer) GreetWithDeadline(ctx context.Context, req *GreetWithDeadlineRequest) (*GreetWithDeadlineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GreetWithDeadline not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_GreetManyTimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetManytimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).GreetManyTimes(m, &greetServiceGreetManyTimesServer{stream})
}

type GreetService_GreetManyTimesServer interface {
	Send(*GreetManytimesResponse) error
	grpc.ServerStream
}

type greetServiceGreetManyTimesServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetManyTimesServer) Send(m *GreetManytimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_LongGreet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).LongGreet(&greetServiceLongGreetServer{stream})
}

type GreetService_LongGreetServer interface {
	SendAndClose(*LongGreetResponse) error
	Recv() (*LongGreetRequest, error)
	grpc.ServerStream
}

type greetServiceLongGreetServer struct {
	grpc.ServerStream
}

func (x *greetServiceLongGreetServer) SendAndClose(m *LongGreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceLongGreetServer) Recv() (*LongGreetRequest, error) {
	m := new(LongGreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_GreetEveryone_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).GreetEveryone(&greetServiceGreetEveryoneServer{stream})
}

type GreetService_GreetEveryoneServer interface {
	Send(*GreetEveryoneResponse) error
	Recv() (*GreetEveryoneRequest, error)
	grpc.ServerStream
}

type greetServiceGreetEveryoneServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetEveryoneServer) Send(m *GreetEveryoneResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceGreetEveryoneServer) Recv() (*GreetEveryoneRequest, error) {
	m := new(GreetEveryoneRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_GreetWithDeadline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetWithDeadlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).GreetWithDeadline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/GreetWithDeadline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).GreetWithDeadline(ctx, req.(*GreetWithDeadlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
		{
			MethodName: "GreetWithDeadline",
			Handler:    _GreetService_GreetWithDeadline_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetManyTimes",
			Handler:       _GreetService_GreetManyTimes_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LongGreet",
			Handler:       _GreetService_LongGreet_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GreetEveryone",
			Handler:       _GreetService_GreetEveryone_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "greet.proto",
}
