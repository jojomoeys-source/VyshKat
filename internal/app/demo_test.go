package app

import (
	"bytes"
	"strings"
	"testing"

	"github.com/jojomoeys-source/VyshKat/internal/service"
)

func TestDemoScenarioShowsAllRequiredScenarios(t *testing.T) {
	fleet := service.NewFleetService(service.NewOperationalServiceCenter())
	var output bytes.Buffer
	demo := NewDemoScenario(fleet, &output)

	if err := demo.Run(); err != nil {
		t.Fatalf("demo failed: %v", err)
	}

	if fleet.ItemCount() != 4 {
		t.Errorf("expected 4 accepted objects, got %d", fleet.ItemCount())
	}
	if total := fleet.TotalDailyEnergyConsumption(); total != 7.0 {
		t.Errorf("expected total energy consumption 7.0, got %f", total)
	}
	if transports := fleet.BeginnerFriendlyTransports(); len(transports) != 1 {
		t.Errorf("expected 1 beginner-friendly transport, got %d", len(transports))
	}

	for _, item := range fleet.Items() {
		if item.Name() == "Broken electric bicycle" {
			t.Error("expected rejected transport to be absent from inventory")
		}
	}

	text := output.String()
	expectedFragments := []string{
		"Демонстрация возможностей ВышКат",
		"Принят: Electric scooter",
		"Отклонён техосмотром: Broken electric bicycle",
		"Инвентарь:",
		"Всего объектов: 4",
		"Суммарное суточное энергопотребление: 7.0 кВт·ч",
		"Транспорт для новичков:",
	}
	for _, fragment := range expectedFragments {
		if !strings.Contains(text, fragment) {
			t.Errorf("expected output to contain %q", fragment)
		}
	}
}
