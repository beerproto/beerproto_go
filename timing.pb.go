//*
// BeerProto
//
// Another beer format, written in protocol buffer.
//
// Copyright (c) 2020 Ross Merrigan

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: beerproto/v1/timing.proto

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

type UseType int32

const (
	UseType_USE_TYPE_UNSPECIFIED UseType = 0
	// add to mash
	UseType_USE_TYPE_ADD_TO_MASH UseType = 1
	// add to boil
	UseType_USE_TYPE_ADD_TO_BOIL UseType = 2
	// add to fermentation
	UseType_USE_TYPE_ADD_TO_FERMENTATION UseType = 3
	// add to package
	UseType_USE_TYPE_ADD_TO_PACKAGE UseType = 4
)

// Enum value maps for UseType.
var (
	UseType_name = map[int32]string{
		0: "USE_TYPE_UNSPECIFIED",
		1: "USE_TYPE_ADD_TO_MASH",
		2: "USE_TYPE_ADD_TO_BOIL",
		3: "USE_TYPE_ADD_TO_FERMENTATION",
		4: "USE_TYPE_ADD_TO_PACKAGE",
	}
	UseType_value = map[string]int32{
		"USE_TYPE_UNSPECIFIED":         0,
		"USE_TYPE_ADD_TO_MASH":         1,
		"USE_TYPE_ADD_TO_BOIL":         2,
		"USE_TYPE_ADD_TO_FERMENTATION": 3,
		"USE_TYPE_ADD_TO_PACKAGE":      4,
	}
)

func (x UseType) Enum() *UseType {
	p := new(UseType)
	*p = x
	return p
}

func (x UseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UseType) Descriptor() protoreflect.EnumDescriptor {
	return file_beerproto_v1_timing_proto_enumTypes[0].Descriptor()
}

func (UseType) Type() protoreflect.EnumType {
	return &file_beerproto_v1_timing_proto_enumTypes[0]
}

func (x UseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UseType.Descriptor instead.
func (UseType) EnumDescriptor() ([]byte, []int) {
	return file_beerproto_v1_timing_proto_rawDescGZIP(), []int{0}
}

// The timing object fully describes the timing of an addition with options for basis on time, gravity, or pH at any process step
type TimingType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// What time during a process step is added, eg a value of 2 days for a dry hop addition would be added 2 days into the fermentation step.
	Time *TimeType `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	// How long an ingredient addition remains, this was referred to as time in the BeerXML standard. E.G. A 40 minute hop boil additions means to boil for 40 minutes, and a 2 day duration for a dry hop means to remove it after 2 days.
	Duration *TimeType `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	// A continuous addition is spread out evenly and added during the entire process step, eg 60 minute IPA by dogfish head takes all ofthe hop additions and adds them throughout the entire boil.
	Continuous bool `protobuf:"varint,3,opt,name=continuous,proto3" json:"continuous,omitempty"`
	// Used to indicate when an addition is added based on a desired specific gravity. E.G. Add dry hop at when SG is 1.018.
	SpecificGravity *GravityType `protobuf:"bytes,4,opt,name=specific_gravity,json=specificGravity,proto3" json:"specific_gravity,omitempty"`
	// Used to indicate when an addition is added based on a desired specific pH. eg Add brett when pH is 3.4.
	Ph *AcidityType `protobuf:"bytes,5,opt,name=ph,proto3" json:"ph,omitempty"`
	// Used to indicate what step this ingredient timing addition is referencing. EG A value of 2 for add_to_fermentation would mean to add during the second fermentation step.
	Step int32   `protobuf:"varint,6,opt,name=step,proto3" json:"step,omitempty"`
	Use  UseType `protobuf:"varint,7,opt,name=use,proto3,enum=beerproto.v1.UseType" json:"use,omitempty"`
}

func (x *TimingType) Reset() {
	*x = TimingType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beerproto_v1_timing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimingType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimingType) ProtoMessage() {}

