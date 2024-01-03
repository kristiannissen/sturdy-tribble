package funwithwords

import "testing"

func TestIsStopWord(t *testing.T) {
	tests := []struct {
		inp string
		res bool
	}{
		{"fremmødte", false},
		{"i", true},
		{"vinterjakker", false},
		{"havde", true},
		{"torsdag", false},
		{"formiddag", false},
		{"trodset", false},
		{"minusgraderne", false},
		{"og", true},
		{"var", true},
		{"mødt", false},
	}

	for _, test := range tests {
		b := IsStopWord(test.inp)
		if b != test.res {
			t.Fatalf("Test %s; want %v got %v", test.inp, test.res, b)
		}
	}
}
