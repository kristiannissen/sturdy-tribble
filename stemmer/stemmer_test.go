package stemmer

import (
	// "bufio"
	// "log"
	// "os"
	// "strings"
	"testing"
)
/**
 * Step 1
 * Search for the longest among the following suffixes in R1, and perform the action indicated.
 */
func TestStep1(t *testing.T) {
	table := []struct {
		got, want string
	}{
		{"herlighedens", "herlig"},
		{"afleverede", "aflev"},
		{"ondsindethed", "ondsind"},
		{"oprigtigheden", "oprigtig"},
		{"seendes", "seend"},
		{"skyldofferets", "skyldoff"},
		{"stridigheder", "stridig"},
		{"syndernes", "synd"},
		{"søsterens", "søst"},
		{"tilbered", "tilb"},
		{"tilflugtsbyerne", "tilflugtsby"},
		{"tjeneres", "tjen"},
		{"tolderen", "told"},
		{"udleveret", "udlev"},
		{"vognenes", "vogn"},
		{"yppigheds", "yppig"},
		{"baalshøjene", "baalshøj"},
		{"evangeliets", "evangeli"},
		{"evighed", "evig"},
		{"evighedernes", "evighed"},
		{"evnet", "evn"},
		{"ezer", "ezer"},
		{"ezraiten", "ezrait"},
	}

	for _, test := range table {
        w := step1(test.got)

        if w != test.want {
			t.Errorf("Got %s, want %s", w, test.want)
		}
	}
}
/**
 * Test hasValidEnding
 */
func TestHasValidEnding(t *testing.T) {
    table := []struct{
        got string
        want bool
    }{
        {"abdas", true},
        {"abdeels", true},
        {"afvis", false},
        {"agabus", false},
    }

    for _, test := range table {
        got := hasValidEnding(test.got)

        if got != test.want {
            t.Errorf("Got %v, want %v", got, test.want)
        }
    }
}

func TestStep2(t *testing.T) {
    table := []struct{
        got, want string
    }{
        {"aflagd", "aflag"},
        {"aflagt", "aflag"},
    }

    for _, test := range table {
        got := step2(test.got)

        if got != test.want {
            t.Errorf("Got %s, want %s", got, test.want)
        }
    }
}

func TestStem(t *testing.T) {
	t.Skip()
}
