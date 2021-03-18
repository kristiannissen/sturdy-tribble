package stemmer

import (
	"log"
	"sort"
	str "strings"
)

var Suffixes []string

func init() {
	log.Println("stem init")
	Suffixes = []string{
		"hed",
		"ethed",
		"ered",
		"e",
		"erede",
		"ende",
		"erende",
		"ene",
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
	word = undouble(other_suffix(constant_pairs(main_suffix(word))))

	return word
}

// FIXME: Will break on short words
func undouble(word string) string {
	chars := str.Split(word, "")

	if chars[len(chars)-2] == chars[len(chars)-1] {
		word = str.TrimSuffix(word, chars[len(chars)-1])
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
			word = str.TrimSuffix(word, s)
			break
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
			break
		}
	}
	return word
}
