package stemmer

import (
	"sort"
	str "strings"
)

type Suffix struct {
	Word     string
	Priority int
}

var Suffixes []Suffix

func init() {
	Suffixes = []Suffix{
		{"hed", 1},
		{"ethed", 1},
		{"ered", 1},
		{"ered", 1},
		{"e", 2},
		{"erede", 1},
		{"erende", 1},
		{"ene", 1},
		{"erne", 1},
		{"ere", 1},
		{"en", 2},
		{"heden", 1},
		{"eren", 1},
		{"er", 1},
		{"heder", 1},
		{"erer", 1},
		{"heds", 1},
		{"es", 0},
		{"endes", 1},
		{"erendes", 1},
		{"enes", 1},
		{"ernes", 1},
		{"eres", 1},
		{"ens", 0},
		{"hedens", 1},
		{"erens", 1},
		{"ers", 1},
		{"ets", 1},
		{"erets", 1},
		{"et", 0},
		{"eret", 1},
		{"else", 1},
	}
	sort.SliceStable(Suffixes, func(i, j int) bool {
		return Suffixes[i].Priority < Suffixes[j].Priority
	})
}

func Stem(word string) string {
	return main_suffix(word)
}

func main_suffix(word string) string {
	word = str.TrimSpace(str.ToLower(word))

	for _, s := range Suffixes {
		if str.HasSuffix(word, s.Word) {
			word = str.TrimRight(word, s.Word)
		}
	}
	return word
}
