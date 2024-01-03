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
func FindSuffix(w string) (string, bool) {
	endings := []string{
		"erendes",
		"erende",
		"hedens",
		"erede",
		"ethed",
		"heden",
		"endes",
		"erets",
		"heder",
		"ernes",
		"erens",
		"ered",
		"ende",
		"erne",
		"eres",
		"eren",
		"eret",
		"erer",
		"enes",
		"heds",
		"ens",
		"ene",
		"ere",
		"ers",
		"ets",
		"hed",
		"es",
		"et",
		"er",
		"en",
		"e",
	}
	// Special case
	if strings.ToLower(w) == "hedens" {
		return "", false
	}
	for _, e := range endings {
		if strings.HasSuffix(strings.ToLower(w), e) && len(w) > len(e) {
			return e, true
		}
	}
	return "", false
}
