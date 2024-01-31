// В файле grep_test.go

package main

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"testing"
)

// Функция для захвата вывода
func getStdout(f func()) string {
	originalOutput := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	go func() {
		defer w.Close()
		f()
	}()

	var buf bytes.Buffer
	buf.ReadFrom(r)

	os.Stdout = originalOutput

	return buf.String()
}

func createTmp(data []byte) (*os.File, error) {
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		return nil, err
	}

	if _, err := tmpfile.Write(data); err != nil {
		return nil, err
	}
	if err := tmpfile.Close(); err != nil {
		return nil, err
	}
	return tmpfile, nil
}

func TestGrepA(t *testing.T) {
	data := []byte("Ilya\nDen\nPasha\nMax\nAnfisa\nZhenya\n")
	file, err := createTmp(data)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	expected := "Max\nAnfisa\nZhenya\n"

	actual := getStdout(func() {
		grepA(file.Name(), "Max", 2)
	})

	if reflect.DeepEqual(expected, actual) {
		fmt.Printf("TEST %v PASSED\n", t.Name())
	} else {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
func TestGrepB(t *testing.T) {
	data := []byte("Ilya\nDen\nPasha\nMax\nAnfisa\nZhenya\n")
	file, err := createTmp(data)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	expected := "Den\nPasha\n"

	actual := getStdout(func() {
		grepB(file.Name(), "Max", 2)
	})

	if reflect.DeepEqual(expected, actual) {
		fmt.Printf("TEST %v PASSED\n", t.Name())
	} else {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
func TestGrepC(t *testing.T) {
	data := []byte("nPasha\nDen\nPasha\nMax\nPasha\nZhenya\n")
	file, err := createTmp(data)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	expected := "matches of pattern : 3 \n"

	actual := getStdout(func() {
		grepC(file.Name(), "Pasha")
	})
	if reflect.DeepEqual(expected, actual) {
		fmt.Printf("TEST %v PASSED\n", t.Name())
	} else {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
func TestGrepI(t *testing.T) {
	data := []byte("Ilya\nDen\nPasha\nMax\nAnfisa\nZhenya\n")
	file, err := createTmp(data)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	expected := "Pasha\n"

	actual := getStdout(func() {
		grepI(file.Name(), "pasha")
	})
	if reflect.DeepEqual(expected, actual) {
		fmt.Printf("TEST %v PASSED\n", t.Name())
	} else {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestGrepV(t *testing.T) {
	data := []byte("Ilya\nDen\nPasha\nMax\nAnfisa\nZhenya\n")
	file, err := createTmp(data)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	expected := "Ilya\nDen\nPasha\nMax\nAnfisa\nZhenya\n"

	actual := getStdout(func() {
		grepV(file.Name(), "Roman")
	})
	if reflect.DeepEqual(expected, actual) {
		fmt.Printf("TEST %v PASSED\n", t.Name())
	} else {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestGrepF(t *testing.T) {
	data := []byte("Ilya Sleep\nDen Work\nPasha Eat\nMax nervous\nAnfisa chiiling\nZhenya swimming\n")
	file, err := createTmp(data)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	expected := ""

	actual := getStdout(func() {
		grepF(file.Name(), "Roman")
	})
	if reflect.DeepEqual(expected, actual) {
		fmt.Printf("TEST %v PASSED\n", t.Name())
	} else {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestGrepN(t *testing.T) {
	data := []byte("Ilya Sleep\nDen Work\nPasha Eat\nMax nervous\nAnfisa chiiling\nZhenya swimming\n")
	file, err := createTmp(data)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	expected := "4.Max nervous \n"

	actual := getStdout(func() {
		grepN(file.Name(), "Max nervous")
	})
	if reflect.DeepEqual(expected, actual) {
		fmt.Printf("TEST %v PASSED\n", t.Name())
	} else {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
