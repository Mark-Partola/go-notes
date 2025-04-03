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

	for v := range transform(ch, func(v int) int {
		return v * v
	}) {
		fmt.Println(v)
	}
}

func transform[T any](ch <-chan T, f func(T) T) chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for v := range ch {
			out <- f(v)
		}
	}()

	return out
}
