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
// source: beerproto/v1/style.proto

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

type StyleCategories int32

const (
	StyleCategories_STYLE_CATEGORIES_UNSPECIFIED StyleCategories = 0
	// beer
	StyleCategories_STYLE_CATEGORIES_BEER StyleCategories = 1
	// cider
	StyleCategories_STYLE_CATEGORIES_CIDER StyleCategories = 2
	// kombucha
	StyleCategories_STYLE_CATEGORIES_KOMBUCHA StyleCategories = 3
	// mead
	StyleCategories_STYLE_CATEGORIES_MEAD StyleCategories = 4
	// soda
	StyleCategories_STYLE_CATEGORIES_SODA StyleCategories = 5
	// wine
	StyleCategories_STYLE_CATEGORIES_WINE StyleCategories = 6
	// other
	StyleCategories_STYLE_CATEGORIES_OTHER StyleCategories = 7
)

// Enum value maps for StyleCategories.
var (
	StyleCategories_name = map[int32]string{
		0: "STYLE_CATEGORIES_UNSPECIFIED",
		1: "STYLE_CATEGORIES_BEER",
		2: "STYLE_CATEGORIES_CIDER",
		3: "STYLE_CATEGORIES_KOMBUCHA",
		4: "STYLE_CATEGORIES_MEAD",
		5: "STYLE_CATEGORIES_SODA",
		6: "STYLE_CATEGORIES_WINE",
		7: "STYLE_CATEGORIES_OTHER",
	}
	StyleCategories_value = map[string]int32{
		"STYLE_CATEGORIES_UNSPECIFIED": 0,
		"STYLE_CATEGORIES_BEER":        1,
		"STYLE_CATEGORIES_CIDER":       2,
		"STYLE_CATEGORIES_KOMBUCHA":    3,
		"STYLE_CATEGORIES_MEAD":        4,
		"STYLE_CATEGORIES_SODA":        5,
		"STYLE_CATEGORIES_WINE":        6,
		"STYLE_CATEGORIES_OTHER":       7,
	}
)

func (x StyleCategories) Enum() *StyleCategories {
	p := new(StyleCategories)
	*p = x
	return p
}

func (x StyleCategories) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StyleCategories) Descriptor() protoreflect.EnumDescriptor {
	return file_beerproto_v1_style_proto_enumTypes[0].Descriptor()
}

func (StyleCategories) Type() protoreflect.EnumType {
	return &file_beerproto_v1_style_proto_enumTypes[0]
}

func (x StyleCategories) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StyleCategories.Descriptor instead.
func (StyleCategories) EnumDescriptor() ([]byte, []int) {
	return file_beerproto_v1_style_proto_rawDescGZIP(), []int{0}
}

// StyleType provide information for Style categorization
type StyleType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                           string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Aroma                        string                `protobuf:"bytes,2,opt,name=aroma,proto3" json:"aroma,omitempty"`
	Ingredients                  string                `protobuf:"bytes,3,opt,name=ingredients,proto3" json:"ingredients,omitempty"`
	CategoryNumber               int32                 `protobuf:"varint,4,opt,name=category_number,json=categoryNumber,proto3" json:"category_number,omitempty"`
	Notes                        string                `protobuf:"bytes,5,opt,name=notes,proto3" json:"notes,omitempty"`
	Flavor                       string                `protobuf:"bytes,6,opt,name=flavor,proto3" json:"flavor,omitempty"`
	Mouthfeel                    string                `protobuf:"bytes,7,opt,name=mouthfeel,proto3" json:"mouthfeel,omitempty"`
	FinalGravity                 *GravityRangeType     `protobuf:"bytes,8,opt,name=final_gravity,json=finalGravity,proto3" json:"final_gravity,omitempty"`
	StyleGuide                   string                `protobuf:"bytes,9,opt,name=style_guide,json=styleGuide,proto3" json:"style_guide,omitempty"`
	Color                        *ColorRangeType       `protobuf:"bytes,10,opt,name=color,proto3" json:"color,omitempty"`
	OriginalGravity              *GravityRangeType     `protobuf:"bytes,11,opt,name=original_gravity,json=originalGravity,proto3" json:"original_gravity,omitempty"`
	Examples                     string                `protobuf:"bytes,12,opt,name=examples,proto3" json:"examples,omitempty"`
	Name                         string                `protobuf:"bytes,13,opt,name=name,proto3" json:"name,omitempty"`
	Carbonation                  *CarbonationRangeType `protobuf:"bytes,14,opt,name=carbonation,proto3" json:"carbonation,omitempty"`
	AlcoholByVolume              *PercentRangeType     `protobuf:"bytes,15,opt,name=alcohol_by_volume,json=alcoholByVolume,proto3" json:"alcohol_by_volume,omitempty"`
	InternationalBitternessUnits *BitternessRangeType  `protobuf:"bytes,16,opt,name=international_bitterness_units,json=internationalBitternessUnits,proto3" json:"international_bitterness_units,omitempty"`
	Appearance                   string                `protobuf:"bytes,17,opt,name=appearance,proto3" json:"appearance,omitempty"`
	Category                     string                `protobuf:"bytes,18,opt,name=category,proto3" json:"category,omitempty"`
	StyleLetter                  string                `protobuf:"bytes,19,opt,name=style_letter,json=styleLetter,proto3" json:"style_letter,omitempty"`
	Type                         StyleCategories       `protobuf:"varint,20,opt,name=type,proto3,enum=beerproto.v1.StyleCategories" json:"type,omitempty"`
	OverallImpression            string                `protobuf:"bytes,21,opt,name=overall_impression,json=overallImpression,proto3" json:"overall_impression,omitempty"`
}

