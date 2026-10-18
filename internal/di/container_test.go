package di

import (
	"testing"

	"github.com/jojomoeys-source/VyshKat/internal/service"
)

func TestBuildContainerResolvesFleetService(t *testing.T) {
	container, err := BuildContainer()
	if err != nil {
		t.Fatalf("failed to build container: %v", err)
	}

	invoked := false
	err = container.Invoke(func(fleet *service.FleetService) {
		invoked = true

		if fleet == nil {
			t.Fatal("expected resolved fleet service, got nil")
		}

		if fleet.ItemCount() != 0 {
			t.Errorf("expected empty fleet, got %d items", fleet.ItemCount())
		}
	})
	if err != nil {
		t.Fatalf("failed to resolve fleet service: %v", err)
	}

	if !invoked {
		t.Error("expected Invoke callback to be called")
	}
}
