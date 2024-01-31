package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern

Стратегия - шаблон поведенческого типа, который предоставляет семейство алгоритмов,
				инкапсулирует каждый из них и обеспечивает их взаимозаменяемость. Позволяет менять выбранный
				алгоритм независимо от объектов-клиентов, которые его используют.
Плюсы -  1. Взаимозаменяемость алгоритмов: Стратегии могут быть легко заменены другими без изменения кода контекста.
			  Это обеспечивает гибкость и возможность поддержания различных вариантов поведения.
		 2. Избежание условных операторов: Паттерн "Стратегия" позволяет избежать больших условных операторов в коде,
		      которые могли бы возникнуть при использовании альтернативных подходов.
Минусы - 1. Увеличение числа объектов: Создание отдельных классов для каждой стратегии может привести к увеличению числа
			 объектов в системе, что потенциально может повлиять на производительность.
		 2. Дублирование кода в стратегиях: Если различные стратегии имеют общие части кода, то может возникнуть дублирование кода.
		 	 В этом случае, возможно, потребуется дополнительная абстракция или рефакторинг.
Для реализации данного шаблона была выбрана сортировка чисел с тремя стратегиями: по возрастанию, убыванию и по модулю
*/

import (
	"sort"
)

// Strategy interface
type SortingStrategy interface {
	Sort(data []int) []int
}

// ConcreteStrategy for ascending order
type AscendingSortStrategy struct{}

func (s *AscendingSortStrategy) Sort(data []int) []int {
	sort.Ints(data)
	return data
}

// ConcreteStrategy for descending order
type DescendingSortStrategy struct{}

func (s *DescendingSortStrategy) Sort(data []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(data)))
	return data
}

// ConcreteStrategy for sorting by absolute value
type AbsoluteSortStrategy struct{}

func (s *AbsoluteSortStrategy) Sort(data []int) []int {
	sort.Slice(data, func(i, j int) bool {
		return abs(data[i]) < abs(data[j])
	})
	return data
}

// Context
type SortContext struct {
	strategy SortingStrategy
}

// SetStrategy allows the client to set a sorting strategy dynamically
func (context *SortContext) SetStrategy(strategy SortingStrategy) {
	context.strategy = strategy
}

// ExecuteSort delegates the sorting to the selected strategy
func (context *SortContext) ExecuteSort(data []int) []int {
	return context.strategy.Sort(data)
}

// Helper function to calculate absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// func main() {
// 	data := []int{5, -1, 3, -8, 2, 0, -6, 4, 1}

// 	// Creating a context with the ascending sort strategy
// 	context := &SortContext{strategy: &AscendingSortStrategy{}}

// 	// Client can dynamically switch the sorting strategy
// 	fmt.Println("Ascending order:", context.ExecuteSort(data))

// 	// Changing the sorting strategy to descending
// 	context.SetStrategy(&DescendingSortStrategy{})
// 	fmt.Println("Descending order:", context.ExecuteSort(data))

// 	// Changing the sorting strategy to absolute value
// 	context.SetStrategy(&AbsoluteSortStrategy{})
// 	fmt.Println("Absolute value order:", context.ExecuteSort(data))
// }
