package funwithwords

import "strings"

func Stem(w string) string {

	// Step 1
	// Search for longest suffix and cut
	suf, has := mainSuffix(w)
	if has {
		// A. cut suffix
		w, _ = strings.CutSuffix(w, suf)
	}

	return w
}
