package domain

import "testing"

var _ InventoryItem = (*Helmet)(nil)

func TestNewHelmetWithValidData(t *testing.T) {
	helmet, err := NewHelmet("Safety helmet", 201)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if helmet.Name() != "Safety helmet" {
		t.Errorf("expected name %q, got %q", "Safety helmet", helmet.Name())
	}
	if helmet.InventoryNumber() != 201 {
		t.Errorf("expected inventory number 201, got %d", helmet.InventoryNumber())
	}
}

func TestNewHelmetRejectsInvalidData(t *testing.T) {
	tests := []struct {
		name            string
		helmetName      string
		inventoryNumber int
	}{
		{name: "empty name", helmetName: "", inventoryNumber: 201},
		{name: "whitespace name", helmetName: "   ", inventoryNumber: 201},
		{name: "non-positive inventory number", helmetName: "Safety helmet", inventoryNumber: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			helmet, err := NewHelmet(test.helmetName, test.inventoryNumber)
			if err == nil {
				t.Error("expected validation error, got nil")
			}
			if helmet != nil {
				t.Errorf("expected helmet to be nil, got %v", helmet)
			}
		})
	}
}