func (x *TimingType) ProtoReflect() protoreflect.Message {
	mi := &file_beerproto_v1_timing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimingType.ProtoReflect.Descriptor instead.
func (*TimingType) Descriptor() ([]byte, []int) {
	return file_beerproto_v1_timing_proto_rawDescGZIP(), []int{0}
}

func (x *TimingType) GetTime() *TimeType {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *TimingType) GetDuration() *TimeType {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *TimingType) GetContinuous() bool {
	if x != nil {
		return x.Continuous
	}
	return false
}

func (x *TimingType) GetSpecificGravity() *GravityType {
	if x != nil {
		return x.SpecificGravity
	}
	return nil
}

func (x *TimingType) GetPh() *AcidityType {
	if x != nil {
		return x.Ph
	}
	return nil
}

func (x *TimingType) GetStep() int32 {
	if x != nil {
		return x.Step
	}
	return 0
}

func (x *TimingType) GetUse() UseType {
	if x != nil {
		return x.Use
	}
	return UseType_USE_TYPE_UNSPECIFIED
}

var File_beerproto_v1_timing_proto protoreflect.FileDescriptor

var file_beerproto_v1_timing_proto_rawDesc = []byte{
	0x0a, 0x19, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x69, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x65, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x62, 0x65, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xba, 0x02, 0x0a, 0x0a, 0x54, 0x69, 0x6d, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62,
	0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x62,
	0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x6f, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x6f, 0x75, 0x73, 0x12, 0x44,
	0x0a, 0x10, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x5f, 0x67, 0x72, 0x61, 0x76, 0x69,
	0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x61, 0x76, 0x69, 0x74, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x47, 0x72, 0x61,
	0x76, 0x69, 0x74, 0x79, 0x12, 0x29, 0x0a, 0x02, 0x70, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x63, 0x69, 0x64, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x02, 0x70, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73,
	0x74, 0x65, 0x70, 0x12, 0x27, 0x0a, 0x03, 0x75, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x15, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x03, 0x75, 0x73, 0x65, 0x2a, 0x96, 0x01, 0x0a,
	0x07, 0x55, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x53, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x55, 0x53, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41,
	0x44, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x4d, 0x41, 0x53, 0x48, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14,
	0x55, 0x53, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x54, 0x4f, 0x5f,
	0x42, 0x4f, 0x49, 0x4c, 0x10, 0x02, 0x12, 0x20, 0x0a, 0x1c, 0x55, 0x53, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x46, 0x45, 0x52, 0x4d, 0x45, 0x4e,
	0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x55, 0x53, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x54, 0x4f, 0x5f, 0x50, 0x41, 0x43, 0x4b,
	0x41, 0x47, 0x45, 0x10, 0x04, 0x42, 0xac, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x65,
	0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x54, 0x69, 0x6d, 0x69,
	0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x6f, 0x2f, 0x62, 0x65, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x42, 0x65,
	0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x42, 0x65, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x42, 0x65, 0x65, 0x72,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_beerproto_v1_timing_proto_rawDescOnce sync.Once
	file_beerproto_v1_timing_proto_rawDescData = file_beerproto_v1_timing_proto_rawDesc
)

func file_beerproto_v1_timing_proto_rawDescGZIP() []byte {
	file_beerproto_v1_timing_proto_rawDescOnce.Do(func() {
		file_beerproto_v1_timing_proto_rawDescData = protoimpl.X.CompressGZIP(file_beerproto_v1_timing_proto_rawDescData)
	})
	return file_beerproto_v1_timing_proto_rawDescData
}

var file_beerproto_v1_timing_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_beerproto_v1_timing_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_beerproto_v1_timing_proto_goTypes = []interface{}{
	(UseType)(0),        // 0: beerproto.v1.UseType
	(*TimingType)(nil),  // 1: beerproto.v1.TimingType
	(*TimeType)(nil),    // 2: beerproto.v1.TimeType
	(*GravityType)(nil), // 3: beerproto.v1.GravityType
	(*AcidityType)(nil), // 4: beerproto.v1.AcidityType
}
var file_beerproto_v1_timing_proto_depIdxs = []int32{
	2, // 0: beerproto.v1.TimingType.time:type_name -> beerproto.v1.TimeType
	2, // 1: beerproto.v1.TimingType.duration:type_name -> beerproto.v1.TimeType
	3, // 2: beerproto.v1.TimingType.specific_gravity:type_name -> beerproto.v1.GravityType
	4, // 3: beerproto.v1.TimingType.ph:type_name -> beerproto.v1.AcidityType
	0, // 4: beerproto.v1.TimingType.use:type_name -> beerproto.v1.UseType
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_beerproto_v1_timing_proto_init() }
func file_beerproto_v1_timing_proto_init() {
	if File_beerproto_v1_timing_proto != nil {
		return
	}
	file_beerproto_v1_measureable_units_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_beerproto_v1_timing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimingType); i {
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
			RawDescriptor: file_beerproto_v1_timing_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_beerproto_v1_timing_proto_goTypes,
		DependencyIndexes: file_beerproto_v1_timing_proto_depIdxs,
		EnumInfos:         file_beerproto_v1_timing_proto_enumTypes,
		MessageInfos:      file_beerproto_v1_timing_proto_msgTypes,
	}.Build()
	File_beerproto_v1_timing_proto = out.File
	file_beerproto_v1_timing_proto_rawDesc = nil
	file_beerproto_v1_timing_proto_goTypes = nil
	file_beerproto_v1_timing_proto_depIdxs = nil
}
