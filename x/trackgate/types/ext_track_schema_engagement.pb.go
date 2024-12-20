// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: junction/trackgate/ext_track_schema_engagement.proto

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

type ExtTrackSchemaEngagement struct {
	ExtTrackStationId   string `protobuf:"bytes,1,opt,name=extTrackStationId,proto3" json:"extTrackStationId,omitempty"`
	EngageBy            string `protobuf:"bytes,2,opt,name=engageBy,proto3" json:"engageBy,omitempty"`
	AcknowledgementHash string `protobuf:"bytes,3,opt,name=acknowledgementHash,proto3" json:"acknowledgementHash,omitempty"`
	PodNumber           uint64 `protobuf:"varint,4,opt,name=podNumber,proto3" json:"podNumber,omitempty"`
	StationName         string `protobuf:"bytes,5,opt,name=stationName,proto3" json:"stationName,omitempty"`
	TrackName           string `protobuf:"bytes,6,opt,name=trackName,proto3" json:"trackName,omitempty"`
	SchemaVersion       string `protobuf:"bytes,7,opt,name=schemaVersion,proto3" json:"schemaVersion,omitempty"`
	SchemaKey           string `protobuf:"bytes,8,opt,name=schemaKey,proto3" json:"schemaKey,omitempty"`
	SchemaObject        []byte `protobuf:"bytes,9,opt,name=schemaObject,proto3" json:"schemaObject,omitempty"`
	SequencerDetails    []byte `protobuf:"bytes,10,opt,name=sequencerDetails,proto3" json:"sequencerDetails,omitempty"`
	IsVerified          bool   `protobuf:"varint,11,opt,name=isVerified,proto3" json:"isVerified,omitempty"`
	VerifiedBy          string `protobuf:"bytes,12,opt,name=verifiedBy,proto3" json:"verifiedBy,omitempty"`
}

func (m *ExtTrackSchemaEngagement) Reset()         { *m = ExtTrackSchemaEngagement{} }
func (m *ExtTrackSchemaEngagement) String() string { return proto.CompactTextString(m) }
func (*ExtTrackSchemaEngagement) ProtoMessage()    {}
func (*ExtTrackSchemaEngagement) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5a267180dcb0ce3, []int{0}
}
func (m *ExtTrackSchemaEngagement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExtTrackSchemaEngagement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExtTrackSchemaEngagement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExtTrackSchemaEngagement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtTrackSchemaEngagement.Merge(m, src)
}
func (m *ExtTrackSchemaEngagement) XXX_Size() int {
	return m.Size()
}
func (m *ExtTrackSchemaEngagement) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtTrackSchemaEngagement.DiscardUnknown(m)
}

var xxx_messageInfo_ExtTrackSchemaEngagement proto.InternalMessageInfo

func (m *ExtTrackSchemaEngagement) GetExtTrackStationId() string {
	if m != nil {
		return m.ExtTrackStationId
	}
	return ""
}

func (m *ExtTrackSchemaEngagement) GetEngageBy() string {
	if m != nil {
		return m.EngageBy
	}
	return ""
}

func (m *ExtTrackSchemaEngagement) GetAcknowledgementHash() string {
	if m != nil {
		return m.AcknowledgementHash
	}
	return ""
}

func (m *ExtTrackSchemaEngagement) GetPodNumber() uint64 {
	if m != nil {
		return m.PodNumber
	}
	return 0
}

func (m *ExtTrackSchemaEngagement) GetStationName() string {
	if m != nil {
		return m.StationName
	}
	return ""
}

func (m *ExtTrackSchemaEngagement) GetTrackName() string {
	if m != nil {
		return m.TrackName
	}
	return ""
}

func (m *ExtTrackSchemaEngagement) GetSchemaVersion() string {
	if m != nil {
		return m.SchemaVersion
	}
	return ""
}

func (m *ExtTrackSchemaEngagement) GetSchemaKey() string {
	if m != nil {
		return m.SchemaKey
	}
	return ""
}

func (m *ExtTrackSchemaEngagement) GetSchemaObject() []byte {
	if m != nil {
		return m.SchemaObject
	}
	return nil
}

func (m *ExtTrackSchemaEngagement) GetSequencerDetails() []byte {
	if m != nil {
		return m.SequencerDetails
	}
	return nil
}

func (m *ExtTrackSchemaEngagement) GetIsVerified() bool {
	if m != nil {
		return m.IsVerified
	}
	return false
}

