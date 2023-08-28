package beerprotov1

import "google.golang.org/protobuf/proto"

func (x VolumeUnit) StringName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x MassUnit) StringName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x DiastaticPowerUnit) StringName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x TemperatureUnit) StringName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x AcidityUnit) StringName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x TimeUnit) StringName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x ColorUnit) StringName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x CarbonationUnit) StringName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x BitternessUnit) StringName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x GravityUnit) StringName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x SpecificHeatUnit) StringName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x ConcentrationUnit) StringName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x SpecificVolumeUnit) StringName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x UnitUnit) StringName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x PercentUnit) StringName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}

func (x ViscosityUnit) StringName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}
