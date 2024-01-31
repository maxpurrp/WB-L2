package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

# Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

# Дополнительное

# Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func ReadFile(filepath string) []string {
	var data []string

	// open the file specified by the filepath
	file, err := os.Open(filepath)
	if err != nil {
		// if there is an error opening the file, print an error message and exit
		fmt.Fprintf(os.Stderr, "erorr opening file: %v \n", err)
		os.Exit(1)
	}
	// ensure the file is closed when the function exits
	defer file.Close()

	// create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	// iterate over each line in the file and append it to the 'data' slice
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}
	// check for any errors that may have occurred during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	// return the slice containing the lines of the file.
	return data
}

func saveFile(data []string, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error with creating file : %v \n", err)
	}
	defer file.Close()

	for _, line := range data {
		fmt.Fprintln(file, strings.TrimSpace(line))
	}
}

func SortByColumn(data []string, columnInd int) {
	fmt.Println(columnInd)

	// create structures of type Line to store the content of each line and its key for sorting
	type Line struct {
		Content string
		Key     string
	}

	var lines []Line

	// split each line into fields and store them in Line structures
	for _, line := range data {
		fields := strings.Fields(line)
		if columnInd > len(fields) {
			// print an error message and exit if the specified column index is out of range
			fmt.Fprintf(os.Stderr, "Error: column %d out of range for line: %s\n", columnInd, line)
			os.Exit(1)
		}
		key := fields[columnInd-1]
		lines = append(lines, Line{Content: line, Key: key})
	}

	// sort the lines based on their keys
	sort.Slice(lines, func(i, j int) bool {
		return lines[i].Key < lines[j].Key
	})

	// create a slice for the sorted lines
	var sortedData []string
	for _, line := range lines {
		sortedData = append(sortedData, line.Content)
	}

	saveFile(sortedData, "result_col.txt")
}

func SortByNumeric(data []string) {
	var intSlice []int
	var strSlice []string

	// convert each string in the data to an integer and append to intSlice
	for _, item := range data {
		num, err := strconv.Atoi(item)
		if err != nil {
			// If there is an error converting to integer, print an error message
			fmt.Fprintf(os.Stderr, "invalid type value: %v", err)
		}
		intSlice = append(intSlice, num)
	}

	// sort the intSlice in ascending order
	sort.Ints(intSlice)

	// convert each integer back to string and append to strSlice
	for _, num := range intSlice {
		strSlice = append(strSlice, strconv.Itoa(num))
	}

	saveFile(strSlice, "result_num.txt")
}

func ReverseSort(data []string) {
	// swap elements from the start to the end of the slice
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}

	saveFile(data, "result_rev.txt")
}

func UniqueSort(data []string) {
	uniqMap := make(map[string]bool)
	var uniqueData []string

	// iterate over each line in the data
	for _, line := range data {
		// if the line is not present in the map, add it to the uniqueData slice
		if !uniqMap[line] {
			uniqMap[line] = true
			uniqueData = append(uniqueData, line)
		}
	}

	// sort the uniqueData slice in ascending order
	sort.Strings(uniqueData)

	saveFile(uniqueData, "result_uniq.txt")
}

func MonthsSort(data []string) {
	sort.Slice(data, func(i, j int) bool {
		// retrieve the numeric representation of months for comparison
		monthI, okI := months[data[i]]
		monthJ, okJ := months[data[j]]

		// if both months are found in the map, compare them based on their numeric values
		if okI && okJ {
			return monthI < monthJ
		}

		// if one or both months are not found in the map, perform lexicographic comparison
		return data[i] < data[j]
	})

	// save the sorted strings to a file named "result_months.txt"
	saveFile(data, "result_months.txt")
}

// trimmedSort sorts the input slice of strings after removing leading and trailing whitespaces
func TrimmedSort(data []string) {
	sort.Slice(data, func(i, j int) bool {
		// trim leading and trailing whitespaces from each string for comparison
		trimmedI := strings.TrimSpace(data[i])
		trimmedJ := strings.TrimSpace(data[j])

		// compare the trimmed strings.
		return trimmedI < trimmedJ
	})

	saveFile(data, "result_trim.txt")
}

func IsSorted(data []string) string {
	if sort.StringsAreSorted(data) {
		return "data already sorted"
	} else {
		return "not sorted"
	}
}

var (
	columnInd int
	numeric   bool
	reverSort bool
	uniqSort  bool
	monthSort bool
	trimSort  bool
	checkSort bool
	months    = map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	}
)

func main() {
	// define command-line flags for sorting options.
	flag.IntVar(&columnInd, "k", 0, "column for sorting (default index - 1)")
	flag.BoolVar(&numeric, "n", false, "sort by numeric")
	flag.BoolVar(&reverSort, "r", false, "sort by reverse order")
	flag.BoolVar(&uniqSort, "u", false, "result of sorting only unique values")
	flag.BoolVar(&monthSort, "m", false, "sort by months")
	flag.BoolVar(&trimSort, "b", false, "sort by ignoring trailing spaces")
	flag.BoolVar(&checkSort, "c", false, "check whether the data is sorted")
	flag.Parse()

	// get the file path and sorting key from command-line arguments.
	filepath := flag.Arg(0)
	key := flag.Arg(1)

	// read data from the specified file.
	data := ReadFile(filepath)
	// perform sorting based on the specified key.
	switch key {
	case "-k":
		if flag.NArg() > 2 {
			// if a column index is provided, extract it and sort by column.
			keyValue := flag.Arg(2)
			columnInd, _ := strconv.Atoi(keyValue)
			SortByColumn(data, columnInd)
		} else {
			// print an error message and exit if the column index is missing.
			fmt.Fprintf(os.Stderr, "Usage: program <filepath> -k <column index>\n")
			os.Exit(1)
		}
	case "-n":
		// sort numerically
		SortByNumeric(data)
	case "-r":
		// sort in reverse order
		ReverseSort(data)
	case "-u":
		// sort and keep only unique values
		UniqueSort(data)
	case "-m":
		// sort by months
		MonthsSort(data)
	case "-b":
		// sort by ignoring trailing spaces
		TrimmedSort(data)
	case "-c":
		// check if the data is already sorted
		res := IsSorted(data)
		fmt.Println(res)
	}
}
