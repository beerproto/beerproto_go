package beerprotov1

import "fmt"

func (x *MassType) DisplayName() string {
	if x == nil {
		return ""
	}
	return fmt.Sprintf("%v %v", x.Value, x.Unit.DisplayName())
}
