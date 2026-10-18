package domain

import "testing"

func TestNewScooterWithValidData(t *testing.T) {
	scooter, err := NewScooter("Electric scooter", 101, 8, true, 2.5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if scooter.Name() != "Electric scooter" {
		t.Errorf("expected name %q, got %q", "Electric scooter", scooter.Name())
	}

	if scooter.InventoryNumber() != 101 {
		t.Errorf("expected inventory number 101, got %d", scooter.InventoryNumber())
	}

	if scooter.BeginnerFriendliness() != 8 {
		t.Errorf("expected beginner friendliness 8, got %d", scooter.BeginnerFriendliness())
	}

	if !scooter.IsOperational() {
		t.Error("expected scooter to be operational")
	}

	if scooter.DailyEnergyConsumption() != 2.5 {
		t.Errorf("expected daily energy consumption 2.5, got %f", scooter.DailyEnergyConsumption())
	}
}

func TestNewScooterWithNegativeEnergyConsumption(t *testing.T) {
	scooter, err := NewScooter("Electric scooter", 101, 8, true, -0.5)

	if err == nil {
		t.Error("expected error for negative energy consumption, got nil")
	}

	if scooter != nil {
		t.Errorf("expected scooter to be nil, got %v", scooter)
	}
}

func TestNewScooterWithInvalidTransportData(t *testing.T) {
	scooter, err := NewScooter("", 101, 8, true, 2.5)

	if err == nil {
		t.Error("expected error for invalid transport data, got nil")
	}

	if scooter != nil {
		t.Errorf("expected scooter to be nil, got %v", scooter)
	}
}
