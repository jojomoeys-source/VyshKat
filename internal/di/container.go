package di

import (
	"fmt"
	"io"
	"os"

	"go.uber.org/dig"

	"github.com/jojomoeys-source/VyshKat/internal/app"
	"github.com/jojomoeys-source/VyshKat/internal/service"
)

func BuildContainer() (*dig.Container, error) {
	container := dig.New()

	if err := container.Provide(
		service.NewOperationalServiceCenter,
		dig.As(new(service.ServiceCenter)),
	); err != nil {
		return nil, fmt.Errorf("provide service center: %w", err)
	}

	if err := container.Provide(service.NewFleetService); err != nil {
		return nil, fmt.Errorf("provide fleet service: %w", err)
	}
	if err := container.Provide(func() io.Reader { return os.Stdin }); err != nil {
		return nil, fmt.Errorf("provide console input: %w", err)
	}
	if err := container.Provide(func() io.Writer { return os.Stdout }); err != nil {
		return nil, fmt.Errorf("provide console output: %w", err)
	}
	if err := container.Provide(app.NewDemoScenario); err != nil {
		return nil, fmt.Errorf("provide demo scenario: %w", err)
	}
	if err := container.Provide(app.NewConsoleApp); err != nil {
		return nil, fmt.Errorf("provide console app: %w", err)
	}

	return container, nil
}
