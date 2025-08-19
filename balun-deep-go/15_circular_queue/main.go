package main

import "sync"

type Queue interface {
	Push(value int) bool
	Pop() bool
	GetFront() int
	GetBack() int
	IsEmpty() bool
	IsFull() bool
}

type CircularQueue struct {
	mu    sync.RWMutex
	start int
	end   int
	size  int
	count int
	queue []int
}

func (c *CircularQueue) Push(value int) bool {
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

func (c *CircularQueue) Pop() bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.isEmptyLocked() {
		return false
	}

	c.start = (c.start + 1) % c.size
	c.count--

	return true
}

func (c *CircularQueue) GetFront() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if !c.isEmptyLocked() {
		return c.queue[c.start]
	}

	return -1
}

func (c *CircularQueue) GetBack() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if !c.isEmptyLocked() {
		end := (c.end - 1 + c.size) % c.size
		return c.queue[end]
	}

	return -1
}

func (c *CircularQueue) IsEmpty() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.isEmptyLocked()
}

func (c *CircularQueue) IsFull() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.isFullLocked()
}

func (c *CircularQueue) isFullLocked() bool {
	return c.count == c.size
}

func (c *CircularQueue) isEmptyLocked() bool {
	return c.count == 0
}

func NewCircularQueue(size int) Queue {
	return &CircularQueue{
		start: 0,
		end:   0,
		size:  size,
		count: 0,
		queue: make([]int, size),
	}
}
