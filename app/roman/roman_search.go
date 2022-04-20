package roman

import (
	"errors"
	"regexp"
	"strings"
)

type Search struct {
	Result []RomanNumber
}

func (r *Search) FindTopValueInString(text string) *RomanNumber {
	r.findInString(text)
	topValue := 0
	var topRomanNumber RomanNumber

	for _, number := range r.Result {
		if number.Value > topValue {
			topValue = number.Value
			topRomanNumber = number
		}
	}

	return &topRomanNumber
}

func (r *Search) findInString(text string) *Search {
	r.Result = []RomanNumber{}

	rgxp, _ := regexp.Compile("[IVXLCDM]+")

	words := rgxp.FindAllString(text, -1)

	for _, word := range words {
		romanNumber, _ := r.createFromString(word)
		r.Result = append(r.Result, *romanNumber)
	}

	return r
}

func (r *Search) createFromString(text string) (*RomanNumber, error) {

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

	rgxp, _ := regexp.Compile("[^IVXLCDM]")
	if rgxp.MatchString(text) {
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
			total += roman[chars]
			chars = ""
		}
	}

	return &RomanNumber{text, total}, nil
}
