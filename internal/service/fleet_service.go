package service

import (
	"errors"

	"github.com/jojomoeys-source/VyshKat/internal/domain"
)

var (
	ErrDuplicateInventoryNumber    = errors.New("inventory number is already in use")
	ErrInspectionFailed            = errors.New("transport failed inspection")
	ErrTransportRequiresInspection = errors.New("transport must be added through inspection")
)

type FleetService struct {
	serviceCenter ServiceCenter
	items         []domain.InventoryItem
}

func NewFleetService(serviceCenter ServiceCenter) *FleetService {
	return &FleetService{
		serviceCenter: serviceCenter,
		items:         make([]domain.InventoryItem, 0),
	}
}

func (f *FleetService) hasInventoryNumber(inventoryNumber int) bool {
	for _, item := range f.items {
		if item.InventoryNumber() == inventoryNumber {
			return true
		}
	}

	return false
}

func (f *FleetService) AddTransport(transport domain.TransportItem) error {
	if f.hasInventoryNumber(transport.InventoryNumber()) {
		return ErrDuplicateInventoryNumber
	}

	if !f.serviceCenter.Inspect(transport) {
		return ErrInspectionFailed
	}

	f.items = append(f.items, transport)

	return nil
}

func (f *FleetService) AddItem(item domain.InventoryItem) error {
	if _, isTransport := item.(domain.TransportItem); isTransport {
		return ErrTransportRequiresInspection
	}

	if f.hasInventoryNumber(item.InventoryNumber()) {
		return ErrDuplicateInventoryNumber
	}

	f.items = append(f.items, item)

	return nil
}

func (f *FleetService) ItemCount() int {
	return len(f.items)
}

func (f *FleetService) Items() []domain.InventoryItem {
	items := make([]domain.InventoryItem, len(f.items))
	copy(items, f.items)

	return items
}

func (f *FleetService) TotalDailyEnergyConsumption() float64 {
	var total float64

	for _, item := range f.items {
		consumer, consumesEnergy := item.(domain.EnergyConsumer)
		if consumesEnergy {
			total += consumer.DailyEnergyConsumption()
		}
	}

	return total
}

func (f *FleetService) BeginnerFriendlyTransports() []domain.TransportItem {
	transports := make([]domain.TransportItem, 0)

	for _, item := range f.items {
		transport, isTransport := item.(domain.TransportItem)
		if isTransport && transport.BeginnerFriendliness() >= 6 {
			transports = append(transports, transport)
		}
	}

	return transports
}
