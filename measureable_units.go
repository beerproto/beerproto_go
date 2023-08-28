package beerprotov1

import "google.golang.org/protobuf/proto"

func (x VolumeUnit) DisplayName() string {
	test := x.Descriptor().Values().ByNumber(x.Number()).Options()
	tt := proto.GetExtension(test, E_StringName)
	return tt.(string)
}
