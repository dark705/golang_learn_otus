package top

import (
	"sort"
	"strings"
)

func Top10(content string) []string {
	type freq struct {
		word  string
		count int
	}

	allWords := strings.Split(content, " ")

	var countWord = make(map[string]int)
	for _, word := range allWords {
		lowerWord := strings.ToLower(word)
		countWord[lowerWord]++
	}

	collection := make([]freq, 0, len(countWord))
	for word, count := range countWord {
		collection = append(collection, freq{word, count})
	}

	sort.Slice(collection, func(i, j int) bool {
		return collection[i].count > collection[j].count
	})

	place := 0
	top10 := make([]string, 0, 10)
	for place < 10 && place < len(collection) {
		top10 = append(top10, collection[place].word)
		place++
	}

	return top10
}
