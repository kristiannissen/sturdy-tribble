package stemmer

import (
	// "sort"
	str "strings"
	// "log"
)

var r1 int
var vowels []string

func init() {
	r1 = 3
	vowels = []string{"a", "e", "i", "o", "u", "y", "æ", "å", "ø"}
}

func Stem(word string) string {
	word = str.TrimSpace(str.ToLower(word))

	word = step1(word)
	// word = step2(word)
	// word = step3(word)
	// word = step4(word)

	return word
}

/**
 * Helper to search within substring
 */
func searchInR1(w string, k string) bool {
	// Avoid slice bounds out of range
	if len(w) < r1 {
		return false
	}

	var word string = w[r1:]
	if len(word) > r1 {
		return str.HasSuffix(word, k)
	}

	return false
}

/**
 * Step 1
 * Search for the longest among the following suffixes in R1, and perform the action indicated.
 */
func step1(w string) string {
	var suffixes []string = []string{"erendes", "erende", "hedens", "erede", "ethed", "heden", "endes", "erets", "heder", "ernes", "erens", "ered", "ende", "erne", "eres", "eren", "eret", "erer", "enes", "heds", "ens", "ene", "ere", "ers", "ets", "hed", "es", "et", "er", "en", "e"}

	for _, suffix := range suffixes {
		if searchInR1(w, suffix) {
			w = str.TrimSuffix(w, suffix)
			break
		}
	}

	// s
	// delete if preceded by a valid s-ending
	if searchInR1(w, "s") {
		if hasValidEnding(w) == true {
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
			return w[0 : len(w)-1]
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

/**
 * Step 4: undouble
 * If the word ends with double consonant in R1, remove one of the consonants
 */
func step4(w string) string {
	ln := len(w)
	if ln < r1 {
		return w
	}

	lastLetter := w[len(w)-1 : len(w)]
	for _, letter := range vowels {
		if lastLetter == letter {
			return w
		}
	}
	beforeLastLetter := w[len(w)-2 : len(w)-1]
	if lastLetter == beforeLastLetter {
		return w[0 : len(w)-1]
	}

	return w
}
