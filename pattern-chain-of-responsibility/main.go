package main

import (
	"fmt"
)

/*
 Плюсы:
	- Паттерн позволяет отделить отправителя запроса от его получателей, что упрощает изменение и расширение кода.
	- Клиенты не знают о внутренней структуре цепочки и не заботятся о том, какой именно обработчик будет обрабатывать запрос.
 Минусы:
	- Сложность отладки.
	- Если цепочка обязанностей динамически модифицируется во время выполнения, управление и поддержка цепочки могут стать более сложными.
*/

// Тип темы
type Topic int

const NoHelpTopic Topic = -1

// Интерфейс обработчика помощи
type HelpHandler interface {
	hasHelp() bool
	handleHelp()
	setNext(next HelpHandler)
}

// Базовый обработчик помощи
type BaseHelpHandler struct {
	next  HelpHandler // Следующий обработчик в цепочке
	topic Topic       // Тема помощи
}

// Проверяет, может ли обработчик предоставить помощь по данной теме
func (h *BaseHelpHandler) hasHelp() bool {
	return h.topic != NoHelpTopic
}

// Устанавливает следующий обработчик в цепочке
func (h *BaseHelpHandler) setNext(next HelpHandler) {
	h.next = next
}

// Обрабатывает запрос на помощь
func (h *BaseHelpHandler) handleHelp() {
	if h.next != nil {
		h.next.handleHelp()
	} else {
		fmt.Println("HelpHandler::handleHelp")
	}
}

// Виджет
type Widget struct {
	BaseHelpHandler
	parent *Widget // Родительский виджет
}

// Кнопка
type Button struct {
	Widget
}

// Диалог
type Dialog struct {
	Widget
}

// Приложение
type Application struct {
	BaseHelpHandler
}

func main() {
	const PrintTopic Topic = 1
	const PaperOrientationTopic Topic = 2
	const ApplicationTopic Topic = 3

	// Создаем цепочку обязанностей
	app := &Application{BaseHelpHandler{topic: ApplicationTopic}}
	dialog := &Dialog{}
	button := &Button{}

	// Устанавливаем связи между обработчиками
	app.setNext(dialog)
	dialog.setNext(button)

	// Запускаем обработку запроса на помощь
	app.handleHelp()
}
