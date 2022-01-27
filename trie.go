package main

import (
	"log"
	"regexp"
	"strings"
)

type TrieNode struct {
	Children [26]*TrieNode
	FileIds  []int
}

func NewTrie() *TrieNode {
	return &TrieNode{}
}

func normalizeWord(key string) string {
	key = strings.ToLower(key)
	re, err := regexp.Compile("[^a-z]+")
	if err != nil {
		log.Fatal(err)
	}
	key = re.ReplaceAllString(key, "")
	return key
}

func (t *TrieNode) Insert(word string, fileId int) {
	word = normalizeWord(word)
	node := t
	for _, char := range word {
		childIndex := char - 'a'
		if node.Children[childIndex] == nil {
			node.Children[childIndex] = &TrieNode{}
		}
		if node.FileIds == nil {
			node.FileIds = make([]int, 0)
		}
		if len(node.FileIds) <= 50 {
			node.FileIds = append(node.FileIds, fileId)
		}
		node = node.Children[childIndex]
	}
	if node.FileIds == nil {
		node.FileIds = make([]int, 0)
	}
	node.FileIds = append(node.FileIds, fileId)
}

func (t *TrieNode) Search(word string) ([]int, bool) {
	word = normalizeWord(word)
	node := t
	for _, char := range word {
		childIndex := char - 'a'
		if node.Children[childIndex] == nil {
			return nil, false
		}
		node = node.Children[childIndex]
	}
	return node.FileIds, true
}
