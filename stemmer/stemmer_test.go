package stemmer

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

func TestStem(t *testing.T) {
	file, err := os.Open("./files/da.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		test := strings.Fields(scanner.Text())
		// test[0] = got, test[1] = want
		got := Stem(strings.TrimSpace(test[0]))
		want := strings.TrimSpace(test[1])

		if got != want {
			t.Errorf("Got %s, want %s", got, want)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
