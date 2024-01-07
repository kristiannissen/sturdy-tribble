package funwithwords

import (
	"strings"
)

func IsValidFuffix(w string) bool {
	return false
}

func isValidSEnding(w string) bool {
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
		"erendes",
		"hedens",
		"ernes",
		"ethed",
		"erede",
		"heden",
		"erens",
		"heder",
		"erets",
		"endes",
		"ende",
		"enes",
		"erer",
		"ered",
		"heds",
		"erne",
		"eren",
		"eres",
		"eret",
		"ens",
		"ers",
		"ets",
		"ene",
		"hed",
		"ede",
		"ere",
		"et",
		"en",
		"er",
		"es",
		"e",
	}
	// Check main suffix
	for _, e := range endings {
		if strings.HasSuffix(w, e) {
			w, b = strings.CutSuffix(w, e)
			break
		}
	}
	// S ending
	if strings.HasSuffix(w, "s") {
		if isValidSEnding(w) == false {
			return w[:len(w)-1], true
		}
	}
	return w, b
}

// Other suffix
func otherSuffix(w string) (string, bool) {
	endings := []string{
		"elig",
		"lig",
		"els",
		"ig",
	}
	if strings.HasSuffix(w, "igst") {
		// remove final st
		w, _ := strings.CutSuffix(w, "st")
		// Find longest ending and cut
		for _, e := range endings {
			if strings.HasSuffix(w, e) {
				return strings.CutSuffix(w, e)
			}
		}
	}
	return w, false
}
