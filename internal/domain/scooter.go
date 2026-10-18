package domain

import "errors"

type Scooter struct {
	*Transport
	dailyEnergyConsumption float64
}

func (s *Scooter) DailyEnergyConsumption() float64 {
	return s.dailyEnergyConsumption
}

func NewScooter(
	name string,
	inventoryNumber int,
	beginnerFriendliness int,
	operational bool,
	dailyEnergyConsumption float64,
) (*Scooter, error) {
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

	return &Scooter{
		Transport:              transport,
		dailyEnergyConsumption: dailyEnergyConsumption,
	}, nil
}
