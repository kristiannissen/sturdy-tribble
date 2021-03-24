package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const dataset string = "./dataset"

func main() {
	chn := make(chan string)

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

	go func(fp []string) {
		for _, f := range fp {
			text, _ := ioutil.ReadFile(f)
			chn <- string(text)
		}
		close(chn)
	}(files)

	for fp := range chn {
		log.Println(fp)
	}

	if err != nil {
		panic(err)
	}
}
