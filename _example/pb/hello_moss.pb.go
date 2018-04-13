// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello_moss.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	hello_moss.proto

It has these top-level messages:
	HelloMoss
	RegisterReq
	RegisterRes
*/
package pb

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

type HelloMoss struct {
	HelloMoss string `protobuf:"bytes,1,opt,name=hello_moss,json=helloMoss" json:"hello_moss,omitempty"`
}

func (m *HelloMoss) Reset()                    { *m = HelloMoss{} }
func (m *HelloMoss) String() string            { return proto.CompactTextString(m) }
func (*HelloMoss) ProtoMessage()               {}
func (*HelloMoss) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloMoss) GetHelloMoss() string {
	if m != nil {
		return m.HelloMoss
	}
	return ""
}

// ===================MOSS REGISTER====================
// user register
type RegisterReq struct {
	UserName     string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	UserPhone    string `protobuf:"bytes,2,opt,name=user_phone,json=userPhone" json:"user_phone,omitempty"`
	UserPassword string `protobuf:"bytes,3,opt,name=user_password,json=userPassword" json:"user_password,omitempty"`
}

func (m *RegisterReq) Reset()                    { *m = RegisterReq{} }
func (m *RegisterReq) String() string            { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()               {}
func (*RegisterReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterReq) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *RegisterReq) GetUserPhone() string {
	if m != nil {
		return m.UserPhone
	}
	return ""
}

func (m *RegisterReq) GetUserPassword() string {
	if m != nil {
		return m.UserPassword
	}
	return ""
}

// return in http.Header "Authorization"= "Bearer "+signedKey
type RegisterRes struct {
	UserName  string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	UserPhone string `protobuf:"bytes,2,opt,name=user_phone,json=userPhone" json:"user_phone,omitempty"`
	UserId    string `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Audience  string `protobuf:"bytes,4,opt,name=audience" json:"audience,omitempty"`
}

func (m *RegisterRes) Reset()                    { *m = RegisterRes{} }
func (m *RegisterRes) String() string            { return proto.CompactTextString(m) }
func (*RegisterRes) ProtoMessage()               {}
func (*RegisterRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RegisterRes) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *RegisterRes) GetUserPhone() string {
	if m != nil {
		return m.UserPhone
	}
	return ""
}

func (m *RegisterRes) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *RegisterRes) GetAudience() string {
	if m != nil {
		return m.Audience
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloMoss)(nil), "pb.HelloMoss")
	proto.RegisterType((*RegisterReq)(nil), "pb.RegisterReq")
	proto.RegisterType((*RegisterRes)(nil), "pb.RegisterRes")
}

func init() { proto.RegisterFile("hello_moss.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x8f, 0xcf, 0xcd, 0x2f, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48,
	0x52, 0xd2, 0xe2, 0xe2, 0xf4, 0x00, 0x89, 0xfb, 0xe6, 0x17, 0x17, 0x0b, 0xc9, 0x72, 0x71, 0x21,
	0x14, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x71, 0x66, 0xc0, 0xa4, 0x95, 0xf2, 0xb8, 0xb8,
	0x83, 0x52, 0xd3, 0x33, 0x8b, 0x4b, 0x52, 0x8b, 0x82, 0x52, 0x0b, 0x85, 0xa4, 0xb9, 0x38, 0x4b,
	0x8b, 0x53, 0x8b, 0xe2, 0xf3, 0x12, 0x73, 0x53, 0xa1, 0x8a, 0x39, 0x40, 0x02, 0x7e, 0x89, 0xb9,
	0xa9, 0x20, 0xa3, 0xc0, 0x92, 0x05, 0x19, 0xf9, 0x79, 0xa9, 0x12, 0x4c, 0x10, 0xa3, 0x40, 0x22,
	0x01, 0x20, 0x01, 0x21, 0x65, 0x2e, 0x5e, 0x88, 0x74, 0x62, 0x71, 0x71, 0x79, 0x7e, 0x51, 0x8a,
	0x04, 0x33, 0x58, 0x05, 0x0f, 0x58, 0x05, 0x54, 0x4c, 0xa9, 0x0e, 0xd9, 0xbe, 0x62, 0x8a, 0xec,
	0x13, 0xe7, 0x62, 0x07, 0x4b, 0x67, 0xc2, 0x6c, 0x62, 0x03, 0x71, 0x3d, 0x53, 0x84, 0xa4, 0xb8,
	0x38, 0x12, 0x4b, 0x53, 0x32, 0x53, 0xf3, 0x92, 0x53, 0x25, 0x58, 0x20, 0x66, 0xc2, 0xf8, 0x49,
	0x6c, 0xe0, 0x60, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x02, 0x9e, 0xed, 0x3a, 0x01,
	0x00, 0x00,
}
