package main

import "strings"

type Index struct {
	fileIds []string
	trie    *TrieNode
}

func NewIndex() *Index {
	return &Index{
		fileIds: make([]string, 0),
		trie:    NewTrie(),
	}
}

func (ix *Index) Add(em *Email) {
	ix.fileIds = append(ix.fileIds, em.Filename)
	fileId := len(ix.fileIds) - 1

	fields := strings.Fields(em.Body)
	for _, field := range fields {
		ix.trie.Insert(field, fileId)
	}

	subjectFields := strings.Fields(em.Subject)
	for _, field := range subjectFields {
		ix.trie.Insert(field, fileId)
	}

	toFields := strings.Fields(em.To)
	for _, field := range toFields {
		ix.trie.Insert(field, fileId)
	}

	fromFields := strings.Fields(em.From)
	for _, field := range fromFields {
		ix.trie.Insert(field, fileId)
	}
}
