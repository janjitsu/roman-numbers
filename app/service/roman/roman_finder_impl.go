package roman

import (
	"regexp"
)

func FindInString(text string) RomanSearch {
	r, _ := regexp.Compile("[IVXLCDM]+")

	search := RomanSearch{Result: []RomanNumber{}}

	words := r.FindAllString(text, -1)

	for _, word := range words {
		romanNumber, _ := CreateFromString(word)
		search.Result = append(search.Result, *romanNumber)
	}

	return search
}
