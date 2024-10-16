// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: junction/trackgate/ext_track_schema.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ExtTrackSchema struct {
	TrackName         string `protobuf:"bytes,1,opt,name=trackName,proto3" json:"trackName,omitempty"`
	ExtTrackStationId string `protobuf:"bytes,2,opt,name=extTrackStationId,proto3" json:"extTrackStationId,omitempty"`
	Version           string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	SchemaKey         string `protobuf:"bytes,4,opt,name=schemaKey,proto3" json:"schemaKey,omitempty"`
	Schema            []byte `protobuf:"bytes,5,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (m *ExtTrackSchema) Reset()         { *m = ExtTrackSchema{} }
func (m *ExtTrackSchema) String() string { return proto.CompactTextString(m) }
func (*ExtTrackSchema) ProtoMessage()    {}
func (*ExtTrackSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_53dc635c68eb7027, []int{0}
}
func (m *ExtTrackSchema) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExtTrackSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExtTrackSchema.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExtTrackSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtTrackSchema.Merge(m, src)
}
func (m *ExtTrackSchema) XXX_Size() int {
	return m.Size()
}
func (m *ExtTrackSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtTrackSchema.DiscardUnknown(m)
}

var xxx_messageInfo_ExtTrackSchema proto.InternalMessageInfo

func (m *ExtTrackSchema) GetTrackName() string {
	if m != nil {
		return m.TrackName
	}
	return ""
}

func (m *ExtTrackSchema) GetExtTrackStationId() string {
	if m != nil {
		return m.ExtTrackStationId
	}
	return ""
}

func (m *ExtTrackSchema) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ExtTrackSchema) GetSchemaKey() string {
	if m != nil {
		return m.SchemaKey
	}
	return ""
}

func (m *ExtTrackSchema) GetSchema() []byte {
	if m != nil {
		return m.Schema
	}
	return nil
}

func init() {
	proto.RegisterType((*ExtTrackSchema)(nil), "junction.trackgate.ExtTrackSchema")
}

func init() {
	proto.RegisterFile("junction/trackgate/ext_track_schema.proto", fileDescriptor_53dc635c68eb7027)
}

var fileDescriptor_53dc635c68eb7027 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcc, 0x2a, 0xcd, 0x4b,
	0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x2f, 0x29, 0x4a, 0x4c, 0xce, 0x4e, 0x4f, 0x2c, 0x49, 0xd5, 0x4f,
	0xad, 0x28, 0x89, 0x07, 0xf3, 0xe2, 0x8b, 0x93, 0x33, 0x52, 0x73, 0x13, 0xf5, 0x0a, 0x8a, 0xf2,
	0x4b, 0xf2, 0x85, 0x84, 0x60, 0x4a, 0xf5, 0xe0, 0x4a, 0x95, 0xd6, 0x30, 0x72, 0xf1, 0xb9, 0x56,
	0x94, 0x84, 0x80, 0x04, 0x82, 0xc1, 0x8a, 0x85, 0x64, 0xb8, 0x38, 0xc1, 0xf2, 0x7e, 0x89, 0xb9,
	0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x08, 0x01, 0x21, 0x1d, 0x2e, 0xc1, 0x54, 0x98,
	0xfa, 0x92, 0x44, 0x90, 0x69, 0x9e, 0x29, 0x12, 0x4c, 0x60, 0x55, 0x98, 0x12, 0x42, 0x12, 0x5c,
	0xec, 0x65, 0xa9, 0x45, 0xc5, 0x99, 0xf9, 0x79, 0x12, 0xcc, 0x60, 0x35, 0x30, 0x2e, 0xc8, 0x16,
	0x88, 0xe3, 0xbc, 0x53, 0x2b, 0x25, 0x58, 0x20, 0xb6, 0xc0, 0x05, 0x84, 0xc4, 0xb8, 0xd8, 0x20,
	0x1c, 0x09, 0x56, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x28, 0xcf, 0x29, 0xf0, 0xc4, 0x23, 0x39, 0xc6,
	0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39,
	0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xcc, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3,
	0x73, 0xf5, 0x13, 0x33, 0x8b, 0x92, 0x33, 0x12, 0x33, 0xf3, 0x8a, 0x75, 0xf3, 0x52, 0x4b, 0xca,
	0xf3, 0x8b, 0xb2, 0xf5, 0xe1, 0x81, 0x54, 0x81, 0x14, 0x4c, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49,
	0x6c, 0xe0, 0xc0, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xd6, 0xaa, 0xfe, 0x49, 0x01,
	0x00, 0x00,
}

func (m *ExtTrackSchema) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtTrackSchema) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExtTrackSchema) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Schema) > 0 {
		i -= len(m.Schema)
		copy(dAtA[i:], m.Schema)
		i = encodeVarintExtTrackSchema(dAtA, i, uint64(len(m.Schema)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SchemaKey) > 0 {
		i -= len(m.SchemaKey)
		copy(dAtA[i:], m.SchemaKey)
		i = encodeVarintExtTrackSchema(dAtA, i, uint64(len(m.SchemaKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintExtTrackSchema(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ExtTrackStationId) > 0 {
		i -= len(m.ExtTrackStationId)
		copy(dAtA[i:], m.ExtTrackStationId)
		i = encodeVarintExtTrackSchema(dAtA, i, uint64(len(m.ExtTrackStationId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TrackName) > 0 {
		i -= len(m.TrackName)
		copy(dAtA[i:], m.TrackName)
		i = encodeVarintExtTrackSchema(dAtA, i, uint64(len(m.TrackName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintExtTrackSchema(dAtA []byte, offset int, v uint64) int {
	offset -= sovExtTrackSchema(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExtTrackSchema) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TrackName)
	if l > 0 {
		n += 1 + l + sovExtTrackSchema(uint64(l))
	}
	l = len(m.ExtTrackStationId)
	if l > 0 {
		n += 1 + l + sovExtTrackSchema(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovExtTrackSchema(uint64(l))
	}
	l = len(m.SchemaKey)
	if l > 0 {
		n += 1 + l + sovExtTrackSchema(uint64(l))
	}
	l = len(m.Schema)
	if l > 0 {
		n += 1 + l + sovExtTrackSchema(uint64(l))
	}
	return n
}

func sovExtTrackSchema(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExtTrackSchema(x uint64) (n int) {
	return sovExtTrackSchema(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExtTrackSchema) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExtTrackSchema
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExtTrackSchema: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtTrackSchema: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrackName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExtTrackSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrackName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtTrackStationId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExtTrackSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExtTrackStationId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExtTrackSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthExtTrackSchema
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SchemaKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Schema", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackSchema
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthExtTrackSchema
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackSchema
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Schema = append(m.Schema[:0], dAtA[iNdEx:postIndex]...)
			if m.Schema == nil {
				m.Schema = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExtTrackSchema(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExtTrackSchema
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
func skipExtTrackSchema(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExtTrackSchema
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
					return 0, ErrIntOverflowExtTrackSchema
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExtTrackSchema
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
			if length < 0 {
				return 0, ErrInvalidLengthExtTrackSchema
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExtTrackSchema
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExtTrackSchema
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExtTrackSchema        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExtTrackSchema          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExtTrackSchema = fmt.Errorf("proto: unexpected end of group")
)
