package service

import (
	"errors"
	"testing"

	"github.com/jojomoeys-source/VyshKat/internal/domain"
)

type approvingServiceCenter struct{}

func (approvingServiceCenter) Inspect(domain.TransportItem) bool {
	return true
}

type rejectingServiceCenter struct{}

func (rejectingServiceCenter) Inspect(domain.TransportItem) bool {
	return false
}

func TestAddTransportWhenInspectionApprovesAddsTransport(t *testing.T) {
	fleet := NewFleetService(approvingServiceCenter{})
	transport := mustNewBicycle(t, "City bicycle", 101)

	err := fleet.AddTransport(transport)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(fleet.items) != 1 {
		t.Fatalf("expected 1 inventory item, got %d", len(fleet.items))
	}

	if fleet.items[0] != transport {
		t.Error("expected accepted transport to be stored in the fleet")
	}
}

func TestAddTransportWhenInspectionRejectsDoesNotAddTransport(t *testing.T) {
	fleet := NewFleetService(rejectingServiceCenter{})
	transport := mustNewBicycle(t, "City bicycle", 101)

	err := fleet.AddTransport(transport)

	if !errors.Is(err, ErrInspectionFailed) {
		t.Fatalf("expected ErrInspectionFailed, got %v", err)
	}

	if len(fleet.items) != 0 {
		t.Errorf("expected empty inventory, got %d items", len(fleet.items))
	}
}

func TestAddTransportWithDuplicateInventoryNumberDoesNotAddSecondTransport(t *testing.T) {
	fleet := NewFleetService(approvingServiceCenter{})
	first := mustNewBicycle(t, "First bicycle", 101)
	second := mustNewBicycle(t, "Second bicycle", 101)

	if err := fleet.AddTransport(first); err != nil {
		t.Fatalf("failed to add first transport: %v", err)
	}

	err := fleet.AddTransport(second)

	if !errors.Is(err, ErrDuplicateInventoryNumber) {
		t.Fatalf("expected ErrDuplicateInventoryNumber, got %v", err)
	}

	if len(fleet.items) != 1 {
		t.Errorf("expected 1 inventory item, got %d", len(fleet.items))
	}
}

func TestAddItemAddsHelmetWithoutInspection(t *testing.T) {
	fleet := NewFleetService(rejectingServiceCenter{})
	helmet := mustNewHelmet(t, "Safety helmet", 201)

	err := fleet.AddItem(helmet)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(fleet.items) != 1 {
		t.Fatalf("expected 1 inventory item, got %d", len(fleet.items))
	}

	if fleet.items[0] != helmet {
		t.Error("expected helmet to be stored in the fleet")
	}
}

func TestAddItemWithDuplicateInventoryNumberDoesNotAddItem(t *testing.T) {
	fleet := NewFleetService(rejectingServiceCenter{})
	helmet := mustNewHelmet(t, "Safety helmet", 201)
	station := mustNewChargingStation(t, "Charging station", 201, 4.5)

	if err := fleet.AddItem(helmet); err != nil {
		t.Fatalf("failed to add first item: %v", err)
	}

	err := fleet.AddItem(station)

	if !errors.Is(err, ErrDuplicateInventoryNumber) {
		t.Fatalf("expected ErrDuplicateInventoryNumber, got %v", err)
	}

	if len(fleet.items) != 1 {
		t.Errorf("expected 1 inventory item, got %d", len(fleet.items))
	}
}

func TestAddItemRejectsTransportThatBypassesInspection(t *testing.T) {
	fleet := NewFleetService(approvingServiceCenter{})
	transport := mustNewBicycle(t, "City bicycle", 101)

	err := fleet.AddItem(transport)

	if !errors.Is(err, ErrTransportRequiresInspection) {
		t.Fatalf("expected ErrTransportRequiresInspection, got %v", err)
	}

	if len(fleet.items) != 0 {
		t.Errorf("expected empty inventory, got %d items", len(fleet.items))
	}
}

func TestItemCountAndItemsReturnAcceptedInventory(t *testing.T) {
	fleet := NewFleetService(approvingServiceCenter{})
	bicycle := mustNewBicycleWithFriendliness(t, "City bicycle", 101, 5)
	helmet := mustNewHelmet(t, "Safety helmet", 201)

	if err := fleet.AddTransport(bicycle); err != nil {
		t.Fatalf("failed to add bicycle: %v", err)
	}
	if err := fleet.AddItem(helmet); err != nil {
		t.Fatalf("failed to add helmet: %v", err)
	}

	if fleet.ItemCount() != 2 {
		t.Errorf("expected item count 2, got %d", fleet.ItemCount())
	}

	items := fleet.Items()
	if len(items) != 2 {
		t.Fatalf("expected 2 returned items, got %d", len(items))
	}

	items[0] = helmet
	if fleet.Items()[0] != bicycle {
		t.Error("expected Items to return a copy of the internal slice")
	}
}

