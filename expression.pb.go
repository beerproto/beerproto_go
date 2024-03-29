//*
// BeerProto
//
// Another beer format, written in protocol buffer.
//
// Copyright (c) 2020 Ross Merrigan

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: beerproto/v1/expression.proto

package beerprotov1

import (
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

type BinaryArithmetic int32

const (
	BinaryArithmetic_BINARY_ARITHMETIC_UNSPECIFIED    BinaryArithmetic = 0
	BinaryArithmetic_BINARY_ARITHMETIC_ADDITION       BinaryArithmetic = 1
	BinaryArithmetic_BINARY_ARITHMETIC_SUBTRACTION    BinaryArithmetic = 2
	BinaryArithmetic_BINARY_ARITHMETIC_MULTIPLICATION BinaryArithmetic = 3
	BinaryArithmetic_BINARY_ARITHMETIC_DIVISION       BinaryArithmetic = 4
	BinaryArithmetic_BINARY_ARITHMETIC_POWER          BinaryArithmetic = 5
)

// Enum value maps for BinaryArithmetic.
var (
	BinaryArithmetic_name = map[int32]string{
		0: "BINARY_ARITHMETIC_UNSPECIFIED",
		1: "BINARY_ARITHMETIC_ADDITION",
		2: "BINARY_ARITHMETIC_SUBTRACTION",
		3: "BINARY_ARITHMETIC_MULTIPLICATION",
		4: "BINARY_ARITHMETIC_DIVISION",
		5: "BINARY_ARITHMETIC_POWER",
	}
	BinaryArithmetic_value = map[string]int32{
		"BINARY_ARITHMETIC_UNSPECIFIED":    0,
		"BINARY_ARITHMETIC_ADDITION":       1,
		"BINARY_ARITHMETIC_SUBTRACTION":    2,
		"BINARY_ARITHMETIC_MULTIPLICATION": 3,
		"BINARY_ARITHMETIC_DIVISION":       4,
		"BINARY_ARITHMETIC_POWER":          5,
	}
)

func (x BinaryArithmetic) Enum() *BinaryArithmetic {
	p := new(BinaryArithmetic)
	*p = x
	return p
}

func (x BinaryArithmetic) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BinaryArithmetic) Descriptor() protoreflect.EnumDescriptor {
	return file_beerproto_v1_expression_proto_enumTypes[0].Descriptor()
}

func (BinaryArithmetic) Type() protoreflect.EnumType {
	return &file_beerproto_v1_expression_proto_enumTypes[0]
}

func (x BinaryArithmetic) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BinaryArithmetic.Descriptor instead.
func (BinaryArithmetic) EnumDescriptor() ([]byte, []int) {
	return file_beerproto_v1_expression_proto_rawDescGZIP(), []int{0}
}

type ExpressionTree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expression *BinaryExpression `protobuf:"bytes,1,opt,name=expression,proto3" json:"expression,omitempty"`
}

