package stemmer

import (
	"testing"
)

func TestStem(t *testing.T) {
	tests := []struct {
		got, want string
	}{
		{"indtage", "indtag"},
		{"indtagelse", "indtag"},
		{"indtager", "indtag"},
		{"indtages", "indtag"},
		{"indtaget", "indtag"},
		{"undertrykt", "undertryk"},
		{"bestemmelse", "bestem"},
		{"friskt", "frisk"},
	}

	for _, test := range tests {
		got := Stem(test.got)

		// Print current test case to stdout
		// t.Logf("Testing %s", got)

		if test.want != got {
			t.Errorf("got %s, want %s", got, test.want)
		}
	}
}
