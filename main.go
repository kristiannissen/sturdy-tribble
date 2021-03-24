package main

import (
	"funwithwords/stopword"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	// "funwithwords/stemmer"
)

const dataset string = "./dataset"

func main() {
	chn := make(chan map[string][]string)

	files := []string{}

	err := filepath.Walk(dataset, func(path string, info os.FileInfo, err error) error {
		abs, _ := filepath.Abs(path)

		if strings.HasSuffix(info.Name(), ".txt") {
			files = append(files, abs)
		}

		if err != nil {
			log.Fatal(err)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	go func(fp []string) {
		for _, f := range fp {
			text, _ := ioutil.ReadFile(f)
			m := make(map[string][]string)
			m[f] = stopword.RemoveStopWords(string(text))

			chn <- m
		}
		close(chn)
	}(files)

	for f := range chn {
		for k, v := range f {
			log.Println(k, v)
		}
	}
}
