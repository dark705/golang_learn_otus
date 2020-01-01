package main

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

func main() {
	task := func() error {
		taskNum := rand.Uint32()
		time.Sleep(time.Millisecond * time.Duration(20+rand.Intn(50)))
		if rand.Intn(100) < 50 {
			errorMessage := "Error on execute Task#" + strconv.Itoa(int(taskNum))
			fmt.Println(errorMessage)
			return errors.New(errorMessage)
		}
		fmt.Printf("Task #%d done work\n", taskNum)
		return nil
	}

	numTasks := 10000
	tasks := make([]func() error, 0, numTasks)
	for i := 0; i < numTasks; i++ {
		tasks = append(tasks, task)
	}
	tStart := time.Now()
	Run(tasks, 1000, 10000)
	fmt.Println("num goroutines", runtime.NumGoroutine())
	tEnd := time.Now()
	fmt.Printf("Elapsed time: %v\n", tEnd.Sub(tStart))
}

func Run(tasks []func() error, N int, M int) error {
	if len(tasks) < N {
		N = len(tasks)
	}

	tasksCh := make(chan func() error, N)
	resCh := make(chan error, N)
	doneWorkGoroutineCh := make(chan struct{}, N)
	shutdown1 := make(chan string)

	//run task in N separate Goroutines
	for i := 1; i <= N; i++ {
		go func() {
			for task := range tasksCh {
				resCh <- task()
			}
			doneWorkGoroutineCh <- struct{}{}
		}()
	}

	//Send tasks to tasksCh and check results
	go func() {
		var suc int
		var err int
		var done bool
		var res string

		//Send first N tasks in to tasksCh
		for i := 0; i < N; i++ {
			tasksCh <- tasks[i]
		}

		addTaskIndex := N
		for resTask := range resCh {
			switch resTask.(type) {
			case error:
				err++
			default:
				suc++
			}
			if (err == M || suc+err == len(tasks)) && !done {
				done = true
				close(tasksCh)
				if err == M {
					res = fmt.Sprintln("Exit by error", "err", err, "suc", suc)
				} else {
					res = fmt.Sprintln("Exit by all done", "err", err, "suc", suc)
				}
			}
			if !done && addTaskIndex < len(tasks) {
				tasksCh <- tasks[addTaskIndex]
				addTaskIndex++
			}
		}
		shutdown1 <- res
	}()

	//Check all Goroutines done work
	go func() {
		var countDoneWorkGoroutines int
		for {
			<-doneWorkGoroutineCh
			countDoneWorkGoroutines++
			if countDoneWorkGoroutines == N {
				close(resCh)
				return
			}
		}
	}()

	fmt.Println(<-shutdown1)
	return nil
}
