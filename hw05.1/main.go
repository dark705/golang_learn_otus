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
	wgWorker := sync.WaitGroup{}
	wgProducer := sync.WaitGroup{}

	//run task in N separate Goroutines
	wgWorker.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wgWorker.Done()
			for task := range tasksCh {
				resCh <- task()
			}
		}()
	}

	//Send tasks to tasksCh and check results
	wgProducer.Add(1)
	var res error
	go func() {
		var suc, err int
		var done bool

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
		wgProducer.Done()
	}()

	//Check all Goroutines done work
	wgWorker.Wait()
	close(resCh)

	//Wait for producer finish, and result ready
	wgProducer.Wait()
	return res
}
