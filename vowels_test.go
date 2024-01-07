package funwithwords

import "testing"

func TestWord(t *testing.T) {
	tests := []struct {
		want, got string
	}{
		{"aber", "aber"},
	}

	for _, test := range tests {
		w := word(test.got)
		if test.want != w {
			t.Fatalf("Test %s; got %s", test.want, w)
		}
	}
}
