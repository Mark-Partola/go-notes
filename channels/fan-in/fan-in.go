package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	create := func() <-chan int {
		ch := make(chan int)

		go func() {
			defer close(ch)
			for i := range 10 {
				ch <- i
			}
		}()

		return ch
	}

	out := fanIn(create(), create(), create())

	for v := range out {
		fmt.Println(v)
	}

	fmt.Print("> Press any key ")
	bufio.NewScanner(os.Stdin).Scan()
}

func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(channels))

	for _, ch := range channels {
		go func() {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
