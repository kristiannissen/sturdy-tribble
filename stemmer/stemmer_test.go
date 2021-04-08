package stemmer

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

var words []struct {
	got, want string
}

func init() {
	file, err := os.Open("./files/da.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		test := strings.Fields(scanner.Text())
		got := strings.TrimSpace(test[0])
		want := strings.TrimSpace(test[1])

		s := struct {
			got, want string
		}{got, want}

		words = append(words, s)
	}
}

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
	table := []struct {
		got  string
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
			t.Errorf("Test: %s. Got %v, want %v", test.got, got, test.want)
		}
	}
}

func TestStep2(t *testing.T) {
	table := []struct {
		got, want string
	}{
		{"aflagd", "aflag"},
		{"aflagt", "aflag"},
	}

	for _, test := range table {
		got := step2(test.got)

		if got != test.want {
			t.Errorf("Test: %s. Got %s, want %s", test.got, got, test.want)
		}
	}
}

func TestStep3(t *testing.T) {
	table := []struct {
		got, want string
	}{
		{"allermægtigst", "allermæg"},
		{"almindelig", "almind"},
		{"alvorlig", "alvor"},
		{"løst", "løs"},
	}

	for _, test := range table {
		got := step3(test.got)

		if got != test.want {
			t.Errorf("Got %s, want %s", got, test.want)
		}
	}
}

func TestSearchInR1(t *testing.T) {
	table := []struct {
		got, want string
		exist     bool
	}{
		{"hingst", "t", true},
		{"hello", "llo", false},
		{"kitty", "sy", false},
		{"pussy", "na", false},
		{"pussies", "s", true},
	}

	for _, test := range table {
		b := searchInR1(test.got, test.want)

		if b != test.exist {
			t.Errorf("Got %s, want %s, %v", test.got, test.want, b)
		}
	}
}

func TestStep4(t *testing.T) {
	table := []struct {
		got, want string
	}{
		{"ødelægg", "ødelæg"},
		{"øjeblikk", "øjeblik"},
	}

	for _, test := range table {
		got := step4(test.got)

		if got != test.want {
			t.Errorf("Test %s. Got %s, want %s", test.got, got, test.want)
		}
	}
}

func TestStem(t *testing.T) {
	file, err := os.Open("./files/da.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		test := strings.Fields(scanner.Text())
		got := Stem(strings.TrimSpace(test[0]))
		want := strings.TrimSpace(test[1])

		if got != want {
			t.Errorf("Test %s. Got %s, want %s", test[0], got, want)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
