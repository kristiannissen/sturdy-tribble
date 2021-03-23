package stopword

import (
	str "strings"
)

var alphabet []byte

var Words []string

func init() {
	Words = []string{
		"ad",
		"af",
		"aldrig",
		"alene",
		"alle",
		"allerede",
		"alligevel",
		"alt",
		"altid",
		"anden",
		"andet",
		"andre",
		"at",
		"bag",
		"bare",
		"begge",
		"bl.a.",
		"blandt",
		"blev",
		"blive",
		"bliver",
		"burde",
		"bør",
		"ca.",
		"da",
		"de",
		"dem",
		"den",
		"denne",
		"dens",
		"der",
		"derefter",
		"deres",
		"derfor",
		"derfra",
		"deri",
		"dermed",
		"derpå",
		"derved",
		"det",
		"dette",
		"dig",
		"din",
		"dine",
		"disse",
		"dit",
		"dog",
		"du",
		"efter",
		"egen",
		"ej",
		"eller",
		"ellers",
		"en",
		"end",
		"endnu",
		"ene",
		"eneste",
		"enhver",
		"ens",
		"enten",
		"er",
		"et",
		"f.eks.",
		"far",
		"fem",
		"fik",
		"fire",
		"flere",
		"flest",
		"fleste",
		"for",
		"foran",
		"fordi",
		"forrige",
		"fra",
		"fx",
		"få",
		"får",
		"før",
		"først",
		"gennem",
		"gjorde",
		"gjort",
		"god",
		"godt",
		"gør",
		"gøre",
		"gørende",
		"ham",
		"han",
		"hans",
		"har",
		"havde",
		"have",
		"hej",
		"hel",
		"heller",
		"helt",
		"hen",
		"hende",
		"hendes",
		"henover",
		"her",
		"herefter",
		"heri",
		"hermed",
		"herpå",
		"hos",
		"hun",
		"hvad",
		"hvem",
		"hver",
		"hvilke",
		"hvilken",
		"hvilkes",
		"hvis",
		"hvor",
		"hvordan",
		"hvorefter",
		"hvorfor",
		"hvorfra",
		"hvorhen",
		"hvori",
		"hvorimod",
		"hvornår",
		"hvorved",
		"i",
		"igen",
		"igennem",
		"ikke",
		"imellem",
		"imens",
		"imod",
		"ind",
		"indtil",
		"ingen",
		"intet",
		"ja",
		"jeg",
		"jer",
		"jeres",
		"jo",
		"kan",
		"kom",
		"komme",
		"kommer",
		"kun",
		"kunne",
		"lad",
		"langs",
		"lav",
		"lave",
		"lavet",
		"lidt",
		"lige",
		"ligesom",
		"lille",
		"længere",
		"man",
		"mand",
		"mange",
		"med",
		"meget",
		"mellem",
		"men",
		"mens",
		"mere",
		"mest",
		"mig",
		"min",
		"mindre",
		"mindst",
		"mine",
		"mit",
		"mod",
		"må",
		"måske",
		"ned",
		"nej",
		"nemlig",
		"ni",
		"nogen",
		"nogensinde",
		"noget",
		"nogle",
		"nok",
		"nu",
		"ny",
		"nyt",
		"når",
		"nær",
		"næste",
		"næsten",
		"og",
		"også",
		"okay",
		"om",
		"omkring",
		"op",
		"os",
		"otte",
		"over",
		"overalt",
		"pga.",
		"på",
		"samme",
		"sammen",
		"se",
		"seks",
		"selv",
		"selvom",
		"senere",
		"ser",
		"ses",
		"siden",
		"sig",
		"sige",
		"sin",
		"sine",
		"sit",
		"skal",
		"skulle",
		"som",
		"stadig",
		"stor",
		"store",
		"synes",
		"syntes",
		"syv",
		"så",
		"sådan",
		"således",
		"tag",
		"tage",
		"temmelig",
		"thi",
		"ti",
		"tidligere",
		"til",
		"tilbage",
		"tit",
		"to",
		"tre",
		"ud",
		"uden",
		"udover",
		"under",
		"undtagen",
		"var",
		"ved",
		"vi",
		"via",
		"vil",
		"ville",
		"vor",
		"vore",
		"vores",
		"vær",
		"være",
		"været",
		"øvrigt",
	}

	alphabet = []byte("abcdefghijklmnopqrstuvwzyxæøå0123456789")
}

func RemoveStopWords(text string) []string {
	var words []string = Tokenize(text)
	res := []string{}

	for _, w := range words {
		if isstopword(w) == false {
			res = append(res, w)
		}
	}

	return res
}

func isstopword(word string) bool {
	for _, w := range Words {
		if w == word {
			return true
		}
	}
	return false
}

func contains(bytes []byte, b byte) bool {
	for _, a := range bytes {
		if a == b {
			return true
		}
	}
	return false
}

func Tokenize(text string) []string {
	var words []string = str.Split(text, " ")

	for i, w := range words {
		bytes := []byte(str.ToLower(w))
		s := []byte{}

		for _, b := range bytes {
			if contains(alphabet, b) {
				s = append(s, b)
			}
		}

		words[i] = string(s)
	}
	return words
}
