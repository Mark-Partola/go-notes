package main

import (
	"fmt"
	"time"
)

func main() {
	w := NewWorker()

	w.Do(func() {
		fmt.Println("working")
		time.Sleep(2 * time.Second)
		fmt.Println("work done")
	})

	time.Sleep(time.Second * 5)
	w.Shutdown()
}

type worker struct {
	closeCh     chan struct{}
	closeDoneCh chan struct{}
}

func NewWorker() *worker {
	return &worker{
		closeCh:     make(chan struct{}),
		closeDoneCh: make(chan struct{}),
	}
}

func (w *worker) Shutdown() {
	close(w.closeCh)
	<-w.closeDoneCh
}

func (w *worker) Do(fn func()) {
	go func() {
		defer close(w.closeDoneCh)

		for {
			// shutdown priority
			select {
			case <-w.closeCh:
				fmt.Println("shutting down")
				return
			default:
			}

			select {
			case <-w.closeCh:
				fmt.Println("shutting down")
				return
			default:
				fn()
			}
		}
	}()
}
