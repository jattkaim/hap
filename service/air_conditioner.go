package service

import (
	"github.com/brutella/hap/characteristic"
)

type AirConditioner struct {
	*S

	Active                      *characteristic.Active
	CurrentHeaterCoolerState    *characteristic.CurrentHeaterCoolerState
	TargetHeaterCoolerState     *characteristic.TargetHeaterCoolerState
	CurrentTemperature          *characteristic.CurrentTemperature
	CoolingThresholdTemperature *characteristic.CoolingThresholdTemperature
	HeatingThresholdTemperature *characteristic.HeatingThresholdTemperature
	SwingMode                   *characteristic.SwingMode
	RotationSpeed               *characteristic.RotationSpeed
	TemperatureDisplayUnits     *characteristic.TemperatureDisplayUnits
	LockPhysicalControls        *characteristic.LockPhysicalControls
}

func NewAirConditioner() *AirConditioner {
	s := AirConditioner{}
	s.S = New(TypeHeaterCooler)

	s.Active = characteristic.NewActive()
	s.AddC(s.Active.C)

	s.CurrentHeaterCoolerState = characteristic.NewCurrentHeaterCoolerState()
	s.AddC(s.CurrentHeaterCoolerState.C)

	s.TargetHeaterCoolerState = characteristic.NewTargetHeaterCoolerState()
	s.AddC(s.TargetHeaterCoolerState.C)

	s.CurrentTemperature = characteristic.NewCurrentTemperature()
	s.AddC(s.CurrentTemperature.C)

	// Add optional characteristics that are commonly used in air conditioners
	s.CoolingThresholdTemperature = characteristic.NewCoolingThresholdTemperature()
	s.AddC(s.CoolingThresholdTemperature.C)

	s.HeatingThresholdTemperature = characteristic.NewHeatingThresholdTemperature()
	s.AddC(s.HeatingThresholdTemperature.C)

	s.SwingMode = characteristic.NewSwingMode()
	s.AddC(s.SwingMode.C)

	s.RotationSpeed = characteristic.NewRotationSpeed()
	s.AddC(s.RotationSpeed.C)

	s.TemperatureDisplayUnits = characteristic.NewTemperatureDisplayUnits()
	s.AddC(s.TemperatureDisplayUnits.C)

	s.LockPhysicalControls = characteristic.NewLockPhysicalControls()
	s.AddC(s.LockPhysicalControls.C)

	return &s
}
