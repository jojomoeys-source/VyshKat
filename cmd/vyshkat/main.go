package main

import (
	"log"

	"github.com/jojomoeys-source/VyshKat/internal/app"
	appdi "github.com/jojomoeys-source/VyshKat/internal/di"
)

func main() {
	container, err := appdi.BuildContainer()
	if err != nil {
		log.Fatal(err)
	}

	err = container.Invoke(func(demo *app.DemoScenario, console *app.ConsoleApp) error {
		if err := demo.Run(); err != nil {
			return err
		}

		return console.Run()
	})
	if err != nil {
		log.Fatal(err)
	}
}
