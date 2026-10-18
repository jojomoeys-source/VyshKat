package domain

import "testing"

func TestNewTransportWithValidData(t *testing.T) {
	transport, err := NewTransport("Car", 1, 5, true)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if transport.Name() != "Car" {
		t.Errorf("expected name 'Car', got '%s'", transport.Name())
	}

	if transport.InventoryNumber() != 1 {
		t.Errorf("expected inventory number 1, got %d", transport.InventoryNumber())
	}

	if transport.BeginnerFriendliness() != 5 {
		t.Errorf("expected beginner friendliness 5, got %d", transport.BeginnerFriendliness())
	}

	if !transport.IsOperational() {
		t.Errorf("expected operational to be true, got false")
	}
}

func TestNewTransportWithInvalidName(t *testing.T) {
	transport, err := NewTransport("", 1, 5, true)

	if err == nil {
		t.Error("expected error for empty name, got nil")
	}

	if transport != nil {
		t.Errorf("expected transport to be nil, got %v", transport)
	}
}

func TestNewTransportWithWhitespaceName(t *testing.T) {
	transport, err := NewTransport("     ", 1, 5, true)

	if err == nil {
		t.Error("expected error for whitespace name, got nil")
	}

	if transport != nil {
		t.Errorf("expected transport to be nil, got %v", transport)
	}
}

func TestNewTransportWithInvalidInventoryNumber(t *testing.T) {
	transport, err := NewTransport("Car", 0, 5, true)

	if err == nil {
		t.Error("expected error for non-positive inventory number, got nil")
	}

	if transport != nil {
		t.Errorf("expected transport to be nil, got %v", transport)
	}
}

func TestNewTransportWithInvalidBeginnerFriendliness(t *testing.T) {
	transport, err := NewTransport("Car", 1, 11, true)

	if err == nil {
		t.Error("expected error for beginner friendliness out of range, got nil")
	}

	if transport != nil {
		t.Errorf("expected transport to be nil, got %v", transport)
	}
}
