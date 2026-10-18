package domain

import "errors"

type ElectricBicycle struct {
	*Transport
	dailyEnergyConsumption float64
}

func (b *ElectricBicycle) DailyEnergyConsumption() float64 {
	return b.dailyEnergyConsumption
}

func NewElectricBicycle(
	name string,
	inventoryNumber int,
	beginnerFriendliness int,
	operational bool,
	dailyEnergyConsumption float64,
) (*ElectricBicycle, error) {
	transport, err := NewTransport(
		name,
		inventoryNumber,
		beginnerFriendliness,
		operational,
	)
	if err != nil {
		return nil, err
	}

	if dailyEnergyConsumption < 0 {
		return nil, errors.New("daily energy consumption cannot be negative")
	}

	return &ElectricBicycle{
		Transport:              transport,
		dailyEnergyConsumption: dailyEnergyConsumption,
	}, nil
}
