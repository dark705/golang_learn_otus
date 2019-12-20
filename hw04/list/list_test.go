package list

import (
	"testing"
)

func TestLink(t *testing.T) {
	l := List{}
	item1 := l.PushBack("s1")
	l.PushBack("s2")
	l.PushFront("s0")

	if l.First().Next() != item1 || l.Last().Prev() != item1 {
		t.Error("Invalid link at item1")
	}

	l.Remove(item1)
	if l.First().Next() != l.Last() || l.Last().Prev() != l.First() {
		t.Error("Invalid link after remove item1")
	}
}

func TestLen(t *testing.T) {
	l := List{}
	l.PushBack("s1")
	l.PushBack(2)
	if l.Len() != 2 {
		t.Error("Invalid len after 3 PushBack")
	}

	l.PushFront("s0")
	item5 := l.PushFront("s-1")
	if l.Len() != 4 {
		t.Error("Invalid len after 1 PushFront")
	}

	l.Remove(item5)
	if l.Len() != 3 {
		t.Error("Invalid len after 1 Remove")
	}
}

func TestValue(t *testing.T) {
	l := List{}
	l.PushBack("s1")
	item2 := l.PushBack("s2")
	l.PushFront("s0")

	if l.First().Value() != "s0" {
		t.Error("Invalid first value")
	}

	if l.Last().Value() != "s2" {
		t.Error("Invalid last value")
	}

	if l.First().Next().Next().Value() != "s2" {
		t.Error("Invalid value at move Next")
	}

	if l.Last().Prev().Prev().Value() != "s0" {
		t.Error("Invalid value at move Prev")
	}

	l.Remove(item2)
	if l.Last().Prev().Value() != "s0" || l.First().Next().Value() != "s1" {
		t.Error("Invalid value after 1 remove")
	}
}
