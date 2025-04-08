package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	resultCh := make(chan struct{})

	wg.Add(10)
	for idx := range 10 {
		go func() {
			defer wg.Done()
			request(ctx, resultCh, idx)
		}()
	}

	<-resultCh
	cancel()

	wg.Wait()
}

func request(ctx context.Context, result chan<- struct{}, idx int) {
	timeout := time.Duration(rand.Intn(5000)) * time.Millisecond
	timer := time.NewTimer(timeout)

	select {
	case <-timer.C:
		result <- struct{}{}
		fmt.Printf("finished %d\n", idx)
	case <-ctx.Done():
		fmt.Printf("cancelled %d\n", idx)
	}
}
