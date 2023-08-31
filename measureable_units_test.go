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
