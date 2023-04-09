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
// source: beerproto/v1/beer.proto

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

type Recipe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// A collection of steps providing process information for common mashing procedures
	Mashes []*MashProcedureType `protobuf:"bytes,2,rep,name=mashes,proto3" json:"mashes,omitempty"`
	// Records containing a minimal collection of the description of ingredients, procedures and other required parameters necessary to recreate a batch of beer
	Recipes []*RecipeType `protobuf:"bytes,3,rep,name=recipes,proto3" json:"recipes,omitempty"`
	// Records for adjuncts which do not contribute to the gravity of the beer
	MiscellaneousIngredients []*MiscellaneousType `protobuf:"bytes,4,rep,name=miscellaneous_ingredients,json=miscellaneousIngredients,proto3" json:"miscellaneous_ingredients,omitempty"`
	// Records detailing the characteristics of the beer styles for which judging guidelines have been established
	Styles []*StyleType `protobuf:"bytes,5,rep,name=styles,proto3" json:"styles,omitempty"`
	// A collection of steps providing process information for common fermentation procedures
	Fermentations []*FermentationProcedureType `protobuf:"bytes,6,rep,name=fermentations,proto3" json:"fermentations,omitempty"`
	// A collection of steps providing process information for common boil procedures
	Boil []*BoilProcedureType `protobuf:"bytes,7,rep,name=boil,proto3" json:"boil,omitempty"`
	// Explicitly encode version within list of records
	Version float64 `protobuf:"fixed64,8,opt,name=version,proto3" json:"version,omitempty"`
	// Records for any ingredient that contributes to the gravity of the beer
	Fermentables []*FermentableType `protobuf:"bytes,9,rep,name=fermentables,proto3" json:"fermentables,omitempty"`
	// Records detailing the wide array of unique cultures
	Cultures []*CultureInformation `protobuf:"bytes,10,rep,name=cultures,proto3" json:"cultures,omitempty"`
	// Provides necessary information for brewing equipment
	Equipments []*EquipmentType `protobuf:"bytes,11,rep,name=equipments,proto3" json:"equipments,omitempty"`
	// A collection of steps providing process information for common packaging procedures
	Packaging []*PackagingProcedureType `protobuf:"bytes,12,rep,name=packaging,proto3" json:"packaging,omitempty"`
	// Records detailing the many properties of unique hop varieties
	HopVarieties []*VarietyInformation `protobuf:"bytes,13,rep,name=hop_varieties,json=hopVarieties,proto3" json:"hop_varieties,omitempty"`
	// Records for water profiles used in brewing
	Profiles []*WaterBase `protobuf:"bytes,14,rep,name=profiles,proto3" json:"profiles,omitempty"`
}

func (x *Recipe) Reset() {
	*x = Recipe{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beerproto_v1_beer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recipe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recipe) ProtoMessage() {}

func (x *Recipe) ProtoReflect() protoreflect.Message {
	mi := &file_beerproto_v1_beer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recipe.ProtoReflect.Descriptor instead.
func (*Recipe) Descriptor() ([]byte, []int) {
	return file_beerproto_v1_beer_proto_rawDescGZIP(), []int{0}
}

func (x *Recipe) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Recipe) GetMashes() []*MashProcedureType {
	if x != nil {
		return x.Mashes
	}
	return nil
}

func (x *Recipe) GetRecipes() []*RecipeType {
	if x != nil {
		return x.Recipes
	}
	return nil
}

func (x *Recipe) GetMiscellaneousIngredients() []*MiscellaneousType {
	if x != nil {
		return x.MiscellaneousIngredients
	}
	return nil
}

func (x *Recipe) GetStyles() []*StyleType {
	if x != nil {
		return x.Styles
	}
	return nil
}

func (x *Recipe) GetFermentations() []*FermentationProcedureType {
	if x != nil {
		return x.Fermentations
	}
	return nil
}

func (x *Recipe) GetBoil() []*BoilProcedureType {
	if x != nil {
		return x.Boil
	}
	return nil
}

func (x *Recipe) GetVersion() float64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Recipe) GetFermentables() []*FermentableType {
	if x != nil {
		return x.Fermentables
	}
	return nil
}

