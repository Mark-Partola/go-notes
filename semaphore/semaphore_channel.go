package main

import (
	"fmt"
	"sync"
	"time"
)

type ChannelSemaphore struct {
	ch chan struct{}
}

func NewChannelSemaphore(n int32) Semaphore {
	return &ChannelSemaphore{
		ch: make(chan struct{}, n),
	}
}

func (s *ChannelSemaphore) Acquire() {
	s.ch <- struct{}{}
}

func (s *ChannelSemaphore) Release() {
	<-s.ch
}

func main() {
	wg := sync.WaitGroup{}
	s := NewChannelSemaphore(3)
	wg.Add(10)

	worker := func() {
		s.Acquire()
		defer s.Release()
		defer wg.Done()

		fmt.Println("start")
		time.Sleep(time.Second)
		fmt.Println("end")
	}

	for range 10 {
		go worker()
	}

	wg.Wait()
}
