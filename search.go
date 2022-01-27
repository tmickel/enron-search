package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func search(term string) {
	start := time.Now()
	fmt.Println("searching for", term)

	if _, err := os.Stat("./index"); os.IsNotExist(err) {
		fmt.Println("index not found. please run ./enron-search -genindex first")
		return
	}

	normalizedTerm := normalizeWord(term)

	path := "./index"
	for _, char := range normalizedTerm {
		path = filepath.Join(path, string(char))
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("no results :( try shortening?")
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	i := 1
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		data, err := os.ReadFile(filepath.Join(path, file.Name()))
		if err != nil {
			log.Fatal(err)
		}

		// alt: just collect into a slice and just return them.
		em := NewEmail(path, string(data))
		fmt.Printf("------- result %d -------", i)
		fmt.Println()
		fmt.Println("from:", em.From)
		fmt.Println("to:", em.To)
		fmt.Println("subject:", em.Subject)
		fmt.Println("body:", em.Body)
		fmt.Println("------------------------")
		fmt.Println()
		i++
	}
	fmt.Printf("found %d emails in %v", i-1, time.Since(start))

}
