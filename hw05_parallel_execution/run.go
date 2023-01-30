package hw05parallelexecution

import (
	"errors"
	"sync"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrNoASingleGoroutine  = errors.New("the number of goroutines to execute tasks <= 0")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) (err error) {
	if n <= 0 {
		return ErrNoASingleGoroutine
	}
	taskChan := make(chan Task)
	errChan := make(chan error)
	quitChan := make(chan struct{})
	wg := sync.WaitGroup{}
	errCount := 0

	go func() {
		defer close(errChan)
		for i := 0; i < n; i++ {
			wg.Add(1)
			go run(taskChan, errChan, quitChan, &wg)
		}
		wg.Wait()
	}()

	go func() {
		defer close(taskChan)
		for i := range tasks {
			select {
			case <-quitChan:
				return
			default:
				taskChan <- tasks[i]
			}
		}
	}()

	for e := range errChan {
		_ = e
		errCount++
		if errCount == m {
			close(quitChan)
			err = ErrErrorsLimitExceeded
		}
	}

	return
}

func run(taskChan <-chan Task, errCh chan<- error, quitChan <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for t := range taskChan {
		select {
		case <-quitChan:
			return
		default:
			if err := t(); err != nil {
				errCh <- err
			}
		}
	}
}
