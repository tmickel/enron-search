package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Index struct {
	symlinkCounts map[string]int
}

func NewIndex() (*Index, error) {
	os.RemoveAll("./index") // xxx: maybe don't remove index every time
	if err := os.Mkdir("./index", 0777); err != nil {
		return nil, err
	}
	return &Index{
		symlinkCounts: make(map[string]int),
	}, nil
}

func (ix *Index) Add(em *Email) error {
	fields := strings.Fields(em.Body)
	for _, field := range fields {
		if err := ix.insertWord(field, em.Filename); err != nil {
			return err
		}
	}

	subjectFields := strings.Fields(em.Subject)
	for _, field := range subjectFields {
		if err := ix.insertWord(field, em.Filename); err != nil {
			return err
		}
	}

	toFields := strings.Fields(em.To)
	for _, field := range toFields {
		if err := ix.insertWord(field, em.Filename); err != nil {
			return err
		}
	}

	fromFields := strings.Fields(em.From)
	for _, field := range fromFields {
		if err := ix.insertWord(field, em.Filename); err != nil {
			return err
		}
	}
	return nil
}

func (ix *Index) insertWord(word string, path string) error {
	word = normalizeWord(word)
	if word == "" {
		return nil
	}
	dstPath := "./index"
	for _, char := range word {
		dstPath = filepath.Join(dstPath, string(char))
		if _, err := os.Stat(dstPath); os.IsNotExist(err) {
			if err := os.MkdirAll(dstPath, 0777); err != nil {
				return err
			}
		}

		count, ok := ix.symlinkCounts[dstPath]
		if !ok {
			ix.symlinkCounts[dstPath] = 1
			count = 1
		} else {
			count++
			ix.symlinkCounts[dstPath]++
		}

		// Only index 50 emails per query
		// I picked this randomly, assuming it would be pretty decent UX.
		// could be adjusted upward or removed depending on appetite for symlinks.
		if count > 50 {
			return nil
		}

		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		src := filepath.Join(wd, path)
		if err := os.Symlink(src, filepath.Join(dstPath, fmt.Sprint(count))); err != nil {
			return err
		}
	}
	return nil
}

func normalizeWord(key string) string {
	// Only index lowercase a-z words
	// This is not really a requirement for this solution, maybe should be deleted.
	key = strings.ToLower(key)
	re, err := regexp.Compile("[^a-z]+")
	if err != nil {
		log.Fatal(err)
	}
	key = re.ReplaceAllString(key, "")
	return key
}
