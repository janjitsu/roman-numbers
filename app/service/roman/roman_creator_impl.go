package roman

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func CreateFromString(text string) (*RomanNumber, error) {

	roman := map[string]int{
		"I":  1,
		"II": 2,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XX": 20,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CC": 200,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
		"MM": 2000,
	}

	r, _ := regexp.Compile("[^IVXLCDM]")
	if r.MatchString(text) {
		return nil, errors.New("Invalid Roman Number word")
	}

	letters := strings.Split(text, "")
	wordsize := len(letters)
	total := 0
	chars := ""

	for i := 0; i < wordsize; i++ {
		chars += letters[i]
		hasnext := i < (wordsize - 1)

		// check if a new value can be formed with next character and skip it
		if hasnext && roman[chars+letters[i+1]] > 0 {
			continue
		}

		if roman[chars] > 0 {
			fmt.Printf("%s = %d\n", chars, roman[chars])
			total += roman[chars]
			chars = ""
		}
	}

	return &RomanNumber{text, total}, nil
}
