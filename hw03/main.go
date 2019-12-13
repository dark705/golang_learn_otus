package main

import (
	"fmt"
	"io/ioutil"

	"github.com/dark705/otus/hw03/top"
)

func main() {
	content := getContentAsStringFromFile("file.txt")
	fmt.Println(top.Top10(content))
}

func getContentAsStringFromFile(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	return string(content)
}
