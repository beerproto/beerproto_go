package beerprotov1

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {

	vol := &VolumeType{
		Unit:  VolumeUnit_VOLUME_UNIT_ML,
		Value: 5,
	}

	got := vol.Convert(VolumeUnit_VOLUME_UNIT_L)

	if got.Value != 0.005 {
		t.Errorf("Convert(5) = %v; want 0.005", got.Value)
	}
	fmt.Print(vol)

}

func Test_DiastaticPowerTypeTo_WK(t *testing.T) {

	vol := &DiastaticPowerType{
		Unit:  DiastaticPowerUnit_DIASTATIC_POWER_UNIT_LINTNER,
		Value: 5,
	}

	got := vol.Convert(DiastaticPowerUnit_DIASTATIC_POWER_UNIT_WK)

	if got.Value != 1.5 {
		t.Errorf("Convert(5) = %v; want 1.5", got.Value)
	}
	fmt.Print(vol)

}

func Test_DiastaticPowerType(t *testing.T) {

	vol := &DiastaticPowerType{
		Unit:  DiastaticPowerUnit_DIASTATIC_POWER_UNIT_WK,
		Value: 5,
	}

	got := vol.Convert(DiastaticPowerUnit_DIASTATIC_POWER_UNIT_LINTNER)

	if got.Value != 6 {
		t.Errorf("Convert(5) = %v; want 6", got.Value)
	}
	fmt.Print(vol)

}
