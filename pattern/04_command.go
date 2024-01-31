package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern


Паттерн комманда представляет собой поведенческий шаблон проектирования,
	который превращает запросы или операции в объекты, позволяя клиентам
	параметризовать клиентские объекты с различными запросами, откладывать
	выполнение запросов, а также поддерживать отмену операций.

Плюсы - 1. Паттерн "Команда" позволяет отделить объект,
		   инициирующий запрос (отправитель), от объекта,
		   выполняющего действие (получатель)
		2. Новые команды могут быть добавлены без изменения кода клиента.
		   Это обеспечивает легкость добавления новой функциональности в систему.

Минусы - 1. Усложнение кода из-за большого количества структур в системе, что может
			сделать код сложнее для понимания и поддержки
		 2. Увеличение объема памяти, комманда может потреблять дополнительную память,
		    особенно если система обрабатывает большое количество мелких команд.
Ниже представлен пример применения паттерна в контексте граф. редкактора
			с реализацией методов Отмена(Undo) и повтор(Redo)
*/

import "fmt"

// Command
type Command interface {
	Execute()
	Undo()
}

// ConcreteCommand
type DrawLineCommand struct {
	Canvas *Canvas
	X1, Y1 int
	X2, Y2 int
}

func (c *DrawLineCommand) Execute() {
	c.Canvas.DrawLine(c.X1, c.Y1, c.X2, c.Y2)
}

func (c *DrawLineCommand) Undo() {
	c.Canvas.ClearLine(c.X1, c.Y1, c.X2, c.Y2)
}

// Receiver
type Canvas struct {
	DrawnLines map[string]bool
}

func NewCanvas() *Canvas {
	return &Canvas{DrawnLines: make(map[string]bool)}
}

func (c *Canvas) DrawLine(x1, y1, x2, y2 int) {
	key := fmt.Sprintf("%d%d%d%d", x1, y1, x2, y2)
	c.DrawnLines[key] = true
	fmt.Printf("drawing line from (%d, %d) to (%d, %d)\n", x1, y1, x2, y2)
}

func (c *Canvas) ClearLine(x1, y1, x2, y2 int) {
	key := fmt.Sprintf("%d%d%d%d", x1, y1, x2, y2)
	delete(c.DrawnLines, key)
	fmt.Printf("clearing line from (%d, %d) to (%d, %d)\n", x1, y1, x2, y2)
}

// Invoker
type History struct {
	commands []Command
}

func (h *History) ExecuteCommand(command Command) {
	command.Execute()
	h.commands = append(h.commands, command)
}

func (h *History) UndoLastCommand() {
	if len(h.commands) > 0 {
		lastCommand := h.commands[len(h.commands)-1]
		lastCommand.Undo()
		h.commands = h.commands[:len(h.commands)-1]
	}
}

// func main() {
// 	canvas := NewCanvas()
// 	history := &History{}

// 	drawLineCommand1 := &DrawLineCommand{Canvas: canvas, X1: 0, Y1: 0, X2: 10, Y2: 10}
// 	drawLineCommand2 := &DrawLineCommand{Canvas: canvas, X1: 5, Y1: 5, X2: 15, Y2: 15}

// 	history.ExecuteCommand(drawLineCommand1) // Draw a line from (0,0) to (10,10)
// 	history.ExecuteCommand(drawLineCommand2) // Draw a line from (5,5) to (15,15)

// 	history.UndoLastCommand() // Undo the last command (remove the line from (5,5) to (15,15))
// }
