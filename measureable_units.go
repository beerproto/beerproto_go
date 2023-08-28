package beerprotov1

import "google.golang.org/protobuf/proto"

func (x VolumeUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x MassUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x DiastaticPowerUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x TemperatureUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x AcidityUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x TimeUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x ColorUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x CarbonationUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x BitternessUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x GravityUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x SpecificHeatUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x ConcentrationUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x SpecificVolumeUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x UnitUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x PercentUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x ViscosityUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}
