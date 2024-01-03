package funwithwords

import "testing"

func TestStem(t *testing.T) {
	tests := []struct {
		got, want string
	}{
		{"indtage", "indtag"},
		{"indtagelse", "indtag"},
		{"indtager", "indtag"},
		{"indtages", "indtag"},
		{"indtaget", "indtag"},
		{"indtil", "indtil"},
		{"indtog", "indtog"},
		{"indtræde", "indtræd"},
		{"indtræffe", "indtræf"},
	}

	for _, test := range tests {
		s := Stem(test.got)
		if s != test.want {
			t.Fatalf("Test %s; want %s got %s", test.got, test.want, s)
		}
	}
}
