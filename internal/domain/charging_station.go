package domain

import (
	"errors"
	"strings"
)

type ChargingStation struct {
	name                   string
	inventoryNumber        int
	dailyEnergyConsumption float64
}

func NewChargingStation(
	name string,
	inventoryNumber int,
	dailyEnergyConsumption float64,
) (*ChargingStation, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name cannot be empty")
	}

	if inventoryNumber <= 0 {
		return nil, errors.New("inventory number must be positive")
	}

	if dailyEnergyConsumption < 0 {
		return nil, errors.New("daily energy consumption cannot be negative")
	}

	return &ChargingStation{
		name:                   name,
		inventoryNumber:        inventoryNumber,
		dailyEnergyConsumption: dailyEnergyConsumption,
	}, nil
}

func (s *ChargingStation) Name() string {
	return s.name
}

func (s *ChargingStation) InventoryNumber() int {
	return s.inventoryNumber
}

func (s *ChargingStation) DailyEnergyConsumption() float64 {
	return s.dailyEnergyConsumption
}
