package funwithwords

import "testing"

func TestStem(t *testing.T) {
	// Create test cases
	tests := []struct {
		got, want string
	}{
		{"uvederhæftighed", "uvederhæftig"}, // hed
		{"utryghed", "utryg"},               // hed
		{"tillukkethed", "tilluk"},          // ethed
		{"tilpassethed", "tilpas"},          // ethed
		{"astringered", "astring"},          // ered
		{"zonegrænse", "zonegræns"},         // e
		{"ynke", "ynk"},                     // e
		{"udspekulerede", "udspekul"},       // erede
	}

	for _, test := range tests {
		s := Stem(test.got)
		// Compare the stemmed word
		if s != test.want {
			t.Fatalf("Test: %s, want %s got %s", test.got, test.want, s)
		}
	}
}
