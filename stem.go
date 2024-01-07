package funwithwords

func Stem(w string) string {
	// Remove main suffix
	if before, has := mainSuffix(w); has {
		w = before
	}
	// Return word
	return w
}
