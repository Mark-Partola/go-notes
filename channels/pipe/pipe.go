package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := range 10 {
			ch <- i
		}
	}()

	process := pipe(
		filter(func(v int) bool {
			return v%2 == 0
		}),
		transform(func(v int) int {
			return v * v
		}),
	)

	for v := range process(ch) {
		fmt.Println(v)
	}
}

func pipe[T any](fns ...func(chan T) chan T) func(ch chan T) chan T {
	return func(ch chan T) chan T {
		out := make(chan T)

		go func() {
			defer close(out)

			for _, fn := range fns {
				ch = fn(ch)
			}

			for v := range ch {
				out <- v
			}
		}()

		return out
	}
}

func filter[T any](f func(T) bool) func(ch chan T) chan T {
	return func(ch chan T) chan T {
		out := make(chan T)

		go func() {
			defer close(out)
			for v := range ch {
				if f(v) {
					out <- v
				}
			}
		}()

		return out
	}
}

func transform[T any](f func(T) T) func(ch chan T) chan T {
	return func(ch chan T) chan T {
		out := make(chan T)

		go func() {
			defer close(out)
			for v := range ch {
				out <- f(v)
			}
		}()

		return out
	}
}
