package main

import (
	"reflect"
	"testing"
)

func TestSortByColumn(t *testing.T) {
	expected := []string{"3 apple", "1 bear", "2 yaml"}
	data := ReadFile("./test_cases/file_col.txt")
	SortByColumn(data, 2)
	result := ReadFile("./result_col.txt")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed: expected %v, got %v", expected, result)
	}
}

func TestSortByNumeric(t *testing.T) {
	expected := []string{"1", "2", "3", "4", "5", "6", "8", "10"}
	data := ReadFile("./test_cases/file_num.txt")
	SortByNumeric(data)
	result := ReadFile("./result_num.txt")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed: expected %v, got %v", expected, result)
	}
}

func TestReverseSort(t *testing.T) {
	expected := []string{"Anfisa", "Max", "Pasha"}
	data := ReadFile("./test_cases/file_rev.txt")
	ReverseSort(data)
	result := ReadFile("./result_rev.txt")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed: expected %v, got %v", expected, result)
	}
}

func TestUniqueSort(t *testing.T) {

	expected := []string{"Eat", "Sleep", "Wake", "Work"}
	data := ReadFile("./test_cases/file_uniq.txt")
	UniqueSort(data)
	result := ReadFile("./result_uniq.txt")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed: expected %v, got %v", expected, result)
	}
}

func TestMonthSort(t *testing.T) {

	expected := []string{"January", "February", "March", "April",
		"May", "June", "July", "August", "September", "October", "November", "December"}
	data := ReadFile("./test_cases/file_months.txt")
	MonthsSort(data)
	result := ReadFile("./result_months.txt")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed: expected %v, got %v", expected, result)
	}
}

func TestTrimmedSort(t *testing.T) {

	expected := []string{"Anagramm", "Apple", "Banana", "Max", "Qwerty", "Sleep"}
	data := ReadFile("./test_cases/file_trim.txt")
	TrimmedSort(data)
	result := ReadFile("./result_trim.txt")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed: expected %v, got %v", expected, result)
	}
}

func TestIsSorted(t *testing.T) {

	expected := "data already sorted"
	data := ReadFile("./test_cases/file_issort.txt")
	result := IsSorted(data)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed: expected %v, got %v", expected, result)
	}
}
