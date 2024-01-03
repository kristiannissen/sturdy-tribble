package funwithwords

import "testing"

func TestUndouble(t *testing.T) {
	tests := []struct {
		got, want string
		res       bool
	}{
		{"bestemm", "bestem", true},
		{"hund", "hund", false},
		{"kat", "kat", false},
		{"indtræff", "indtræf", true},
		{"sø", "sø", false},
		{"ø", "ø", false},
	}

	for _, test := range tests {
		s, _ := undouble(test.got)
		if test.want != s {
			t.Fatalf("Test %s; want %s got %s", test.got, test.want, s)
		}
	}
}
