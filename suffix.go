package funwithwords

import "strings"

func IsValidFuffix(w string) bool {
	return false
}

func IsValidSEnding(w string) bool {
	s := w[len(w)-2:] // Gives us ns from Digterens
	e := []string{
		"es",
		"rs",
		"ns",
		"ts",
		"is",
		"ds",
		"ms",
		"ts",
		"es",
		"øs",
	}

	for _, f := range e {
		if s == f {
			return true
		}
	}
	return false
}

// Finds erendes in studerendes
func mainSuffix(w string) (string, bool) {
	var b bool = false
	endings := []string{
		"ethed",
		"erede",
		"ende",
		"ered",
		"ene",
		"hed",
		"ede",
		"et",
		"e",
	}
	for _, e := range endings {
		if strings.HasSuffix(strings.ToLower(w), e) && len(w) > len(e) {
			w, b = strings.CutSuffix(w, e)
		}
	}
	return w, b
}
