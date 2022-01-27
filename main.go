package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	log.Println("starting index")
	index := NewIndex()

	err := filepath.Walk("./raw-data",
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

			index.Add(em)
			log.Println(len(index.fileIds))
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("done")
}
