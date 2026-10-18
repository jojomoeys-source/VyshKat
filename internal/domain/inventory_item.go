package domain

type InventoryItem interface {
	Name() string
	InventoryNumber() int
}
