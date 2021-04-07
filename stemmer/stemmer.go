package stemmer

import (
	// "sort"
	str "strings"
	// "log"
)

var r1 int

func init() {
	r1 = 3
}

func Stem(word string) string {
	word = str.TrimSpace(str.ToLower(word))

	return word
}

/**
 * Helper to search within substring
 */
func searchInR1(w string, k string) bool {
	var word string = w[r1:]
	return str.HasSuffix(word, k)
}

/**
 * Step 1
 * Search for the longest among the following suffixes in R1, and perform the action indicated.
 */
func step1(w string) string {
	var suffixes []string = []string{"erendes", "erende", "hedens", "erede", "ethed", "heden", "endes", "erets", "heder", "ernes", "erens", "ered", "ende", "erne", "eres", "eren", "eret", "erer", "enes", "heds", "ens", "ene", "ere", "ers", "ets", "hed", "es", "et", "er", "en", "e"}

	for _, suffix := range suffixes {
		if str.HasSuffix(w, suffix) {
			w = str.TrimSuffix(w, suffix)
			break
		}
	}

	// s
	// delete if preceded by a valid s-ending
	if str.HasSuffix(w, "s") {
		if hasValidEnding(w) == false {
			w = str.TrimSuffix(w, "s")
		}
	}

	return w
}

/**
 * Define a valid s-ending as one of
 * a   b   c   d   f   g   h   j   k   l   m   n   o   p   r   t   v   y   z   å
 */
func hasValidEnding(w string) bool {
	var letters []string = []string{"a", "b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "o", "p", "r", "t", "v", "y", "z", "å"}
	var word string = w[0 : len(w)-1]

	for _, l := range letters {
		if str.HasSuffix(word, l) {
			return true
		}
	}

	return false
}

/**
 * Step 2
 * Search for one of the following suffixes in R1, and if found delete the last letter
 * gd   dt   gt   kt
 */
func step2(w string) string {
	var endings []string = []string{"gd", "dt", "gt", "kt"}

	for _, e := range endings {
		if searchInR1(w, e) {
			if str.HasSuffix(w, e) {
				return w[0 : len(w)-1]
			}
		}
	}

	return w
}

/**
 * Step3
 */
func step3(w string) string {
	// If the word ends igst, remove the final st
	if str.HasSuffix(w, "igst") {
		w = str.TrimSuffix(w, "st")
	}

	// Search for the longest among the following suffixes in R1, and perform the action indicated
	// ig   lig   elig   els
	//      delete, and then repeat step 2
	var suffixes []string = []string{"elig", "lig", "ig", "els"}

	for _, suffix := range suffixes {
		if searchInR1(w, suffix) {
			w = str.TrimSuffix(w, suffix)
			w = step2(w)
		}
	}

	// løst
	// replace with løs
	if str.HasSuffix(w, "løst") {
		w = w[0 : len(w)-1]
	}

	return w
}
