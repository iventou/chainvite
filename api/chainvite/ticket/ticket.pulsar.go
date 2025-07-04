// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package ticket

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
	md_Ticket            protoreflect.MessageDescriptor
	fd_Ticket_index      protoreflect.FieldDescriptor
	fd_Ticket_ticketId   protoreflect.FieldDescriptor
	fd_Ticket_eventId    protoreflect.FieldDescriptor
	fd_Ticket_ticketType protoreflect.FieldDescriptor
	fd_Ticket_valid      protoreflect.FieldDescriptor
	fd_Ticket_owner      protoreflect.FieldDescriptor
	fd_Ticket_metadata   protoreflect.FieldDescriptor
	fd_Ticket_creator    protoreflect.FieldDescriptor
)

func init() {
	file_chainvite_ticket_ticket_proto_init()
	md_Ticket = File_chainvite_ticket_ticket_proto.Messages().ByName("Ticket")
	fd_Ticket_index = md_Ticket.Fields().ByName("index")
	fd_Ticket_ticketId = md_Ticket.Fields().ByName("ticketId")
	fd_Ticket_eventId = md_Ticket.Fields().ByName("eventId")
	fd_Ticket_ticketType = md_Ticket.Fields().ByName("ticketType")
	fd_Ticket_valid = md_Ticket.Fields().ByName("valid")
	fd_Ticket_owner = md_Ticket.Fields().ByName("owner")
	fd_Ticket_metadata = md_Ticket.Fields().ByName("metadata")
	fd_Ticket_creator = md_Ticket.Fields().ByName("creator")
}

var _ protoreflect.Message = (*fastReflection_Ticket)(nil)

type fastReflection_Ticket Ticket

func (x *Ticket) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Ticket)(x)
}

