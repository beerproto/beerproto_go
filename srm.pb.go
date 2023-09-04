//*
// BeerProto
//
// Another beer format, written in protocol buffer.
//
// Copyright (c) 2020 Ross Merrigan

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: beerproto/v1/srm.proto

package beerprotov1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SRM int32

const (
	SRM_SRM_UNSPECIFIED SRM = 0
	SRM_SRM_1           SRM = 1
	SRM_SRM_2           SRM = 2
	SRM_SRM_3           SRM = 3
	SRM_SRM_4           SRM = 4
	SRM_SRM_5           SRM = 5
	SRM_SRM_6           SRM = 6
	SRM_SRM_7           SRM = 7
	SRM_SRM_8           SRM = 8
	SRM_SRM_9           SRM = 9
	SRM_SRM_10          SRM = 10
	SRM_SRM_11          SRM = 11
	SRM_SRM_12          SRM = 12
	SRM_SRM_13          SRM = 13
	SRM_SRM_14          SRM = 14
	SRM_SRM_15          SRM = 15
	SRM_SRM_16          SRM = 16
	SRM_SRM_17          SRM = 17
	SRM_SRM_18          SRM = 18
	SRM_SRM_19          SRM = 19
	SRM_SRM_20          SRM = 20
	SRM_SRM_21          SRM = 21
	SRM_SRM_22          SRM = 22
	SRM_SRM_23          SRM = 23
	SRM_SRM_24          SRM = 24
	SRM_SRM_25          SRM = 25
	SRM_SRM_26          SRM = 26
	SRM_SRM_27          SRM = 27
	SRM_SRM_28          SRM = 28
	SRM_SRM_29          SRM = 29
	SRM_SRM_30          SRM = 30
)

// Enum value maps for SRM.
var (
	SRM_name = map[int32]string{
		0:  "SRM_UNSPECIFIED",
		1:  "SRM_1",
		2:  "SRM_2",
		3:  "SRM_3",
		4:  "SRM_4",
		5:  "SRM_5",
		6:  "SRM_6",
		7:  "SRM_7",
		8:  "SRM_8",
		9:  "SRM_9",
		10: "SRM_10",
		11: "SRM_11",
		12: "SRM_12",
		13: "SRM_13",
		14: "SRM_14",
		15: "SRM_15",
		16: "SRM_16",
		17: "SRM_17",
		18: "SRM_18",
		19: "SRM_19",
		20: "SRM_20",
		21: "SRM_21",
		22: "SRM_22",
		23: "SRM_23",
		24: "SRM_24",
		25: "SRM_25",
		26: "SRM_26",
		27: "SRM_27",
		28: "SRM_28",
		29: "SRM_29",
		30: "SRM_30",
	}
	SRM_value = map[string]int32{
		"SRM_UNSPECIFIED": 0,
		"SRM_1":           1,
		"SRM_2":           2,
		"SRM_3":           3,
		"SRM_4":           4,
		"SRM_5":           5,
		"SRM_6":           6,
		"SRM_7":           7,
		"SRM_8":           8,
		"SRM_9":           9,
		"SRM_10":          10,
		"SRM_11":          11,
		"SRM_12":          12,
		"SRM_13":          13,
		"SRM_14":          14,
		"SRM_15":          15,
		"SRM_16":          16,
		"SRM_17":          17,
		"SRM_18":          18,
		"SRM_19":          19,
		"SRM_20":          20,
		"SRM_21":          21,
		"SRM_22":          22,
		"SRM_23":          23,
		"SRM_24":          24,
		"SRM_25":          25,
		"SRM_26":          26,
		"SRM_27":          27,
		"SRM_28":          28,
		"SRM_29":          29,
		"SRM_30":          30,
	}
)

func (x SRM) Enum() *SRM {
	p := new(SRM)
	*p = x
	return p
}

func (x SRM) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SRM) Descriptor() protoreflect.EnumDescriptor {
	return file_beerproto_v1_srm_proto_enumTypes[0].Descriptor()
}

func (SRM) Type() protoreflect.EnumType {
	return &file_beerproto_v1_srm_proto_enumTypes[0]
}

func (x SRM) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SRM.Descriptor instead.
func (SRM) EnumDescriptor() ([]byte, []int) {
	return file_beerproto_v1_srm_proto_rawDescGZIP(), []int{0}
}

