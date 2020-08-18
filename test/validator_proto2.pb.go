// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: test/validator_proto2.proto

package validatortest

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type EnumProto2 int32

const (
	EnumProto2_alpha2 EnumProto2 = 0
	EnumProto2_beta2  EnumProto2 = 1
)

// Enum value maps for EnumProto2.
var (
	EnumProto2_name = map[int32]string{
		0: "alpha2",
		1: "beta2",
	}
	EnumProto2_value = map[string]int32{
		"alpha2": 0,
		"beta2":  1,
	}
)

func (x EnumProto2) Enum() *EnumProto2 {
	p := new(EnumProto2)
	*p = x
	return p
}

func (x EnumProto2) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumProto2) Descriptor() protoreflect.EnumDescriptor {
	return file_test_validator_proto2_proto_enumTypes[0].Descriptor()
}

func (EnumProto2) Type() protoreflect.EnumType {
	return &file_test_validator_proto2_proto_enumTypes[0]
}

func (x EnumProto2) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EnumProto2) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EnumProto2(num)
	return nil
}

// Deprecated: Use EnumProto2.Descriptor instead.
func (EnumProto2) EnumDescriptor() ([]byte, []int) {
	return file_test_validator_proto2_proto_rawDescGZIP(), []int{0}
}

type ValidatorMessage_EmbeddedEnum int32

const (
	ValidatorMessage_zero ValidatorMessage_EmbeddedEnum = 0
	ValidatorMessage_one  ValidatorMessage_EmbeddedEnum = 1
)

// Enum value maps for ValidatorMessage_EmbeddedEnum.
var (
	ValidatorMessage_EmbeddedEnum_name = map[int32]string{
		0: "zero",
		1: "one",
	}
	ValidatorMessage_EmbeddedEnum_value = map[string]int32{
		"zero": 0,
		"one":  1,
	}
)

func (x ValidatorMessage_EmbeddedEnum) Enum() *ValidatorMessage_EmbeddedEnum {
	p := new(ValidatorMessage_EmbeddedEnum)
	*p = x
	return p
}

func (x ValidatorMessage_EmbeddedEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ValidatorMessage_EmbeddedEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_test_validator_proto2_proto_enumTypes[1].Descriptor()
}

func (ValidatorMessage_EmbeddedEnum) Type() protoreflect.EnumType {
	return &file_test_validator_proto2_proto_enumTypes[1]
}

func (x ValidatorMessage_EmbeddedEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ValidatorMessage_EmbeddedEnum) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ValidatorMessage_EmbeddedEnum(num)
	return nil
}

// Deprecated: Use ValidatorMessage_EmbeddedEnum.Descriptor instead.
func (ValidatorMessage_EmbeddedEnum) EnumDescriptor() ([]byte, []int) {
	return file_test_validator_proto2_proto_rawDescGZIP(), []int{0, 0}
}

type ValidatorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StringReq           *string                             `protobuf:"bytes,1,req,name=StringReq" json:"StringReq,omitempty"`
	StringOpt           *string                             `protobuf:"bytes,2,opt,name=StringOpt" json:"StringOpt,omitempty"`
	StringNoQuotes      *string                             `protobuf:"bytes,3,req,name=StringNoQuotes" json:"StringNoQuotes,omitempty"`
	StringUnescaped     *string                             `protobuf:"bytes,4,req,name=StringUnescaped" json:"StringUnescaped,omitempty"`
	IntReq              *uint32                             `protobuf:"varint,5,req,name=IntReq" json:"IntReq,omitempty"`
	IntRep              []uint32                            `protobuf:"varint,6,rep,name=IntRep" json:"IntRep,omitempty"`
	EmbeddedReq         *ValidatorMessage_EmbeddedMessage   `protobuf:"bytes,7,req,name=embeddedReq" json:"embeddedReq,omitempty"`
	EmbeddedRep         []*ValidatorMessage_EmbeddedMessage `protobuf:"bytes,8,rep,name=embeddedRep" json:"embeddedRep,omitempty"`
	CustomErrorInt      *int32                              `protobuf:"varint,9,opt,name=CustomErrorInt" json:"CustomErrorInt,omitempty"`
	StrictSomeDoubleReq *float64                            `protobuf:"fixed64,10,req,name=StrictSomeDoubleReq" json:"StrictSomeDoubleReq,omitempty"`
	StrictSomeDoubleRep []float64                           `protobuf:"fixed64,11,rep,name=StrictSomeDoubleRep" json:"StrictSomeDoubleRep,omitempty"`
	StrictSomeFloatReq  *float32                            `protobuf:"fixed32,12,req,name=StrictSomeFloatReq" json:"StrictSomeFloatReq,omitempty"`
	StrictSomeFloatRep  []float32                           `protobuf:"fixed32,13,rep,name=StrictSomeFloatRep" json:"StrictSomeFloatRep,omitempty"`
	SomeDoubleReq       *float64                            `protobuf:"fixed64,14,req,name=SomeDoubleReq" json:"SomeDoubleReq,omitempty"`
	SomeDoubleRep       []float64                           `protobuf:"fixed64,15,rep,name=SomeDoubleRep" json:"SomeDoubleRep,omitempty"`
	SomeFloatReq        *float32                            `protobuf:"fixed32,16,req,name=SomeFloatReq" json:"SomeFloatReq,omitempty"`
	SomeFloatRep        []float32                           `protobuf:"fixed32,17,rep,name=SomeFloatRep" json:"SomeFloatRep,omitempty"`
	SomeNonEmptyString  *string                             `protobuf:"bytes,18,req,name=SomeNonEmptyString" json:"SomeNonEmptyString,omitempty"`
	RepeatedBaseType    []int32                             `protobuf:"varint,19,rep,name=RepeatedBaseType" json:"RepeatedBaseType,omitempty"`
	Repeated            []int32                             `protobuf:"varint,20,rep,name=Repeated" json:"Repeated,omitempty"`
	SomeStringLtReq     *string                             `protobuf:"bytes,21,opt,name=SomeStringLtReq" json:"SomeStringLtReq,omitempty"`
	SomeStringGtReq     *string                             `protobuf:"bytes,22,opt,name=SomeStringGtReq" json:"SomeStringGtReq,omitempty"`
	SomeStringEqReq     *string                             `protobuf:"bytes,23,opt,name=SomeStringEqReq" json:"SomeStringEqReq,omitempty"`
	SomeBytesLtReq      []byte                              `protobuf:"bytes,24,opt,name=SomeBytesLtReq" json:"SomeBytesLtReq,omitempty"`
	SomeBytesGtReq      []byte                              `protobuf:"bytes,25,opt,name=SomeBytesGtReq" json:"SomeBytesGtReq,omitempty"`
	SomeBytesEqReq      []byte                              `protobuf:"bytes,26,opt,name=SomeBytesEqReq" json:"SomeBytesEqReq,omitempty"`
	UUIDAny             *string                             `protobuf:"bytes,27,opt,name=UUIDAny" json:"UUIDAny,omitempty"`
	UUID4NotEmpty       *string                             `protobuf:"bytes,28,req,name=UUID4NotEmpty" json:"UUID4NotEmpty,omitempty"`
	SomeEnum            *EnumProto2                         `protobuf:"varint,29,req,name=someEnum,enum=validatortest.EnumProto2" json:"someEnum,omitempty"`
	SomeEmbeddedEnum    *ValidatorMessage_EmbeddedEnum      `protobuf:"varint,30,req,name=someEmbeddedEnum,enum=validatortest.ValidatorMessage_EmbeddedEnum" json:"someEmbeddedEnum,omitempty"`
}

func (x *ValidatorMessage) Reset() {
	*x = ValidatorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_validator_proto2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorMessage) ProtoMessage() {}

