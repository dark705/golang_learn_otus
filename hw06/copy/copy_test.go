package copy

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
)

const (
	testFileIn  = "testIn.txt"
	testFileOut = "testOut.txt"
)

func TestCopyIdentical(t *testing.T) {
	defer func() {
		_ = os.Remove(testFileIn)
		_ = os.Remove(testFileOut)
	}()

	rnd := make([]byte, 10)
	rand.Read(rnd)

	err := ioutil.WriteFile(testFileIn, rnd, 0644)
	if err != nil {
		t.Error("Can't write file for test")
	}
	err = Copy(testFileIn, testFileOut, 0, 0)
	if err != nil {
		t.Error("Unexpected error on run Copy")
	}

	bIn, _ := ioutil.ReadFile(testFileIn)
	bOut, _ := ioutil.ReadFile(testFileOut)

	if bytes.Compare(bIn, bOut) != 0 {
		t.Error("Files not ident")
	}
}

func TestCopyLimit(t *testing.T) {
	defer func() {
		_ = os.Remove(testFileIn)
		_ = os.Remove(testFileOut)
	}()

	err := ioutil.WriteFile(testFileIn, []byte("123456"), 0644)
	if err != nil {
		t.Error("Can't write file for test")
	}
	err = Copy(testFileIn, testFileOut, 2, 0)
	if err != nil {
		t.Error("Unexpected error on run Copy")
	}

	bOut, _ := ioutil.ReadFile(testFileOut)
	if string(bOut) != "12" {
		t.Error("Error on limit copy")
	}
}

func TestCopyOffset(t *testing.T) {
	defer func() {
		_ = os.Remove(testFileIn)
		_ = os.Remove(testFileOut)
	}()

	err := ioutil.WriteFile(testFileIn, []byte("1234"), 0644)
	if err != nil {
		t.Error("Can't write file for test")
	}
	err = Copy(testFileIn, testFileOut, 3, 2)
	if err != nil {
		t.Error("Unexpected error on run Copy")
	}

	bOut, _ := ioutil.ReadFile(testFileOut)
	if string(bOut) != "34" {
		t.Error("Error on offset copy")
	}
}

func TestCopyOffsetAndLimit(t *testing.T) {
	defer func() {
		_ = os.Remove(testFileIn)
		_ = os.Remove(testFileOut)
	}()

	err := ioutil.WriteFile(testFileIn, []byte("123456"), 0644)
	if err != nil {
		t.Error("Can't write file for test")
	}
	err = Copy(testFileIn, testFileOut, 2, 2)
	if err != nil {
		t.Error("Unexpected error on run Copy")
	}

	bOut, _ := ioutil.ReadFile(testFileOut)
	if string(bOut) != "34" {
		t.Error("Error on offset and limit copy")
	}
}

func TestCopyOffsetIsBiggerThenFile(t *testing.T) {
	defer func() {
		_ = os.Remove(testFileIn)
	}()

	err := ioutil.WriteFile(testFileIn, []byte("123456"), 0644)
	if err != nil {
		t.Error("Can't write file for test")
	}
	err = Copy(testFileIn, testFileOut, 2, 100000)
	if err == nil {
		t.Error("Copy not return error for offset is bigger then input file len")
	}
}
