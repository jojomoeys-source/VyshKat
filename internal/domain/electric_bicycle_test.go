package domain

import "testing"

var (
	_ InventoryItem  = (*ElectricBicycle)(nil)
	_ TransportItem  = (*ElectricBicycle)(nil)
	_ EnergyConsumer = (*ElectricBicycle)(nil)
)

func TestNewElectricBicycleWithValidData(t *testing.T) {
	bicycle, err := NewElectricBicycle("Electric bicycle", 102, 8, true, 3.25)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if bicycle.Name() != "Electric bicycle" {
		t.Errorf("expected name %q, got %q", "Electric bicycle", bicycle.Name())
	}
	if bicycle.InventoryNumber() != 102 {
		t.Errorf("expected inventory number 102, got %d", bicycle.InventoryNumber())
	}
	if bicycle.BeginnerFriendliness() != 8 {
		t.Errorf("expected beginner friendliness 8, got %d", bicycle.BeginnerFriendliness())
	}
	if !bicycle.IsOperational() {
		t.Error("expected electric bicycle to be operational")
	}
	if bicycle.DailyEnergyConsumption() != 3.25 {
		t.Errorf("expected daily energy consumption 3.25, got %f", bicycle.DailyEnergyConsumption())
	}
}

func TestNewElectricBicycleWithNegativeEnergyConsumption(t *testing.T) {
	bicycle, err := NewElectricBicycle("Electric bicycle", 102, 8, true, -1)

	if err == nil {
		t.Error("expected error for negative energy consumption, got nil")
	}
	if bicycle != nil {
		t.Errorf("expected electric bicycle to be nil, got %v", bicycle)
	}
}

func TestNewElectricBicyclePropagatesTransportValidationError(t *testing.T) {
	bicycle, err := NewElectricBicycle("", 102, 8, true, 3.25)

	if err == nil {
		t.Error("expected error for invalid transport data, got nil")
	}
	if bicycle != nil {
		t.Errorf("expected electric bicycle to be nil, got %v", bicycle)
	}
}
