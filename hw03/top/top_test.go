package top

import (
	"testing"
)

func TestTop10(t *testing.T) {

	type pair struct {
		content string
		top     []string
	}

	var test = []pair{
		{"два слова слова", []string{"слова", "два"}},
		{"много много слов", []string{"много", "слов"}},
		{"очень очень очень много много слов", []string{"очень", "много", "слов"}},
		{"Яблоко яблоко на на на", []string{"на", "яблоко"}},
	}

	for _, c := range test {
		result := Top10(c.content)
		expected := c.top

		for i := range result {
			if result[i] != expected[i] {
				t.Error("Test fail for string:", c.content, "\n result:", result, "\n expected:", expected)
			}
		}
	}
}
