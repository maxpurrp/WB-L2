package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file %s: %s\n", filename, err)
		os.Exit(1)

	}
	return file
}

func grepA(filename, pattern string, linesAfter int) {
	var buff []string
	var match bool

	// open the specified file
	file := openFile(filename)
	defer file.Close()

	// create a scanner to read lines from the file
	scanner := bufio.NewScanner(file)

	// iterate through each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		// if a match is found, start collecting lines into the buffer
		if match {
			buff = append(buff, line)
			linesAfter--
			if linesAfter == 0 {
				break
			}
		}

		// check if the line contains the specified pattern
		if strings.Contains(line, pattern) {
			match = true
			buff = append(buff, line)
		}
	}

	printResult(buff)
}

func grepB(filename, pattern string, linesBefore int) {
	var buff []string
	// open the specified file
	file := openFile(filename)
	defer file.Close()

	// create a scanner to read lines from the file
	scanner := bufio.NewScanner(file)

	// iterate through each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		// if a match is found, we output the saved strings before the match
		if strings.Contains(line, pattern) {
			printResult(buff)
		}

		// saving lines to the buffer
		buff = append(buff, line)

		// limit the number of rows to be saved
		if len(buff) > linesBefore {
			buff = buff[1:]
		}
	}
}

func grepC(filename, pattern string) {
	var count int
	// open the specified file
	file := openFile(filename)
	defer file.Close()

	// create a scanner to read lines from the file
	scanner := bufio.NewScanner(file)

	// iterate through each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		// check if the line contains the specified pattern
		if strings.Contains(line, pattern) {
			count++
		}
	}
	// print the total count of pattern matches
	fmt.Printf("matches of pattern : %v \n", count)
}

func grepI(filename, pattern string) {
	var buff []string
	// open the specified file
	file := openFile(filename)
	defer file.Close()

	// create a scanner to read lines from the file
	scanner := bufio.NewScanner(file)

	// iterate through each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		// Check if the lowercase version of the line
		// contains the lowercase version of the pattern
		if strings.Contains(strings.ToLower(line), strings.ToLower(pattern)) {
			buff = append(buff, line)
		}
	}
	printResult(buff)
}

func grepV(filename, pattern string) {
	var buff []string
	// open the specified file
	file := openFile(filename)
	defer file.Close()

	// create a scanner to read lines from the file
	scanner := bufio.NewScanner(file)

	// iterate through each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		// Check if the line contains the specified pattern
		if strings.Contains(line, pattern) {
			continue
		} else {
			buff = append(buff, line)
		}
	}
	printResult(buff)
}

func grepF(filename, pattern string) {
	var buff []string
	// open the specified file
	file := openFile(filename)
	defer file.Close()

	// create a scanner to read lines from the file
	scanner := bufio.NewScanner(file)

	// iterate through each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		// check if the line is equal to the specified pattern
		if reflect.DeepEqual(line, pattern) {
			buff = append(buff, line)
		}
	}
	printResult(buff)
}

func grepN(filename, pattern string) {
	var ind int = 1
	// open the specified file
	file := openFile(filename)
	defer file.Close()

	// create a scanner to read lines from the file
	scanner := bufio.NewScanner(file)

	// iterate through each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		// check if the line is equal to the specified pattern
		if strings.Contains(line, pattern) {
			// print the line number and the matching line
			fmt.Printf("%v.%v \n", ind, line)
		}
		ind++
	}
}

func printResult(data []string) {
	for _, line := range data {
		fmt.Println(line)
	}
}

var (
	afterFlag       int
	beforeFlag      int
	contextFlag     int
	countFlag       bool
	ignoreCaseFlag  bool
	invertFlag      bool
	fixedStringFlag bool
	lineNumberFlag  bool
)

func main() {
	// set flags
	flag.IntVar(&afterFlag, "A", 0, "output N lines after a match")
	flag.IntVar(&beforeFlag, "B", 0, "output N lines before a match")
	flag.IntVar(&contextFlag, "C", 0, "output N lines after and before a match")
	flag.BoolVar(&countFlag, "c", false, "output count of pattern mathes")
	flag.BoolVar(&ignoreCaseFlag, "i", false, "output lines that match the pattern ignoring case")
	flag.BoolVar(&invertFlag, "v", false, "output lines that did not match the pattern")
	flag.BoolVar(&fixedStringFlag, "F", false, "output strings that are equal to a string as a pattern")
	flag.BoolVar(&lineNumberFlag, "n", false, "Output the line number and the line that matches the pattern")
	flag.Parse()

	args := flag.Args()
	// extract filepath and pattern from arguments
	pattern := args[0]
	filepath := args[1]

	if afterFlag != 0 {
		grepA(filepath, pattern, afterFlag)
	}

	if beforeFlag != 0 {
		grepB(filepath, pattern, beforeFlag)
	}

	if contextFlag != 0 {
		grepB(filepath, pattern, contextFlag)
		grepA(filepath, pattern, contextFlag)
	}
	if countFlag {
		grepC(filepath, pattern)
	}
	if ignoreCaseFlag {
		grepI(filepath, pattern)
	}
	if invertFlag {
		grepV(filepath, pattern)
	}
	if fixedStringFlag {
		grepF(filepath, pattern)
	}
	if lineNumberFlag {
		grepN(filepath, pattern)
	}

}
