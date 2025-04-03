package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := range 20 {
			ch <- i
		}
	}()

	out := fanOut(ch, 3)

	wg := sync.WaitGroup{}
	wg.Add(len(out))
	for idx, ch := range out {
		go func() {
			defer wg.Done()
			printCh(idx, ch)
		}()
	}

	wg.Wait()
	fmt.Println("> Press any key ")
	bufio.NewScanner(os.Stdin).Scan()
}

func printCh[T any](idx int, ch <-chan T) {
	for v := range ch {
		fmt.Printf("[%d]: %v\n", idx, v)
	}
}

func fanOut[T any](ch <-chan T, n int) []chan T {
	out := make([]chan T, n)
	for i := range out {
		out[i] = make(chan T)
	}

	go func() {
		idx := 0
		for v := range ch {
			out[idx] <- v
			idx = (idx + 1) % n

			/** non blocking variant
			func() {
				for {
					idx = (idx + 1) % n
					select {
					case out[idx] <- v:
						return
					default:
					}
				}
			}()
			*/
		}

		for _, ch := range out {
			close(ch)
		}
	}()

	return out
}
