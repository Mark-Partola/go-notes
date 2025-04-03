package main

import (
	"fmt"
	"sync"
	"time"
)

type Barrier struct {
	mx    sync.Mutex
	n     int
	count int
	ch    chan struct{}
}

func NewBarrier(n int) *Barrier {
	return &Barrier{n: n, ch: make(chan struct{}, n)}
}

func (b *Barrier) Wait() {
	b.mx.Lock()

	b.count++
	if b.count == b.n {
		for range b.n {
			b.ch <- struct{}{}
		}
		b.count = 0
	}

	b.mx.Unlock()

	<-b.ch
}

func main() {
	workers := 3
	barrier := NewBarrier(workers)

	for range workers {
		go func() {
			for {
				barrier.Wait()
				bootstrap()
				barrier.Wait()
				work()
			}
		}()
	}

	select {}
}

func bootstrap() {
	time.Sleep(1 * time.Second)
	fmt.Println("bootstrap")
}

func work() {
	time.Sleep(2 * time.Second)
	fmt.Println("work")
}