func (x *StyleType) Reset() {
	*x = StyleType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beerproto_v1_style_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StyleType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StyleType) ProtoMessage() {}

func (x *StyleType) ProtoReflect() protoreflect.Message {
	mi := &file_beerproto_v1_style_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StyleType.ProtoReflect.Descriptor instead.
func (*StyleType) Descriptor() ([]byte, []int) {
	return file_beerproto_v1_style_proto_rawDescGZIP(), []int{0}
}

func (x *StyleType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StyleType) GetAroma() string {
	if x != nil {
		return x.Aroma
	}
	return ""
}

func (x *StyleType) GetIngredients() string {
	if x != nil {
		return x.Ingredients
	}
	return ""
}

func (x *StyleType) GetCategoryNumber() int32 {
	if x != nil {
		return x.CategoryNumber
	}
	return 0
}

func (x *StyleType) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *StyleType) GetFlavor() string {
	if x != nil {
		return x.Flavor
	}
	return ""
}

func (x *StyleType) GetMouthfeel() string {
	if x != nil {
		return x.Mouthfeel
	}
	return ""
}

func (x *StyleType) GetFinalGravity() *GravityRangeType {
	if x != nil {
		return x.FinalGravity
	}
	return nil
}

func (x *StyleType) GetStyleGuide() string {
	if x != nil {
		return x.StyleGuide
	}
	return ""
}

func (x *StyleType) GetColor() *ColorRangeType {
	if x != nil {
		return x.Color
	}
	return nil
}

func (x *StyleType) GetOriginalGravity() *GravityRangeType {
	if x != nil {
		return x.OriginalGravity
	}
	return nil
}

func (x *StyleType) GetExamples() string {
	if x != nil {
		return x.Examples
	}
	return ""
}

func (x *StyleType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StyleType) GetCarbonation() *CarbonationRangeType {
	if x != nil {
		return x.Carbonation
	}
	return nil
}

func (x *StyleType) GetAlcoholByVolume() *PercentRangeType {
	if x != nil {
		return x.AlcoholByVolume
	}
	return nil
}

func (x *StyleType) GetInternationalBitternessUnits() *BitternessRangeType {
	if x != nil {
		return x.InternationalBitternessUnits
	}
	return nil
}

func (x *StyleType) GetAppearance() string {
	if x != nil {
		return x.Appearance
	}
	return ""
}

func (x *StyleType) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *StyleType) GetStyleLetter() string {
	if x != nil {
		return x.StyleLetter
	}
	return ""
}

func (x *StyleType) GetType() StyleCategories {
	if x != nil {
		return x.Type
	}
	return StyleCategories_STYLE_CATEGORIES_UNSPECIFIED
}

func (x *StyleType) GetOverallImpression() string {
	if x != nil {
		return x.OverallImpression
	}
	return ""
}