func (x *ExpressionTree) Reset() {
	*x = ExpressionTree{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beerproto_v1_expression_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpressionTree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpressionTree) ProtoMessage() {}

func (x *ExpressionTree) ProtoReflect() protoreflect.Message {
	mi := &file_beerproto_v1_expression_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpressionTree.ProtoReflect.Descriptor instead.
func (*ExpressionTree) Descriptor() ([]byte, []int) {
	return file_beerproto_v1_expression_proto_rawDescGZIP(), []int{0}
}

func (x *ExpressionTree) GetExpression() *BinaryExpression {
	if x != nil {
		return x.Expression
	}
	return nil
}

type UnaryExpression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Double float64 `protobuf:"fixed64,1,opt,name=double,proto3" json:"double,omitempty"`
}

func (x *UnaryExpression) Reset() {
	*x = UnaryExpression{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beerproto_v1_expression_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnaryExpression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnaryExpression) ProtoMessage() {}

func (x *UnaryExpression) ProtoReflect() protoreflect.Message {
	mi := &file_beerproto_v1_expression_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnaryExpression.ProtoReflect.Descriptor instead.
func (*UnaryExpression) Descriptor() ([]byte, []int) {
	return file_beerproto_v1_expression_proto_rawDescGZIP(), []int{1}
}

func (x *UnaryExpression) GetDouble() float64 {
	if x != nil {
		return x.Double
	}
	return 0
}

type BinaryExpression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operator BinaryArithmetic `protobuf:"varint,1,opt,name=operator,proto3,enum=beerproto.v1.BinaryArithmetic" json:"operator,omitempty"`
	// Types that are assignable to Left:
	//
	//	*BinaryExpression_BinaryLeft
	//	*BinaryExpression_UnaryLeft
	//	*BinaryExpression_ParameterLeft
	Left isBinaryExpression_Left `protobuf_oneof:"left"`
	// Types that are assignable to Right:
	//
	//	*BinaryExpression_BinaryRight
	//	*BinaryExpression_UnaryRight
	//	*BinaryExpression_ParameterRight
	Right isBinaryExpression_Right `protobuf_oneof:"right"`
}

func (x *BinaryExpression) Reset() {
	*x = BinaryExpression{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beerproto_v1_expression_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryExpression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryExpression) ProtoMessage() {}

func (x *BinaryExpression) ProtoReflect() protoreflect.Message {
	mi := &file_beerproto_v1_expression_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryExpression.ProtoReflect.Descriptor instead.
func (*BinaryExpression) Descriptor() ([]byte, []int) {
	return file_beerproto_v1_expression_proto_rawDescGZIP(), []int{2}
}

func (x *BinaryExpression) GetOperator() BinaryArithmetic {
	if x != nil {
		return x.Operator
	}
	return BinaryArithmetic_BINARY_ARITHMETIC_UNSPECIFIED
}

func (m *BinaryExpression) GetLeft() isBinaryExpression_Left {
	if m != nil {
		return m.Left
	}
	return nil
}

func (x *BinaryExpression) GetBinaryLeft() *BinaryExpression {
	if x, ok := x.GetLeft().(*BinaryExpression_BinaryLeft); ok {
		return x.BinaryLeft
	}
	return nil
}

func (x *BinaryExpression) GetUnaryLeft() *UnaryExpression {
	if x, ok := x.GetLeft().(*BinaryExpression_UnaryLeft); ok {
		return x.UnaryLeft
	}
	return nil
}

func (x *BinaryExpression) GetParameterLeft() *ParameterExpression {
	if x, ok := x.GetLeft().(*BinaryExpression_ParameterLeft); ok {
		return x.ParameterLeft
	}
	return nil
}

func (m *BinaryExpression) GetRight() isBinaryExpression_Right {
	if m != nil {
		return m.Right
	}
	return nil
}

func (x *BinaryExpression) GetBinaryRight() *BinaryExpression {
	if x, ok := x.GetRight().(*BinaryExpression_BinaryRight); ok {
		return x.BinaryRight
	}
	return nil
}

func (x *BinaryExpression) GetUnaryRight() *UnaryExpression {
	if x, ok := x.GetRight().(*BinaryExpression_UnaryRight); ok {
		return x.UnaryRight
	}
	return nil
}

func (x *BinaryExpression) GetParameterRight() *ParameterExpression {
	if x, ok := x.GetRight().(*BinaryExpression_ParameterRight); ok {
		return x.ParameterRight
	}
	return nil
}

type isBinaryExpression_Left interface {
	isBinaryExpression_Left()
}

type BinaryExpression_BinaryLeft struct {
	BinaryLeft *BinaryExpression `protobuf:"bytes,10,opt,name=binary_left,json=binaryLeft,proto3,oneof"`
}

type BinaryExpression_UnaryLeft struct {
	UnaryLeft *UnaryExpression `protobuf:"bytes,11,opt,name=unary_left,json=unaryLeft,proto3,oneof"`
}

type BinaryExpression_ParameterLeft struct {
	ParameterLeft *ParameterExpression `protobuf:"bytes,12,opt,name=parameter_left,json=parameterLeft,proto3,oneof"`
}

func (*BinaryExpression_BinaryLeft) isBinaryExpression_Left() {}

func (*BinaryExpression_UnaryLeft) isBinaryExpression_Left() {}

func (*BinaryExpression_ParameterLeft) isBinaryExpression_Left() {}

type isBinaryExpression_Right interface {
	isBinaryExpression_Right()
}

type BinaryExpression_BinaryRight struct {
	BinaryRight *BinaryExpression `protobuf:"bytes,20,opt,name=binary_right,json=binaryRight,proto3,oneof"`
}

type BinaryExpression_UnaryRight struct {
	UnaryRight *UnaryExpression `protobuf:"bytes,21,opt,name=unary_right,json=unaryRight,proto3,oneof"`
}

type BinaryExpression_ParameterRight struct {
	ParameterRight *ParameterExpression `protobuf:"bytes,22,opt,name=parameter_right,json=parameterRight,proto3,oneof"`
}

func (*BinaryExpression_BinaryRight) isBinaryExpression_Right() {}

func (*BinaryExpression_UnaryRight) isBinaryExpression_Right() {}

func (*BinaryExpression_ParameterRight) isBinaryExpression_Right() {}

type ParameterExpression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parameter string `protobuf:"bytes,1,opt,name=parameter,proto3" json:"parameter,omitempty"`
}

func (x *ParameterExpression) Reset() {
	*x = ParameterExpression{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beerproto_v1_expression_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParameterExpression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParameterExpression) ProtoMessage() {}

func (x *ParameterExpression) ProtoReflect() protoreflect.Message {
	mi := &file_beerproto_v1_expression_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParameterExpression.ProtoReflect.Descriptor instead.
func (*ParameterExpression) Descriptor() ([]byte, []int) {
	return file_beerproto_v1_expression_proto_rawDescGZIP(), []int{3}
}

func (x *ParameterExpression) GetParameter() string {
	if x != nil {
		return x.Parameter
	}
	return ""
}

var File_beerproto_v1_expression_proto protoreflect.FileDescriptor

var file_beerproto_v1_expression_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x22, 0x50, 0x0a,
	0x0e, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x72, 0x65, 0x65, 0x12,
	0x3e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x29, 0x0a, 0x0f, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x22, 0x83, 0x04, 0x0a, 0x10, 0x42,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x3a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1e, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x41, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69,
	0x63, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x41, 0x0a, 0x0b, 0x62,
	0x69, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x0a, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x4c, 0x65, 0x66, 0x74, 0x12, 0x3e,
	0x0a, 0x0a, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x48, 0x00, 0x52, 0x09, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x4c, 0x65, 0x66, 0x74, 0x12, 0x4a,
	0x0a, 0x0e, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x66, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x45,
	0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0d, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x66, 0x74, 0x12, 0x43, 0x0a, 0x0c, 0x62, 0x69,
	0x6e, 0x61, 0x72, 0x79, 0x5f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x48, 0x01, 0x52, 0x0b, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x40, 0x0a, 0x0b, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x5f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x48, 0x01, 0x52, 0x0a, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x4c, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x5f, 0x72,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x65, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x01, 0x52,
	0x0e, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x52, 0x69, 0x67, 0x68, 0x74, 0x42,
	0x06, 0x0a, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x42, 0x07, 0x0a, 0x05, 0x72, 0x69, 0x67, 0x68, 0x74,
	0x22, 0x33, 0x0a, 0x13, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x45, 0x78, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x2a, 0xdb, 0x01, 0x0a, 0x10, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x41, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x65, 0x74, 0x69, 0x63, 0x12, 0x21, 0x0a, 0x1d, 0x42, 0x49,
	0x4e, 0x41, 0x52, 0x59, 0x5f, 0x41, 0x52, 0x49, 0x54, 0x48, 0x4d, 0x45, 0x54, 0x49, 0x43, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a,
	0x1a, 0x42, 0x49, 0x4e, 0x41, 0x52, 0x59, 0x5f, 0x41, 0x52, 0x49, 0x54, 0x48, 0x4d, 0x45, 0x54,
	0x49, 0x43, 0x5f, 0x41, 0x44, 0x44, 0x49, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x21, 0x0a,
	0x1d, 0x42, 0x49, 0x4e, 0x41, 0x52, 0x59, 0x5f, 0x41, 0x52, 0x49, 0x54, 0x48, 0x4d, 0x45, 0x54,
	0x49, 0x43, 0x5f, 0x53, 0x55, 0x42, 0x54, 0x52, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02,
	0x12, 0x24, 0x0a, 0x20, 0x42, 0x49, 0x4e, 0x41, 0x52, 0x59, 0x5f, 0x41, 0x52, 0x49, 0x54, 0x48,
	0x4d, 0x45, 0x54, 0x49, 0x43, 0x5f, 0x4d, 0x55, 0x4c, 0x54, 0x49, 0x50, 0x4c, 0x49, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x1e, 0x0a, 0x1a, 0x42, 0x49, 0x4e, 0x41, 0x52, 0x59,
	0x5f, 0x41, 0x52, 0x49, 0x54, 0x48, 0x4d, 0x45, 0x54, 0x49, 0x43, 0x5f, 0x44, 0x49, 0x56, 0x49,
	0x53, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x42, 0x49, 0x4e, 0x41, 0x52, 0x59,
	0x5f, 0x41, 0x52, 0x49, 0x54, 0x48, 0x4d, 0x45, 0x54, 0x49, 0x43, 0x5f, 0x50, 0x4f, 0x57, 0x45,
	0x52, 0x10, 0x05, 0x42, 0xb0, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x65, 0x65, 0x72,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x6f, 0x2f, 0x62,
	0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x65, 0x65, 0x72,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x0c,
	0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x42,
	0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x42, 0x65,
	0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_beerproto_v1_expression_proto_rawDescOnce sync.Once
	file_beerproto_v1_expression_proto_rawDescData = file_beerproto_v1_expression_proto_rawDesc
)

func file_beerproto_v1_expression_proto_rawDescGZIP() []byte {
	file_beerproto_v1_expression_proto_rawDescOnce.Do(func() {
		file_beerproto_v1_expression_proto_rawDescData = protoimpl.X.CompressGZIP(file_beerproto_v1_expression_proto_rawDescData)
	})
	return file_beerproto_v1_expression_proto_rawDescData
}

var file_beerproto_v1_expression_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_beerproto_v1_expression_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_beerproto_v1_expression_proto_goTypes = []interface{}{
	(BinaryArithmetic)(0),       // 0: beerproto.v1.BinaryArithmetic
	(*ExpressionTree)(nil),      // 1: beerproto.v1.ExpressionTree
	(*UnaryExpression)(nil),     // 2: beerproto.v1.UnaryExpression
	(*BinaryExpression)(nil),    // 3: beerproto.v1.BinaryExpression
	(*ParameterExpression)(nil), // 4: beerproto.v1.ParameterExpression
}
var file_beerproto_v1_expression_proto_depIdxs = []int32{
	3, // 0: beerproto.v1.ExpressionTree.expression:type_name -> beerproto.v1.BinaryExpression
	0, // 1: beerproto.v1.BinaryExpression.operator:type_name -> beerproto.v1.BinaryArithmetic
	3, // 2: beerproto.v1.BinaryExpression.binary_left:type_name -> beerproto.v1.BinaryExpression
	2, // 3: beerproto.v1.BinaryExpression.unary_left:type_name -> beerproto.v1.UnaryExpression
	4, // 4: beerproto.v1.BinaryExpression.parameter_left:type_name -> beerproto.v1.ParameterExpression
	3, // 5: beerproto.v1.BinaryExpression.binary_right:type_name -> beerproto.v1.BinaryExpression
	2, // 6: beerproto.v1.BinaryExpression.unary_right:type_name -> beerproto.v1.UnaryExpression
	4, // 7: beerproto.v1.BinaryExpression.parameter_right:type_name -> beerproto.v1.ParameterExpression
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_beerproto_v1_expression_proto_init() }
func file_beerproto_v1_expression_proto_init() {
	if File_beerproto_v1_expression_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_beerproto_v1_expression_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpressionTree); i {
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
		file_beerproto_v1_expression_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnaryExpression); i {
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
		file_beerproto_v1_expression_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryExpression); i {
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
		file_beerproto_v1_expression_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParameterExpression); i {
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
	file_beerproto_v1_expression_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*BinaryExpression_BinaryLeft)(nil),
		(*BinaryExpression_UnaryLeft)(nil),
		(*BinaryExpression_ParameterLeft)(nil),
		(*BinaryExpression_BinaryRight)(nil),
		(*BinaryExpression_UnaryRight)(nil),
		(*BinaryExpression_ParameterRight)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_beerproto_v1_expression_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_beerproto_v1_expression_proto_goTypes,
		DependencyIndexes: file_beerproto_v1_expression_proto_depIdxs,
		EnumInfos:         file_beerproto_v1_expression_proto_enumTypes,
		MessageInfos:      file_beerproto_v1_expression_proto_msgTypes,
	}.Build()
	File_beerproto_v1_expression_proto = out.File
	file_beerproto_v1_expression_proto_rawDesc = nil
	file_beerproto_v1_expression_proto_goTypes = nil
	file_beerproto_v1_expression_proto_depIdxs = nil
}
