package funwithwords

func Stem(w string) string {
	// Remove main suffix
	if before, has := mainSuffix(w); has {
		return before
	}
	return w
}
