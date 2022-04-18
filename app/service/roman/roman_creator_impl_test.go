package roman

import (
	"fmt"
	"testing"
)

func TestCreateFromString(t *testing.T) {
	tests := []struct {
		word     string
		expected int
	}{
		{"III", 3},
		{"MDC", 1600},
		{"CIV", 104},
		{"CDXX", 420},
		{"CMXCII", 992},
		{"MMXXXIV", 2034},
		{"CMXCIX", 999},
	}

	for _, wordtest := range tests {
		testname := fmt.Sprintf("word %s should have value of %d", wordtest.word, wordtest.expected)
		t.Run(testname, func(t *testing.T) {
			romanNumber, _ := CreateFromString(wordtest.word)
			if romanNumber.value != wordtest.expected {
				t.Errorf("Error: %s", testname)
			}
		})
	}
}

func TestCreatFromStringThrwosError(t *testing.T) {
	_, err := CreateFromString("ABCXII")

	if err == nil {
		t.Errorf("Should throw error if invalid word")
	}
}
