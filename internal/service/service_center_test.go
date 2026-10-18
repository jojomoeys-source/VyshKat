package service

import (
	"testing"

	"github.com/jojomoeys-source/VyshKat/internal/domain"
)

func TestOperationalServiceCenterAcceptsOperationalTransport(t *testing.T) {
	serviceCenter := NewOperationalServiceCenter()
	transport := mustNewBicycleWithOperationalState(t, "Operational bicycle", 101, true)

	if !serviceCenter.Inspect(transport) {
		t.Error("expected operational transport to pass inspection")
	}
}

func TestOperationalServiceCenterRejectsNonOperationalTransport(t *testing.T) {
	serviceCenter := NewOperationalServiceCenter()
	transport := mustNewBicycleWithOperationalState(t, "Broken bicycle", 102, false)

	if serviceCenter.Inspect(transport) {
		t.Error("expected non-operational transport to fail inspection")
	}
}

func mustNewBicycleWithOperationalState(
	t *testing.T,
	name string,
	inventoryNumber int,
	operational bool,
) *domain.Bicycle {
	t.Helper()

	bicycle, err := domain.NewBicycle(name, inventoryNumber, 5, operational)
	if err != nil {
		t.Fatalf("failed to create test bicycle: %v", err)
	}

	return bicycle
}
