package main

import "fmt"

/*
 Плюсы:
	- Позволяет выбирать алгоритмы на лету, что увеличивает гибкость системы.
	- Позволяет легко добавлять новые алгоритмы без изменения существующего кода.
 Минусы:
	- Введение дополнительных интерфейсов и классов может усложнить архитектуру программы.
	- Для простых случаев с небольшим количеством алгоритмов использование паттерна может быть излишним.
*/

// Интерфейс поведения торможения
type IBrakeBehavior interface {
	brake()
}

// Торможение с ABS
type BrakeWithABS struct{}

func (b *BrakeWithABS) brake() {
	fmt.Println("Tорможение с ABS применено")
}

// Простое торможение
type Brake struct{}

func (b *Brake) brake() {
	fmt.Println("Простое торможение применено")
}

// Автомобиль, использующий алгоритмы торможения
type Car struct {
	brakeBehavior IBrakeBehavior
}

func NewCar(brakeBehavior IBrakeBehavior) *Car {
	return &Car{
		brakeBehavior: brakeBehavior,
	}
}

func (c *Car) ApplyBrake() {
	c.brakeBehavior.brake()
}

func (c *Car) SetBrakeBehavior(brakeBehavior IBrakeBehavior) {
	c.brakeBehavior = brakeBehavior
}

// Седан, использующий простое торможение
type Sedan struct {
	Car
}

func NewSedan() *Sedan {
	return &Sedan{
		Car: Car{
			brakeBehavior: &Brake{},
		},
	}
}

// SUV, использующий торможение с ABS
type SUV struct {
	Car
}

func NewSUV() *SUV {
	return &SUV{
		Car: Car{
			brakeBehavior: &BrakeWithABS{},
		},
	}
}

func main() {
	sedanCar := NewSedan()
	sedanCar.ApplyBrake() // Это вызовет класс "Brake"

	suvCar := NewSUV()
	suvCar.ApplyBrake() // Это вызовет класс "BrakeWithABS"

	// Динамическое установление поведения торможения
	suvCar.SetBrakeBehavior(&Brake{})
	suvCar.ApplyBrake() // Это вызовет класс "Brake"
}
