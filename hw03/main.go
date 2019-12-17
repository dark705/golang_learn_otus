package main

import (
	"fmt"
	"io/ioutil"
	"main/top"
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
