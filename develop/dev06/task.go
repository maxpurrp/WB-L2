package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func cut(columns []int, delimeter string, onlySeparated bool) {
	// create scanner for read stdin
	scanner := bufio.NewScanner(os.Stdin)

	// iterate over each line from stdin
	for scanner.Scan() {
		// read cte current line
		line := scanner.Text()

		// if onlySeparated is true and the line doesn't contain the delimiter, skip the line
		if onlySeparated && !strings.Contains(line, delimeter) {
			continue
		}

		// split the line into parts based on the delimiter
		parts := strings.Split(line, delimeter)

		// create an array to store the selected output parts based on the specified columns
		var outputParts []string

		// iterate over the specified columns
		for _, col := range columns {
			// if the column index is valid, append the
			// corresponding part to the outputParts array
			// Otherwise, append an empty string
			if col >= 0 && col < len(parts) {
				outputParts = append(outputParts, parts[col])
			} else {
				outputParts = append(outputParts, "")
			}
		}
		// print the selected output parts, joined by a tab character
		fmt.Println(strings.Join(outputParts, "\t"))
	}

	// check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading stdin : %v \n", err)
		os.Exit(1)
	}
}

func parseColumn(column string) (int, error) {
	var col int
	// use fmt.Sscanf to parse the column string as an integer
	_, err := fmt.Sscanf(column, "%d", &col)
	if err != nil {
		// if an error occurs during parsing, return -1 and the error
		return -1, err
	}
	// subtract 1 from the parsed value to convert it to a zero-based index
	return col - 1, nil
}

func parseColumns(column string) []int {
	var columns []int

	// split the input string into individual column representations
	parts := strings.Split(column, ",")

	// iterate over each part and parse it as a column
	for _, part := range parts {
		col, err := parseColumn(part)
		if err != nil {
			// if an error occurs during parsing, print an error message to standard error
			fmt.Fprintf(os.Stderr, "error on parsing column: %v \n", err)
		}
		// append the parsed column to the columns array
		columns = append(columns, col)

	}
	return columns
}

var (
	columnStr     string
	delimiter     string
	onlySeparated bool
)

func main() {
	// define command-line flags for columnStr, delimiter, and onlySeparated
	flag.StringVar(&columnStr, "f", "", "columns for stdout")
	flag.StringVar(&delimiter, "d", "\t", "column separator(default - \t)")
	flag.BoolVar(&onlySeparated, "s", false, "only separated lines")

	// Parse the command-line arguments
	flag.Parse()

	// if columnStr is empty, print an error message to
	//  standard error and exit with a non-zero status code
	if columnStr == "" {
		fmt.Fprintf(os.Stderr, "Specify the column numbers to output")
		os.Exit(1)
	}

	// parse the columnStr into an array of column indices
	columns := parseColumns(columnStr)

	// call the cut function with the specified columns, delimiter, and onlySeparated flag
	cut(columns, delimiter, onlySeparated)
}
