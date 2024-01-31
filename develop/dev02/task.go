package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func strWriter(char rune, builder *strings.Builder, count int) {
	for i := 0; i < count; i++ {
		builder.WriteString(string(char))
	}
}

func Unpack(s string) (string, error) {
	// convert the input string to a rune slice for efficient character manipulation
	arr := []rune(s)
	var ind int
	var builder strings.Builder
	// iterate through the characters in the rune slice
	for ind < len(arr) {
		curChar := arr[ind]
		//// if the current character is a letter, append it to the result
		if unicode.IsLetter(curChar) {
			builder.WriteString(string(curChar))
			//// if the current character is a digit, unpack the previous letter using the digit count
		} else if unicode.IsDigit(curChar) {
			count := 0
			j := ind
			prevInd := ind
			// parse the digit count
			for j < len(arr) && unicode.IsDigit(arr[j]) {
				num, _ := strconv.Atoi(string(arr[j]))
				count = count*10 + num
				j++
				ind++
			}
			// write the repeated letters to the result using strWriter
			if prevInd > 0 {
				strWriter(arr[prevInd-1], &builder, count-1)
			} else {
				return "", errors.New("invalid string")
			}
			continue
			// if the current character is a backslash, append the next character to the result
		} else if string(curChar) == `\` {
			if ind < len(arr)-1 {
				builder.WriteString(string(arr[ind+1]))
				ind++
			} else {
				return "", errors.New("invalid string")
			}
		}
		ind++
	}
	return builder.String(), nil
}

func main() {
	str := `a4bc2d5e`
	result, _ := Unpack(str)
	fmt.Println(result)
}