func (m *ExtTrackSchemaEngagement) GetVerifiedBy() string {
	if m != nil {
		return m.VerifiedBy
	}
	return ""
}

func init() {
	proto.RegisterType((*ExtTrackSchemaEngagement)(nil), "junction.trackgate.ExtTrackSchemaEngagement")
}

func init() {
	proto.RegisterFile("junction/trackgate/ext_track_schema_engagement.proto", fileDescriptor_a5a267180dcb0ce3)
}

var fileDescriptor_a5a267180dcb0ce3 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcd, 0xaa, 0xd3, 0x40,
	0x14, 0xc7, 0x3b, 0xde, 0xeb, 0xb5, 0x9d, 0x5b, 0x41, 0xc7, 0xcd, 0x20, 0x12, 0x42, 0x71, 0x11,
	0x44, 0x1b, 0x41, 0xc1, 0x7d, 0xb1, 0xa0, 0x08, 0x15, 0xa3, 0x74, 0xe1, 0xa6, 0x4c, 0x26, 0xc7,
	0x64, 0xda, 0x66, 0xa6, 0xce, 0x4c, 0x6c, 0xf3, 0x16, 0x3e, 0x56, 0x97, 0x5d, 0xba, 0x94, 0xf6,
	0x45, 0x24, 0x13, 0x9b, 0xa6, 0xf4, 0xee, 0x72, 0x7e, 0xff, 0x8f, 0x93, 0x84, 0x83, 0xdf, 0xce,
	0x0b, 0xc9, 0xad, 0x50, 0x32, 0xb4, 0x9a, 0xf1, 0x45, 0xca, 0x2c, 0x84, 0xb0, 0xb1, 0x33, 0x37,
	0xcd, 0x0c, 0xcf, 0x20, 0x67, 0x33, 0x90, 0x29, 0x4b, 0x21, 0x07, 0x69, 0x87, 0x2b, 0xad, 0xac,
	0x22, 0xe4, 0x98, 0x1a, 0x36, 0xa9, 0xc1, 0xf6, 0x0a, 0xd3, 0xf1, 0xc6, 0x7e, 0xab, 0xc0, 0x57,
	0x97, 0x1b, 0x37, 0x31, 0xf2, 0x12, 0x3f, 0x86, 0xa3, 0x66, 0x59, 0x95, 0xfc, 0x98, 0x50, 0xe4,
	0xa3, 0xa0, 0x17, 0x5d, 0x0a, 0xe4, 0x29, 0xee, 0xd6, 0x2b, 0x47, 0x25, 0xbd, 0xe7, 0x4c, 0xcd,
	0x4c, 0x5e, 0xe3, 0x27, 0x8c, 0x2f, 0xa4, 0x5a, 0x2f, 0x21, 0xa9, 0xcb, 0x3f, 0x30, 0x93, 0xd1,
	0x2b, 0x67, 0xbb, 0x4b, 0x22, 0xcf, 0x70, 0x6f, 0xa5, 0x92, 0x49, 0x91, 0xc7, 0xa0, 0xe9, 0xb5,
	0x8f, 0x82, 0xeb, 0xe8, 0x04, 0x88, 0x8f, 0x6f, 0x4d, 0xbd, 0x78, 0xc2, 0x72, 0xa0, 0xf7, 0x5d,
	0x4f, 0x1b, 0x55, 0x79, 0xf7, 0x95, 0x4e, 0xbf, 0x71, 0xfa, 0x09, 0x90, 0xe7, 0xf8, 0x61, 0xfd,
	0x97, 0xa6, 0xa0, 0x8d, 0x50, 0x92, 0x3e, 0x70, 0x8e, 0x73, 0x58, 0x75, 0xd4, 0xe0, 0x13, 0x94,
	0xb4, 0x5b, 0x77, 0x34, 0x80, 0x0c, 0x70, 0xbf, 0x1e, 0x3e, 0xc7, 0x73, 0xe0, 0x96, 0xf6, 0x7c,
	0x14, 0xf4, 0xa3, 0x33, 0x46, 0x5e, 0xe0, 0x47, 0x06, 0x7e, 0x16, 0x20, 0x39, 0xe8, 0xf7, 0x60,
	0x99, 0x58, 0x1a, 0x8a, 0x9d, 0xef, 0x82, 0x13, 0x0f, 0x63, 0x61, 0xa6, 0xa0, 0xc5, 0x0f, 0x01,
	0x09, 0xbd, 0xf5, 0x51, 0xd0, 0x8d, 0x5a, 0xa4, 0xd2, 0x7f, 0xfd, 0x7f, 0x1e, 0x95, 0xb4, 0xef,
	0x5e, 0xa7, 0x45, 0x46, 0x5f, 0xb6, 0x7b, 0x0f, 0xed, 0xf6, 0x1e, 0xfa, 0xbb, 0xf7, 0xd0, 0xef,
	0x83, 0xd7, 0xd9, 0x1d, 0xbc, 0xce, 0x9f, 0x83, 0xd7, 0xf9, 0xfe, 0x2e, 0x15, 0x36, 0x2b, 0xe2,
	0x21, 0x57, 0x79, 0xc8, 0x84, 0xe6, 0x19, 0x13, 0xd2, 0xbc, 0x92, 0x60, 0xd7, 0x4a, 0x2f, 0xc2,
	0xe6, 0x96, 0x36, 0xad, 0x6b, 0xb2, 0xe5, 0x0a, 0x4c, 0x7c, 0xe3, 0x0e, 0xe7, 0xcd, 0xbf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xd0, 0xc7, 0x5f, 0x1a, 0x70, 0x02, 0x00, 0x00,
}

