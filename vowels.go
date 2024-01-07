package funwithwords

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'y', 'æ', 'å', 'ø':
		return true
	}
	return false
}
func word(w string) string {
	return w
}
