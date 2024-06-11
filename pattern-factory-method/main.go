package main

import (
	"fmt"
)

/*
 Плюсы:
	- Паттерн позволяет скрыть детали создания объектов от клиента, что упрощает изменение процесса создания объектов без необходимости изменения клиентского кода.
	- Паттерн позволяет легко добавлять новые типы продуктов без изменения существующего кода, что соответствует принципу открытости/закрытости.
 Минусы:
	- Введение дополнительных абстракций и классов может усложнить архитектуру и понимание кода, особенно если система уже достаточно сложна.
*/

// Идентификатор продукта
type ProductID string

// Интерфейс продукта
type Product interface {
	print()
}

// Конкретный продукт MINE
type ConcreteProductMINE struct{}

func (p *ConcreteProductMINE) print() {
	fmt.Println("Print MINE")
}

// Конкретный продукт YOURS
type ConcreteProductYOURS struct{}

func (p *ConcreteProductYOURS) print() {
	fmt.Println("Print YOURS")
}

// Создатель
type Creator struct{}

// Фабричный метод
func (c *Creator) create(id ProductID) Product {
	switch id {
	case "MINE":
		return &ConcreteProductMINE{}
	case "YOURS":
		return &ConcreteProductYOURS{}
	default:
		return nil
	}
}

func main() {
	creator := &Creator{}

	product := creator.create("MINE")
	product.print()

	product = creator.create("YOURS")
	product.print()
}