func (x *ValidatorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_test_validator_proto2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatorMessage.ProtoReflect.Descriptor instead.
func (*ValidatorMessage) Descriptor() ([]byte, []int) {
	return file_test_validator_proto2_proto_rawDescGZIP(), []int{0}
}

func (x *ValidatorMessage) GetStringReq() string {
	if x != nil && x.StringReq != nil {
		return *x.StringReq
	}
	return ""
}

func (x *ValidatorMessage) GetStringOpt() string {
	if x != nil && x.StringOpt != nil {
		return *x.StringOpt
	}
	return ""
}

func (x *ValidatorMessage) GetStringNoQuotes() string {
	if x != nil && x.StringNoQuotes != nil {
		return *x.StringNoQuotes
	}
	return ""
}

func (x *ValidatorMessage) GetStringUnescaped() string {
	if x != nil && x.StringUnescaped != nil {
		return *x.StringUnescaped
	}
	return ""
}

func (x *ValidatorMessage) GetIntReq() uint32 {
	if x != nil && x.IntReq != nil {
		return *x.IntReq
	}
	return 0
}

func (x *ValidatorMessage) GetIntRep() []uint32 {
	if x != nil {
		return x.IntRep
	}
	return nil
}

func (x *ValidatorMessage) GetEmbeddedReq() *ValidatorMessage_EmbeddedMessage {
	if x != nil {
		return x.EmbeddedReq
	}
	return nil
}

func (x *ValidatorMessage) GetEmbeddedRep() []*ValidatorMessage_EmbeddedMessage {
	if x != nil {
		return x.EmbeddedRep
	}
	return nil
}

func (x *ValidatorMessage) GetCustomErrorInt() int32 {
	if x != nil && x.CustomErrorInt != nil {
		return *x.CustomErrorInt
	}
	return 0
}

func (x *ValidatorMessage) GetStrictSomeDoubleReq() float64 {
	if x != nil && x.StrictSomeDoubleReq != nil {
		return *x.StrictSomeDoubleReq
	}
	return 0
}

func (x *ValidatorMessage) GetStrictSomeDoubleRep() []float64 {
	if x != nil {
		return x.StrictSomeDoubleRep
	}
	return nil
}

func (x *ValidatorMessage) GetStrictSomeFloatReq() float32 {
	if x != nil && x.StrictSomeFloatReq != nil {
		return *x.StrictSomeFloatReq
	}
	return 0
}

func (x *ValidatorMessage) GetStrictSomeFloatRep() []float32 {
	if x != nil {
		return x.StrictSomeFloatRep
	}
	return nil
}

func (x *ValidatorMessage) GetSomeDoubleReq() float64 {
	if x != nil && x.SomeDoubleReq != nil {
		return *x.SomeDoubleReq
	}
	return 0
}

func (x *ValidatorMessage) GetSomeDoubleRep() []float64 {
	if x != nil {
		return x.SomeDoubleRep
	}
	return nil
}

func (x *ValidatorMessage) GetSomeFloatReq() float32 {
	if x != nil && x.SomeFloatReq != nil {
		return *x.SomeFloatReq
	}
	return 0
}

func (x *ValidatorMessage) GetSomeFloatRep() []float32 {
	if x != nil {
		return x.SomeFloatRep
	}
	return nil
}

func (x *ValidatorMessage) GetSomeNonEmptyString() string {
	if x != nil && x.SomeNonEmptyString != nil {
		return *x.SomeNonEmptyString
	}
	return ""
}

func (x *ValidatorMessage) GetRepeatedBaseType() []int32 {
	if x != nil {
		return x.RepeatedBaseType
	}
	return nil
}

func (x *ValidatorMessage) GetRepeated() []int32 {
	if x != nil {
		return x.Repeated
	}
	return nil
}

func (x *ValidatorMessage) GetSomeStringLtReq() string {
	if x != nil && x.SomeStringLtReq != nil {
		return *x.SomeStringLtReq
	}
	return ""
}

func (x *ValidatorMessage) GetSomeStringGtReq() string {
	if x != nil && x.SomeStringGtReq != nil {
		return *x.SomeStringGtReq
	}
	return ""
}

func (x *ValidatorMessage) GetSomeStringEqReq() string {
	if x != nil && x.SomeStringEqReq != nil {
		return *x.SomeStringEqReq
	}
	return ""
}

func (x *ValidatorMessage) GetSomeBytesLtReq() []byte {
	if x != nil {
		return x.SomeBytesLtReq
	}
	return nil
}

func (x *ValidatorMessage) GetSomeBytesGtReq() []byte {
	if x != nil {
		return x.SomeBytesGtReq
	}
	return nil
}

func (x *ValidatorMessage) GetSomeBytesEqReq() []byte {
	if x != nil {
		return x.SomeBytesEqReq
	}
	return nil
}

func (x *ValidatorMessage) GetUUIDAny() string {
	if x != nil && x.UUIDAny != nil {
		return *x.UUIDAny
	}
	return ""
}

func (x *ValidatorMessage) GetUUID4NotEmpty() string {
	if x != nil && x.UUID4NotEmpty != nil {
		return *x.UUID4NotEmpty
	}
	return ""
}

func (x *ValidatorMessage) GetSomeEnum() EnumProto2 {
	if x != nil && x.SomeEnum != nil {
		return *x.SomeEnum
	}
	return EnumProto2_alpha2
}

func (x *ValidatorMessage) GetSomeEmbeddedEnum() ValidatorMessage_EmbeddedEnum {
	if x != nil && x.SomeEmbeddedEnum != nil {
		return *x.SomeEmbeddedEnum
	}
	return ValidatorMessage_zero
}

type ValidatorMessage_EmbeddedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier *string `protobuf:"bytes,1,opt,name=Identifier" json:"Identifier,omitempty"`
	SomeValue  *int64  `protobuf:"varint,2,req,name=SomeValue" json:"SomeValue,omitempty"`
}

