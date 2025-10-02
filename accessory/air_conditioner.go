package accessory

import (
	"github.com/brutella/hap/service"
)

type AirConditioner struct {
	*A
	AirConditioner *service.AirConditioner
}

// NewAirConditioner returns an air conditioner accessory.
func NewAirConditioner(info Info) *AirConditioner {
	a := AirConditioner{}
	a.A = New(info, TypeAirConditioner)

	a.AirConditioner = service.NewAirConditioner()
	a.AddS(a.AirConditioner.S)

	return &a
}
