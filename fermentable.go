package beerprotov1

func (x *FermentableAdditionType) DisplayName() string {
	if x == nil {
		return ""
	}
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
