package funwithwords

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestStem(t *testing.T) {
	f, err := os.Open("da.txt")
	defer f.Close()

	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		c := strings.Split(s.Text(), ",")

		stem := Stem(c[0])
		if stem != c[1] {
			t.Fatalf("Test %s; want %s got %s", c[0], c[1], stem)
		}
	}
}
