package main

import (
	"fmt"
	"time"
)

func main() {
	f := Future(func() string {
		time.Sleep(time.Second)
		return "result"
	})
	r := f.Get()
	fmt.Println(r)
}

type future[T any] struct {
	ch chan T
}

func Future[T any](fn func() T) *future[T] {
	if fn == nil {
		panic("[future]: fn is nil")
	}

	future := &future[T]{
		ch: make(chan T),
	}

	go func() {
		future.ch <- fn()
	}()

	return future
}

func (future *future[T]) Get() T {
	return <-future.ch
}
