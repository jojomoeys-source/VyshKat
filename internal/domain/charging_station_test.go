package domain

import "testing"

var (
	_ InventoryItem  = (*ChargingStation)(nil)
	_ EnergyConsumer = (*ChargingStation)(nil)
)

func TestNewChargingStationWithValidData(t *testing.T) {
	station, err := NewChargingStation("Charging station", 301, 4.5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if station.Name() != "Charging station" {
		t.Errorf("expected name %q, got %q", "Charging station", station.Name())
	}
	if station.InventoryNumber() != 301 {
		t.Errorf("expected inventory number 301, got %d", station.InventoryNumber())
	}
	if station.DailyEnergyConsumption() != 4.5 {
		t.Errorf("expected daily energy consumption 4.5, got %f", station.DailyEnergyConsumption())
	}
}

func TestNewChargingStationRejectsInvalidData(t *testing.T) {
	tests := []struct {
		name                   string
		stationName            string
		inventoryNumber        int
		dailyEnergyConsumption float64
	}{
		{name: "empty name", stationName: "", inventoryNumber: 301, dailyEnergyConsumption: 4.5},
		{name: "whitespace name", stationName: "   ", inventoryNumber: 301, dailyEnergyConsumption: 4.5},
		{name: "non-positive inventory number", stationName: "Charging station", inventoryNumber: 0, dailyEnergyConsumption: 4.5},
		{name: "negative energy consumption", stationName: "Charging station", inventoryNumber: 301, dailyEnergyConsumption: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			station, err := NewChargingStation(
				test.stationName,
				test.inventoryNumber,
				test.dailyEnergyConsumption,
			)
			if err == nil {
				t.Error("expected validation error, got nil")
			}
			if station != nil {
				t.Errorf("expected charging station to be nil, got %v", station)
			}
		})
	}
}
