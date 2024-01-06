package funwithwords

import "testing"

func TestStem(t *testing.T) {
	// Create test cases
	tests := []struct {
		got, want string
	}{
		{"uvederhæftighed", "uvederhæftig"},             // hed
		{"utryghed", "utryg"},                           // hed
		{"tillukkethed", "tilluk"},                      // ethed
		{"tilpassethed", "tilpas"},                      // ethed
		{"astringered", "astring"},                      // ered
		{"zonegrænse", "zonegræns"},                     // e
		{"ynke", "ynk"},                                 // e
		{"udspekulerede", "udspekul"},                   // erede
		{"vismændene", "vismænd"},                       // ene
		{"yderverdnerne", "yderverdn"},                  // erne
		{"vulkanisere", "vulkanis"},                     // ere
		{"wienerschnitzelen", "wienerschnitzel"},        // en
		{"imponeretheden", "imponeret"},                 // heden
		{"versemageren", "versemag"},                    // eren
		{"ynglingealder", "ynglingeald"},                // er
		{"tronstridigheder", "tronstridig"},             // heder
		{"regredierer", "regredi"},                      // erer
		{"toværelserslejligheds", "toværelserslejlig"},  // heds
		{"ypperstes", "ypperst"},                        // es
		{"videnskabeliggørendes", "videnskabeliggør"},   // endes
		{"øverstkommanderendes", "øverstkommand"},       // erendes
		{"yngstemændenes", "yngstemænd"},                // enes
		{"wienerschnitzlernes", "wienerschnitzl"},       // ernes
		{"velopdragneres", "velopdragn"},                // eres
		{"wienerschnitzelens", "wienerschnitzel"},       // ens
		{"cateringvirksomhedens", "cateringvirksom"},    // hedens
		{"svigerfaderens", "svigerfad"},                 // erens
		{"wienerschnitzlers", "wienerschnitzl"},         // ers
		{"whiteboardets", "whiteboard"},                 // ets
		{"undervisningscenterets", "undervisningscent"}, // erets
		{"vurderingskriteriet", "vurderingskriteri"},    // et
		{"velproportioneret", "velproportion"},          // eret
	}

	for _, test := range tests {
		s := Stem(test.got)
		// Compare the stemmed word
		if s != test.want {
			t.Fatalf("Test: %s, want %s got %s", test.got, test.want, s)
		}
	}
}
