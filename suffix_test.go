package funwithwords

import "testing"

func TestIsValidSEnding(t *testing.T) {
	tests := []struct {
		word string
		res  bool
	}{
		{"digterens", true},
		{"barndommens", true},
		{"tag", false},
	}

	for _, test := range tests {
		b := IsValidSEnding(test.word)
		if b != test.res {
			t.Fatalf("Test %s, want %v got %v", test.word, test.res, b)
		}
	}
}

func TestFindLongestSuffix(t *testing.T) {
	tests := []struct {
		word, ending string
		res          bool
	}{
		{"studerendes", "erendes", true},
		{"sikkerheds", "heds", true},
		{"ende", "e", true},
		{"hedens", "", false},
		{"cyklisten", "en", true},
	}

	for _, test := range tests {
		e, _ := FindLongestSuffix(test.word)
		if e != test.ending {
			t.Fatalf("Test %s, want %s got %s", test.word, test.ending, e)
		}
	}
}