func (x *Recipe) GetCultures() []*CultureInformation {
	if x != nil {
		return x.Cultures
	}
	return nil
}

func (x *Recipe) GetEquipments() []*EquipmentType {
	if x != nil {
		return x.Equipments
	}
	return nil
}

func (x *Recipe) GetPackaging() []*PackagingProcedureType {
	if x != nil {
		return x.Packaging
	}
	return nil
}

func (x *Recipe) GetHopVarieties() []*VarietyInformation {
	if x != nil {
		return x.HopVarieties
	}
	return nil
}

func (x *Recipe) GetProfiles() []*WaterBase {
	if x != nil {
		return x.Profiles
	}
	return nil
}

var File_beerproto_v1_beer_proto protoreflect.FileDescriptor

var file_beerproto_v1_beer_proto_rawDesc = []byte{
	0x0a, 0x17, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x62,
	0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x65, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x62, 0x65, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x72,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f,
	0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x6c, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x16, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x68,
	0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xb0, 0x06, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x37, 0x0a,
	0x06, 0x6d, 0x61, 0x73, 0x68, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x73,
	0x68, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06,
	0x6d, 0x61, 0x73, 0x68, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x07, 0x72, 0x65, 0x63, 0x69, 0x70, 0x65, 0x73, 0x12, 0x5c, 0x0a, 0x19, 0x6d, 0x69,
	0x73, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x6e, 0x65, 0x6f, 0x75, 0x73, 0x5f, 0x69, 0x6e, 0x67, 0x72,
	0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x69, 0x73,
	0x63, 0x65, 0x6c, 0x6c, 0x61, 0x6e, 0x65, 0x6f, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x18,
	0x6d, 0x69, 0x73, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x6e, 0x65, 0x6f, 0x75, 0x73, 0x49, 0x6e, 0x67,
	0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x79, 0x6c,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x06, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x73, 0x12, 0x4d, 0x0a, 0x0d, 0x66, 0x65, 0x72,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x65, 0x72, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x64, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x66, 0x65, 0x72, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x33, 0x0a, 0x04, 0x62, 0x6f, 0x69, 0x6c,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x69, 0x6c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64,
	0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x62, 0x6f, 0x69, 0x6c, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x0c, 0x66, 0x65, 0x72, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x72,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x66, 0x65,
	0x72, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x08, 0x63, 0x75,
	0x6c, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62,
	0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x6c, 0x74,
	0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x63, 0x75, 0x6c, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x0a, 0x65, 0x71, 0x75, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62,
	0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x71, 0x75, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x65, 0x71, 0x75, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x42, 0x0a, 0x09, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x45, 0x0a, 0x0d, 0x68, 0x6f, 0x70,
	0x5f, 0x76, 0x61, 0x72, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x68, 0x6f, 0x70, 0x56, 0x61, 0x72, 0x69, 0x65, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x33, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x57, 0x61, 0x74, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x42, 0xaa, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x65,
	0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x42, 0x65, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x65,
	0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x6f, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x42, 0x65, 0x65, 0x72,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x42, 0x65, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0d, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_beerproto_v1_beer_proto_rawDescOnce sync.Once
	file_beerproto_v1_beer_proto_rawDescData = file_beerproto_v1_beer_proto_rawDesc
)

func file_beerproto_v1_beer_proto_rawDescGZIP() []byte {
	file_beerproto_v1_beer_proto_rawDescOnce.Do(func() {
		file_beerproto_v1_beer_proto_rawDescData = protoimpl.X.CompressGZIP(file_beerproto_v1_beer_proto_rawDescData)
	})
	return file_beerproto_v1_beer_proto_rawDescData
}

