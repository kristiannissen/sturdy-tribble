package funwithwords

func Stem(w string) string {
	// Remove main suffix
	if before, has := mainSuffix(w); has {
		w = before
	}
	// Check undouble
	if und, has := undouble(w); has {
		w = und
	}
	// Return word
	return w
}
