package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func sortedWord(word string) string {
	wordChars := strings.Split(word, "")
	sort.Strings(wordChars)
	return strings.Join(wordChars, "")

}

func appendUnique(list []string, elem string) []string {
	// iterate through the list, checking for the presence of the element
	for _, el := range list {
		if el == elem {
			// the element is already in the list, return the original list
			return list
		}
	}

	// the element is not in the list, append it
	return append(list, elem)
}

func FindAnagrams(elems []string) map[string][]string {
	// tmp for temporarily store words with the same sorted form
	tmp := make(map[string][]string)
	// annagramSet will store the final sets of anagrams
	annamgramSet := make(map[string][]string)

	for _, word := range elems {
		// convert the word to lowercase
		lowerWord := strings.ToLower(word)
		// sort the characters in the word
		sortedWord := sortedWord(lowerWord)

		// append the original word to the temporary map
		tmp[sortedWord] = append(tmp[sortedWord], word)

		// check if a set with the same sorted form already exists
		// if it exists, add the current word to the anagram set
		arr, exists := tmp[sortedWord]
		if exists {
			annamgramSet[arr[0]] = appendUnique(annamgramSet[arr[0]], word)
		}
	}
	// sort the arrays in each set in ascending order
	for key, value := range annamgramSet {
		if len(value) == 1 {
			delete(annamgramSet, key)
			continue
		}
		sort.Strings(annamgramSet[key])

	}

	return annamgramSet
}

func main() {
	elements := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	// find and display the sets of anagrams.
	annagramSet := FindAnagrams(elements)
	for key, value := range annagramSet {
		fmt.Printf("%s: %v\n", key, value)
	}
}
