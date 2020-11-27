package main

import (
	"fmt"
	"index/suffixarray"
	"sort"
)

func main() {
	docs := []string{
		"hello world",
		"worldly goods",
		"yello",
		"lowly",
	}

	var data []byte
	var offsets []int

	for _, d := range docs {
		data = append(data, []byte(d)...)
		offsets = append(offsets, len(data))
	}
	sfx := suffixarray.New(data)

	query := "ello"

	idxs := sfx.Lookup([]byte(query), -1)
	var results []int
	for _, idx := range idxs {
		i := sort.Search(len(offsets), func(i int) bool { return offsets[i] > idx })
		if idx+len(query) <= offsets[i] {
			results = append(results, i)
		}
	}

	fmt.Printf("%q is in documents %v\n", query, results)
}