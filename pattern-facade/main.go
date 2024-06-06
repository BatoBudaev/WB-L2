package main

import "fmt"

/*
 Плюсы:
	- Фасад предоставляет простой интерфейс для сложной подсистемы, что облегчает её использование.
	- Пользователям системы не нужно знать о внутренних деталях и взаимодействиях между компонентами.
	- Изменения в одной подсистеме не требуют изменений в других, если интерфейс фасада остается неизменным.
	- Фасад позволяет легко добавлять или заменять компоненты подсистемы без изменения кода, который использует фасад.
 Минусы:
	- Добавление лишнего уровня абстракции.
	- Скрытие функциональности.
*/

// CPU структура
type CPU struct{}

func (c *CPU) Freeze() {
	fmt.Println("CPU: Freeze")
}

func (c *CPU) Jump(position int64) {
	fmt.Printf("CPU: Jump to %d\n", position)
}

func (c *CPU) Execute() {
	fmt.Println("CPU: Execute")
}

// HardDrive структура
type HardDrive struct{}

func (hd *HardDrive) Read(lba int64, size int) string {
	// Вернем строку вместо массива байтов
	return "data from hard drive"
}

// Memory структура
type Memory struct{}

func (m *Memory) Load(position int64, data string) {
	fmt.Printf("Memory: Load data '%s' to position %d\n", data, position)
}

// ComputerFacade структура
type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

const (
	kBootAddress = 0x7C00 // Адрес загрузки
	kBootSector  = 0      // Загрузочный сектор
	kSectorSize  = 512    // Размер сектора
)

// Конструктор для ComputerFacade
func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

// Метод Start для запуска компьютера
func (cf *ComputerFacade) Start() {
	cf.cpu.Freeze()
	data := cf.hardDrive.Read(kBootSector, kSectorSize)
	cf.memory.Load(kBootAddress, data)
	cf.cpu.Jump(kBootAddress)
	cf.cpu.Execute()
}

func main() {
	// Создаем фасад компьютера и запускаем его
	computer := NewComputerFacade()
	computer.Start()
}
