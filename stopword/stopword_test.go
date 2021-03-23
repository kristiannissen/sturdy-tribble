package stopword

import (
	str "strings"
	"testing"
)

func TestRemoveStopWords(t *testing.T) {
	table := []struct {
		got, want string
	}{
		{"Hiroshi Sasaki, der indtil torsdag", "hiroshi sasaki torsdag"},
		{"Årsagen er, at han er kommet med nedsættende bemærkninger, øh sagde han", "årsagen kommet nedsættende bemærkninger øh sagde"},
	}

	for _, test := range table {
		got := str.Join(RemoveStopWords(test.got), " ")

		if got != test.want {
			t.Fatalf("Got %s, want %s", got, test.want)
		}
	}
}

func TestTokenize(t *testing.T) {
	table := []struct {
		got, want string
	}{
		{"’Olympig’", "olympig"},
		{"Reuters.", "reuters"},
		{"Sasaki,", "sasaki"},
		{"Hiroshi Sasaki, der indtil torsdag", "hiroshi sasaki der indtil torsdag"},
	}

	for _, test := range table {
		got := str.Join(Tokenize(test.got), " ")

		if got != test.want {
			t.Fatalf("Got %s, want %s, test %s", got, test.want, test.got)
		}
	}
}
