package main

import (
	"fmt"
)

/*
 Плюсы:
	- Легко добавлять новые команды без изменения существующего кода.
	- Можно создавать последовательности команд, используя систему очередей.
 Минусы:
	- Большое количество классов и объектов работает вместе для достижения цели, что требует внимательного разработки этих классов.
	- Если ваше приложение не требует функциональности отмены/повтора и вы не планируете поддерживать такие функции в будущем, паттерн Команда может оказаться излишним усложнением.
*/

// Интерфейс команды
type Command interface {
	execute()
}

// Структура получателя
type MyClass struct{}

// Метод действия для MyClass
func (m *MyClass) action() {
	fmt.Println("MyClass::action")
}

// Простая команда
type SimpleCommand struct {
	receiver *MyClass       // Получатель
	action   func(*MyClass) // Действие
}

// Конструктор для создания экземпляров SimpleCommand
func NewSimpleCommand(receiver *MyClass, action func(*MyClass)) Command {
	return &SimpleCommand{receiver: receiver, action: action}
}

// Реализация метода execute для SimpleCommand
func (s *SimpleCommand) execute() {
	s.action(s.receiver) // Вызов действия
}

func main() {
	// Создание экземпляра MyClass
	receiver := &MyClass{}

	// Создание новой простой команды с получателем и его действием
	command := NewSimpleCommand(receiver, (*MyClass).action)

	command.execute()
}
