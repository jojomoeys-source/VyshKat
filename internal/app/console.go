package app

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/jojomoeys-source/VyshKat/internal/domain"
	"github.com/jojomoeys-source/VyshKat/internal/service"
)

type ConsoleApp struct {
	fleet   *service.FleetService
	scanner *bufio.Scanner
	output  io.Writer
}

func NewConsoleApp(
	fleet *service.FleetService,
	input io.Reader,
	output io.Writer,
) *ConsoleApp {
	return &ConsoleApp{
		fleet:   fleet,
		scanner: bufio.NewScanner(input),
		output:  output,
	}
}

func (c *ConsoleApp) Run() error {
	for {
		if err := c.printMenu(); err != nil {
			return err
		}

		choice, err := c.readInt("Выберите действие: ")
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}

		switch choice {
		case 1:
			err = c.addScooter()
		case 2:
			err = c.addBicycle()
		case 3:
			err = c.addElectricBicycle()
		case 4:
			err = c.addHelmet()
		case 5:
			err = c.addChargingStation()
		case 6:
			err = c.printInventory()
		case 7:
			err = c.writef("Всего объектов: %d\n", c.fleet.ItemCount())
		case 8:
			err = c.writef(
				"Суммарное суточное энергопотребление: %.1f кВт·ч\n",
				c.fleet.TotalDailyEnergyConsumption(),
			)
		case 9:
			err = c.printBeginnerFriendlyTransports()
		case 0:
			return c.writef("До свидания!\n")
		default:
			err = c.writef("Неизвестный пункт меню.\n")
		}

		if err != nil {
			return err
		}
	}
}

func (c *ConsoleApp) printMenu() error {
	return c.writef(`
Главное меню 
1. Добавить самокат
2. Добавить велосипед
3. Добавить электровелосипед
4. Добавить шлем
5. Добавить зарядную станцию
6. Показать весь инвентарь
7. Показать количество объектов
8. Показать суммарное энергопотребление
9. Показать транспорт для новичков
0. Выход
`)
}

func (c *ConsoleApp) addScooter() error {
	name, inventoryNumber, friendliness, operational, err := c.readTransportData()
	if err != nil {
		return err
	}
	energy, err := c.readNonNegativeFloat("Суточное энергопотребление (кВт·ч): ")
	if err != nil {
		return err
	}

	scooter, err := domain.NewScooter(name, inventoryNumber, friendliness, operational, energy)
	if err != nil {
		return c.writef("Ошибка создания самоката: %v\n", err)
	}

	return c.reportTransportAddition(c.fleet.AddTransport(scooter))
}

func (c *ConsoleApp) addBicycle() error {
	name, inventoryNumber, friendliness, operational, err := c.readTransportData()
	if err != nil {
		return err
	}

	bicycle, err := domain.NewBicycle(name, inventoryNumber, friendliness, operational)
	if err != nil {
		return c.writef("Ошибка создания велосипеда: %v\n", err)
	}

	return c.reportTransportAddition(c.fleet.AddTransport(bicycle))
}

func (c *ConsoleApp) addElectricBicycle() error {
	name, inventoryNumber, friendliness, operational, err := c.readTransportData()
	if err != nil {
		return err
	}
	energy, err := c.readNonNegativeFloat("Суточное энергопотребление (кВт·ч): ")
	if err != nil {
		return err
	}

	bicycle, err := domain.NewElectricBicycle(
		name,
		inventoryNumber,
		friendliness,
		operational,
		energy,
	)
	if err != nil {
		return c.writef("Ошибка создания электровелосипеда: %v\n", err)
	}

	return c.reportTransportAddition(c.fleet.AddTransport(bicycle))
}

func (c *ConsoleApp) addHelmet() error {
	name, err := c.readNonEmptyLine("Название: ")
	if err != nil {
		return err
	}
	inventoryNumber, err := c.readPositiveInt("Инвентарный номер: ")
	if err != nil {
		return err
	}

	helmet, err := domain.NewHelmet(name, inventoryNumber)
	if err != nil {
		return c.writef("Ошибка создания шлема: %v\n", err)
	}

	return c.reportItemAddition(c.fleet.AddItem(helmet))
}

func (c *ConsoleApp) addChargingStation() error {
	name, err := c.readNonEmptyLine("Название: ")
	if err != nil {
		return err
	}
	inventoryNumber, err := c.readPositiveInt("Инвентарный номер: ")
	if err != nil {
		return err
	}
	energy, err := c.readNonNegativeFloat("Суточное энергопотребление (кВт·ч): ")
	if err != nil {
		return err
	}

	station, err := domain.NewChargingStation(name, inventoryNumber, energy)
	if err != nil {
		return c.writef("Ошибка создания зарядной станции: %v\n", err)
	}

	return c.reportItemAddition(c.fleet.AddItem(station))
}

func (c *ConsoleApp) readTransportData() (string, int, int, bool, error) {
	name, err := c.readNonEmptyLine("Название: ")
	if err != nil {
		return "", 0, 0, false, err
	}
	inventoryNumber, err := c.readPositiveInt("Инвентарный номер: ")
	if err != nil {
		return "", 0, 0, false, err
	}
	friendliness, err := c.readFriendliness()
	if err != nil {
		return "", 0, 0, false, err
	}
	operational, err := c.readBool("Транспорт исправен? (да/нет): ")
	if err != nil {
		return "", 0, 0, false, err
	}

	return name, inventoryNumber, friendliness, operational, nil
}

