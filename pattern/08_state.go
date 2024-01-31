package pattern

import (
	"fmt"
	"time"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern

Паттерн состояние - шаблон, при котором объект должен менять свое поведение в зависимости от состояния
Плюсы -  1. Четкое разделение ответственности: Паттерн позволяет разделить ответственность за обработку различных состояний между объектами,
			 что способствует уменьшению зависимости между кодом для каждого состояния.
		 2. Простота добавления новых состояний: Добавление новых состояний — это относительно простая задача.
			 Новые состояния могут быть добавлены, не затрагивая существующий код.
		 3. Соблюдение принципов SOLID: Применение паттерна "Состояние" способствует соблюдению принципов Single Responsibility Principle
Минусы - 1. Увеличение числа классов: Внедрение состояний может привести к увеличению числа
			 классов в системе, что может усложнить структуру кода.
		 2. Потенциальная сложность в реализации переходов: Реализация переходов между состояниями может
		     потребовать дополнительных усилий, особенно в случае сложных логик переходов.
Реализация данного шаблона может выглядеть как светофор на улицах.
*/

// State interface defines the methods that concrete states should implement
type State interface {
	HandleRequest(trafficSignal *TrafficSignal)
}

// ConcreteState for the "Green" state
type GreenState struct{}

// HandleRequest method implements the behavior for the Green state.
func (s *GreenState) HandleRequest(trafficSignal *TrafficSignal) {
	fmt.Println("Switching to Green")
	// Perform actions related to the Green state
	// For simplicity, let's just wait for a few seconds to simulate time passing
	time.Sleep(3 * time.Second)
	// Transition to the next state
	trafficSignal.currentState = &YellowState{}
}

// ConcreteState for the "Yellow" state
type YellowState struct{}

// HandleRequest method implements the behavior for the Yellow state
func (s *YellowState) HandleRequest(trafficSignal *TrafficSignal) {
	fmt.Println("Switching to Yellow")
	// Perform actions related to the Yellow state
	time.Sleep(1 * time.Second)
	// Transition to the next state
	trafficSignal.currentState = &RedState{}
}

// ConcreteState for the "Red" state
type RedState struct{}

// HandleRequest method implements the behavior for the Red state
func (s *RedState) HandleRequest(trafficSignal *TrafficSignal) {
	fmt.Println("Switching to Red")
	// Perform actions related to the Red state
	time.Sleep(2 * time.Second)
	// Transition to the next state
	trafficSignal.currentState = &GreenState{}
}

// Context represents the context that contains a reference to the current state
type TrafficSignal struct {
	currentState State
}

// Request method is used to trigger the state transition
func (signal *TrafficSignal) Request() {
	signal.currentState.HandleRequest(signal)
}

// func main() {
// 	trafficSignal := &TrafficSignal{currentState: &GreenState{}}

// 	// Simulate multiple state transitions
// 	for i := 0; i < 5; i++ {
// 		trafficSignal.Request()
// 	}
// }
