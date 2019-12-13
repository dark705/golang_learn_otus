package main

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	intStart = 48 //int 0
)

func getIntFromRune(r rune) int {
	if unicode.IsDigit(r) {
		return int(r) - intStart
	}
	return 0
}

func Unpack(in string) (out string) {
	var prevElement rune
	var doubleSlash bool

	for _, element := range in {
		if element == '\\' && prevElement == '\\' {
			doubleSlash = true
		}

		if unicode.IsDigit(element) && doubleSlash {
			needRepeatCount := getIntFromRune(element)
			out += strings.Repeat(string(prevElement), needRepeatCount)
			doubleSlash = false
			continue
		}

		if unicode.IsDigit(element) && prevElement != '\\' {
			needRepeatCount := getIntFromRune(element)
			if prevElement == 0 {
				return ""
			}
			out += strings.Repeat(string(prevElement), needRepeatCount-1)
			continue
		}

		prevElement = element

		if element != '\\' {
			out += string(element)
		}
	}

	return out
}

func main() {
	type pair struct {
		in  string
		out string
	}

	test := []pair{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{`qwe\4\5`, `qwe45`},
		{`qwe\45`, `qwe44444`},
		{`qwe\\5`, `qwe\\\\\`},
	}

	for _, t := range test {
		if t.out == Unpack(t.in) {
			fmt.Printf("%s - %s\n", t.in, "OK")
		} else {
			fmt.Printf("%s - %s\n", t.in, "FAIL")
		}
	}

}
