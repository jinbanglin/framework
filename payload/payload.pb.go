// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: payload.proto

/*
	Package payload is a generated protocol buffer package.

	It is generated from these files:
		payload.proto

	It has these top-level messages:
		MossPacket
		MossMessage
*/
package payload

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/lyft/protoc-gen-validate/validate"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type MossPacket struct {
	ServiceCode uint32       `protobuf:"varint,1,opt,name=service_code,json=serviceCode,proto3" json:"service_code,omitempty"`
	Payload     []byte       `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	MossMessage *MossMessage `protobuf:"bytes,3,opt,name=moss_message,json=mossMessage" json:"moss_message,omitempty"`
	UserId      string       `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ClientIp    string       `protobuf:"bytes,5,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
}

func (m *MossPacket) Reset()                    { *m = MossPacket{} }
func (m *MossPacket) String() string            { return proto.CompactTextString(m) }
func (*MossPacket) ProtoMessage()               {}
func (*MossPacket) Descriptor() ([]byte, []int) { return fileDescriptorPayload, []int{0} }

func (m *MossPacket) GetServiceCode() uint32 {
	if m != nil {
		return m.ServiceCode
	}
	return 0
}

func (m *MossPacket) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *MossPacket) GetMossMessage() *MossMessage {
	if m != nil {
		return m.MossMessage
	}
	return nil
}

func (m *MossPacket) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *MossPacket) GetClientIp() string {
	if m != nil {
		return m.ClientIp
	}
	return ""
}

type MossMessage struct {
	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *MossMessage) Reset()                    { *m = MossMessage{} }
func (m *MossMessage) String() string            { return proto.CompactTextString(m) }
func (*MossMessage) ProtoMessage()               {}
func (*MossMessage) Descriptor() ([]byte, []int) { return fileDescriptorPayload, []int{1} }

func (m *MossMessage) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *MossMessage) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*MossPacket)(nil), "payload.MossPacket")
	proto.RegisterType((*MossMessage)(nil), "payload.MossMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Invoking service

type InvokingClient interface {
	Invoking(ctx context.Context, in *MossPacket, opts ...grpc.CallOption) (*MossPacket, error)
}

type invokingClient struct {
	cc *grpc.ClientConn
}

func NewInvokingClient(cc *grpc.ClientConn) InvokingClient {
	return &invokingClient{cc}
}

func (c *invokingClient) Invoking(ctx context.Context, in *MossPacket, opts ...grpc.CallOption) (*MossPacket, error) {
	out := new(MossPacket)
	err := grpc.Invoke(ctx, "/payload.Invoking/Invoking", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Invoking service

type InvokingServer interface {
	Invoking(context.Context, *MossPacket) (*MossPacket, error)
}

func RegisterInvokingServer(s *grpc.Server, srv InvokingServer) {
	s.RegisterService(&_Invoking_serviceDesc, srv)
}

func _Invoking_Invoking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MossPacket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvokingServer).Invoking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payload.Invoking/Invoking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvokingServer).Invoking(ctx, req.(*MossPacket))
	}
	return interceptor(ctx, in, info, handler)
}

var _Invoking_serviceDesc = grpc.ServiceDesc{
	ServiceName: "payload.Invoking",
	HandlerType: (*InvokingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoking",
			Handler:    _Invoking_Invoking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payload.proto",
}

func (m *MossPacket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MossPacket) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ServiceCode != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPayload(dAtA, i, uint64(m.ServiceCode))
	}
	if len(m.Payload) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPayload(dAtA, i, uint64(len(m.Payload)))
		i += copy(dAtA[i:], m.Payload)
	}
	if m.MossMessage != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPayload(dAtA, i, uint64(m.MossMessage.Size()))
		n1, err := m.MossMessage.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.UserId) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintPayload(dAtA, i, uint64(len(m.UserId)))
		i += copy(dAtA[i:], m.UserId)
	}
	if len(m.ClientIp) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintPayload(dAtA, i, uint64(len(m.ClientIp)))
		i += copy(dAtA[i:], m.ClientIp)
	}
	return i, nil
}

func (m *MossMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MossMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPayload(dAtA, i, uint64(m.Code))
	}
	if len(m.Msg) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPayload(dAtA, i, uint64(len(m.Msg)))
		i += copy(dAtA[i:], m.Msg)
	}
	return i, nil
}

func encodeVarintPayload(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MossPacket) Size() (n int) {
	var l int
	_ = l
	if m.ServiceCode != 0 {
		n += 1 + sovPayload(uint64(m.ServiceCode))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovPayload(uint64(l))
	}
	if m.MossMessage != nil {
		l = m.MossMessage.Size()
		n += 1 + l + sovPayload(uint64(l))
	}
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovPayload(uint64(l))
	}
	l = len(m.ClientIp)
	if l > 0 {
		n += 1 + l + sovPayload(uint64(l))
	}
	return n
}

func (m *MossMessage) Size() (n int) {
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovPayload(uint64(m.Code))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovPayload(uint64(l))
	}
	return n
}

func sovPayload(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPayload(x uint64) (n int) {
	return sovPayload(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MossPacket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayload
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MossPacket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MossPacket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceCode", wireType)
			}
			m.ServiceCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServiceCode |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MossMessage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MossMessage == nil {
				m.MossMessage = &MossMessage{}
			}
			if err := m.MossMessage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientIp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientIp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPayload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPayload
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MossMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPayload
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MossMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MossMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPayload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPayload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPayload
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPayload(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPayload
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPayload
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthPayload
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPayload
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipPayload(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthPayload = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPayload   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("payload.proto", fileDescriptorPayload) }

var fileDescriptorPayload = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0xbf, 0x69, 0xfa, 0x2f, 0x37, 0x2d, 0x94, 0xf9, 0x04, 0x43, 0x85, 0x10, 0xea, 0x26,
	0x08, 0x69, 0xa1, 0x05, 0x5d, 0xb8, 0x8b, 0xab, 0x2e, 0x2a, 0x92, 0x17, 0x08, 0x69, 0x32, 0x1d,
	0x87, 0x26, 0x99, 0xd0, 0x49, 0x0b, 0xf5, 0x29, 0x7c, 0x1e, 0x57, 0x2e, 0xbb, 0xd3, 0x47, 0xd0,
	0xee, 0x7c, 0x0b, 0xc9, 0x4c, 0x4b, 0x23, 0xba, 0x3b, 0xe7, 0xfe, 0xee, 0x9d, 0x39, 0xf7, 0x42,
	0x37, 0x0f, 0xb7, 0x09, 0x0f, 0xe3, 0x61, 0xbe, 0xe2, 0x05, 0xc7, 0xad, 0x83, 0xed, 0xbb, 0x94,
	0x15, 0x8f, 0xeb, 0xf9, 0x30, 0xe2, 0xe9, 0x88, 0x72, 0xca, 0x47, 0x92, 0xcf, 0xd7, 0x0b, 0xe9,
	0xa4, 0x91, 0x4a, 0xcd, 0xf5, 0x6f, 0x2b, 0xed, 0xc9, 0x76, 0x51, 0xa8, 0xf6, 0xc8, 0xa5, 0x24,
	0x73, 0x37, 0x61, 0xc2, 0xe2, 0xb0, 0x20, 0xa3, 0x5f, 0x42, 0x0d, 0x0f, 0xde, 0x10, 0xc0, 0x8c,
	0x0b, 0xf1, 0x10, 0x46, 0x4b, 0x52, 0x60, 0x17, 0x3a, 0x82, 0xac, 0x36, 0x2c, 0x22, 0x41, 0xc4,
	0x63, 0x62, 0x22, 0x1b, 0x39, 0x5d, 0x0f, 0x5e, 0xbe, 0x5e, 0xb5, 0xc6, 0x95, 0x66, 0x3f, 0xdf,
	0xfb, 0xc6, 0x81, 0xdf, 0xf1, 0x98, 0xe0, 0x4b, 0x38, 0x86, 0x36, 0x6b, 0x36, 0x72, 0x3a, 0x9e,
	0x5e, 0x76, 0xd6, 0x9f, 0x6a, 0x3d, 0xe4, 0x1f, 0x09, 0xbe, 0x81, 0x4e, 0xca, 0x85, 0x08, 0x52,
	0x22, 0x44, 0x48, 0x89, 0xa9, 0xd9, 0xc8, 0x31, 0xc6, 0x67, 0xc3, 0xe3, 0xf6, 0xe5, 0xf7, 0x33,
	0xc5, 0x7c, 0x23, 0x3d, 0x19, 0x7c, 0x0e, 0xad, 0xb5, 0x20, 0xab, 0x80, 0xc5, 0x66, 0xdd, 0x46,
	0x8e, 0xee, 0x37, 0x4b, 0x3b, 0x8d, 0xf1, 0x05, 0xe8, 0x51, 0xc2, 0x48, 0x56, 0x04, 0x2c, 0x37,
	0x1b, 0x12, 0xb5, 0x55, 0x61, 0x9a, 0x0f, 0x26, 0x60, 0x54, 0x5e, 0xc4, 0x18, 0xea, 0xa7, 0x4d,
	0x7c, 0xa9, 0x71, 0x0f, 0xb4, 0x54, 0x50, 0x19, 0x59, 0xf7, 0x4b, 0x39, 0xf6, 0xa0, 0x3d, 0xcd,
	0x36, 0x7c, 0xc9, 0x32, 0x8a, 0xaf, 0x2b, 0xfa, 0xff, 0x8f, 0x94, 0xea, 0x48, 0xfd, 0xbf, 0x8a,
	0x83, 0x7f, 0x5e, 0x6f, 0xf7, 0x69, 0xa1, 0xdd, 0xde, 0x42, 0xef, 0x7b, 0x0b, 0x7d, 0xec, 0x2d,
	0x34, 0x6f, 0xca, 0x1b, 0x4f, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x84, 0x1e, 0x0e, 0xe9,
	0x01, 0x00, 0x00,
}
