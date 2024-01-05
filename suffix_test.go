package funwithwords

import "testing"

func TestMainSuffix(t *testing.T) {
	tests := []struct {
		got, want string
	}{
		{"vidtgående", "vidtgå"},
		{"vuggende", "vugg"},
		{"visaene", "visa"},
		{"workshoppene", "workshopp"},
	}

	for _, test := range tests {
		s, _ := mainSuffix(test.got)
		if s != test.want {
			t.Fatalf("Test %s; got %s want %s", test.got, s, test.want)
		}
	}
}
