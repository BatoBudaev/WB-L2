package main

import "fmt"

/*
 Плюсы:
	- Позволяет изменять внутреннее представление продукта.
	- Инкапсулирует код для построения и представления.
	- Обеспечивает контроль над этапами строительного процесса.
 Минусы:
	- Для каждого типа продукта необходимо создать отдельный ConcreteBuilder.
	- Классы Builder должны быть изменяемыми.
	- Может затруднить/усложнить внедрение зависимостей.
*/

// Bicycle представляет продукт, создаваемый строителем.
type Bicycle struct {
	Make   string
	Model  string
	Height int
	Colour string
}

// BicycleBuilder - интерфейс строителя велосипеда.
type BicycleBuilder interface {
	SetMake(make string) BicycleBuilder
	SetModel(model string) BicycleBuilder
	SetHeight(height int) BicycleBuilder
	SetColour(colour string) BicycleBuilder
	GetResult() Bicycle
}

// GTBuilder - конкретная реализация строителя для велосипеда марки GT.
type GTBuilder struct {
	make   string
	model  string
	height int
	colour string
}

func (b *GTBuilder) SetMake(make string) BicycleBuilder {
	b.make = make
	return b
}

func (b *GTBuilder) SetModel(model string) BicycleBuilder {
	b.model = model
	return b
}

func (b *GTBuilder) SetHeight(height int) BicycleBuilder {
	b.height = height
	return b
}

func (b *GTBuilder) SetColour(colour string) BicycleBuilder {
	b.colour = colour
	return b
}

func (b *GTBuilder) GetResult() Bicycle {
	return Bicycle{
		Make:   b.make,
		Model:  b.model,
		Height: b.height,
		Colour: b.colour,
	}
}

// MountainBikeBuildDirector - директор, управляющий процессом строительства.
type MountainBikeBuildDirector struct {
	builder BicycleBuilder
}

func NewMountainBikeBuildDirector(builder BicycleBuilder) *MountainBikeBuildDirector {
	return &MountainBikeBuildDirector{builder: builder}
}

func (d *MountainBikeBuildDirector) Construct() {
	d.builder.SetMake("GT").SetModel("Avalanche").SetColour("Red").SetHeight(29)
}

func (d *MountainBikeBuildDirector) GetResult() Bicycle {
	return d.builder.GetResult()
}

func main() {
	builder := &GTBuilder{}
	director := NewMountainBikeBuildDirector(builder)

	// Директор управляет процессом создания продукта и возвращает результат.
	director.Construct()
	bicycle := director.GetResult()

	fmt.Printf("Bicycle: %+v\n", bicycle)
}
