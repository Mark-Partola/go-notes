package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		ticker := time.NewTicker(1 * time.Second)

		for v := range ticker.C {
			ch <- v.Second()
		}
	}()

	done := time.After(5 * time.Second)
	for v := range OrDone(done, ch) {
		fmt.Println(v)
	}
}

func OrDone[D any, T any](done <-chan D, ch chan T) chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			default:
			}

			select {
			case <-done:
				return
			case v, ok := <-ch:
				if !ok {
					return
				}

				out <- v
			}
		}
	}()

	return out
}
