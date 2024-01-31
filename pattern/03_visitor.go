package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern



Паттерн посетитель позволяет добавлять новую функциональность объектам, не изменяя их структуры
	Применяется, когда нужно для ряда классов сделать похожую операцию

Плюсы -  1. Новая функциональность у нескольких классов добавляется сразу, не изменяя код этих классов
		 2. Позволяет получить информацию о типе объекта

Минусы - 1. При изменении обслуживаемого класса нужно поменять код у шаблона
		 2. Затруднено добавление новых классов, поскольку нужно обновлять иерархию посетителя и его сыновей
*/

import "fmt"

// Element
type Shape interface {
	Accept(Visitor)
}

// ConcreteElement: Circle
type Circle struct {
	Radius float64
}

func (c *Circle) Accept(visitor Visitor) {
	visitor.VisitCircle(c)
}

// ConcreteElement: Rectangle
type Rectangle struct {
	Width, Height float64
}

func (r *Rectangle) Accept(visitor Visitor) {
	visitor.VisitRectangle(r)
}

// ConcreteElement: Triangle
type Triangle struct {
	SideA, SideB, SideC float64
}

func (t *Triangle) Accept(visitor Visitor) {
	visitor.VisitTriangle(t)
}

// Visitor
type Visitor interface {
	VisitCircle(*Circle)
	VisitRectangle(*Rectangle)
	VisitTriangle(*Triangle)
}

// ConcreteVisitor: AreaCalculator
type AreaCalculator struct {
	TotalArea float64
}

func (a *AreaCalculator) VisitCircle(c *Circle) {
	area := 3.14 * c.Radius * c.Radius
	fmt.Printf("Calculating area for Circle: %.2f\n", area)
	a.TotalArea += area
}

func (a *AreaCalculator) VisitRectangle(r *Rectangle) {
	area := r.Width * r.Height
	fmt.Printf("Calculating area for Rectangle: %.2f\n", area)
	a.TotalArea += area
}

func (a *AreaCalculator) VisitTriangle(t *Triangle) {
	fmt.Println("Cannot calculate area for Triangle in this example")
}

// ConcreteVisitor: PerimeterCalculator
type PerimeterCalculator struct {
	TotalPerimeter float64
}

func (p *PerimeterCalculator) VisitCircle(c *Circle) {
	perimeter := 2 * 3.14 * c.Radius
	fmt.Printf("Calculating perimeter for Circle: %.2f\n", perimeter)
	p.TotalPerimeter += perimeter
}

func (p *PerimeterCalculator) VisitRectangle(r *Rectangle) {
	perimeter := 2 * (r.Width + r.Height)
	fmt.Printf("Calculating perimeter for Rectangle: %.2f\n", perimeter)
	p.TotalPerimeter += perimeter
}

func (p *PerimeterCalculator) VisitTriangle(t *Triangle) {
	fmt.Println("Cannot calculate perimeter for Triangle in this example")
}

// func main() {
// 	shapes := []Shape{
// 		&Circle{Radius: 5},
// 		&Rectangle{Width: 3, Height: 4},
// 		&Triangle{SideA: 5, SideB: 12, SideC: 13},
// 	}

// 	areaCalculator := &AreaCalculator{}
// 	perimeterCalculator := &PerimeterCalculator{}

// 	for _, shape := range shapes {
// 		shape.Accept(areaCalculator)
// 		shape.Accept(perimeterCalculator)
// 	}

// 	fmt.Printf("Total Area: %.2f\n", areaCalculator.TotalArea)
// 	fmt.Printf("Total Perimeter: %.2f\n", perimeterCalculator.TotalPerimeter)
// }
