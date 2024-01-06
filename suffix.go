package funwithwords

import (
	"strings"
)

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
	for _, e := range endings {
		if strings.HasSuffix(w, e) {
			w, b = strings.CutSuffix(w, e)
			break
		}
	}
	return w, b
}
