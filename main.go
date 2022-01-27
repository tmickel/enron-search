package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	index := NewIndex()

	emails := make([]*Email, 0)
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
			emails = append(emails, em)
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}

	for _, email := range emails {
		fields := strings.Fields(email.Body)
		for _, field := range fields {
			index.Add(field)
		}
	}
	log.Println(len(index.data))
	log.Println(index.data["app"])
}