func (x *Ticket) slowProtoReflect() protoreflect.Message {
	mi := &file_chainvite_ticket_ticket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Ticket_messageType fastReflection_Ticket_messageType
var _ protoreflect.MessageType = fastReflection_Ticket_messageType{}

type fastReflection_Ticket_messageType struct{}

func (x fastReflection_Ticket_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Ticket)(nil)
}
func (x fastReflection_Ticket_messageType) New() protoreflect.Message {
	return new(fastReflection_Ticket)
}
func (x fastReflection_Ticket_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Ticket
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Ticket) Descriptor() protoreflect.MessageDescriptor {
	return md_Ticket
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Ticket) Type() protoreflect.MessageType {
	return _fastReflection_Ticket_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Ticket) New() protoreflect.Message {
	return new(fastReflection_Ticket)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Ticket) Interface() protoreflect.ProtoMessage {
	return (*Ticket)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Ticket) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Index != "" {
		value := protoreflect.ValueOfString(x.Index)
		if !f(fd_Ticket_index, value) {
			return
		}
	}
	if x.TicketId != "" {
		value := protoreflect.ValueOfString(x.TicketId)
		if !f(fd_Ticket_ticketId, value) {
			return
		}
	}
	if x.EventId != "" {
		value := protoreflect.ValueOfString(x.EventId)
		if !f(fd_Ticket_eventId, value) {
			return
		}
	}
	if x.TicketType != "" {
		value := protoreflect.ValueOfString(x.TicketType)
		if !f(fd_Ticket_ticketType, value) {
			return
		}
	}
	if x.Valid != false {
		value := protoreflect.ValueOfBool(x.Valid)
		if !f(fd_Ticket_valid, value) {
			return
		}
	}
	if x.Owner != "" {
		value := protoreflect.ValueOfString(x.Owner)
		if !f(fd_Ticket_owner, value) {
			return
		}
	}
	if x.Metadata != "" {
		value := protoreflect.ValueOfString(x.Metadata)
		if !f(fd_Ticket_metadata, value) {
			return
		}
	}
	if x.Creator != "" {
		value := protoreflect.ValueOfString(x.Creator)
		if !f(fd_Ticket_creator, value) {
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
func (x *fastReflection_Ticket) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "chainvite.ticket.Ticket.index":
		return x.Index != ""
	case "chainvite.ticket.Ticket.ticketId":
		return x.TicketId != ""
	case "chainvite.ticket.Ticket.eventId":
		return x.EventId != ""
	case "chainvite.ticket.Ticket.ticketType":
		return x.TicketType != ""
	case "chainvite.ticket.Ticket.valid":
		return x.Valid != false
	case "chainvite.ticket.Ticket.owner":
		return x.Owner != ""
	case "chainvite.ticket.Ticket.metadata":
		return x.Metadata != ""
	case "chainvite.ticket.Ticket.creator":
		return x.Creator != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: chainvite.ticket.Ticket"))
		}
		panic(fmt.Errorf("message chainvite.ticket.Ticket does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Ticket) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "chainvite.ticket.Ticket.index":
		x.Index = ""
	case "chainvite.ticket.Ticket.ticketId":
		x.TicketId = ""
	case "chainvite.ticket.Ticket.eventId":
		x.EventId = ""
	case "chainvite.ticket.Ticket.ticketType":
		x.TicketType = ""
	case "chainvite.ticket.Ticket.valid":
		x.Valid = false
	case "chainvite.ticket.Ticket.owner":
		x.Owner = ""
	case "chainvite.ticket.Ticket.metadata":
		x.Metadata = ""
	case "chainvite.ticket.Ticket.creator":
		x.Creator = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: chainvite.ticket.Ticket"))
		}
		panic(fmt.Errorf("message chainvite.ticket.Ticket does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Ticket) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "chainvite.ticket.Ticket.index":
		value := x.Index
		return protoreflect.ValueOfString(value)
	case "chainvite.ticket.Ticket.ticketId":
		value := x.TicketId
		return protoreflect.ValueOfString(value)
	case "chainvite.ticket.Ticket.eventId":
		value := x.EventId
		return protoreflect.ValueOfString(value)
	case "chainvite.ticket.Ticket.ticketType":
		value := x.TicketType
		return protoreflect.ValueOfString(value)
	case "chainvite.ticket.Ticket.valid":
		value := x.Valid
		return protoreflect.ValueOfBool(value)
	case "chainvite.ticket.Ticket.owner":
		value := x.Owner
		return protoreflect.ValueOfString(value)
	case "chainvite.ticket.Ticket.metadata":
		value := x.Metadata
		return protoreflect.ValueOfString(value)
	case "chainvite.ticket.Ticket.creator":
		value := x.Creator
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: chainvite.ticket.Ticket"))
		}
		panic(fmt.Errorf("message chainvite.ticket.Ticket does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Ticket) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "chainvite.ticket.Ticket.index":
		x.Index = value.Interface().(string)
	case "chainvite.ticket.Ticket.ticketId":
		x.TicketId = value.Interface().(string)
	case "chainvite.ticket.Ticket.eventId":
		x.EventId = value.Interface().(string)
	case "chainvite.ticket.Ticket.ticketType":
		x.TicketType = value.Interface().(string)
	case "chainvite.ticket.Ticket.valid":
		x.Valid = value.Bool()
	case "chainvite.ticket.Ticket.owner":
		x.Owner = value.Interface().(string)
	case "chainvite.ticket.Ticket.metadata":
		x.Metadata = value.Interface().(string)
	case "chainvite.ticket.Ticket.creator":
		x.Creator = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: chainvite.ticket.Ticket"))
		}
		panic(fmt.Errorf("message chainvite.ticket.Ticket does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Ticket) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "chainvite.ticket.Ticket.index":
		panic(fmt.Errorf("field index of message chainvite.ticket.Ticket is not mutable"))
	case "chainvite.ticket.Ticket.ticketId":
		panic(fmt.Errorf("field ticketId of message chainvite.ticket.Ticket is not mutable"))
	case "chainvite.ticket.Ticket.eventId":
		panic(fmt.Errorf("field eventId of message chainvite.ticket.Ticket is not mutable"))
	case "chainvite.ticket.Ticket.ticketType":
		panic(fmt.Errorf("field ticketType of message chainvite.ticket.Ticket is not mutable"))
	case "chainvite.ticket.Ticket.valid":
		panic(fmt.Errorf("field valid of message chainvite.ticket.Ticket is not mutable"))
	case "chainvite.ticket.Ticket.owner":
		panic(fmt.Errorf("field owner of message chainvite.ticket.Ticket is not mutable"))
	case "chainvite.ticket.Ticket.metadata":
		panic(fmt.Errorf("field metadata of message chainvite.ticket.Ticket is not mutable"))
	case "chainvite.ticket.Ticket.creator":
		panic(fmt.Errorf("field creator of message chainvite.ticket.Ticket is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: chainvite.ticket.Ticket"))
		}
		panic(fmt.Errorf("message chainvite.ticket.Ticket does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Ticket) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "chainvite.ticket.Ticket.index":
		return protoreflect.ValueOfString("")
	case "chainvite.ticket.Ticket.ticketId":
		return protoreflect.ValueOfString("")
	case "chainvite.ticket.Ticket.eventId":
		return protoreflect.ValueOfString("")
	case "chainvite.ticket.Ticket.ticketType":
		return protoreflect.ValueOfString("")
	case "chainvite.ticket.Ticket.valid":
		return protoreflect.ValueOfBool(false)
	case "chainvite.ticket.Ticket.owner":
		return protoreflect.ValueOfString("")
	case "chainvite.ticket.Ticket.metadata":
		return protoreflect.ValueOfString("")
	case "chainvite.ticket.Ticket.creator":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: chainvite.ticket.Ticket"))
		}
		panic(fmt.Errorf("message chainvite.ticket.Ticket does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Ticket) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in chainvite.ticket.Ticket", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Ticket) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Ticket) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Ticket) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Ticket) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Ticket)
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
		l = len(x.Index)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.TicketId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.EventId)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.TicketType)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Valid {
			n += 2
		}
		l = len(x.Owner)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Metadata)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Creator)
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
		x := input.Message.Interface().(*Ticket)
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
		if len(x.Creator) > 0 {
			i -= len(x.Creator)
			copy(dAtA[i:], x.Creator)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Creator)))
			i--
			dAtA[i] = 0x42
		}
		if len(x.Metadata) > 0 {
			i -= len(x.Metadata)
			copy(dAtA[i:], x.Metadata)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Metadata)))
			i--
			dAtA[i] = 0x3a
		}
		if len(x.Owner) > 0 {
			i -= len(x.Owner)
			copy(dAtA[i:], x.Owner)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Owner)))
			i--
			dAtA[i] = 0x32
		}
		if x.Valid {
			i--
			if x.Valid {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x28
		}
		if len(x.TicketType) > 0 {
			i -= len(x.TicketType)
			copy(dAtA[i:], x.TicketType)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TicketType)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.EventId) > 0 {
			i -= len(x.EventId)
			copy(dAtA[i:], x.EventId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.EventId)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.TicketId) > 0 {
			i -= len(x.TicketId)
			copy(dAtA[i:], x.TicketId)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TicketId)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Index) > 0 {
			i -= len(x.Index)
			copy(dAtA[i:], x.Index)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Index)))
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
		x := input.Message.Interface().(*Ticket)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Ticket: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Ticket: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
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
				x.Index = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TicketId", wireType)
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
				x.TicketId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field EventId", wireType)
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
				x.EventId = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TicketType", wireType)
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
				x.TicketType = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Valid", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.Valid = bool(v != 0)
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
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
				x.Owner = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
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
				x.Metadata = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 8:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
				x.Creator = string(dAtA[iNdEx:postIndex])
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
// source: chainvite/ticket/ticket.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Ticket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index      string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	TicketId   string `protobuf:"bytes,2,opt,name=ticketId,proto3" json:"ticketId,omitempty"`
	EventId    string `protobuf:"bytes,3,opt,name=eventId,proto3" json:"eventId,omitempty"`
	TicketType string `protobuf:"bytes,4,opt,name=ticketType,proto3" json:"ticketType,omitempty"`
	Valid      bool   `protobuf:"varint,5,opt,name=valid,proto3" json:"valid,omitempty"`
	Owner      string `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
	Metadata   string `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Creator    string `protobuf:"bytes,8,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (x *Ticket) Reset() {
	*x = Ticket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainvite_ticket_ticket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ticket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ticket) ProtoMessage() {}

// Deprecated: Use Ticket.ProtoReflect.Descriptor instead.
func (*Ticket) Descriptor() ([]byte, []int) {
	return file_chainvite_ticket_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *Ticket) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

func (x *Ticket) GetTicketId() string {
	if x != nil {
		return x.TicketId
	}
	return ""
}

func (x *Ticket) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *Ticket) GetTicketType() string {
	if x != nil {
		return x.TicketType
	}
	return ""
}

func (x *Ticket) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *Ticket) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Ticket) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *Ticket) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

var File_chainvite_ticket_ticket_proto protoreflect.FileDescriptor

var file_chainvite_ticket_ticket_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x2f, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x22, 0xd6, 0x01, 0x0a, 0x06, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x42, 0xb7, 0x01, 0x0a, 0x14, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x2e, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x42, 0x0b, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x75, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x2f, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0xa2, 0x02, 0x03, 0x43, 0x54, 0x58, 0xaa, 0x02, 0x10, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0xca, 0x02,
	0x10, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x5c, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0xe2, 0x02, 0x1c, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x5c, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x11, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x3a, 0x3a, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chainvite_ticket_ticket_proto_rawDescOnce sync.Once
	file_chainvite_ticket_ticket_proto_rawDescData = file_chainvite_ticket_ticket_proto_rawDesc
)

func file_chainvite_ticket_ticket_proto_rawDescGZIP() []byte {
	file_chainvite_ticket_ticket_proto_rawDescOnce.Do(func() {
		file_chainvite_ticket_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_chainvite_ticket_ticket_proto_rawDescData)
	})
	return file_chainvite_ticket_ticket_proto_rawDescData
}

var file_chainvite_ticket_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_chainvite_ticket_ticket_proto_goTypes = []interface{}{
	(*Ticket)(nil), // 0: chainvite.ticket.Ticket
}
var file_chainvite_ticket_ticket_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chainvite_ticket_ticket_proto_init() }
func file_chainvite_ticket_ticket_proto_init() {
	if File_chainvite_ticket_ticket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chainvite_ticket_ticket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ticket); i {
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
			RawDescriptor: file_chainvite_ticket_ticket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chainvite_ticket_ticket_proto_goTypes,
		DependencyIndexes: file_chainvite_ticket_ticket_proto_depIdxs,
		MessageInfos:      file_chainvite_ticket_ticket_proto_msgTypes,
	}.Build()
	File_chainvite_ticket_ticket_proto = out.File
	file_chainvite_ticket_ticket_proto_rawDesc = nil
	file_chainvite_ticket_ticket_proto_goTypes = nil
	file_chainvite_ticket_ticket_proto_depIdxs = nil
}
