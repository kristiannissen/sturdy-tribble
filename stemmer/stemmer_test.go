package stemmer

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

var testTable []struct {
	got, want string
}
/**
 * Setup
 */
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

		testTable = append(testTable, s)
	}
}

/**
 * Step 1
 * Search for the longest among the following suffixes in R1, and perform the action indicated.
 */
func TestStep1(t *testing.T) {
	var suffixes []string = []string{"erendes", "erende", "hedens", "erede", "ethed", "heden", "endes", "erets", "heder", "ernes", "erens", "ered", "ende", "erne", "eres", "eren", "eret", "erer", "enes", "heds", "ens", "ene", "ere", "ers", "ets", "hed", "es", "et", "er", "en", "e"}

    for _, test := range testTable {
        for _, suffix := range suffixes {
            if strings.HasSuffix(test.got, suffix) {
                got := step1(test.got)

                if got != test.want {
                    t.Errorf("Test %s. Got %s want %s", test.got, got, test.want)
                }
            }
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
		{"øvet", "et", false},
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
