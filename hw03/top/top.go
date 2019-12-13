package top

import (
	"sort"
	"strings"
)

func Top10(content string) (top10 []string) {
	type words struct {
		word  string
		count int
	}

	allWords := strings.Split(content, " ")

	var countWord = make(map[string]int)
	for _, word := range allWords {
		lowerWord := strings.ToLower(word)
		countWord[lowerWord]++
	}

	var collection []words
	for word, count := range countWord {
		collection = append(collection, words{word, count})
	}

	sort.Slice(collection, func(i, j int) bool {
		return collection[i].count > collection[j].count
	})

	place := 0
	for place < 10 && place < len(collection) {
		top10 = append(top10, collection[place].word)
		//fmt.Println(collection[place].word, collection[place].count)
		place++
	}

	return top10
}
