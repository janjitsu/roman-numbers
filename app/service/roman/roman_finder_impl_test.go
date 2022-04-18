package roman

import (
	"testing"
)

func TestFindInString(t *testing.T) {
	text := "ABIIIZWIX"
	romanSearch := FindInString(text)

	if len(romanSearch.Result) != 2 {
		t.Errorf("Search for roman numbers in %s show return 2 results", text)
	}
}