type Color struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A uint32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	R uint32 `protobuf:"varint,2,opt,name=r,proto3" json:"r,omitempty"`
	G uint32 `protobuf:"varint,3,opt,name=g,proto3" json:"g,omitempty"`
	B uint32 `protobuf:"varint,4,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *Color) Reset() {
	*x = Color{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beerproto_v1_srm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Color) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Color) ProtoMessage() {}

func (x *Color) ProtoReflect() protoreflect.Message {
	mi := &file_beerproto_v1_srm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Color.ProtoReflect.Descriptor instead.
func (*Color) Descriptor() ([]byte, []int) {
	return file_beerproto_v1_srm_proto_rawDescGZIP(), []int{0}
}

func (x *Color) GetA() uint32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *Color) GetR() uint32 {
	if x != nil {
		return x.R
	}
	return 0
}

func (x *Color) GetG() uint32 {
	if x != nil {
		return x.G
	}
	return 0
}

func (x *Color) GetB() uint32 {
	if x != nil {
		return x.B
	}
	return 0
}

var file_beerproto_v1_srm_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*Color)(nil),
		Field:         123456783,
		Name:          "beerproto.v1.color",
		Tag:           "bytes,123456783,opt,name=color",
		Filename:      "beerproto/v1/srm.proto",
	},
}

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional beerproto.v1.Color color = 123456783;
	E_Color = &file_beerproto_v1_srm_proto_extTypes[0]
)

var File_beerproto_v1_srm_proto protoreflect.FileDescriptor

var file_beerproto_v1_srm_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f,
	0x72, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x61, 0x12,
	0x0c, 0x0a, 0x01, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x72, 0x12, 0x0c, 0x0a,
	0x01, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x67, 0x12, 0x0c, 0x0a, 0x01, 0x62,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x62, 0x2a, 0xa4, 0x07, 0x0a, 0x03, 0x53, 0x52,
	0x4d, 0x12, 0x24, 0x0a, 0x0f, 0x53, 0x52, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x0f, 0xfa, 0xd0, 0xf9, 0xd6, 0x03, 0x09, 0x10, 0xff,
	0x01, 0x18, 0xff, 0x01, 0x20, 0xff, 0x01, 0x12, 0x1d, 0x0a, 0x05, 0x53, 0x52, 0x4d, 0x5f, 0x31,
	0x10, 0x01, 0x1a, 0x12, 0xfa, 0xd0, 0xf9, 0xd6, 0x03, 0x0c, 0x08, 0xff, 0x01, 0x10, 0xf3, 0x01,
	0x18, 0xf9, 0x01, 0x20, 0x93, 0x01, 0x12, 0x1c, 0x0a, 0x05, 0x53, 0x52, 0x4d, 0x5f, 0x32, 0x10,
	0x02, 0x1a, 0x11, 0xfa, 0xd0, 0xf9, 0xd6, 0x03, 0x0b, 0x08, 0xff, 0x01, 0x10, 0xf5, 0x01, 0x18,
	0xf7, 0x01, 0x20, 0x5c, 0x12, 0x1c, 0x0a, 0x05, 0x53, 0x52, 0x4d, 0x5f, 0x33, 0x10, 0x03, 0x1a,
	0x11, 0xfa, 0xd0, 0xf9, 0xd6, 0x03, 0x0b, 0x08, 0xff, 0x01, 0x10, 0xf6, 0x01, 0x18, 0xf5, 0x01,
	0x20, 0x13, 0x12, 0x1c, 0x0a, 0x05, 0x53, 0x52, 0x4d, 0x5f, 0x34, 0x10, 0x04, 0x1a, 0x11, 0xfa,
	0xd0, 0xf9, 0xd6, 0x03, 0x0b, 0x08, 0xff, 0x01, 0x10, 0xea, 0x01, 0x18, 0xe6, 0x01, 0x20, 0x15,
	0x12, 0x1c, 0x0a, 0x05, 0x53, 0x52, 0x4d, 0x5f, 0x35, 0x10, 0x05, 0x1a, 0x11, 0xfa, 0xd0, 0xf9,
	0xd6, 0x03, 0x0b, 0x08, 0xff, 0x01, 0x10, 0xe0, 0x01, 0x18, 0xd0, 0x01, 0x20, 0x1b, 0x12, 0x1c,
	0x0a, 0x05, 0x53, 0x52, 0x4d, 0x5f, 0x36, 0x10, 0x06, 0x1a, 0x11, 0xfa, 0xd0, 0xf9, 0xd6, 0x03,
	0x0b, 0x08, 0xff, 0x01, 0x10, 0xd5, 0x01, 0x18, 0xbc, 0x01, 0x20, 0x26, 0x12, 0x1c, 0x0a, 0x05,
	0x53, 0x52, 0x4d, 0x5f, 0x37, 0x10, 0x07, 0x1a, 0x11, 0xfa, 0xd0, 0xf9, 0xd6, 0x03, 0x0b, 0x08,
	0xff, 0x01, 0x10, 0xcd, 0x01, 0x18, 0xaa, 0x01, 0x20, 0x37, 0x12, 0x1c, 0x0a, 0x05, 0x53, 0x52,
	0x4d, 0x5f, 0x38, 0x10, 0x08, 0x1a, 0x11, 0xfa, 0xd0, 0xf9, 0xd6, 0x03, 0x0b, 0x08, 0xff, 0x01,
	0x10, 0xc1, 0x01, 0x18, 0x96, 0x01, 0x20, 0x3c, 0x12, 0x1c, 0x0a, 0x05, 0x53, 0x52, 0x4d, 0x5f,
	0x39, 0x10, 0x09, 0x1a, 0x11, 0xfa, 0xd0, 0xf9, 0xd6, 0x03, 0x0b, 0x08, 0xff, 0x01, 0x10, 0xbe,
	0x01, 0x18, 0x8c, 0x01, 0x20, 0x3a, 0x12, 0x1d, 0x0a, 0x06, 0x53, 0x52, 0x4d, 0x5f, 0x31, 0x30,
	0x10, 0x0a, 0x1a, 0x11, 0xfa, 0xd0, 0xf9, 0xd6, 0x03, 0x0b, 0x08, 0xff, 0x01, 0x10, 0xbe, 0x01,
	0x18, 0x82, 0x01, 0x20, 0x3a, 0x12, 0x1c, 0x0a, 0x06, 0x53, 0x52, 0x4d, 0x5f, 0x31, 0x31, 0x10,
	0x0b, 0x1a, 0x10, 0xfa, 0xd0, 0xf9, 0xd6, 0x03, 0x0a, 0x08, 0xff, 0x01, 0x10, 0xc1, 0x01, 0x18,
	0x7a, 0x20, 0x37, 0x12, 0x1c, 0x0a, 0x06, 0x53, 0x52, 0x4d, 0x5f, 0x31, 0x32, 0x10, 0x0c, 0x1a,
	0x10, 0xfa, 0xd0, 0xf9, 0xd6, 0x03, 0x0a, 0x08, 0xff, 0x01, 0x10, 0xbf, 0x01, 0x18, 0x71, 0x20,
	0x38, 0x12, 0x1c, 0x0a, 0x06, 0x53, 0x52, 0x4d, 0x5f, 0x31, 0x33, 0x10, 0x0d, 0x1a, 0x10, 0xfa,
	0xd0, 0xf9, 0xd6, 0x03, 0x0a, 0x08, 0xff, 0x01, 0x10, 0xbc, 0x01, 0x18, 0x67, 0x20, 0x33, 0x12,
	0x1c, 0x0a, 0x06, 0x53, 0x52, 0x4d, 0x5f, 0x31, 0x34, 0x10, 0x0e, 0x1a, 0x10, 0xfa, 0xd0, 0xf9,
	0xd6, 0x03, 0x0a, 0x08, 0xff, 0x01, 0x10, 0xb2, 0x01, 0x18, 0x60, 0x20, 0x33, 0x12, 0x1c, 0x0a,
	0x06, 0x53, 0x52, 0x4d, 0x5f, 0x31, 0x35, 0x10, 0x0f, 0x1a, 0x10, 0xfa, 0xd0, 0xf9, 0xd6, 0x03,
	0x0a, 0x08, 0xff, 0x01, 0x10, 0xa8, 0x01, 0x18, 0x58, 0x20, 0x39, 0x12, 0x1c, 0x0a, 0x06, 0x53,
	0x52, 0x4d, 0x5f, 0x31, 0x36, 0x10, 0x10, 0x1a, 0x10, 0xfa, 0xd0, 0xf9, 0xd6, 0x03, 0x0a, 0x08,
	0xff, 0x01, 0x10, 0x98, 0x01, 0x18, 0x53, 0x20, 0x36, 0x12, 0x1c, 0x0a, 0x06, 0x53, 0x52, 0x4d,
	0x5f, 0x31, 0x37, 0x10, 0x11, 0x1a, 0x10, 0xfa, 0xd0, 0xf9, 0xd6, 0x03, 0x0a, 0x08, 0xff, 0x01,
	0x10, 0x8d, 0x01, 0x18, 0x4c, 0x20, 0x32, 0x12, 0x1b, 0x0a, 0x06, 0x53, 0x52, 0x4d, 0x5f, 0x31,
	0x38, 0x10, 0x12, 0x1a, 0x0f, 0xfa, 0xd0, 0xf9, 0xd6, 0x03, 0x09, 0x08, 0xff, 0x01, 0x10, 0x7c,
	0x18, 0x45, 0x20, 0x2d, 0x12, 0x1b, 0x0a, 0x06, 0x53, 0x52, 0x4d, 0x5f, 0x31, 0x39, 0x10, 0x13,
	0x1a, 0x0f, 0xfa, 0xd0, 0xf9, 0xd6, 0x03, 0x09, 0x08, 0xff, 0x01, 0x10, 0x6b, 0x18, 0x3a, 0x20,
	0x1e, 0x12, 0x1b, 0x0a, 0x06, 0x53, 0x52, 0x4d, 0x5f, 0x32, 0x30, 0x10, 0x14, 0x1a, 0x0f, 0xfa,
	0xd0, 0xf9, 0xd6, 0x03, 0x09, 0x08, 0xff, 0x01, 0x10, 0x5d, 0x18, 0x34, 0x20, 0x1a, 0x12, 0x1b,
	0x0a, 0x06, 0x53, 0x52, 0x4d, 0x5f, 0x32, 0x31, 0x10, 0x15, 0x1a, 0x0f, 0xfa, 0xd0, 0xf9, 0xd6,
	0x03, 0x09, 0x08, 0xff, 0x01, 0x10, 0x4e, 0x18, 0x2a, 0x20, 0x0c, 0x12, 0x1b, 0x0a, 0x06, 0x53,
	0x52, 0x4d, 0x5f, 0x32, 0x32, 0x10, 0x16, 0x1a, 0x0f, 0xfa, 0xd0, 0xf9, 0xd6, 0x03, 0x09, 0x08,
	0xff, 0x01, 0x10, 0x4a, 0x18, 0x27, 0x20, 0x27, 0x12, 0x1b, 0x0a, 0x06, 0x53, 0x52, 0x4d, 0x5f,
	0x32, 0x33, 0x10, 0x17, 0x1a, 0x0f, 0xfa, 0xd0, 0xf9, 0xd6, 0x03, 0x09, 0x08, 0xff, 0x01, 0x10,
	0x36, 0x18, 0x1f, 0x20, 0x1b, 0x12, 0x1b, 0x0a, 0x06, 0x53, 0x52, 0x4d, 0x5f, 0x32, 0x34, 0x10,
	0x18, 0x1a, 0x0f, 0xfa, 0xd0, 0xf9, 0xd6, 0x03, 0x09, 0x08, 0xff, 0x01, 0x10, 0x26, 0x18, 0x17,
	0x20, 0x16, 0x12, 0x1b, 0x0a, 0x06, 0x53, 0x52, 0x4d, 0x5f, 0x32, 0x35, 0x10, 0x19, 0x1a, 0x0f,
	0xfa, 0xd0, 0xf9, 0xd6, 0x03, 0x09, 0x08, 0xff, 0x01, 0x10, 0x23, 0x18, 0x17, 0x20, 0x16, 0x12,
	0x1b, 0x0a, 0x06, 0x53, 0x52, 0x4d, 0x5f, 0x32, 0x36, 0x10, 0x1a, 0x1a, 0x0f, 0xfa, 0xd0, 0xf9,
	0xd6, 0x03, 0x09, 0x08, 0xff, 0x01, 0x10, 0x19, 0x18, 0x10, 0x20, 0x0f, 0x12, 0x1b, 0x0a, 0x06,
	0x53, 0x52, 0x4d, 0x5f, 0x32, 0x37, 0x10, 0x1b, 0x1a, 0x0f, 0xfa, 0xd0, 0xf9, 0xd6, 0x03, 0x09,
	0x08, 0xff, 0x01, 0x10, 0x16, 0x18, 0x10, 0x20, 0x0f, 0x12, 0x1b, 0x0a, 0x06, 0x53, 0x52, 0x4d,
	0x5f, 0x32, 0x38, 0x10, 0x1c, 0x1a, 0x0f, 0xfa, 0xd0, 0xf9, 0xd6, 0x03, 0x09, 0x08, 0xff, 0x01,
	0x10, 0x12, 0x18, 0x0d, 0x20, 0x0c, 0x12, 0x1b, 0x0a, 0x06, 0x53, 0x52, 0x4d, 0x5f, 0x32, 0x39,
	0x10, 0x1d, 0x1a, 0x0f, 0xfa, 0xd0, 0xf9, 0xd6, 0x03, 0x09, 0x08, 0xff, 0x01, 0x10, 0x10, 0x18,
	0x0b, 0x20, 0x0a, 0x12, 0x1b, 0x0a, 0x06, 0x53, 0x52, 0x4d, 0x5f, 0x33, 0x30, 0x10, 0x1e, 0x1a,
	0x0f, 0xfa, 0xd0, 0xf9, 0xd6, 0x03, 0x09, 0x08, 0xff, 0x01, 0x10, 0x05, 0x18, 0x0b, 0x20, 0x0a,
	0x3a, 0x52, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8f, 0x9a, 0xef,
	0x3a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x88, 0x01, 0x01, 0x42, 0xa9, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x65, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x53, 0x72, 0x6d, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x65, 0x65, 0x72,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x6f, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0d, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_beerproto_v1_srm_proto_rawDescOnce sync.Once
	file_beerproto_v1_srm_proto_rawDescData = file_beerproto_v1_srm_proto_rawDesc
)

func file_beerproto_v1_srm_proto_rawDescGZIP() []byte {
	file_beerproto_v1_srm_proto_rawDescOnce.Do(func() {
		file_beerproto_v1_srm_proto_rawDescData = protoimpl.X.CompressGZIP(file_beerproto_v1_srm_proto_rawDescData)
	})
	return file_beerproto_v1_srm_proto_rawDescData
}

var file_beerproto_v1_srm_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_beerproto_v1_srm_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_beerproto_v1_srm_proto_goTypes = []interface{}{
	(SRM)(0),                              // 0: beerproto.v1.SRM
	(*Color)(nil),                         // 1: beerproto.v1.Color
	(*descriptorpb.EnumValueOptions)(nil), // 2: google.protobuf.EnumValueOptions
}
var file_beerproto_v1_srm_proto_depIdxs = []int32{
	2, // 0: beerproto.v1.color:extendee -> google.protobuf.EnumValueOptions
	1, // 1: beerproto.v1.color:type_name -> beerproto.v1.Color
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_beerproto_v1_srm_proto_init() }
func file_beerproto_v1_srm_proto_init() {
	if File_beerproto_v1_srm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_beerproto_v1_srm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Color); i {
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
			RawDescriptor: file_beerproto_v1_srm_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_beerproto_v1_srm_proto_goTypes,
		DependencyIndexes: file_beerproto_v1_srm_proto_depIdxs,
		EnumInfos:         file_beerproto_v1_srm_proto_enumTypes,
		MessageInfos:      file_beerproto_v1_srm_proto_msgTypes,
		ExtensionInfos:    file_beerproto_v1_srm_proto_extTypes,
	}.Build()
	File_beerproto_v1_srm_proto = out.File
	file_beerproto_v1_srm_proto_rawDesc = nil
	file_beerproto_v1_srm_proto_goTypes = nil
	file_beerproto_v1_srm_proto_depIdxs = nil
}
