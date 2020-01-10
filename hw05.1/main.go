package main

import (
	"errors"
	"fmt"
	"sync"
)

func Run(tasks []func() error, N int, M int) error {
	if len(tasks) < N {
		N = len(tasks)
	}

	tasksCh := make(chan func() error, N)
	resCh := make(chan error, N)
	returnCh := make(chan error)
	wg := sync.WaitGroup{}

	//run task in N separate Goroutines
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			for task := range tasksCh {
				resCh <- task()
			}
		}()
	}

	//Send tasks to tasksCh and check results
	go func() {
		var suc, err int
		var done bool
		var res error

		//Send first N tasks in to tasksCh
		for i := 0; i < N; i++ {
			tasksCh <- tasks[i]
		}

		addTaskIndex := N //index of task which will be send next
		for resTask := range resCh {
			if resTask != nil {
				err++
			} else {
				suc++
			}
			if done {
				continue
			}
			if err > 0 && err == M || suc+err == len(tasks) {
				done = true
				close(tasksCh)
				if err > 0 && err == M {
					message := fmt.Sprintln("Exit by N errors limit", "err", err, "suc", suc)
					res = errors.New(message)
				} else {
					res = nil
				}
			}
			if !done && addTaskIndex < len(tasks) {
				//send Additional task for run, after some task done
				tasksCh <- tasks[addTaskIndex]
				addTaskIndex++
			}
		}
		returnCh <- res
	}()

	//Check all Goroutines done work
	wg.Wait()
	close(resCh)

	return <-returnCh
}
