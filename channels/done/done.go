package main

import (
	"fmt"
	"time"
)

/**
* - If we need just to send cancellation signal - it's good practice to use context
* - If there is a necessity to wait for a task to be shut down,
* it's better to use 2 channels
 */
func main() {
	done := make(chan struct{})
	go func() {
		time.Sleep(2 * time.Second)
		close(done)
	}()
	closed := work(done)
	<-closed
}

func work(signal chan struct{}) chan struct{} {
	ch := make(chan struct{})

	go func() {
		defer close(ch)

		for {
			select {
			case <-signal:
				fmt.Println("shutting down")
				return
			default:
				fmt.Println("sleep")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	return ch
}
