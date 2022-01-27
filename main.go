package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	genIndex := flag.Bool("genindex", false, "regenerate the index")
	searchTerm := flag.String("search", "", "search the index")
	flag.Parse()

	if *genIndex {
		log.Println("starting index")
		index, err := NewIndex()
		if err != nil {
			log.Fatal(err)
		}
		i := 0

		err = filepath.Walk("./raw-data",
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}
				data, err := os.ReadFile(path)
				if err != nil {
					return err
				}

				em := NewEmail(path, string(data))

				if err := index.Add(em); err != nil {
					log.Fatal(err)
				}
				log.Println(fmt.Sprintf("%f%% done (%d files)", float64(i*100)/float64(517440), i))
				i++
				return nil
			})
		if err != nil {
			log.Fatal(err)
		}

		log.Println("done")
		return
	}

	if *searchTerm != "" {
		term := *searchTerm
		search(term)
		return
	}

	flag.PrintDefaults()
}
