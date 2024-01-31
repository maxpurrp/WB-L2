package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные
	примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern

	Структурный паттерн проектирования, при котором реализуется унифицированный
интерфейс для набора интрефейсов в подсистеме. С помощью даного паттерна
упрощается взаимодействие с подсистемами, делаяя их более доступыми и понятными.
Плюсы - 1. Унифицированный интерфейс с набором разрозненных реализаций
								подсистем без нежелательного связывания с этой подсистемой

		   2. Фасад упрощает внесение изменений в систему, поскольку все изменения
					внутри подсистемы могут быть легко абстрагированы через унифицированный интерфейс фасада

Минусы - 1. Ограниченность реализации.  Фасад может предоставлять только ограниченный набор функций подсистемы.
											 Если клиенту требуется более сложная или специфичная функциональность,
											ему придется обращаться к непосредственным компонентам подсистемы,
													что может уменьшить преимущества фасада

			2. Дополнительная абстракция: Фасад добавляет еще один уровень абстракции в систему.
								В некоторых случаях это может усложнить понимание кода, особенно если фасад
						не является должным образом документированным или если клиенту требуется понимание деталей реализации.
*/
// SubSystem of Company
type FrontDepartment struct {
	workers int
}

func (s *FrontDepartment) ReadDocumentation() {
	fmt.Printf("%v workers read documantation right now \n", s.workers/2)
}

func (s *FrontDepartment) HtmlCoding() {
	fmt.Printf("%v workers creating some awesome features \n", s.workers/2)
}

// SubSystem of Company
type BackDepartmant struct {
	developers  int
	testEngneer int
}

func (b *BackDepartmant) Develop() {
	fmt.Printf("all %v Developers fixing bugs \n", b.developers)
}

func (b *BackDepartmant) Testing() {
	fmt.Printf("all %v QA Engineers testing bugs \n", b.testEngneer)
}

// MainSystem
type Company struct {
	Front *FrontDepartment
	Back  *BackDepartmant
}

func NewFacade() *Company {
	return &Company{
		Front: &FrontDepartment{
			workers: 8,
		},
		Back: &BackDepartmant{
			developers:  10,
			testEngneer: 6,
		},
	}
}

func (f *Company) Operation() {
	fmt.Println("Facade: Operation")
	f.Front.HtmlCoding()
	f.Front.ReadDocumentation()
	f.Back.Develop()
	f.Back.Testing()
}

// func main() {
// 	facade := NewFacade()
// 	facade.Operation()
// }
