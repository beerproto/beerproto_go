package beerprotov1

func (x *HopAdditionType) DisplayName() string {
	mass := x.GetMass()
	if mass != nil {
		return mass.DisplayName()
	}
	volume := x.GetVolume()
	if volume != nil {
		return volume.DisplayName()
	}

	return ""
}