func TestTotalDailyEnergyConsumptionSumsOnlyEnergyConsumers(t *testing.T) {
	fleet := NewFleetService(approvingServiceCenter{})
	scooter := mustNewScooter(t, "Electric scooter", 101, 8, 2.5)
	bicycle := mustNewBicycleWithFriendliness(t, "City bicycle", 102, 5)
	helmet := mustNewHelmet(t, "Safety helmet", 201)
	station := mustNewChargingStation(t, "Charging station", 202, 4.5)

	if err := fleet.AddTransport(scooter); err != nil {
		t.Fatalf("failed to add scooter: %v", err)
	}
	if err := fleet.AddTransport(bicycle); err != nil {
		t.Fatalf("failed to add bicycle: %v", err)
	}
	if err := fleet.AddItem(helmet); err != nil {
		t.Fatalf("failed to add helmet: %v", err)
	}
	if err := fleet.AddItem(station); err != nil {
		t.Fatalf("failed to add charging station: %v", err)
	}

	if total := fleet.TotalDailyEnergyConsumption(); total != 7.0 {
		t.Errorf("expected total energy consumption 7.0, got %f", total)
	}
}

func TestBeginnerFriendlyTransportsReturnsOnlyTransportsWithScoreAtLeastSix(t *testing.T) {
	fleet := NewFleetService(approvingServiceCenter{})
	difficult := mustNewBicycleWithFriendliness(t, "Difficult bicycle", 101, 5)
	boundary := mustNewBicycleWithFriendliness(t, "Beginner bicycle", 102, 6)
	easy := mustNewScooter(t, "Easy scooter", 103, 9, 2.5)
	helmet := mustNewHelmet(t, "Safety helmet", 201)

	for _, transport := range []domain.TransportItem{difficult, boundary, easy} {
		if err := fleet.AddTransport(transport); err != nil {
			t.Fatalf("failed to add transport: %v", err)
		}
	}
	if err := fleet.AddItem(helmet); err != nil {
		t.Fatalf("failed to add helmet: %v", err)
	}

	transports := fleet.BeginnerFriendlyTransports()
	if len(transports) != 2 {
		t.Fatalf("expected 2 beginner-friendly transports, got %d", len(transports))
	}

	if transports[0] != boundary || transports[1] != easy {
		t.Error("expected transports with beginner friendliness 6 and 9")
	}
}

func mustNewBicycle(t *testing.T, name string, inventoryNumber int) *domain.Bicycle {
	return mustNewBicycleWithFriendliness(t, name, inventoryNumber, 5)
}

func mustNewBicycleWithFriendliness(
	t *testing.T,
	name string,
	inventoryNumber int,
	beginnerFriendliness int,
) *domain.Bicycle {
	t.Helper()

	bicycle, err := domain.NewBicycle(name, inventoryNumber, beginnerFriendliness, true)
	if err != nil {
		t.Fatalf("failed to create test bicycle: %v", err)
	}

	return bicycle
}

func mustNewScooter(
	t *testing.T,
	name string,
	inventoryNumber int,
	beginnerFriendliness int,
	dailyEnergyConsumption float64,
) *domain.Scooter {
	t.Helper()

	scooter, err := domain.NewScooter(
		name,
		inventoryNumber,
		beginnerFriendliness,
		true,
		dailyEnergyConsumption,
	)
	if err != nil {
		t.Fatalf("failed to create test scooter: %v", err)
	}

	return scooter
}

func mustNewHelmet(t *testing.T, name string, inventoryNumber int) *domain.Helmet {
	t.Helper()

	helmet, err := domain.NewHelmet(name, inventoryNumber)
	if err != nil {
		t.Fatalf("failed to create test helmet: %v", err)
	}

	return helmet
}

func mustNewChargingStation(
	t *testing.T,
	name string,
	inventoryNumber int,
	dailyEnergyConsumption float64,
) *domain.ChargingStation {
	t.Helper()

	station, err := domain.NewChargingStation(name, inventoryNumber, dailyEnergyConsumption)
	if err != nil {
		t.Fatalf("failed to create test charging station: %v", err)
	}

	return station
}
