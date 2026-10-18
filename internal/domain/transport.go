package domain

import (
	"errors"
	"strings"
)

type Transport struct {
	name                 string
	inventoryNumber      int
	beginnerFriendliness int
	operational          bool
}

func (t *Transport) Name() string {
	return t.name
}

func (t *Transport) InventoryNumber() int {
	return t.inventoryNumber
}

func (t *Transport) BeginnerFriendliness() int {
	return t.beginnerFriendliness
}

func (t *Transport) IsOperational() bool {
	return t.operational
}

func NewTransport(name string, inventoryNumber int, beginnerFriendliness int, operational bool) (*Transport, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name cannot be empty")
	}
	if inventoryNumber <= 0 {
		return nil, errors.New("inventory number must be positive")
	}
	if beginnerFriendliness < 1 || beginnerFriendliness > 10 {
		return nil, errors.New("beginner friendliness must be between 1 and 10")
	}
	return &Transport{
		name:                 name,
		inventoryNumber:      inventoryNumber,
		beginnerFriendliness: beginnerFriendliness,
		operational:          operational,
	}, nil
}

type TransportItem interface {
	InventoryItem
	BeginnerFriendliness() int
	IsOperational() bool
}
