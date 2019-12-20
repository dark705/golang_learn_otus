package main

import "fmt"
import "main/list"

func main() {
	l := list.List{}
	l.PushBack("s1")
	someItem := l.PushBack(2)
	l.PushBack([]int{3})
	l.PushFront("s0")

	for item := l.Last(); item != nil; item = item.Prev() {
		fmt.Println(item.Value())
	}

	for item := l.First(); item != nil; item = item.Next() {
		fmt.Println(item.Value())
	}

	l.Remove(someItem)
	for item := l.First(); item != nil; item = item.Next() {
		fmt.Println(item.Value())
	}
}
