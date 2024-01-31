package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestCut(t *testing.T) {
	inputFile, err := os.CreateTemp("", "input_test_*.txt")
	if err != nil {
		t.Fatalf("Error creating temporary input file: %v", err)
	}
	defer os.Remove(inputFile.Name())

	testData := "price\tweight\tcount\n100\t15g\t3\n250\t5g\t5\n"
	if _, err := inputFile.Write([]byte(testData)); err != nil {
		t.Fatalf("Error writing to temporary input file: %v", err)
	}
	if err := inputFile.Close(); err != nil {
		t.Fatalf("Error closing temporary input file: %v", err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin, err = os.Open(inputFile.Name())
	if err != nil {
		t.Fatalf("Error opening temporary input file for reading: %v", err)
	}
	defer os.Stdin.Close()

	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()
	r, w, _ := os.Pipe()
	os.Stdout = w

	cut([]int{0, 2}, "\t", false)

	w.Close()

	var outputBuf bytes.Buffer
	_, err = io.Copy(&outputBuf, r)
	if err != nil {
		t.Fatalf("Error reading from captured output: %v", err)
	}

	expectedOutput := "price\tcount\n100\t3\n250\t5\n"

	if outputBuf.String() != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, outputBuf.String())
	}
}
