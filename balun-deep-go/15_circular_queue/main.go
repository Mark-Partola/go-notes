package main

import "sync"

type Queue[T comparable] interface {
	Push(value T) bool
	Pop() bool
	GetFront() (T, bool)
	GetBack() (T, bool)
	IsEmpty() bool
	IsFull() bool
}

type CircularQueue[T comparable] struct {
	mu    sync.RWMutex
	start int
	end   int
	size  int
	count int
	queue []T
}

func (c *CircularQueue[T]) Push(value T) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.isFullLocked() {
		return false
	}

	c.queue[c.end] = value
	c.end = (c.end + 1) % c.size
	c.count++

	return true
}

func (c *CircularQueue[T]) Pop() bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.isEmptyLocked() {
		return false
	}

	c.start = (c.start + 1) % c.size
	c.count--

	return true
}

func (c *CircularQueue[T]) GetFront() (T, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if !c.isEmptyLocked() {
		return c.queue[c.start], true
	}

	var zero T
	return zero, false
}

func (c *CircularQueue[T]) GetBack() (T, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if !c.isEmptyLocked() {
		end := (c.end - 1 + c.size) % c.size
		return c.queue[end], true
	}

	var zero T
	return zero, false
}

func (c *CircularQueue[T]) IsEmpty() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.isEmptyLocked()
}

func (c *CircularQueue[T]) IsFull() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.isFullLocked()
}

func (c *CircularQueue[T]) isFullLocked() bool {
	return c.count == c.size
}

func (c *CircularQueue[T]) isEmptyLocked() bool {
	return c.count == 0
}

func NewCircularQueue[T comparable](size int) Queue[T] {
	return &CircularQueue[T]{
		start: 0,
		end:   0,
		count: 0,
		size:  size,
		queue: make([]T, size),
	}
}
