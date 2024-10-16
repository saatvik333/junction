// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package trackgate

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_ExtTrackSchema                   protoreflect.MessageDescriptor
	fd_ExtTrackSchema_trackName         protoreflect.FieldDescriptor
	fd_ExtTrackSchema_extTrackStationId protoreflect.FieldDescriptor
	fd_ExtTrackSchema_version           protoreflect.FieldDescriptor
	fd_ExtTrackSchema_schemaKey         protoreflect.FieldDescriptor
	fd_ExtTrackSchema_schema            protoreflect.FieldDescriptor
)

func init() {
	file_junction_trackgate_ext_track_schema_proto_init()
	md_ExtTrackSchema = File_junction_trackgate_ext_track_schema_proto.Messages().ByName("ExtTrackSchema")
	fd_ExtTrackSchema_trackName = md_ExtTrackSchema.Fields().ByName("trackName")
	fd_ExtTrackSchema_extTrackStationId = md_ExtTrackSchema.Fields().ByName("extTrackStationId")
	fd_ExtTrackSchema_version = md_ExtTrackSchema.Fields().ByName("version")
	fd_ExtTrackSchema_schemaKey = md_ExtTrackSchema.Fields().ByName("schemaKey")
	fd_ExtTrackSchema_schema = md_ExtTrackSchema.Fields().ByName("schema")
}

var _ protoreflect.Message = (*fastReflection_ExtTrackSchema)(nil)

type fastReflection_ExtTrackSchema ExtTrackSchema

func (x *ExtTrackSchema) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ExtTrackSchema)(x)
}

