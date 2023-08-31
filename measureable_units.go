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

func (x VolumeUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_DisplayName)
	return tt.(string)
}

func (x MassUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_DisplayName)
	return tt.(string)
}

func (x DiastaticPowerUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_DisplayName)
	return tt.(string)
}

func (x TemperatureUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_DisplayName)
	return tt.(string)
}

func (x AcidityUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_DisplayName)
	return tt.(string)
}

func (x TimeUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_DisplayName)
	return tt.(string)
}

func (x ColorUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_DisplayName)
	return tt.(string)
}

func (x CarbonationUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_DisplayName)
	return tt.(string)
}

func (x BitternessUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_DisplayName)
	return tt.(string)
}

func (x GravityUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_DisplayName)
	return tt.(string)
}

func (x SpecificHeatUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_DisplayName)
	return tt.(string)
}

func (x ConcentrationUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_DisplayName)
	return tt.(string)
}

func (x SpecificVolumeUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_DisplayName)
	return tt.(string)
}

func (x UnitUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_DisplayName)
	return tt.(string)
}

func (x PercentUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_DisplayName)
	return tt.(string)
}

func (x ViscosityUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_DisplayName)
	return tt.(string)
}

func (x *VolumeType) Convert(unit VolumeUnit) *VolumeType {
	test := x.Unit.Descriptor().Values().ByNumber(x.Unit.Number()).Options()
	tt := proto.GetExtension(test, E_ConversionVolumeUnit)
	response := &VolumeType{
		Unit:  unit,
		Value: x.Value,
	}
	c, ok := tt.(*ConversionVolumeUnit)
	if ok {
		for _, r := range c.Rates {
			if r.Target == unit {
				switch r.Operator {
				case ArithmeticOperators_ARITHMETIC_OPERATORS_ADDITION:
					response.Value = x.Value + r.Value
				case ArithmeticOperators_ARITHMETIC_OPERATORS_SUBTRACTION:
					response.Value = x.Value - r.Value
				case ArithmeticOperators_ARITHMETIC_OPERATORS_MULTIPLICATION:
					response.Value = x.Value * r.Value
				case ArithmeticOperators_ARITHMETIC_OPERATORS_DIVISION:
					response.Value = x.Value / r.Value
				case ArithmeticOperators_ARITHMETIC_OPERATORS_UNSPECIFIED:
					response.Value = x.Value
				}
				break
			}
		}
	}

	return response
}
