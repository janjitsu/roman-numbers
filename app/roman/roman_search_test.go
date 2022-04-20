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
		{"CX", 110},
		{"XC", 90},
	}

	search := Search{}

	for _, wordtest := range tests {
		testname := fmt.Sprintf("word %s should have value of %d", wordtest.word, wordtest.expected)
		t.Run(testname, func(t *testing.T) {
			romanNumber, _ := search.createFromString(wordtest.word)
			if romanNumber.Value != wordtest.expected {
				t.Errorf("Error: %s", testname)
			}
		})
	}
}

func TestCreateFromStringThrowsError(t *testing.T) {
	search := Search{}
	_, err := search.createFromString("ABCXII")

	if err == nil {
		t.Errorf("Should throw error if invalid word")
	}
}

func TestFindInString(t *testing.T) {
	tests := []struct {
		word     string
		expected int
	}{
		{"ABIIIZWIX", 2},
		{"ABJJJZWRY", 0},
	}

	search := Search{}
	for _, wordtest := range tests {
		testname := fmt.Sprintf("Search for roman numbers in %s should get %d results", wordtest.word, wordtest.expected)
		t.Run(testname, func(t *testing.T) {
			result := search.findInString(wordtest.word)
			if len(result.Result) != wordtest.expected {
				t.Errorf("Error: %s", testname)
			}
		})
	}
}

func TestFindTopValueInString(t *testing.T) {
	tests := []struct {
		word     string
		expected int
	}{
		{"ABIIIZWIX", 9},
		{"ABJJJZWRYMC", 1100},
		{"ABIIIYZIVABC", 100},
		{"ABIIIABIVABVABVIABVII", 7},
		{"CXABIXZZXC", 110},
	}

	search := Search{}
	for _, wordtest := range tests {
		testname := fmt.Sprintf("Search for roman numbers in %s should get top value %d", wordtest.word, wordtest.expected)
		t.Run(testname, func(t *testing.T) {
			romanNumber := search.FindTopValueInString(wordtest.word)
			if romanNumber.Value != wordtest.expected {
				t.Errorf("Error: %s got %d instead", testname, romanNumber.Value)
			}
		})
	}
}