func (m *ExtTrackSchemaEngagement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtTrackSchemaEngagement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExtTrackSchemaEngagement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VerifiedBy) > 0 {
		i -= len(m.VerifiedBy)
		copy(dAtA[i:], m.VerifiedBy)
		i = encodeVarintExtTrackSchemaEngagement(dAtA, i, uint64(len(m.VerifiedBy)))
		i--
		dAtA[i] = 0x62
	}
	if m.IsVerified {
		i--
		if m.IsVerified {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x58
	}
	if len(m.SequencerDetails) > 0 {
		i -= len(m.SequencerDetails)
		copy(dAtA[i:], m.SequencerDetails)
		i = encodeVarintExtTrackSchemaEngagement(dAtA, i, uint64(len(m.SequencerDetails)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.SchemaObject) > 0 {
		i -= len(m.SchemaObject)
		copy(dAtA[i:], m.SchemaObject)
		i = encodeVarintExtTrackSchemaEngagement(dAtA, i, uint64(len(m.SchemaObject)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.SchemaKey) > 0 {
		i -= len(m.SchemaKey)
		copy(dAtA[i:], m.SchemaKey)
		i = encodeVarintExtTrackSchemaEngagement(dAtA, i, uint64(len(m.SchemaKey)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.SchemaVersion) > 0 {
		i -= len(m.SchemaVersion)
		copy(dAtA[i:], m.SchemaVersion)
		i = encodeVarintExtTrackSchemaEngagement(dAtA, i, uint64(len(m.SchemaVersion)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.TrackName) > 0 {
		i -= len(m.TrackName)
		copy(dAtA[i:], m.TrackName)
		i = encodeVarintExtTrackSchemaEngagement(dAtA, i, uint64(len(m.TrackName)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.StationName) > 0 {
		i -= len(m.StationName)
		copy(dAtA[i:], m.StationName)
		i = encodeVarintExtTrackSchemaEngagement(dAtA, i, uint64(len(m.StationName)))
		i--
		dAtA[i] = 0x2a
	}
	if m.PodNumber != 0 {
		i = encodeVarintExtTrackSchemaEngagement(dAtA, i, uint64(m.PodNumber))
		i--
		dAtA[i] = 0x20
	}
	if len(m.AcknowledgementHash) > 0 {
		i -= len(m.AcknowledgementHash)
		copy(dAtA[i:], m.AcknowledgementHash)
		i = encodeVarintExtTrackSchemaEngagement(dAtA, i, uint64(len(m.AcknowledgementHash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.EngageBy) > 0 {
		i -= len(m.EngageBy)
		copy(dAtA[i:], m.EngageBy)
		i = encodeVarintExtTrackSchemaEngagement(dAtA, i, uint64(len(m.EngageBy)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ExtTrackStationId) > 0 {
		i -= len(m.ExtTrackStationId)
		copy(dAtA[i:], m.ExtTrackStationId)
		i = encodeVarintExtTrackSchemaEngagement(dAtA, i, uint64(len(m.ExtTrackStationId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintExtTrackSchemaEngagement(dAtA []byte, offset int, v uint64) int {
	offset -= sovExtTrackSchemaEngagement(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExtTrackSchemaEngagement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ExtTrackStationId)
	if l > 0 {
		n += 1 + l + sovExtTrackSchemaEngagement(uint64(l))
	}
	l = len(m.EngageBy)
	if l > 0 {
		n += 1 + l + sovExtTrackSchemaEngagement(uint64(l))
	}
	l = len(m.AcknowledgementHash)
	if l > 0 {
		n += 1 + l + sovExtTrackSchemaEngagement(uint64(l))
	}
	if m.PodNumber != 0 {
		n += 1 + sovExtTrackSchemaEngagement(uint64(m.PodNumber))
	}
	l = len(m.StationName)
	if l > 0 {
		n += 1 + l + sovExtTrackSchemaEngagement(uint64(l))
	}
	l = len(m.TrackName)
	if l > 0 {
		n += 1 + l + sovExtTrackSchemaEngagement(uint64(l))
	}
	l = len(m.SchemaVersion)
	if l > 0 {
		n += 1 + l + sovExtTrackSchemaEngagement(uint64(l))
	}
	l = len(m.SchemaKey)
	if l > 0 {
		n += 1 + l + sovExtTrackSchemaEngagement(uint64(l))
	}
	l = len(m.SchemaObject)
	if l > 0 {
		n += 1 + l + sovExtTrackSchemaEngagement(uint64(l))
	}
	l = len(m.SequencerDetails)
	if l > 0 {
		n += 1 + l + sovExtTrackSchemaEngagement(uint64(l))
	}
	if m.IsVerified {
		n += 2
	}
	l = len(m.VerifiedBy)
	if l > 0 {
		n += 1 + l + sovExtTrackSchemaEngagement(uint64(l))
	}
	return n
}

func sovExtTrackSchemaEngagement(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExtTrackSchemaEngagement(x uint64) (n int) {
	return sovExtTrackSchemaEngagement(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExtTrackSchemaEngagement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExtTrackSchemaEngagement
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
			return fmt.Errorf("proto: ExtTrackSchemaEngagement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtTrackSchemaEngagement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtTrackStationId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackSchemaEngagement
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
				return ErrInvalidLengthExtTrackSchemaEngagement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackSchemaEngagement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExtTrackStationId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EngageBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackSchemaEngagement
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
				return ErrInvalidLengthExtTrackSchemaEngagement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackSchemaEngagement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EngageBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AcknowledgementHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackSchemaEngagement
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
				return ErrInvalidLengthExtTrackSchemaEngagement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackSchemaEngagement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AcknowledgementHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PodNumber", wireType)
			}
			m.PodNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackSchemaEngagement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PodNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StationName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackSchemaEngagement
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
				return ErrInvalidLengthExtTrackSchemaEngagement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackSchemaEngagement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StationName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrackName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackSchemaEngagement
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
				return ErrInvalidLengthExtTrackSchemaEngagement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackSchemaEngagement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrackName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackSchemaEngagement
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
				return ErrInvalidLengthExtTrackSchemaEngagement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackSchemaEngagement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SchemaVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackSchemaEngagement
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
				return ErrInvalidLengthExtTrackSchemaEngagement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackSchemaEngagement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SchemaKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaObject", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackSchemaEngagement
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
				return ErrInvalidLengthExtTrackSchemaEngagement
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackSchemaEngagement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SchemaObject = append(m.SchemaObject[:0], dAtA[iNdEx:postIndex]...)
			if m.SchemaObject == nil {
				m.SchemaObject = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequencerDetails", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackSchemaEngagement
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
				return ErrInvalidLengthExtTrackSchemaEngagement
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackSchemaEngagement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SequencerDetails = append(m.SequencerDetails[:0], dAtA[iNdEx:postIndex]...)
			if m.SequencerDetails == nil {
				m.SequencerDetails = []byte{}
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsVerified", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackSchemaEngagement
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsVerified = bool(v != 0)
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerifiedBy", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtTrackSchemaEngagement
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
				return ErrInvalidLengthExtTrackSchemaEngagement
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtTrackSchemaEngagement
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VerifiedBy = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExtTrackSchemaEngagement(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExtTrackSchemaEngagement
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
func skipExtTrackSchemaEngagement(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExtTrackSchemaEngagement
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
					return 0, ErrIntOverflowExtTrackSchemaEngagement
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
					return 0, ErrIntOverflowExtTrackSchemaEngagement
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
				return 0, ErrInvalidLengthExtTrackSchemaEngagement
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExtTrackSchemaEngagement
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExtTrackSchemaEngagement
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExtTrackSchemaEngagement        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExtTrackSchemaEngagement          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExtTrackSchemaEngagement = fmt.Errorf("proto: unexpected end of group")
)