func (x *ExtTrackSchema) slowProtoReflect() protoreflect.Message {
	mi := &file_junction_trackgate_ext_track_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ExtTrackSchema_messageType fastReflection_ExtTrackSchema_messageType
var _ protoreflect.MessageType = fastReflection_ExtTrackSchema_messageType{}

type fastReflection_ExtTrackSchema_messageType struct{}

func (x fastReflection_ExtTrackSchema_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ExtTrackSchema)(nil)
}
func (x fastReflection_ExtTrackSchema_messageType) New() protoreflect.Message {
	return new(fastReflection_ExtTrackSchema)
}
func (x fastReflection_ExtTrackSchema_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ExtTrackSchema
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ExtTrackSchema) Descriptor() protoreflect.MessageDescriptor {
	return md_ExtTrackSchema
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ExtTrackSchema) Type() protoreflect.MessageType {
	return _fastReflection_ExtTrackSchema_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ExtTrackSchema) New() protoreflect.Message {
	return new(fastReflection_ExtTrackSchema)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ExtTrackSchema) Interface() protoreflect.ProtoMessage {
	return (*ExtTrackSchema)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ExtTrackSchema) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.TrackName != "" {
		value := protoreflect.ValueOfString(x.TrackName)
		if !f(fd_ExtTrackSchema_trackName, value) {
			return
		}
	}
	if x.ExtTrackStationId != "" {
		value := protoreflect.ValueOfString(x.ExtTrackStationId)
		if !f(fd_ExtTrackSchema_extTrackStationId, value) {
			return
		}
	}
	if x.Version != "" {
		value := protoreflect.ValueOfString(x.Version)
		if !f(fd_ExtTrackSchema_version, value) {
			return
		}
	}
	if x.SchemaKey != "" {
		value := protoreflect.ValueOfString(x.SchemaKey)
		if !f(fd_ExtTrackSchema_schemaKey, value) {
			return
		}
	}
	if len(x.Schema) != 0 {
		value := protoreflect.ValueOfBytes(x.Schema)
		if !f(fd_ExtTrackSchema_schema, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ExtTrackSchema) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "junction.trackgate.ExtTrackSchema.trackName":
		return x.TrackName != ""
	case "junction.trackgate.ExtTrackSchema.extTrackStationId":
		return x.ExtTrackStationId != ""
	case "junction.trackgate.ExtTrackSchema.version":
		return x.Version != ""
	case "junction.trackgate.ExtTrackSchema.schemaKey":
		return x.SchemaKey != ""
	case "junction.trackgate.ExtTrackSchema.schema":
		return len(x.Schema) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: junction.trackgate.ExtTrackSchema"))
		}
		panic(fmt.Errorf("message junction.trackgate.ExtTrackSchema does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ExtTrackSchema) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "junction.trackgate.ExtTrackSchema.trackName":
		x.TrackName = ""
	case "junction.trackgate.ExtTrackSchema.extTrackStationId":
		x.ExtTrackStationId = ""
	case "junction.trackgate.ExtTrackSchema.version":
		x.Version = ""
	case "junction.trackgate.ExtTrackSchema.schemaKey":
		x.SchemaKey = ""
	case "junction.trackgate.ExtTrackSchema.schema":
		x.Schema = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: junction.trackgate.ExtTrackSchema"))
		}
		panic(fmt.Errorf("message junction.trackgate.ExtTrackSchema does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ExtTrackSchema) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "junction.trackgate.ExtTrackSchema.trackName":
		value := x.TrackName
		return protoreflect.ValueOfString(value)
	case "junction.trackgate.ExtTrackSchema.extTrackStationId":
		value := x.ExtTrackStationId
		return protoreflect.ValueOfString(value)
	case "junction.trackgate.ExtTrackSchema.version":
		value := x.Version
		return protoreflect.ValueOfString(value)
	case "junction.trackgate.ExtTrackSchema.schemaKey":
		value := x.SchemaKey
		return protoreflect.ValueOfString(value)
	case "junction.trackgate.ExtTrackSchema.schema":
		value := x.Schema
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: junction.trackgate.ExtTrackSchema"))
		}
		panic(fmt.Errorf("message junction.trackgate.ExtTrackSchema does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ExtTrackSchema) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "junction.trackgate.ExtTrackSchema.trackName":
		x.TrackName = value.Interface().(string)
	case "junction.trackgate.ExtTrackSchema.extTrackStationId":
		x.ExtTrackStationId = value.Interface().(string)
	case "junction.trackgate.ExtTrackSchema.version":
		x.Version = value.Interface().(string)
	case "junction.trackgate.ExtTrackSchema.schemaKey":
		x.SchemaKey = value.Interface().(string)
	case "junction.trackgate.ExtTrackSchema.schema":
		x.Schema = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: junction.trackgate.ExtTrackSchema"))
		}
		panic(fmt.Errorf("message junction.trackgate.ExtTrackSchema does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ExtTrackSchema) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "junction.trackgate.ExtTrackSchema.trackName":
		panic(fmt.Errorf("field trackName of message junction.trackgate.ExtTrackSchema is not mutable"))
	case "junction.trackgate.ExtTrackSchema.extTrackStationId":
		panic(fmt.Errorf("field extTrackStationId of message junction.trackgate.ExtTrackSchema is not mutable"))
	case "junction.trackgate.ExtTrackSchema.version":
		panic(fmt.Errorf("field version of message junction.trackgate.ExtTrackSchema is not mutable"))
	case "junction.trackgate.ExtTrackSchema.schemaKey":
		panic(fmt.Errorf("field schemaKey of message junction.trackgate.ExtTrackSchema is not mutable"))
	case "junction.trackgate.ExtTrackSchema.schema":
		panic(fmt.Errorf("field schema of message junction.trackgate.ExtTrackSchema is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: junction.trackgate.ExtTrackSchema"))
		}
		panic(fmt.Errorf("message junction.trackgate.ExtTrackSchema does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ExtTrackSchema) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "junction.trackgate.ExtTrackSchema.trackName":
		return protoreflect.ValueOfString("")
	case "junction.trackgate.ExtTrackSchema.extTrackStationId":
		return protoreflect.ValueOfString("")
	case "junction.trackgate.ExtTrackSchema.version":
		return protoreflect.ValueOfString("")
	case "junction.trackgate.ExtTrackSchema.schemaKey":
		return protoreflect.ValueOfString("")
	case "junction.trackgate.ExtTrackSchema.schema":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: junction.trackgate.ExtTrackSchema"))
		}
		panic(fmt.Errorf("message junction.trackgate.ExtTrackSchema does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ExtTrackSchema) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in junction.trackgate.ExtTrackSchema", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ExtTrackSchema) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ExtTrackSchema) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ExtTrackSchema) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ExtTrackSchema) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ExtTrackSchema)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.TrackName)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.ExtTrackStationId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Version)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.SchemaKey)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Schema)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ExtTrackSchema)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Schema) > 0 {
			i -= len(x.Schema)
			copy(dAtA[i:], x.Schema)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Schema)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.SchemaKey) > 0 {
			i -= len(x.SchemaKey)
			copy(dAtA[i:], x.SchemaKey)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.SchemaKey)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Version) > 0 {
			i -= len(x.Version)
			copy(dAtA[i:], x.Version)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Version)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.ExtTrackStationId) > 0 {
			i -= len(x.ExtTrackStationId)
			copy(dAtA[i:], x.ExtTrackStationId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ExtTrackStationId)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.TrackName) > 0 {
			i -= len(x.TrackName)
			copy(dAtA[i:], x.TrackName)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TrackName)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ExtTrackSchema)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ExtTrackSchema: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ExtTrackSchema: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TrackName", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.TrackName = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ExtTrackStationId", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ExtTrackStationId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Version = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SchemaKey", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.SchemaKey = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Schema", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Schema = append(x.Schema[:0], dAtA[iNdEx:postIndex]...)
				if x.Schema == nil {
					x.Schema = []byte{}
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: junction/trackgate/ext_track_schema.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExtTrackSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrackName         string `protobuf:"bytes,1,opt,name=trackName,proto3" json:"trackName,omitempty"`
	ExtTrackStationId string `protobuf:"bytes,2,opt,name=extTrackStationId,proto3" json:"extTrackStationId,omitempty"`
	Version           string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	SchemaKey         string `protobuf:"bytes,4,opt,name=schemaKey,proto3" json:"schemaKey,omitempty"`
	Schema            []byte `protobuf:"bytes,5,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (x *ExtTrackSchema) Reset() {
	*x = ExtTrackSchema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_junction_trackgate_ext_track_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtTrackSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtTrackSchema) ProtoMessage() {}

// Deprecated: Use ExtTrackSchema.ProtoReflect.Descriptor instead.
func (*ExtTrackSchema) Descriptor() ([]byte, []int) {
	return file_junction_trackgate_ext_track_schema_proto_rawDescGZIP(), []int{0}
}

func (x *ExtTrackSchema) GetTrackName() string {
	if x != nil {
		return x.TrackName
	}
	return ""
}

func (x *ExtTrackSchema) GetExtTrackStationId() string {
	if x != nil {
		return x.ExtTrackStationId
	}
	return ""
}

func (x *ExtTrackSchema) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ExtTrackSchema) GetSchemaKey() string {
	if x != nil {
		return x.SchemaKey
	}
	return ""
}

func (x *ExtTrackSchema) GetSchema() []byte {
	if x != nil {
		return x.Schema
	}
	return nil
}

var File_junction_trackgate_ext_track_schema_proto protoreflect.FileDescriptor

var file_junction_trackgate_ext_track_schema_proto_rawDesc = []byte{
	0x0a, 0x29, 0x6a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x67, 0x61, 0x74, 0x65, 0x2f, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6a, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x67, 0x61, 0x74, 0x65, 0x22,
	0xac, 0x01, 0x0a, 0x0e, 0x45, 0x78, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x53, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2c, 0x0a, 0x11, 0x65, 0x78, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x65, 0x78, 0x74,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x42, 0xbb,
	0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x6a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x67, 0x61, 0x74, 0x65, 0x42, 0x13, 0x45, 0x78, 0x74, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x23, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x67, 0x61, 0x74, 0x65, 0xa2, 0x02, 0x03, 0x4a, 0x54, 0x58, 0xaa, 0x02, 0x12, 0x4a, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x67, 0x61, 0x74, 0x65,
	0xca, 0x02, 0x12, 0x4a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x67, 0x61, 0x74, 0x65, 0xe2, 0x02, 0x1e, 0x4a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5c, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x67, 0x61, 0x74, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x4a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x3a, 0x3a, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x67, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_junction_trackgate_ext_track_schema_proto_rawDescOnce sync.Once
	file_junction_trackgate_ext_track_schema_proto_rawDescData = file_junction_trackgate_ext_track_schema_proto_rawDesc
)

func file_junction_trackgate_ext_track_schema_proto_rawDescGZIP() []byte {
	file_junction_trackgate_ext_track_schema_proto_rawDescOnce.Do(func() {
		file_junction_trackgate_ext_track_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_junction_trackgate_ext_track_schema_proto_rawDescData)
	})
	return file_junction_trackgate_ext_track_schema_proto_rawDescData
}

var file_junction_trackgate_ext_track_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_junction_trackgate_ext_track_schema_proto_goTypes = []interface{}{
	(*ExtTrackSchema)(nil), // 0: junction.trackgate.ExtTrackSchema
}
var file_junction_trackgate_ext_track_schema_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_junction_trackgate_ext_track_schema_proto_init() }
func file_junction_trackgate_ext_track_schema_proto_init() {
	if File_junction_trackgate_ext_track_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_junction_trackgate_ext_track_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtTrackSchema); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_junction_trackgate_ext_track_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_junction_trackgate_ext_track_schema_proto_goTypes,
		DependencyIndexes: file_junction_trackgate_ext_track_schema_proto_depIdxs,
		MessageInfos:      file_junction_trackgate_ext_track_schema_proto_msgTypes,
	}.Build()
	File_junction_trackgate_ext_track_schema_proto = out.File
	file_junction_trackgate_ext_track_schema_proto_rawDesc = nil
	file_junction_trackgate_ext_track_schema_proto_goTypes = nil
	file_junction_trackgate_ext_track_schema_proto_depIdxs = nil
}
