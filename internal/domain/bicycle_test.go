package domain

import "testing"

var (
	_ InventoryItem = (*Bicycle)(nil)
	_ TransportItem = (*Bicycle)(nil)
)

func TestNewBicycleWithValidData(t *testing.T) {
	bicycle, err := NewBicycle("City bicycle", 101, 6, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if bicycle.Name() != "City bicycle" {
		t.Errorf("expected name %q, got %q", "City bicycle", bicycle.Name())
	}
	if bicycle.InventoryNumber() != 101 {
		t.Errorf("expected inventory number 101, got %d", bicycle.InventoryNumber())
	}
	if bicycle.BeginnerFriendliness() != 6 {
		t.Errorf("expected beginner friendliness 6, got %d", bicycle.BeginnerFriendliness())
	}
	if !bicycle.IsOperational() {
		t.Error("expected bicycle to be operational")
	}
}

func TestNewBicyclePropagatesTransportValidationError(t *testing.T) {
	bicycle, err := NewBicycle("", 101, 6, true)

	if err == nil {
		t.Error("expected error for invalid transport data, got nil")
	}
	if bicycle != nil {
		t.Errorf("expected bicycle to be nil, got %v", bicycle)
	}
}
