package app

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/jojomoeys-source/VyshKat/internal/domain"
	"github.com/jojomoeys-source/VyshKat/internal/service"
)

type DemoScenario struct {
	fleet  *service.FleetService
	output io.Writer
}

func NewDemoScenario(fleet *service.FleetService, output io.Writer) *DemoScenario {
	return &DemoScenario{
		fleet:  fleet,
		output: output,
	}
}

func (d *DemoScenario) Run() error {
	var report strings.Builder

	fmt.Fprintln(&report, "Демонстрация возможностей ВышКат")

	scooter, err := domain.NewScooter("Electric scooter", 101, 8, true, 2.5)
	if err != nil {
		return fmt.Errorf("create demo scooter: %w", err)
	}
	if err := d.fleet.AddTransport(scooter); err != nil {
		return fmt.Errorf("add demo scooter: %w", err)
	}
	fmt.Fprintf(&report, "Принят: %s (№%d)\n", scooter.Name(), scooter.InventoryNumber())

	brokenBicycle, err := domain.NewElectricBicycle(
		"Broken electric bicycle",
		102,
		7,
		false,
		3.0,
	)
	if err != nil {
		return fmt.Errorf("create broken demo bicycle: %w", err)
	}
	inspectionErr := d.fleet.AddTransport(brokenBicycle)
	if inspectionErr == nil {
		return errors.New("broken demo bicycle unexpectedly passed inspection")
	}
	if !errors.Is(inspectionErr, service.ErrInspectionFailed) {
		return fmt.Errorf("inspect broken demo bicycle: %w", inspectionErr)
	}
	fmt.Fprintf(
		&report,
		"Отклонён техосмотром: %s (№%d)\n",
		brokenBicycle.Name(),
		brokenBicycle.InventoryNumber(),
	)

	bicycle, err := domain.NewBicycle("City bicycle", 103, 5, true)
	if err != nil {
		return fmt.Errorf("create demo bicycle: %w", err)
	}
	if err := d.fleet.AddTransport(bicycle); err != nil {
		return fmt.Errorf("add demo bicycle: %w", err)
	}

	helmet, err := domain.NewHelmet("Safety helmet", 201)
	if err != nil {
		return fmt.Errorf("create demo helmet: %w", err)
	}
	if err := d.fleet.AddItem(helmet); err != nil {
		return fmt.Errorf("add demo helmet: %w", err)
	}

	station, err := domain.NewChargingStation("Charging station", 202, 4.5)
	if err != nil {
		return fmt.Errorf("create demo charging station: %w", err)
	}
	if err := d.fleet.AddItem(station); err != nil {
		return fmt.Errorf("add demo charging station: %w", err)
	}

	fmt.Fprintln(&report, "\nИнвентарь:")
	for _, item := range d.fleet.Items() {
		fmt.Fprintf(&report, "- %s (№%d)\n", item.Name(), item.InventoryNumber())
	}

	fmt.Fprintf(&report, "\nВсего объектов: %d\n", d.fleet.ItemCount())
	fmt.Fprintf(
		&report,
		"Суммарное суточное энергопотребление: %.1f кВт·ч\n",
		d.fleet.TotalDailyEnergyConsumption(),
	)

	fmt.Fprintln(&report, "\nТранспорт для новичков:")
	for _, transport := range d.fleet.BeginnerFriendlyTransports() {
		fmt.Fprintf(
			&report,
			"- %s (№%d, простота: %d)\n",
			transport.Name(),
			transport.InventoryNumber(),
			transport.BeginnerFriendliness(),
		)
	}

	if _, err := io.WriteString(d.output, report.String()); err != nil {
		return fmt.Errorf("write demo report: %w", err)
	}

	return nil
}