func (x *ValidatorMessage_EmbeddedMessage) Reset() {
	*x = ValidatorMessage_EmbeddedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_validator_proto2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidatorMessage_EmbeddedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidatorMessage_EmbeddedMessage) ProtoMessage() {}

func (x *ValidatorMessage_EmbeddedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_test_validator_proto2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidatorMessage_EmbeddedMessage.ProtoReflect.Descriptor instead.
func (*ValidatorMessage_EmbeddedMessage) Descriptor() ([]byte, []int) {
	return file_test_validator_proto2_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ValidatorMessage_EmbeddedMessage) GetIdentifier() string {
	if x != nil && x.Identifier != nil {
		return *x.Identifier
	}
	return ""
}

func (x *ValidatorMessage_EmbeddedMessage) GetSomeValue() int64 {
	if x != nil && x.SomeValue != nil {
		return *x.SomeValue
	}
	return 0
}

var File_test_validator_proto2_proto protoreflect.FileDescriptor

var file_test_validator_proto2_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77,
	0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x0f, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x42, 0x0e, 0xe2,
	0xdf, 0x1f, 0x0a, 0x0a, 0x08, 0x5e, 0x2e, 0x7b, 0x32, 0x2c, 0x35, 0x7d, 0x24, 0x52, 0x09, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x2c, 0x0a, 0x09, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x4f, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0e, 0xe2, 0xdf, 0x1f,
	0x0a, 0x0a, 0x08, 0x5e, 0x2e, 0x7b, 0x32, 0x2c, 0x35, 0x7d, 0x24, 0x52, 0x09, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x4f, 0x70, 0x74, 0x12, 0x39, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x4e, 0x6f, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x02, 0x28, 0x09, 0x42, 0x11,
	0xe2, 0xdf, 0x1f, 0x0d, 0x0a, 0x0b, 0x5e, 0x5b, 0x5e, 0x22, 0x5d, 0x7b, 0x32, 0x2c, 0x35, 0x7d,
	0x24, 0x52, 0x0e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x73, 0x12, 0x61, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x55, 0x6e, 0x65, 0x73, 0x63,
	0x61, 0x70, 0x65, 0x64, 0x18, 0x04, 0x20, 0x02, 0x28, 0x09, 0x42, 0x37, 0xe2, 0xdf, 0x1f, 0x33,
	0x0a, 0x31, 0x5b, 0x5c, 0x70, 0x7b, 0x4c, 0x7d, 0x5c, 0x70, 0x7b, 0x4e, 0x7d, 0x5d, 0x28, 0x7b,
	0x5c, 0x70, 0x7b, 0x4c, 0x7d, 0x5c, 0x70, 0x7b, 0x4e, 0x7d, 0x5f, 0x2d, 0x20, 0x5d, 0x7b, 0x30,
	0x2c, 0x32, 0x38, 0x7d, 0x5b, 0x5c, 0x70, 0x7b, 0x4c, 0x7d, 0x5c, 0x70, 0x7b, 0x4e, 0x7d, 0x5d,
	0x29, 0x3f, 0x2e, 0x52, 0x0f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x55, 0x6e, 0x65, 0x73, 0x63,
	0x61, 0x70, 0x65, 0x64, 0x12, 0x1e, 0x0a, 0x06, 0x49, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x18, 0x05,
	0x20, 0x02, 0x28, 0x0d, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10, 0x0a, 0x52, 0x06, 0x49, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x06, 0x49, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0d, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10, 0x0a, 0x52, 0x06, 0x49, 0x6e,
	0x74, 0x52, 0x65, 0x70, 0x12, 0x51, 0x0a, 0x0b, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64,
	0x52, 0x65, 0x71, 0x18, 0x07, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64,
	0x64, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x65, 0x6d, 0x62, 0x65,
	0x64, 0x64, 0x65, 0x64, 0x52, 0x65, 0x71, 0x12, 0x51, 0x0a, 0x0b, 0x65, 0x6d, 0x62, 0x65, 0x64,
	0x64, 0x65, 0x64, 0x52, 0x65, 0x70, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6d,
	0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x65,
	0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x52, 0x65, 0x70, 0x12, 0x3f, 0x0a, 0x0e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x17, 0xe2, 0xdf, 0x1f, 0x13, 0x10, 0x0a, 0x2a, 0x0f, 0x4d, 0x79, 0x20, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x20, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x0e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x74, 0x12, 0x51, 0x0a, 0x13, 0x53,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x53, 0x6f, 0x6d, 0x65, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x18, 0x0a, 0x20, 0x02, 0x28, 0x01, 0x42, 0x1f, 0xe2, 0xdf, 0x1f, 0x1b, 0x31, 0x66,
	0x66, 0x66, 0x66, 0x66, 0x66, 0xd6, 0x3f, 0x39, 0xcd, 0xcc, 0xcc, 0xcc, 0xcc, 0xcc, 0xe4, 0x3f,
	0x41, 0x9a, 0x99, 0x99, 0x99, 0x99, 0x99, 0xa9, 0x3f, 0x52, 0x13, 0x53, 0x74, 0x72, 0x69, 0x63,
	0x74, 0x53, 0x6f, 0x6d, 0x65, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x51,
	0x0a, 0x13, 0x53, 0x74, 0x72, 0x69, 0x63, 0x74, 0x53, 0x6f, 0x6d, 0x65, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x52, 0x65, 0x70, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x01, 0x42, 0x1f, 0xe2, 0xdf, 0x1f,
	0x1b, 0x31, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0xd6, 0x3f, 0x39, 0xcd, 0xcc, 0xcc, 0xcc, 0xcc,
	0xcc, 0xe4, 0x3f, 0x41, 0x9a, 0x99, 0x99, 0x99, 0x99, 0x99, 0xa9, 0x3f, 0x52, 0x13, 0x53, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x53, 0x6f, 0x6d, 0x65, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x52, 0x65,
	0x70, 0x12, 0x4f, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x69, 0x63, 0x74, 0x53, 0x6f, 0x6d, 0x65, 0x46,
	0x6c, 0x6f, 0x61, 0x74, 0x52, 0x65, 0x71, 0x18, 0x0c, 0x20, 0x02, 0x28, 0x02, 0x42, 0x1f, 0xe2,
	0xdf, 0x1f, 0x1b, 0x31, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0xd6, 0x3f, 0x39, 0xcd, 0xcc, 0xcc,
	0xcc, 0xcc, 0xcc, 0xe4, 0x3f, 0x41, 0x9a, 0x99, 0x99, 0x99, 0x99, 0x99, 0xa9, 0x3f, 0x52, 0x12,
	0x53, 0x74, 0x72, 0x69, 0x63, 0x74, 0x53, 0x6f, 0x6d, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x4f, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x69, 0x63, 0x74, 0x53, 0x6f, 0x6d, 0x65,
	0x46, 0x6c, 0x6f, 0x61, 0x74, 0x52, 0x65, 0x70, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x02, 0x42, 0x1f,
	0xe2, 0xdf, 0x1f, 0x1b, 0x31, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0xd6, 0x3f, 0x39, 0xcd, 0xcc,
	0xcc, 0xcc, 0xcc, 0xcc, 0xe4, 0x3f, 0x41, 0x9a, 0x99, 0x99, 0x99, 0x99, 0x99, 0xa9, 0x3f, 0x52,
	0x12, 0x53, 0x74, 0x72, 0x69, 0x63, 0x74, 0x53, 0x6f, 0x6d, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74,
	0x52, 0x65, 0x70, 0x12, 0x3c, 0x0a, 0x0d, 0x53, 0x6f, 0x6d, 0x65, 0x44, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x18, 0x0e, 0x20, 0x02, 0x28, 0x01, 0x42, 0x16, 0xe2, 0xdf, 0x1f, 0x12,
	0x49, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xd0, 0x3f, 0x51, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0xe8, 0x3f, 0x52, 0x0d, 0x53, 0x6f, 0x6d, 0x65, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x3c, 0x0a, 0x0d, 0x53, 0x6f, 0x6d, 0x65, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x70, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x01, 0x42, 0x16, 0xe2, 0xdf, 0x1f, 0x12, 0x49, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0xd0, 0x3f, 0x51, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe8, 0x3f,
	0x52, 0x0d, 0x53, 0x6f, 0x6d, 0x65, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x12,
	0x3a, 0x0a, 0x0c, 0x53, 0x6f, 0x6d, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x52, 0x65, 0x71, 0x18,
	0x10, 0x20, 0x02, 0x28, 0x02, 0x42, 0x16, 0xe2, 0xdf, 0x1f, 0x12, 0x49, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0xd0, 0x3f, 0x51, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe8, 0x3f, 0x52, 0x0c, 0x53,
	0x6f, 0x6d, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x3a, 0x0a, 0x0c, 0x53,
	0x6f, 0x6d, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x52, 0x65, 0x70, 0x18, 0x11, 0x20, 0x03, 0x28,
	0x02, 0x42, 0x16, 0xe2, 0xdf, 0x1f, 0x12, 0x49, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xd0, 0x3f,
	0x51, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe8, 0x3f, 0x52, 0x0c, 0x53, 0x6f, 0x6d, 0x65, 0x46,
	0x6c, 0x6f, 0x61, 0x74, 0x52, 0x65, 0x70, 0x12, 0x36, 0x0a, 0x12, 0x53, 0x6f, 0x6d, 0x65, 0x4e,
	0x6f, 0x6e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x12, 0x20,
	0x02, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x12, 0x53, 0x6f, 0x6d,
	0x65, 0x4e, 0x6f, 0x6e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x2a, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x61, 0x73, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x13, 0x20, 0x03, 0x28, 0x05, 0x52, 0x10, 0x52, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x61, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x52,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x14, 0x20, 0x03, 0x28, 0x05, 0x42, 0x08, 0xe2,
	0xdf, 0x1f, 0x04, 0x60, 0x02, 0x68, 0x05, 0x52, 0x08, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x30, 0x0a, 0x0f, 0x53, 0x6f, 0x6d, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c,
	0x74, 0x52, 0x65, 0x71, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02,
	0x70, 0x02, 0x52, 0x0f, 0x53, 0x6f, 0x6d, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x30, 0x0a, 0x0f, 0x53, 0x6f, 0x6d, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x47, 0x74, 0x52, 0x65, 0x71, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf,
	0x1f, 0x02, 0x78, 0x0c, 0x52, 0x0f, 0x53, 0x6f, 0x6d, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x47, 0x74, 0x52, 0x65, 0x71, 0x12, 0x31, 0x0a, 0x0f, 0x53, 0x6f, 0x6d, 0x65, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x45, 0x71, 0x52, 0x65, 0x71, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x0a, 0x52, 0x0f, 0x53, 0x6f, 0x6d, 0x65, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x45, 0x71, 0x52, 0x65, 0x71, 0x12, 0x2e, 0x0a, 0x0e, 0x53, 0x6f, 0x6d, 0x65,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x4c, 0x74, 0x52, 0x65, 0x71, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0c,
	0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x70, 0x05, 0x52, 0x0e, 0x53, 0x6f, 0x6d, 0x65, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x4c, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2e, 0x0a, 0x0e, 0x53, 0x6f, 0x6d, 0x65,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x47, 0x74, 0x52, 0x65, 0x71, 0x18, 0x19, 0x20, 0x01, 0x28, 0x0c,
	0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x78, 0x14, 0x52, 0x0e, 0x53, 0x6f, 0x6d, 0x65, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x47, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2f, 0x0a, 0x0e, 0x53, 0x6f, 0x6d, 0x65,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x45, 0x71, 0x52, 0x65, 0x71, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x0c,
	0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x80, 0x01, 0x0c, 0x52, 0x0e, 0x53, 0x6f, 0x6d, 0x65, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x45, 0x71, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x07, 0x55, 0x55, 0x49,
	0x44, 0x41, 0x6e, 0x79, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03,
	0x90, 0x01, 0x00, 0x52, 0x07, 0x55, 0x55, 0x49, 0x44, 0x41, 0x6e, 0x79, 0x12, 0x2d, 0x0a, 0x0d,
	0x55, 0x55, 0x49, 0x44, 0x34, 0x4e, 0x6f, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x1c, 0x20,
	0x02, 0x28, 0x09, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x90, 0x01, 0x04, 0x52, 0x0d, 0x55, 0x55,
	0x49, 0x44, 0x34, 0x4e, 0x6f, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3e, 0x0a, 0x08, 0x73,
	0x6f, 0x6d, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x18, 0x1d, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x19, 0x2e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x88, 0x01,
	0x01, 0x52, 0x08, 0x73, 0x6f, 0x6d, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x61, 0x0a, 0x10, 0x73,
	0x6f, 0x6d, 0x65, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x18,
	0x1e, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x45,
	0x6e, 0x75, 0x6d, 0x42, 0x07, 0xe2, 0xdf, 0x1f, 0x03, 0x88, 0x01, 0x01, 0x52, 0x10, 0x73, 0x6f,
	0x6d, 0x65, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x1a, 0x6d,
	0x0a, 0x0f, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x32, 0x0a, 0x0a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0xe2, 0xdf, 0x1f, 0x0e, 0x0a, 0x0c, 0x5e, 0x5b, 0x61,
	0x2d, 0x7a, 0x5d, 0x7b, 0x32, 0x2c, 0x35, 0x7d, 0x24, 0x52, 0x0a, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x09, 0x53, 0x6f, 0x6d, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x03, 0x42, 0x08, 0xe2, 0xdf, 0x1f, 0x04, 0x10, 0x00,
	0x18, 0x64, 0x52, 0x09, 0x53, 0x6f, 0x6d, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x21, 0x0a,
	0x0c, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x08, 0x0a,
	0x04, 0x7a, 0x65, 0x72, 0x6f, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x6f, 0x6e, 0x65, 0x10, 0x01,
	0x2a, 0x23, 0x0a, 0x0a, 0x45, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x12, 0x0a,
	0x0a, 0x06, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x62, 0x65,
	0x74, 0x61, 0x32, 0x10, 0x01,
}

