package main

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestRunAllGoroutinesFinishedByDoneAllTasks(t *testing.T) {
	testCh := make(chan string, 1000)
	tasks := getTasks(1000, testCh)

	_ = Run(tasks, 100, 10)
	n := runtime.NumGoroutine() - 2 //-1 main + -1 test itself
	if n != 0 {
		t.Error("Not all goroutines finished after Run() done by M errors:", n)
	}
}

func TestRunAllGoroutinesFinishedByMErrors(t *testing.T) {
	testCh := make(chan string, 1000)
	tasks := getTasks(1000, testCh)

	_ = Run(tasks, 100, 10)
	n := runtime.NumGoroutine() - 2 //-1 main + -1 test itself
	if n != 0 {
		t.Error("Not all goroutines finished after Run() done by M errors:", n)
	}
}

func TestRunRes(t *testing.T) {
	testCh := make(chan string, 1000)
	tasks := getTasks(1000, testCh)
	_ = Run(tasks, 100, 10000)
}

func getTask(testCh chan string) func() error {
	return func() error {
		taskNum := rand.Uint32()
		time.Sleep(time.Millisecond * time.Duration(20+rand.Intn(50)))
		if rand.Intn(100) < 50 {
			errorMessage := "Error on execute Task#" + strconv.Itoa(int(taskNum))
			testCh <- fmt.Sprintf(errorMessage)
			return errors.New(errorMessage)
		}
		testCh <- fmt.Sprintf("Task #%d done work\n", taskNum)
		return nil
	}
}

func getTasks(num int, testCh chan string) []func() error {
	tasks := make([]func() error, 0, num)
	for i := 0; i < num; i++ {
		tasks = append(tasks, getTask(testCh))
	}
	return tasks
}