// RecipeStyleType defines style information stored in a recipe record
type RecipeStyleType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type           StyleCategories `protobuf:"varint,1,opt,name=type,proto3,enum=beerproto.v1.StyleCategories" json:"type,omitempty"`
	Name           string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Category       string          `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	CategoryNumber int32           `protobuf:"varint,4,opt,name=category_number,json=categoryNumber,proto3" json:"category_number,omitempty"`
	StyleLetter    string          `protobuf:"bytes,5,opt,name=style_letter,json=styleLetter,proto3" json:"style_letter,omitempty"`
	StyleGuide     string          `protobuf:"bytes,6,opt,name=style_guide,json=styleGuide,proto3" json:"style_guide,omitempty"`
}

func (x *RecipeStyleType) Reset() {
	*x = RecipeStyleType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beerproto_v1_style_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecipeStyleType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecipeStyleType) ProtoMessage() {}

func (x *RecipeStyleType) ProtoReflect() protoreflect.Message {
	mi := &file_beerproto_v1_style_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecipeStyleType.ProtoReflect.Descriptor instead.
func (*RecipeStyleType) Descriptor() ([]byte, []int) {
	return file_beerproto_v1_style_proto_rawDescGZIP(), []int{1}
}

func (x *RecipeStyleType) GetType() StyleCategories {
	if x != nil {
		return x.Type
	}
	return StyleCategories_STYLE_CATEGORIES_UNSPECIFIED
}

func (x *RecipeStyleType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RecipeStyleType) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *RecipeStyleType) GetCategoryNumber() int32 {
	if x != nil {
		return x.CategoryNumber
	}
	return 0
}

func (x *RecipeStyleType) GetStyleLetter() string {
	if x != nil {
		return x.StyleLetter
	}
	return ""
}

func (x *RecipeStyleType) GetStyleGuide() string {
	if x != nil {
		return x.StyleGuide
	}
	return ""
}

var File_beerproto_v1_style_proto protoreflect.FileDescriptor

var file_beerproto_v1_style_proto_rawDesc = []byte{
	0x0a, 0x18, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x74, 0x79, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62, 0x65, 0x65, 0x72,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99,
	0x07, 0x0a, 0x09, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x72, 0x6f, 0x6d, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x72, 0x6f,
	0x6d, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x64, 0x69,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f,
	0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6c, 0x61, 0x76, 0x6f, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6c, 0x61, 0x76, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6d,
	0x6f, 0x75, 0x74, 0x68, 0x66, 0x65, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x6f, 0x75, 0x74, 0x68, 0x66, 0x65, 0x65, 0x6c, 0x12, 0x43, 0x0a, 0x0d, 0x66, 0x69, 0x6e,
	0x61, 0x6c, 0x5f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x72, 0x61, 0x76, 0x69, 0x74, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x0c, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x47, 0x72, 0x61, 0x76, 0x69, 0x74, 0x79, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x12,
	0x32, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6c, 0x6f, 0x72, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x12, 0x49, 0x0a, 0x10, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f,
	0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x61,
	0x76, 0x69, 0x74, 0x79, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x47, 0x72, 0x61, 0x76, 0x69, 0x74, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x44,
	0x0a, 0x0b, 0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x11, 0x61, 0x6c, 0x63, 0x6f, 0x68, 0x6f, 0x6c, 0x5f,
	0x62, 0x79, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0f, 0x61, 0x6c, 0x63, 0x6f, 0x68, 0x6f, 0x6c, 0x42, 0x79, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x12, 0x67, 0x0a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x5f, 0x62, 0x69, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x75, 0x6e, 0x69,
	0x74, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x65,
	0x73, 0x73, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x1c, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x42, 0x69, 0x74, 0x74, 0x65, 0x72,
	0x6e, 0x65, 0x73, 0x73, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x70, 0x70,
	0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61,
	0x70, 0x70, 0x65, 0x61, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x5f, 0x6c,
	0x65, 0x74, 0x74, 0x65, 0x72, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x79,
	0x6c, 0x65, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x6f,
	0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c,
	0x49, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xe1, 0x01, 0x0a, 0x0f, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x65, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x31,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x62,
	0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x79, 0x6c,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74,
	0x79, 0x6c, 0x65, 0x5f, 0x6c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x2a, 0xf6,
	0x01, 0x0a, 0x0f, 0x53, 0x74, 0x79, 0x6c, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x54, 0x59, 0x4c, 0x45, 0x5f, 0x43, 0x41, 0x54, 0x45,
	0x47, 0x4f, 0x52, 0x49, 0x45, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x54, 0x59, 0x4c, 0x45, 0x5f, 0x43, 0x41,
	0x54, 0x45, 0x47, 0x4f, 0x52, 0x49, 0x45, 0x53, 0x5f, 0x42, 0x45, 0x45, 0x52, 0x10, 0x01, 0x12,
	0x1a, 0x0a, 0x16, 0x53, 0x54, 0x59, 0x4c, 0x45, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52,
	0x49, 0x45, 0x53, 0x5f, 0x43, 0x49, 0x44, 0x45, 0x52, 0x10, 0x02, 0x12, 0x1d, 0x0a, 0x19, 0x53,
	0x54, 0x59, 0x4c, 0x45, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x49, 0x45, 0x53, 0x5f,
	0x4b, 0x4f, 0x4d, 0x42, 0x55, 0x43, 0x48, 0x41, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x54,
	0x59, 0x4c, 0x45, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x49, 0x45, 0x53, 0x5f, 0x4d,
	0x45, 0x41, 0x44, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x54, 0x59, 0x4c, 0x45, 0x5f, 0x43,
	0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x49, 0x45, 0x53, 0x5f, 0x53, 0x4f, 0x44, 0x41, 0x10, 0x05,
	0x12, 0x19, 0x0a, 0x15, 0x53, 0x54, 0x59, 0x4c, 0x45, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f,
	0x52, 0x49, 0x45, 0x53, 0x5f, 0x57, 0x49, 0x4e, 0x45, 0x10, 0x06, 0x12, 0x1a, 0x0a, 0x16, 0x53,
	0x54, 0x59, 0x4c, 0x45, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x49, 0x45, 0x53, 0x5f,
	0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x07, 0x42, 0xab, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e,
	0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x53, 0x74,
	0x79, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x62, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x6f, 0x2f, 0x62, 0x65,
	0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x65, 0x65, 0x72, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x42,
	0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x42, 0x65,
	0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x42, 0x65, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x42, 0x65, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_beerproto_v1_style_proto_rawDescOnce sync.Once
	file_beerproto_v1_style_proto_rawDescData = file_beerproto_v1_style_proto_rawDesc
)

func file_beerproto_v1_style_proto_rawDescGZIP() []byte {
	file_beerproto_v1_style_proto_rawDescOnce.Do(func() {
		file_beerproto_v1_style_proto_rawDescData = protoimpl.X.CompressGZIP(file_beerproto_v1_style_proto_rawDescData)
	})
	return file_beerproto_v1_style_proto_rawDescData
}

var file_beerproto_v1_style_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_beerproto_v1_style_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_beerproto_v1_style_proto_goTypes = []interface{}{
	(StyleCategories)(0),         // 0: beerproto.v1.StyleCategories
	(*StyleType)(nil),            // 1: beerproto.v1.StyleType
	(*RecipeStyleType)(nil),      // 2: beerproto.v1.RecipeStyleType
	(*GravityRangeType)(nil),     // 3: beerproto.v1.GravityRangeType
	(*ColorRangeType)(nil),       // 4: beerproto.v1.ColorRangeType
	(*CarbonationRangeType)(nil), // 5: beerproto.v1.CarbonationRangeType
	(*PercentRangeType)(nil),     // 6: beerproto.v1.PercentRangeType
	(*BitternessRangeType)(nil),  // 7: beerproto.v1.BitternessRangeType
}
var file_beerproto_v1_style_proto_depIdxs = []int32{
	3, // 0: beerproto.v1.StyleType.final_gravity:type_name -> beerproto.v1.GravityRangeType
	4, // 1: beerproto.v1.StyleType.color:type_name -> beerproto.v1.ColorRangeType
	3, // 2: beerproto.v1.StyleType.original_gravity:type_name -> beerproto.v1.GravityRangeType
	5, // 3: beerproto.v1.StyleType.carbonation:type_name -> beerproto.v1.CarbonationRangeType
	6, // 4: beerproto.v1.StyleType.alcohol_by_volume:type_name -> beerproto.v1.PercentRangeType
	7, // 5: beerproto.v1.StyleType.international_bitterness_units:type_name -> beerproto.v1.BitternessRangeType
	0, // 6: beerproto.v1.StyleType.type:type_name -> beerproto.v1.StyleCategories
	0, // 7: beerproto.v1.RecipeStyleType.type:type_name -> beerproto.v1.StyleCategories
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_beerproto_v1_style_proto_init() }
func file_beerproto_v1_style_proto_init() {
	if File_beerproto_v1_style_proto != nil {
		return
	}
	file_beerproto_v1_measureable_units_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_beerproto_v1_style_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StyleType); i {
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
		file_beerproto_v1_style_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecipeStyleType); i {
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
			RawDescriptor: file_beerproto_v1_style_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_beerproto_v1_style_proto_goTypes,
		DependencyIndexes: file_beerproto_v1_style_proto_depIdxs,
		EnumInfos:         file_beerproto_v1_style_proto_enumTypes,
		MessageInfos:      file_beerproto_v1_style_proto_msgTypes,
	}.Build()
	File_beerproto_v1_style_proto = out.File
	file_beerproto_v1_style_proto_rawDesc = nil
	file_beerproto_v1_style_proto_goTypes = nil
	file_beerproto_v1_style_proto_depIdxs = nil
}
