package main

import (
	"fmt"
)

/*
 Плюсы:
	- Каждое состояние инкапсулировано в своем собственном объекте, что делает код более читаемым и модульным.
	- Уменьшает количество условных операторов, заменяя их переходами между состояниями.
	- Легко добавляются новые состояния без изменения существующего кода.
 Минусы:
	- Может привести к увеличению числа классов, особенно если у вас много состояний.
	- В некоторых случаях может возникнуть сложность с управлением памятью, особенно если состояния создаются динамически и часто меняются.
*/

// Интерфейс State определяет общие методы для всех состояний.
type State interface {
	Freeze(context *StateContext)
	Heat(context *StateContext)
	GetName() string // Метод для получения названия состояния.
}

// StateContext хранит текущее состояние и предоставляет методы для его изменения.
type StateContext struct {
	state State // Текущее состояние.
}

// Создает новый контекст состояния с начальным состоянием.
func NewStateContext(initialState State) *StateContext {
	return &StateContext{state: initialState}
}

// Замораживает текущее состояние.
func (sc *StateContext) Freeze() {
	fmt.Println("Замораживаем", sc.state.GetName())
	sc.state.Freeze(sc)
}

// Нагревает текущее состояние.
func (sc *StateContext) Heat() {
	fmt.Println("Нагреваем", sc.state.GetName())
	sc.state.Heat(sc)
}

// Изменяет текущее состояние на новое.
func (sc *StateContext) SetState(newState State) {
	fmt.Println("Меняем состояние с", sc.state.GetName(), "на", newState.GetName())
	sc.state = newState
}

// Конкретные реализации состояний.
type SolidState struct{} // Твердое состояние.

func (ss *SolidState) Freeze(_ *StateContext) { fmt.Println("Ничего не происходит") }
func (ss *SolidState) Heat(sc *StateContext)  { sc.SetState(&LiquidState{}) }
func (ss *SolidState) GetName() string        { return "Твердое" }

type LiquidState struct{} // Жидкое состояние.

func (ls *LiquidState) Freeze(sc *StateContext) { sc.SetState(&SolidState{}) }
func (ls *LiquidState) Heat(sc *StateContext)   { sc.SetState(&GasState{}) }
func (ls *LiquidState) GetName() string         { return "Жидкое" }

type GasState struct{} // Газообразное состояние.

func (gs *GasState) Freeze(sc *StateContext) { sc.SetState(&LiquidState{}) }
func (gs *GasState) Heat(_ *StateContext)    { fmt.Println("Ничего не происходит") }
func (gs *GasState) GetName() string         { return "Газообразное" }

func main() {
	// Создаем контекст состояния с начальным твердым состоянием.
	sc := NewStateContext(&SolidState{})

	// Последовательность действий над состоянием.
	sc.Heat()
	sc.Heat()
	sc.Heat()
	sc.Freeze()
	sc.Freeze()
	sc.Freeze()
}
