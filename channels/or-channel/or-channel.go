package main

import (
	"fmt"
	"time"
)

func main() {
	after := func(after time.Duration) chan struct{} {
		ch := make(chan struct{})
		go func() {
			defer close(ch)
			time.Sleep(after)
		}()
		return ch
	}

	start := time.Now()

	<-or(
		after(2*time.Hour),
		after(2*time.Second),
		after(10*time.Second),
		after(500*time.Millisecond),
	)

	fmt.Printf("called after: %v\n", time.Since(start))
}

func or[T any](in ...chan T) chan struct{} {
	if len(in) == 0 {
		// <-or([]) == nil
		// <- nil == blocking
		return nil
	}

	out := make(chan struct{})

	go func() {
		select {
		case <-in[0]:
			return
		case <-or(in[1:]...):
			return
		}
	}()

	return out
}

/**

V. Balun's variant

func or[T any](in ...chan T) chan T {
	switch len(in) {
	case 0:
		return nil
	case 1:
		return in[0]
	}

	out := make(chan T)

	go func() {
		defer close(out)

		switch len(in) {
		case 2:
			select {
			case <-in[0]:
			case <-in[1]:
			}
		default:
			select {
			case <-in[0]:
			case <-in[1]:
			case <-in[2]:
			case <-or(append([]chan T{out}, in[3:]...)...):
			}
		}
	}()

	return out
}
*/
