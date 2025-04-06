package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := range 10 {
			ch <- i
		}
	}()

	for v := range Filter(ch, func(v int) bool {
		return v%2 == 0
	}) {
		fmt.Println(v)
	}
}

func Filter(ch <-chan int, f func(int) bool) <-chan int {
	out := make(chan int)

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
