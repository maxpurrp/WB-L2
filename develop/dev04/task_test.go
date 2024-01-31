package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAnagrams(t *testing.T) {
	input := []string{"ток", "кот", "куб", "грот", "бук", "торг"}
	expected := map[string][]string{
		"ток":  {"ток", "кот"},
		"куб":  {"куб", "бук"},
		"грот": {"грот", "торг"},
	}
	result := FindAnagrams(input)

	for k := range expected {
		expectedSlice := expected[k]
		resultSlice := result[k]
		assert.ElementsMatch(t, expectedSlice, resultSlice)
	}

}
