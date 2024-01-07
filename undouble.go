package funwithwords

import (
	"strings"
)

func undouble(w string) (string, bool) {
	// Split string into chars
	s := strings.Split(w, "")
	// Handle error
	if len(s) < 3 {
		return w, false
	}
	// Compare the last 2
	if strings.Compare(s[len(s)-2], s[len(s)-1]) == 0 {
		// Return string without trailing character
		return strings.Join(s[:len(s)-1], ""), true
	}
	// Otherwise return w and false
	return w, false
}
