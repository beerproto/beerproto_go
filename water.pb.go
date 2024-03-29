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
// source: beerproto/v1/water.proto

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

// WaterBase provides unique properties to identify individual records of  brewing water
type WaterBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Calcium occurs naturally in most water sources and is the principle cause of hardness.
	Calcium *ConcentrationType `protobuf:"bytes,2,opt,name=calcium,proto3" json:"calcium,omitempty"`
	// Nitrite occurs naturally at low levels in some waters but is removed by treatment. It is sometime produced as a by-product when chloramine is used as a disinfectant.
	Nitrite *ConcentrationType `protobuf:"bytes,3,opt,name=nitrite,proto3" json:"nitrite,omitempty"`
	// Occurs naturally in water sources and is derived through contact with rocks
	Chloride  *ConcentrationType `protobuf:"bytes,4,opt,name=chloride,proto3" json:"chloride,omitempty"`
	Name      string             `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Potassium *ConcentrationType `protobuf:"bytes,6,opt,name=potassium,proto3" json:"potassium,omitempty"`
	Carbonate *ConcentrationType `protobuf:"bytes,7,opt,name=carbonate,proto3" json:"carbonate,omitempty"`
	// Iron occurs naturally in some water. High levels are treated to reduce the iron content. A number of water mains are made of iron. Brown discolouration complaints are associated with corroding iron mains. Iron is not harmful to health.
	Iron *ConcentrationType `protobuf:"bytes,8,opt,name=iron,proto3" json:"iron,omitempty"`
	// Fluoride salts typically have distinctive bitter tastes, and are odorless.
	Flouride *ConcentrationType `protobuf:"bytes,9,opt,name=flouride,proto3" json:"flouride,omitempty"`
	Sulfate  *ConcentrationType `protobuf:"bytes,10,opt,name=sulfate,proto3" json:"sulfate,omitempty"`
	// At levels of 10-30mg/l it is an important yeast nutrient, but above 30mg/l it can cause a sour/bitter taste to the beer.
	Magnesium   *ConcentrationType `protobuf:"bytes,11,opt,name=magnesium,proto3" json:"magnesium,omitempty"`
	Producer    string             `protobuf:"bytes,12,opt,name=producer,proto3" json:"producer,omitempty"`
	Bicarbonate *ConcentrationType `protobuf:"bytes,13,opt,name=bicarbonate,proto3" json:"bicarbonate,omitempty"`
	// Nitrate occurs naturally in most source waters but concentrations can be increased as a result of fertiliser use. Where necessary concentrations in drinking water can be reduced by diluting with sources where nitrate levels are low or through specific treatment.
	Nitrate *ConcentrationType `protobuf:"bytes,14,opt,name=nitrate,proto3" json:"nitrate,omitempty"`
	// Sodium is naturally present in many water sources. Domestic water softeners can increase the sodium concentration.
	Sodium *ConcentrationType `protobuf:"bytes,15,opt,name=sodium,proto3" json:"sodium,omitempty"`
}

func (x *WaterBase) Reset() {
	*x = WaterBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beerproto_v1_water_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WaterBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaterBase) ProtoMessage() {}

func (x *WaterBase) ProtoReflect() protoreflect.Message {
	mi := &file_beerproto_v1_water_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaterBase.ProtoReflect.Descriptor instead.
func (*WaterBase) Descriptor() ([]byte, []int) {
	return file_beerproto_v1_water_proto_rawDescGZIP(), []int{0}
}

func (x *WaterBase) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WaterBase) GetCalcium() *ConcentrationType {
	if x != nil {
		return x.Calcium
	}
	return nil
}

func (x *WaterBase) GetNitrite() *ConcentrationType {
	if x != nil {
		return x.Nitrite
	}
	return nil
}

func (x *WaterBase) GetChloride() *ConcentrationType {
	if x != nil {
		return x.Chloride
	}
	return nil
}

func (x *WaterBase) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WaterBase) GetPotassium() *ConcentrationType {
	if x != nil {
		return x.Potassium
	}
	return nil
}

func (x *WaterBase) GetCarbonate() *ConcentrationType {
	if x != nil {
		return x.Carbonate
	}
	return nil
}

func (x *WaterBase) GetIron() *ConcentrationType {
	if x != nil {
		return x.Iron
	}
	return nil
}

func (x *WaterBase) GetFlouride() *ConcentrationType {
	if x != nil {
		return x.Flouride
	}
	return nil
}

func (x *WaterBase) GetSulfate() *ConcentrationType {
	if x != nil {
		return x.Sulfate
	}
	return nil
}

func (x *WaterBase) GetMagnesium() *ConcentrationType {
	if x != nil {
		return x.Magnesium
	}
	return nil
}

func (x *WaterBase) GetProducer() string {
	if x != nil {
		return x.Producer
	}
	return ""
}

func (x *WaterBase) GetBicarbonate() *ConcentrationType {
	if x != nil {
		return x.Bicarbonate
	}
	return nil
}

func (x *WaterBase) GetNitrate() *ConcentrationType {
	if x != nil {
		return x.Nitrate
	}
	return nil
}

func (x *WaterBase) GetSodium() *ConcentrationType {
	if x != nil {
		return x.Sodium
	}
	return nil
}

// WaterAdditionType collects the attributes of each water addition for use in a recipe
type WaterAdditionType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Flouride    *ConcentrationType `protobuf:"bytes,2,opt,name=flouride,proto3" json:"flouride,omitempty"`
	Sulfate     *ConcentrationType `protobuf:"bytes,3,opt,name=sulfate,proto3" json:"sulfate,omitempty"`
	Producer    string             `protobuf:"bytes,4,opt,name=producer,proto3" json:"producer,omitempty"`
	Nitrate     *ConcentrationType `protobuf:"bytes,5,opt,name=nitrate,proto3" json:"nitrate,omitempty"`
	Nitrite     *ConcentrationType `protobuf:"bytes,6,opt,name=nitrite,proto3" json:"nitrite,omitempty"`
	Chloride    *ConcentrationType `protobuf:"bytes,7,opt,name=chloride,proto3" json:"chloride,omitempty"`
	Amount      *VolumeType        `protobuf:"bytes,8,opt,name=amount,proto3" json:"amount,omitempty"`
	Name        string             `protobuf:"bytes,89,opt,name=name,proto3" json:"name,omitempty"`
	Potassium   *ConcentrationType `protobuf:"bytes,10,opt,name=potassium,proto3" json:"potassium,omitempty"`
	Magnesium   *ConcentrationType `protobuf:"bytes,11,opt,name=magnesium,proto3" json:"magnesium,omitempty"`
	Iron        *ConcentrationType `protobuf:"bytes,12,opt,name=iron,proto3" json:"iron,omitempty"`
	Bicarbonate *ConcentrationType `protobuf:"bytes,13,opt,name=bicarbonate,proto3" json:"bicarbonate,omitempty"`
	Calcium     *ConcentrationType `protobuf:"bytes,14,opt,name=calcium,proto3" json:"calcium,omitempty"`
	Carbonate   *ConcentrationType `protobuf:"bytes,15,opt,name=carbonate,proto3" json:"carbonate,omitempty"`
	Sodium      *ConcentrationType `protobuf:"bytes,16,opt,name=sodium,proto3" json:"sodium,omitempty"`
}

func (x *WaterAdditionType) Reset() {
	*x = WaterAdditionType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beerproto_v1_water_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WaterAdditionType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WaterAdditionType) ProtoMessage() {}

func (x *WaterAdditionType) ProtoReflect() protoreflect.Message {
	mi := &file_beerproto_v1_water_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WaterAdditionType.ProtoReflect.Descriptor instead.
func (*WaterAdditionType) Descriptor() ([]byte, []int) {
	return file_beerproto_v1_water_proto_rawDescGZIP(), []int{1}
}

func (x *WaterAdditionType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WaterAdditionType) GetFlouride() *ConcentrationType {
	if x != nil {
		return x.Flouride
	}
	return nil
}

func (x *WaterAdditionType) GetSulfate() *ConcentrationType {
	if x != nil {
		return x.Sulfate
	}
	return nil
}

func (x *WaterAdditionType) GetProducer() string {
	if x != nil {
		return x.Producer
	}
	return ""
}

func (x *WaterAdditionType) GetNitrate() *ConcentrationType {
	if x != nil {
		return x.Nitrate
	}
	return nil
}

func (x *WaterAdditionType) GetNitrite() *ConcentrationType {
	if x != nil {
		return x.Nitrite
	}
	return nil
}

func (x *WaterAdditionType) GetChloride() *ConcentrationType {
	if x != nil {
		return x.Chloride
	}
	return nil
}

func (x *WaterAdditionType) GetAmount() *VolumeType {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *WaterAdditionType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WaterAdditionType) GetPotassium() *ConcentrationType {
	if x != nil {
		return x.Potassium
	}
	return nil
}

func (x *WaterAdditionType) GetMagnesium() *ConcentrationType {
	if x != nil {
		return x.Magnesium
	}
	return nil
}

func (x *WaterAdditionType) GetIron() *ConcentrationType {
	if x != nil {
		return x.Iron
	}
	return nil
}

func (x *WaterAdditionType) GetBicarbonate() *ConcentrationType {
	if x != nil {
		return x.Bicarbonate
	}
	return nil
}

func (x *WaterAdditionType) GetCalcium() *ConcentrationType {
	if x != nil {
		return x.Calcium
	}
	return nil
}

func (x *WaterAdditionType) GetCarbonate() *ConcentrationType {
	if x != nil {
		return x.Carbonate
	}
	return nil
}

func (x *WaterAdditionType) GetSodium() *ConcentrationType {
	if x != nil {
		return x.Sodium
	}
	return nil
}

var File_beerproto_v1_water_proto protoreflect.FileDescriptor

var file_beerproto_v1_water_proto_rawDesc = []byte{
	0x0a, 0x18, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x77,
	0x61, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x65, 0x65, 0x72,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f,
	0x06, 0x0a, 0x09, 0x57, 0x61, 0x74, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x07,
	0x63, 0x61, 0x6c, 0x63, 0x69, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07,
	0x63, 0x61, 0x6c, 0x63, 0x69, 0x75, 0x6d, 0x12, 0x39, 0x0a, 0x07, 0x6e, 0x69, 0x74, 0x72, 0x69,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6e, 0x69, 0x74, 0x72, 0x69,
	0x74, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x63, 0x68, 0x6c, 0x6f, 0x72, 0x69, 0x64, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x63, 0x68, 0x6c, 0x6f, 0x72, 0x69, 0x64, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x70, 0x6f, 0x74, 0x61, 0x73, 0x73, 0x69, 0x75, 0x6d,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x70, 0x6f, 0x74, 0x61, 0x73, 0x73, 0x69,
	0x75, 0x6d, 0x12, 0x3d, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x61, 0x74,
	0x65, 0x12, 0x33, 0x0a, 0x04, 0x69, 0x72, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x69, 0x72, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x08, 0x66, 0x6c, 0x6f, 0x75, 0x72, 0x69,
	0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x75, 0x72,
	0x69, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x73, 0x75, 0x6c, 0x66, 0x61, 0x74, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x73, 0x75, 0x6c, 0x66, 0x61, 0x74, 0x65, 0x12, 0x3d,
	0x0a, 0x09, 0x6d, 0x61, 0x67, 0x6e, 0x65, 0x73, 0x69, 0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x6d, 0x61, 0x67, 0x6e, 0x65, 0x73, 0x69, 0x75, 0x6d, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x0b, 0x62, 0x69, 0x63,
	0x61, 0x72, 0x62, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0b, 0x62, 0x69, 0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x07,
	0x6e, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07,
	0x6e, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x6f, 0x64, 0x69, 0x75,
	0x6d, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x64, 0x69, 0x75, 0x6d,
	0x22, 0xd9, 0x06, 0x0a, 0x11, 0x57, 0x61, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x08, 0x66, 0x6c, 0x6f, 0x75, 0x72, 0x69,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x75, 0x72,
	0x69, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x73, 0x75, 0x6c, 0x66, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x73, 0x75, 0x6c, 0x66, 0x61, 0x74, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x07, 0x6e, 0x69,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x65,
	0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65,
	0x6e, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6e, 0x69,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x74, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6e, 0x69, 0x74, 0x72, 0x69, 0x74, 0x65,
	0x12, 0x3b, 0x0a, 0x08, 0x63, 0x68, 0x6c, 0x6f, 0x72, 0x69, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x08, 0x63, 0x68, 0x6c, 0x6f, 0x72, 0x69, 0x64, 0x65, 0x12, 0x30, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x59, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x70, 0x6f, 0x74, 0x61, 0x73, 0x73, 0x69, 0x75, 0x6d,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x70, 0x6f, 0x74, 0x61, 0x73, 0x73, 0x69,
	0x75, 0x6d, 0x12, 0x3d, 0x0a, 0x09, 0x6d, 0x61, 0x67, 0x6e, 0x65, 0x73, 0x69, 0x75, 0x6d, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x6d, 0x61, 0x67, 0x6e, 0x65, 0x73, 0x69, 0x75,
	0x6d, 0x12, 0x33, 0x0a, 0x04, 0x69, 0x72, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x69, 0x72, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x0b, 0x62, 0x69, 0x63, 0x61, 0x72, 0x62,
	0x6f, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x65,
	0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65,
	0x6e, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x62, 0x69,
	0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x63, 0x61, 0x6c,
	0x63, 0x69, 0x75, 0x6d, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x65, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x6e,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x63, 0x61, 0x6c,
	0x63, 0x69, 0x75, 0x6d, 0x12, 0x3d, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x61, 0x74,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e,
	0x61, 0x74, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x6f, 0x64, 0x69, 0x75, 0x6d, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x64, 0x69, 0x75, 0x6d, 0x42, 0xab, 0x01, 0x0a,
	0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x42, 0x0a, 0x57, 0x61, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x65, 0x72,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f,
	0x67, 0x6f, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b,
	0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x58,
	0x58, 0xaa, 0x02, 0x0c, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0c, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x18, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x42, 0x65, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_beerproto_v1_water_proto_rawDescOnce sync.Once
	file_beerproto_v1_water_proto_rawDescData = file_beerproto_v1_water_proto_rawDesc
)

func file_beerproto_v1_water_proto_rawDescGZIP() []byte {
	file_beerproto_v1_water_proto_rawDescOnce.Do(func() {
		file_beerproto_v1_water_proto_rawDescData = protoimpl.X.CompressGZIP(file_beerproto_v1_water_proto_rawDescData)
	})
	return file_beerproto_v1_water_proto_rawDescData
}

var file_beerproto_v1_water_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_beerproto_v1_water_proto_goTypes = []interface{}{
	(*WaterBase)(nil),         // 0: beerproto.v1.WaterBase
	(*WaterAdditionType)(nil), // 1: beerproto.v1.WaterAdditionType
	(*ConcentrationType)(nil), // 2: beerproto.v1.ConcentrationType
	(*VolumeType)(nil),        // 3: beerproto.v1.VolumeType
}
var file_beerproto_v1_water_proto_depIdxs = []int32{
	2,  // 0: beerproto.v1.WaterBase.calcium:type_name -> beerproto.v1.ConcentrationType
	2,  // 1: beerproto.v1.WaterBase.nitrite:type_name -> beerproto.v1.ConcentrationType
	2,  // 2: beerproto.v1.WaterBase.chloride:type_name -> beerproto.v1.ConcentrationType
	2,  // 3: beerproto.v1.WaterBase.potassium:type_name -> beerproto.v1.ConcentrationType
	2,  // 4: beerproto.v1.WaterBase.carbonate:type_name -> beerproto.v1.ConcentrationType
	2,  // 5: beerproto.v1.WaterBase.iron:type_name -> beerproto.v1.ConcentrationType
	2,  // 6: beerproto.v1.WaterBase.flouride:type_name -> beerproto.v1.ConcentrationType
	2,  // 7: beerproto.v1.WaterBase.sulfate:type_name -> beerproto.v1.ConcentrationType
	2,  // 8: beerproto.v1.WaterBase.magnesium:type_name -> beerproto.v1.ConcentrationType
	2,  // 9: beerproto.v1.WaterBase.bicarbonate:type_name -> beerproto.v1.ConcentrationType
	2,  // 10: beerproto.v1.WaterBase.nitrate:type_name -> beerproto.v1.ConcentrationType
	2,  // 11: beerproto.v1.WaterBase.sodium:type_name -> beerproto.v1.ConcentrationType
	2,  // 12: beerproto.v1.WaterAdditionType.flouride:type_name -> beerproto.v1.ConcentrationType
	2,  // 13: beerproto.v1.WaterAdditionType.sulfate:type_name -> beerproto.v1.ConcentrationType
	2,  // 14: beerproto.v1.WaterAdditionType.nitrate:type_name -> beerproto.v1.ConcentrationType
	2,  // 15: beerproto.v1.WaterAdditionType.nitrite:type_name -> beerproto.v1.ConcentrationType
	2,  // 16: beerproto.v1.WaterAdditionType.chloride:type_name -> beerproto.v1.ConcentrationType
	3,  // 17: beerproto.v1.WaterAdditionType.amount:type_name -> beerproto.v1.VolumeType
	2,  // 18: beerproto.v1.WaterAdditionType.potassium:type_name -> beerproto.v1.ConcentrationType
	2,  // 19: beerproto.v1.WaterAdditionType.magnesium:type_name -> beerproto.v1.ConcentrationType
	2,  // 20: beerproto.v1.WaterAdditionType.iron:type_name -> beerproto.v1.ConcentrationType
	2,  // 21: beerproto.v1.WaterAdditionType.bicarbonate:type_name -> beerproto.v1.ConcentrationType
	2,  // 22: beerproto.v1.WaterAdditionType.calcium:type_name -> beerproto.v1.ConcentrationType
	2,  // 23: beerproto.v1.WaterAdditionType.carbonate:type_name -> beerproto.v1.ConcentrationType
	2,  // 24: beerproto.v1.WaterAdditionType.sodium:type_name -> beerproto.v1.ConcentrationType
	25, // [25:25] is the sub-list for method output_type
	25, // [25:25] is the sub-list for method input_type
	25, // [25:25] is the sub-list for extension type_name
	25, // [25:25] is the sub-list for extension extendee
	0,  // [0:25] is the sub-list for field type_name
}

func init() { file_beerproto_v1_water_proto_init() }
func file_beerproto_v1_water_proto_init() {
	if File_beerproto_v1_water_proto != nil {
		return
	}
	file_beerproto_v1_measureable_units_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_beerproto_v1_water_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WaterBase); i {
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
		file_beerproto_v1_water_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WaterAdditionType); i {
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
			RawDescriptor: file_beerproto_v1_water_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_beerproto_v1_water_proto_goTypes,
		DependencyIndexes: file_beerproto_v1_water_proto_depIdxs,
		MessageInfos:      file_beerproto_v1_water_proto_msgTypes,
	}.Build()
	File_beerproto_v1_water_proto = out.File
	file_beerproto_v1_water_proto_rawDesc = nil
	file_beerproto_v1_water_proto_goTypes = nil
	file_beerproto_v1_water_proto_depIdxs = nil
}
