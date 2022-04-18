package roman

type Creator interface {
	CreateFromString(text string) (*RomanNumber, error)
}

type Finder interface {
	FindInString(text string) []RomanNumber
}

type RomanNumber struct {
	number string `json:"number"`
	value  int    `json:"value"`
}

type RomanSearch struct {
	Result []RomanNumber
}
