package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	actual := "abcd"
	expected := "abcd"
	er := false

	result, err := Unpack(actual)

	if err != nil && er {
		t.Errorf("Test failed for string '%s': unexpected error %v", actual, err)
	}

	if result != expected {
		t.Errorf("Test failed for string '%s': expected '%s', got '%s'", actual, expected, result)
	}
}
func TestUnpack_1(t *testing.T) {
	actual := "45"
	expected := ""
	er := true

	result, err := Unpack(actual)

	if err != nil && !er {
		t.Errorf("Test failed for string '%s': unexpected error %v", actual, err)
	}

	if result != expected {
		t.Errorf("Test failed for string '%s': expected '%s', got '%s'", actual, expected, result)
	}
}
func TestUnpack_2(t *testing.T) {
	actual := ""
	expected := ""
	er := false

	result, err := Unpack(actual)

	if err != nil && !er {
		t.Errorf("Test failed for string '%s': unexpected error %v", actual, err)
	}

	if result != expected {
		t.Errorf("Test failed for string '%s': expected '%s', got '%s'", actual, expected, result)
	}
}
func TestUnpack_3(t *testing.T) {
	actual := `qwe\4\5`
	expected := "qwe45"
	er := false

	result, err := Unpack(actual)

	if err != nil && !er {
		t.Errorf("Test failed for string '%s': unexpected error %v", actual, err)
	}

	if result != expected {
		t.Errorf("Test failed for string '%s': expected '%s', got '%s'", actual, expected, result)
	}
}
func TestUnpack_4(t *testing.T) {
	actual := `qwe\45`
	expected := "qwe44444"
	er := false

	result, err := Unpack(actual)

	if err != nil && !er {
		t.Errorf("Test failed for string '%s': unexpected error %v", actual, err)
	}

	if result != expected {
		t.Errorf("Test failed for string '%s': expected '%s', got '%s'", actual, expected, result)
	}
}
func TestUnpack_5(t *testing.T) {
	actual := `qwe\\5`
	expected := `qwe\\\\\`
	er := false

	result, err := Unpack(actual)

	if err != nil && !er {
		t.Errorf("Test failed for string '%s': unexpected error %v", actual, err)
	}

	if result != expected {
		t.Errorf("Test failed for string '%s': expected '%s', got '%s'", actual, expected, result)
	}
}