var file_beerproto_v1_beer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_beerproto_v1_beer_proto_goTypes = []interface{}{
	(*Recipe)(nil),                    // 0: beerproto.v1.Recipe
	(*MashProcedureType)(nil),         // 1: beerproto.v1.MashProcedureType
	(*RecipeType)(nil),                // 2: beerproto.v1.RecipeType
	(*MiscellaneousType)(nil),         // 3: beerproto.v1.MiscellaneousType
	(*StyleType)(nil),                 // 4: beerproto.v1.StyleType
	(*FermentationProcedureType)(nil), // 5: beerproto.v1.FermentationProcedureType
	(*BoilProcedureType)(nil),         // 6: beerproto.v1.BoilProcedureType
	(*FermentableType)(nil),           // 7: beerproto.v1.FermentableType
	(*CultureInformation)(nil),        // 8: beerproto.v1.CultureInformation
	(*EquipmentType)(nil),             // 9: beerproto.v1.EquipmentType
	(*PackagingProcedureType)(nil),    // 10: beerproto.v1.PackagingProcedureType
	(*VarietyInformation)(nil),        // 11: beerproto.v1.VarietyInformation
	(*WaterBase)(nil),                 // 12: beerproto.v1.WaterBase
}
var file_beerproto_v1_beer_proto_depIdxs = []int32{
	1,  // 0: beerproto.v1.Recipe.mashes:type_name -> beerproto.v1.MashProcedureType
	2,  // 1: beerproto.v1.Recipe.recipes:type_name -> beerproto.v1.RecipeType
	3,  // 2: beerproto.v1.Recipe.miscellaneous_ingredients:type_name -> beerproto.v1.MiscellaneousType
	4,  // 3: beerproto.v1.Recipe.styles:type_name -> beerproto.v1.StyleType
	5,  // 4: beerproto.v1.Recipe.fermentations:type_name -> beerproto.v1.FermentationProcedureType
	6,  // 5: beerproto.v1.Recipe.boil:type_name -> beerproto.v1.BoilProcedureType
	7,  // 6: beerproto.v1.Recipe.fermentables:type_name -> beerproto.v1.FermentableType
	8,  // 7: beerproto.v1.Recipe.cultures:type_name -> beerproto.v1.CultureInformation
	9,  // 8: beerproto.v1.Recipe.equipments:type_name -> beerproto.v1.EquipmentType
	10, // 9: beerproto.v1.Recipe.packaging:type_name -> beerproto.v1.PackagingProcedureType
	11, // 10: beerproto.v1.Recipe.hop_varieties:type_name -> beerproto.v1.VarietyInformation
	12, // 11: beerproto.v1.Recipe.profiles:type_name -> beerproto.v1.WaterBase
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_beerproto_v1_beer_proto_init() }
func file_beerproto_v1_beer_proto_init() {
	if File_beerproto_v1_beer_proto != nil {
		return
	}
	file_beerproto_v1_mash_proto_init()
	file_beerproto_v1_recipe_proto_init()
	file_beerproto_v1_misc_proto_init()
	file_beerproto_v1_style_proto_init()
	file_beerproto_v1_fermentation_proto_init()
	file_beerproto_v1_boil_proto_init()
	file_beerproto_v1_fermentable_proto_init()
	file_beerproto_v1_culture_proto_init()
	file_beerproto_v1_equipment_proto_init()
	file_beerproto_v1_packaging_proto_init()
	file_beerproto_v1_hop_proto_init()
	file_beerproto_v1_water_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_beerproto_v1_beer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recipe); i {
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
			RawDescriptor: file_beerproto_v1_beer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_beerproto_v1_beer_proto_goTypes,
		DependencyIndexes: file_beerproto_v1_beer_proto_depIdxs,
		MessageInfos:      file_beerproto_v1_beer_proto_msgTypes,
	}.Build()
	File_beerproto_v1_beer_proto = out.File
	file_beerproto_v1_beer_proto_rawDesc = nil
	file_beerproto_v1_beer_proto_goTypes = nil
	file_beerproto_v1_beer_proto_depIdxs = nil
}
