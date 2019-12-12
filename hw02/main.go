package main

import (
	"fmt"
)

const (
	intStart = 48 //int 0
	intEnd   = 57 //int 9
)

func isIntRune(r rune) bool {
	return (intStart <= int(r)) && (int(r) <= intEnd)
}

func getIntFromRune(r rune) int {
	if isIntRune(r) {
		return int(r) - intStart
	} else {
		return 0
	}
}

func repeatRuneToString(r rune, i int) (out string) {
	if r == 0 {
		return ""
	}

	for i > 1 {
		out += string(r)
		i--
	}
	return
}

func Unpack(in string) (out string) {
	var prevElement rune
	var doubleSlash bool

	for _, element := range in {
		if element == '\\' && prevElement == '\\' {
			doubleSlash = true
		}

		if isIntRune(element) && doubleSlash {
			needRepeatCount := getIntFromRune(element)
			out += repeatRuneToString(prevElement, needRepeatCount+1)
			doubleSlash = false
			continue
		}

		if isIntRune(element) && prevElement != '\\' {
			needRepeatCount := getIntFromRune(element)
			out += repeatRuneToString(prevElement, needRepeatCount)
			continue
		}

		prevElement = element

		if element != '\\' {
			out += string(element)
		}
	}

	return
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
