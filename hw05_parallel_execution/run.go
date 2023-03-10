package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	// Place your code here.
	var err error
	var errCounter int

	ch := make(chan Task)
	mx := sync.Mutex{}
	wg := sync.WaitGroup{}

	for x := 0; x < n; x++ {
		wg.Add(1)
		go func(chT <-chan Task, wg *sync.WaitGroup, mtx *sync.Mutex, errCounter *int) {
			for {
				tsk, ok := <-chT
				if !ok {
					break
				}
				if tsk() != nil {
					mtx.Lock()
					(*errCounter)++
					mtx.Unlock()
				}
			}
			wg.Done()
		}(ch, &wg, &mx, &errCounter)
	}
	for _, tsk := range tasks {
		ch <- tsk
		if m > 0 {
			mx.Lock()
			if errCounter >= m {
				err = ErrErrorsLimitExceeded
			}
			mx.Unlock()
			if err != nil {
				break
			}
		}
	}
	close(ch)
	wg.Wait()
	return err
}
