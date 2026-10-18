package app

import (
	"bytes"
	"strings"
	"testing"

	"github.com/jojomoeys-source/VyshKat/internal/service"
)

func TestConsoleAppAddsHelmetAndShowsInventory(t *testing.T) {
	fleet := service.NewFleetService(service.NewOperationalServiceCenter())
	input := strings.NewReader("4\nInteractive helmet\n301\n6\n0\n")
	var output bytes.Buffer
	console := NewConsoleApp(fleet, input, &output)

	if err := console.Run(); err != nil {
		t.Fatalf("console app failed: %v", err)
	}

	if fleet.ItemCount() != 1 {
		t.Errorf("expected 1 inventory item, got %d", fleet.ItemCount())
	}

	text := output.String()
	for _, fragment := range []string{
		"Вещь добавлена на баланс.",
		"Interactive helmet (№301)",
		"До свидания!",
	} {
		if !strings.Contains(text, fragment) {
			t.Errorf("expected output to contain %q", fragment)
		}
	}
}

func TestConsoleAppAddsScooterAndShowsReports(t *testing.T) {
	fleet, text := runConsoleScenario(t, strings.Join([]string{
		"1",
		"Interactive scooter",
		"401",
		"8",
		"да",
		"2,5",
		"7",
		"8",
		"9",
		"0",
		"",
	}, "\n"))

	if fleet.ItemCount() != 1 {
		t.Errorf("expected 1 inventory item, got %d", fleet.ItemCount())
	}

	for _, fragment := range []string{
		"Транспорт принят в парк.",
		"Всего объектов: 1",
		"Суммарное суточное энергопотребление: 2.5 кВт·ч",
		"Interactive scooter (№401, простота: 8)",
	} {
		if !strings.Contains(text, fragment) {
			t.Errorf("expected output to contain %q", fragment)
		}
	}
}

func TestConsoleAppRejectsBrokenBicycle(t *testing.T) {
	fleet, text := runConsoleScenario(t, strings.Join([]string{
		"2",
		"Broken bicycle",
		"402",
		"5",
		"нет",
		"6",
		"0",
		"",
	}, "\n"))

	if fleet.ItemCount() != 0 {
		t.Errorf("expected rejected bicycle not to be stored, got %d items", fleet.ItemCount())
	}

	for _, fragment := range []string{
		"Транспорт не прошёл техосмотр.",
		"Инвентарь пуст.",
	} {
		if !strings.Contains(text, fragment) {
			t.Errorf("expected output to contain %q", fragment)
		}
	}
}

func TestConsoleAppAddsElectricBicycleAndChargingStation(t *testing.T) {
	fleet, text := runConsoleScenario(t, strings.Join([]string{
		"3",
		"Electric bicycle",
		"403",
		"7",
		"yes",
		"3.5",
		"5",
		"Interactive station",
		"501",
		"4.5",
		"8",
		"0",
		"",
	}, "\n"))

	if fleet.ItemCount() != 2 {
		t.Errorf("expected 2 inventory items, got %d", fleet.ItemCount())
	}
	if total := fleet.TotalDailyEnergyConsumption(); total != 8.0 {
		t.Errorf("expected total energy consumption 8.0, got %f", total)
	}

	for _, fragment := range []string{
		"Транспорт принят в парк.",
		"Вещь добавлена на баланс.",
		"Суммарное суточное энергопотребление: 8.0 кВт·ч",
	} {
		if !strings.Contains(text, fragment) {
			t.Errorf("expected output to contain %q", fragment)
		}
	}
}

func TestConsoleAppRepromptsAfterInvalidInput(t *testing.T) {
	fleet, text := runConsoleScenario(t, strings.Join([]string{
		"unknown",
		"42",
		"4",
		"   ",
		"Validated helmet",
		"not-a-number",
		"0",
		"601",
		"0",
		"",
	}, "\n"))

	if fleet.ItemCount() != 1 {
		t.Errorf("expected 1 inventory item after corrected input, got %d", fleet.ItemCount())
	}

	for _, fragment := range []string{
		"Введите целое число.",
		"Неизвестный пункт меню.",
		"Значение не может быть пустым.",
		"Введите положительное число.",
		"Вещь добавлена на баланс.",
	} {
		if !strings.Contains(text, fragment) {
			t.Errorf("expected output to contain %q", fragment)
		}
	}
}

func runConsoleScenario(t *testing.T, inputText string) (*service.FleetService, string) {
	t.Helper()

	fleet := service.NewFleetService(service.NewOperationalServiceCenter())
	input := strings.NewReader(inputText)
	var output bytes.Buffer
	console := NewConsoleApp(fleet, input, &output)

	if err := console.Run(); err != nil {
		t.Fatalf("console app failed: %v", err)
	}

	return fleet, output.String()
}