func (c *ConsoleApp) printInventory() error {
	items := c.fleet.Items()
	if len(items) == 0 {
		return c.writef("Инвентарь пуст.\n")
	}

	if err := c.writef("Инвентарь:\n"); err != nil {
		return err
	}
	for _, item := range items {
		if err := c.writef("- %s (№%d)\n", item.Name(), item.InventoryNumber()); err != nil {
			return err
		}
	}

	return nil
}

func (c *ConsoleApp) printBeginnerFriendlyTransports() error {
	transports := c.fleet.BeginnerFriendlyTransports()
	if len(transports) == 0 {
		return c.writef("Подходящего транспорта нет.\n")
	}

	if err := c.writef("Транспорт для новичков:\n"); err != nil {
		return err
	}
	for _, transport := range transports {
		if err := c.writef(
			"- %s (№%d, простота: %d)\n",
			transport.Name(),
			transport.InventoryNumber(),
			transport.BeginnerFriendliness(),
		); err != nil {
			return err
		}
	}

	return nil
}

func (c *ConsoleApp) reportTransportAddition(err error) error {
	switch {
	case err == nil:
		return c.writef("Транспорт принят в парк.\n")
	case errors.Is(err, service.ErrInspectionFailed):
		return c.writef("Транспорт не прошёл техосмотр.\n")
	case errors.Is(err, service.ErrDuplicateInventoryNumber):
		return c.writef("Инвентарный номер уже используется.\n")
	default:
		return fmt.Errorf("add transport: %w", err)
	}
}

func (c *ConsoleApp) reportItemAddition(err error) error {
	switch {
	case err == nil:
		return c.writef("Вещь добавлена на баланс.\n")
	case errors.Is(err, service.ErrDuplicateInventoryNumber):
		return c.writef("Инвентарный номер уже используется.\n")
	default:
		return fmt.Errorf("add inventory item: %w", err)
	}
}

func (c *ConsoleApp) readNonEmptyLine(prompt string) (string, error) {
	for {
		line, err := c.readLine(prompt)
		if err != nil {
			return "", err
		}
		line = strings.TrimSpace(line)
		if line != "" {
			return line, nil
		}
		if err := c.writef("Значение не может быть пустым.\n"); err != nil {
			return "", err
		}
	}
}

func (c *ConsoleApp) readPositiveInt(prompt string) (int, error) {
	for {
		value, err := c.readInt(prompt)
		if err != nil {
			return 0, err
		}
		if value > 0 {
			return value, nil
		}
		if err := c.writef("Введите положительное число.\n"); err != nil {
			return 0, err
		}
	}
}

func (c *ConsoleApp) readFriendliness() (int, error) {
	for {
		value, err := c.readInt("Простота для новичка (1–10): ")
		if err != nil {
			return 0, err
		}
		if value >= 1 && value <= 10 {
			return value, nil
		}
		if err := c.writef("Введите число от 1 до 10.\n"); err != nil {
			return 0, err
		}
	}
}

func (c *ConsoleApp) readInt(prompt string) (int, error) {
	for {
		line, err := c.readLine(prompt)
		if err != nil {
			return 0, err
		}
		value, conversionErr := strconv.Atoi(strings.TrimSpace(line))
		if conversionErr == nil {
			return value, nil
		}
		if err := c.writef("Введите целое число.\n"); err != nil {
			return 0, err
		}
	}
}

func (c *ConsoleApp) readNonNegativeFloat(prompt string) (float64, error) {
	for {
		line, err := c.readLine(prompt)
		if err != nil {
			return 0, err
		}
		line = strings.ReplaceAll(strings.TrimSpace(line), ",", ".")
		value, conversionErr := strconv.ParseFloat(line, 64)
		if conversionErr == nil && value >= 0 {
			return value, nil
		}
		if err := c.writef("Введите неотрицательное число.\n"); err != nil {
			return 0, err
		}
	}
}

func (c *ConsoleApp) readBool(prompt string) (bool, error) {
	for {
		line, err := c.readLine(prompt)
		if err != nil {
			return false, err
		}
		switch strings.ToLower(strings.TrimSpace(line)) {
		case "да", "д", "yes", "y", "1":
			return true, nil
		case "нет", "н", "no", "n", "0":
			return false, nil
		default:
			if err := c.writef("Введите да или нет.\n"); err != nil {
				return false, err
			}
		}
	}
}

func (c *ConsoleApp) readLine(prompt string) (string, error) {
	if err := c.writef("%s", prompt); err != nil {
		return "", err
	}
	if c.scanner.Scan() {
		return c.scanner.Text(), nil
	}
	if err := c.scanner.Err(); err != nil {
		return "", fmt.Errorf("read console input: %w", err)
	}

	return "", io.EOF
}

func (c *ConsoleApp) writef(format string, args ...any) error {
	if _, err := fmt.Fprintf(c.output, format, args...); err != nil {
		return fmt.Errorf("write console output: %w", err)
	}

	return nil
}
