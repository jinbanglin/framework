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
	ServiceCode  uint32            `protobuf:"varint,1,opt,name=service_code,json=serviceCode,proto3" json:"service_code,omitempty"`
	Payload      []byte            `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	MossMessage  *MossMessage      `protobuf:"bytes,3,opt,name=moss_message,json=mossMessage" json:"moss_message,omitempty"`
	MossMetadata map[string]string `protobuf:"bytes,4,rep,name=moss_metadata,json=mossMetadata" json:"moss_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
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

func (m *MossPacket) GetMossMetadata() map[string]string {
	if m != nil {
		return m.MossMetadata
	}
	return nil
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
	if len(m.MossMetadata) > 0 {
		for k, _ := range m.MossMetadata {
			dAtA[i] = 0x22
			i++
			v := m.MossMetadata[k]
			mapSize := 1 + len(k) + sovPayload(uint64(len(k))) + 1 + len(v) + sovPayload(uint64(len(v)))
			i = encodeVarintPayload(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintPayload(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintPayload(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
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
	if len(m.MossMetadata) > 0 {
		for k, v := range m.MossMetadata {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovPayload(uint64(len(k))) + 1 + len(v) + sovPayload(uint64(len(v)))
			n += mapEntrySize + 1 + sovPayload(uint64(mapEntrySize))
		}
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
				return fmt.Errorf("proto: wrong wireType = %d for field MossMetadata", wireType)
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
			if m.MossMetadata == nil {
				m.MossMetadata = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPayload
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthPayload
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowPayload
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthPayload
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipPayload(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthPayload
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.MossMetadata[mapkey] = mapvalue
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
				return fmt.Errorf("proto: wrong wireType = %d for field PrepID", wireType)
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
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x7d, 0x93, 0xb4, 0xef, 0xbd, 0x4e, 0x52, 0xe8, 0x9b, 0xd7, 0x45, 0xc8, 0x22, 0x84, 0x8a,
	0x10, 0x84, 0xb4, 0xd0, 0x82, 0x8a, 0x2e, 0x84, 0x88, 0x0b, 0x85, 0x8a, 0xe4, 0x07, 0xca, 0x34,
	0x99, 0xc6, 0xd0, 0x24, 0x53, 0x3a, 0xd3, 0x40, 0xfc, 0x0a, 0xbf, 0xc2, 0x8f, 0x70, 0xe5, 0xb2,
	0x4b, 0x3f, 0x41, 0xbb, 0xf3, 0x2f, 0x24, 0x33, 0x29, 0x8d, 0xb4, 0xbb, 0x73, 0xee, 0x39, 0xe7,
	0xe6, 0x9e, 0x0c, 0x6c, 0x2f, 0x70, 0x91, 0x50, 0x1c, 0xf6, 0x17, 0x4b, 0xca, 0x29, 0xfa, 0x53,
	0x51, 0xd3, 0x8d, 0x62, 0xfe, 0xb8, 0x9a, 0xf6, 0x03, 0x9a, 0x0e, 0x22, 0x1a, 0xd1, 0x81, 0xd0,
	0xa7, 0xab, 0x99, 0x60, 0x82, 0x08, 0x24, 0x73, 0xe6, 0x65, 0xcd, 0x9e, 0x14, 0x33, 0x2e, 0xed,
	0x81, 0x1b, 0x91, 0xcc, 0xcd, 0x71, 0x12, 0x87, 0x98, 0x93, 0xc1, 0x1e, 0x90, 0xe1, 0xde, 0x8b,
	0x02, 0xe1, 0x98, 0x32, 0xf6, 0x80, 0x83, 0x39, 0xe1, 0xc8, 0x85, 0x3a, 0x23, 0xcb, 0x3c, 0x0e,
	0xc8, 0x24, 0xa0, 0x21, 0x31, 0x80, 0x0d, 0x9c, 0xb6, 0x07, 0x5f, 0xbf, 0xde, 0xd4, 0xe6, 0x89,
	0x6a, 0x3f, 0xdf, 0xfb, 0x5a, 0xa5, 0x5f, 0xd3, 0x90, 0xa0, 0x23, 0xb8, 0x3d, 0xda, 0x50, 0x6c,
	0xe0, 0xe8, 0x5e, 0xab, 0x74, 0x36, 0x9e, 0x94, 0x0e, 0xf0, 0xb7, 0x0a, 0x3a, 0x83, 0x7a, 0x4a,
	0x19, 0x9b, 0xa4, 0x84, 0x31, 0x1c, 0x11, 0x43, 0xb5, 0x81, 0xa3, 0x0d, 0xbb, 0xfd, 0x6d, 0xfb,
	0xf2, 0xf3, 0x63, 0xa9, 0xf9, 0x5a, 0xba, 0x23, 0xe8, 0x0e, 0xb6, 0xab, 0x20, 0xc7, 0x21, 0xe6,
	0xd8, 0x68, 0xd8, 0xaa, 0xa3, 0x0d, 0x8f, 0x7f, 0x24, 0xe5, 0xe1, 0xd5, 0x12, 0xe9, 0xbb, 0xc9,
	0xf8, 0xb2, 0xf0, 0xf5, 0xb4, 0x36, 0x32, 0xaf, 0xe0, 0xbf, 0x3d, 0x0b, 0xea, 0x40, 0x75, 0x4e,
	0x0a, 0x51, 0xb2, 0xe5, 0x97, 0x10, 0x75, 0x61, 0x33, 0xc7, 0xc9, 0x8a, 0x88, 0x3a, 0x2d, 0x5f,
	0x92, 0x0b, 0xe5, 0x1c, 0xf4, 0x46, 0x50, 0xab, 0x1d, 0x8a, 0x10, 0x6c, 0xec, 0x7e, 0x90, 0x2f,
	0x70, 0xb9, 0x2e, 0x65, 0x51, 0x15, 0x2d, 0xe1, 0xd0, 0x83, 0x7f, 0x6f, 0xb3, 0x9c, 0xce, 0xe3,
	0x2c, 0x42, 0xa7, 0x35, 0xfc, 0xff, 0x40, 0x05, 0xf3, 0xd0, 0xb0, 0xf7, 0xcb, 0xeb, 0xac, 0x3f,
	0x2d, 0xb0, 0xde, 0x58, 0xe0, 0x7d, 0x63, 0x81, 0x8f, 0x8d, 0x05, 0xa6, 0xbf, 0xc5, 0xd3, 0x8d,
	0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x9c, 0xb5, 0x7c, 0x40, 0x02, 0x00, 0x00,
}
