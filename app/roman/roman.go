package roman

type Finder interface {
	FindTopValueInString(text string) *RomanNumber
}

type RomanNumber struct {
	Number string `json:"number"`
	Value  int    `json:"value"`
}
