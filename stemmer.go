package stemmer

import (
	"sort"
	str "strings"
)

var Suffixes []string

func init() {
	Suffixes = []string{
		"hed",
		"ethed",
		"ered",
		"e",
		"erede",
		"erende",
		"ene",
		"ende",
		"erne",
		"ere",
		"en",
		"heden",
		"eren",
		"er",
		"heder",
		"erer",
		"heds",
		"es",
		"endes",
		"erendes",
		"enes",
		"ernes",
		"eres",
		"ens",
		"hedens",
		"erens",
		"ers",
		"ets",
		"erets",
		"et",
		"eret",
		"else",
	}
	SortSuffixes()
}

func SortSuffixes() {
	sort.SliceStable(Suffixes, func(i, j int) bool {
		return len(Suffixes[i]) > len(Suffixes[j])
	})
}

func Stem(word string) string {
	word = str.TrimSpace(str.ToLower(word))
	return other_suffix(undouble(constant_pairs(main_suffix(word))))
}

// FIXME: Will break on short words
func undouble(word string) string {
	l1 := word[len(word)-1:]
	l2 := word[len(word)-2 : len(word)-1]
	if l1 == l2 {
		word = word[:len(word)-1]
	}
	return word
}

func constant_pairs(word string) string {
	pairs := make(map[string]string)
	pairs["kt"] = "t"
	pairs["dt"] = "t"
	pairs["gd"] = "d"
	pairs["gt"] = "t"

	for k, v := range pairs {
		if str.HasSuffix(word, k) {
			word = str.TrimRight(word, v)
		}
	}
	return word
}

func main_suffix(word string) string {
	for _, s := range Suffixes {
		if str.HasSuffix(word, s) {
			word = str.TrimRight(word, s)
		}
	}
	return word
}

func other_suffix(word string) string {
	pairs := make(map[string]string)
	pairs["els"] = "els"
	pairs["ig"] = "ig"

	for k, v := range pairs {
		if str.HasSuffix(word, k) {
			word = str.TrimRight(word, v)
		}
	}
	return word
}
