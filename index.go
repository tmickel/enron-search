package main

type Index struct {
	data map[string]int
}

func NewIndex() *Index {
	return &Index{
		data: make(map[string]int, 0),
	}
}

func (ix *Index) Add(word string) {
	for i := 1; i < len(word); i++ {
		prefix := word[0:i]
		if _, ok := ix.data[prefix]; ok {
			ix.data[prefix]++
		} else {
			ix.data[prefix] = 1
		}
	}
}
