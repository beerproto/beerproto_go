package beerprotov1

func (x *MiscellaneousAdditionType) DisplayName() string {
	mass := x.GetMass()
	if mass != nil {
		return mass.DisplayName()
	}
	volume := x.GetVolume()
	if volume != nil {
		return volume.DisplayName()
	}
	unit := x.GetUnit()

	if unit != nil {
		return unit.DisplayName()
	}
	return ""
}
