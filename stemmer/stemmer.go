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
	word = valid_ending(undouble(other_suffix(consonant_pairs(main_suffix(word)))))

	return word
}

func undouble(word string) string {
	chars := str.Split(word, "")

	if len(chars) > 2 {
		if chars[len(chars)-2] == chars[len(chars)-1] {
			word = str.TrimSuffix(word, chars[len(chars)-1])
		}
	}
	return word
}

func consonant_pairs(word string) string {
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

// If the word ends in 'igst', remove the final st.
// Search for the longest among the following suffixes,
// and perform the action indicated
func other_suffix(word string) string {
	if str.HasSuffix(word, "igst") {
		word = str.TrimSuffix(word, "st")
	}

	endings := []string{"elig", "lig", "els", "ig"}

	for _, e := range endings {
		if str.HasSuffix(word, e) {
			word = str.TrimRight(word, e)
			break
		}
	}
	return word
}

func valid_ending(word string) string {
	if len(word) > 2 {
		chars := str.Split("abcdfghjklmnoprtvyzå", "")
		l := word[len(word)-2 : len(word)-1]

		if str.HasSuffix(word, "s") {
			for _, c := range chars {
				if l == c {
					word = str.TrimSuffix(word, "s")
				}
			}
		}
	}
	return word
}
