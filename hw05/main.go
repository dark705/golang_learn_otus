package main

import (
	"errors"
	"fmt"
)

func Run(tasks []func() error, N int, M int) error {
	if len(tasks) < N {
		N = len(tasks)
	}

	tasksCh := make(chan func() error, N)
	resCh := make(chan error, N)
	doneWorkGoroutineCh := make(chan struct{}, N)
	returnCh := make(chan error)

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
		var suc, err int
		var done bool
		var res error

		//Send first N tasks in to tasksCh
		for i := 0; i < N; i++ {
			tasksCh <- tasks[i]
		}

		addTaskIndex := N //index of task which will be send next
		for resTask := range resCh {
			switch resTask.(type) {
			case error:
				err++
			default:
				suc++
			}
			if (err > 0 && err == M || suc+err == len(tasks)) && !done {
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

	return <-returnCh
}