var (
	file_test_validator_proto2_proto_rawDescOnce sync.Once
	file_test_validator_proto2_proto_rawDescData = file_test_validator_proto2_proto_rawDesc
)

func file_test_validator_proto2_proto_rawDescGZIP() []byte {
	file_test_validator_proto2_proto_rawDescOnce.Do(func() {
		file_test_validator_proto2_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_validator_proto2_proto_rawDescData)
	})
	return file_test_validator_proto2_proto_rawDescData
}

var file_test_validator_proto2_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_test_validator_proto2_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test_validator_proto2_proto_goTypes = []interface{}{
	(EnumProto2)(0),                          // 0: validatortest.EnumProto2
	(ValidatorMessage_EmbeddedEnum)(0),       // 1: validatortest.ValidatorMessage.EmbeddedEnum
	(*ValidatorMessage)(nil),                 // 2: validatortest.ValidatorMessage
	(*ValidatorMessage_EmbeddedMessage)(nil), // 3: validatortest.ValidatorMessage.EmbeddedMessage
}
var file_test_validator_proto2_proto_depIdxs = []int32{
	3, // 0: validatortest.ValidatorMessage.embeddedReq:type_name -> validatortest.ValidatorMessage.EmbeddedMessage
	3, // 1: validatortest.ValidatorMessage.embeddedRep:type_name -> validatortest.ValidatorMessage.EmbeddedMessage
	0, // 2: validatortest.ValidatorMessage.someEnum:type_name -> validatortest.EnumProto2
	1, // 3: validatortest.ValidatorMessage.someEmbeddedEnum:type_name -> validatortest.ValidatorMessage.EmbeddedEnum
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_test_validator_proto2_proto_init() }
func file_test_validator_proto2_proto_init() {
	if File_test_validator_proto2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_validator_proto2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorMessage); i {
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
		file_test_validator_proto2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidatorMessage_EmbeddedMessage); i {
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
			RawDescriptor: file_test_validator_proto2_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_validator_proto2_proto_goTypes,
		DependencyIndexes: file_test_validator_proto2_proto_depIdxs,
		EnumInfos:         file_test_validator_proto2_proto_enumTypes,
		MessageInfos:      file_test_validator_proto2_proto_msgTypes,
	}.Build()
	File_test_validator_proto2_proto = out.File
	file_test_validator_proto2_proto_rawDesc = nil
	file_test_validator_proto2_proto_goTypes = nil
	file_test_validator_proto2_proto_depIdxs = nil
}