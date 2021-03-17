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
		{"indtil", "indtil"},
		{"indtog", "indtog"},
		{"indtraf", "indtraf"},
		{"indtryk", "indtryk"},
		{"indtræde", "indtræd"},
		{"indtræder", "indtræd"},
		{"indtræffe", "indtræf"},
		{"indtræffer", "indtræf"},
		{"indtrængende", "indtræng"},
		{"indtægt", "indtæg"},
		{"indtægter", "indtæg"},
		{"indvandrede", "indvandred"},
		{"indvandret", "indvandr"},
		{"indvender", "indvend"},
		{"indvendig", "indvend"},
		{"indvendige", "indvend"},
		{"indvendigt", "indvend"},
		{"indvending", "indvending"},
		{"indvendingerne", "indvending"},
		{"indvie", "indvi"},
		{"indviede", "indvied"},
		{"indvielse", "indvi"},
		{"indvielsen", "indvi"},
		{"indvielsesløfte", "indvielsesløft"},
		{"indvielsestid", "indvielsestid"},
		{"indvier", "indvi"},
		{"indvies", "indvi"},
		{"indviet", "indvi"},
		{"indvikle", "indvikl"},
		{"indvikler", "indvikl"},
		{"indvolde", "indvold"},
		{"indvoldene", "indvold"},
		{"indvortes", "indvort"},
		{"indånde", "indånd"},
		{"indåndede", "indånded"},
		{"underste", "underst"},
		{"undersåtter", "undersåt"},
		{"undersåtters", "undersåt"},
		{"undersøg", "undersøg"},
		{"undersøge", "undersøg"},
		{"undersøgelse", "undersøg"},
		{"undersøgelsen", "undersøg"},
		{"undersøger", "undersøg"},
		{"undersøgt", "undersøg"},
		{"undersøgte", "undersøg"},
		{"undertryk", "undertryk"},
		{"undertrykke", "undertryk"},
		{"undertrykkelse", "undertryk"},
		{"undertrykker", "undertryk"},
		{"undertrykkere", "undertryk"},
		{"undertrykkeren", "undertryk"},
		{"undertrykkerens", "undertryk"},
		{"undertrykkeres", "undertryk"},
		{"undertrykkes", "undertryk"},
		{"undertrykt", "undertryk"},
		{"undertrykte", "undertryk"},
		{"undertryktes", "undertryk"},
		{"undertvang", "undertvang"},
		{"undertvunget", "undertvung"},
		{"undertvungne", "undertvungn"},
		{"undervejs", "undervej"},
		{"underverdenen", "underverden"},
		{"undervise", "undervis"},
		{"underviser", "undervis"},
		{"undervises", "undervis"},
		{"undervisning", "undervisning"},
		{"undervisningen", "undervisning"},
		{"undervist", "undervist"},
		{"underviste", "undervist"},
		{"underværk", "underværk"},
		{"underværker", "underværk"},
		{"undevise", "undevis"},
		{"undeviste", "undevist"},
		{"undfange", "undfang"},
		{"undfanged", "undfanged"},
	}

	for _, test := range tests {
		got := Stem(test.got)

		// log.Printf("Got %s want %s stem %s", test.got, test.want, got)

		if test.want != got {
			t.Errorf("Got %s, want %s, test %s", got, test.want, test.got)
		}
	}
}
