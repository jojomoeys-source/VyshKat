package domain

import (
	"errors"
	"strings"
)

type Helmet struct {
	name            string
	inventoryNumber int
}

func NewHelmet(name string, inventoryNumber int) (*Helmet, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name cannot be empty")
	}

	if inventoryNumber <= 0 {
		return nil, errors.New("inventory number must be positive")
	}

	return &Helmet{
		name:            name,
		inventoryNumber: inventoryNumber,
	}, nil
}

func (h *Helmet) Name() string {
	return h.name
}

func (h *Helmet) InventoryNumber() int {
	return h.inventoryNumber
}
