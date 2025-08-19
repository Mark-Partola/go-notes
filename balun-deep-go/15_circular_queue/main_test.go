package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue(queueSize)

	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsFull())

	assert.Equal(t, -1, queue.GetFront())
	assert.Equal(t, -1, queue.GetBack())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.False(t, queue.IsEmpty())
	assert.True(t, queue.IsFull())

	assert.Equal(t, 1, queue.GetFront())
	assert.Equal(t, 3, queue.GetBack())

	assert.True(t, queue.Pop())
	assert.False(t, queue.IsEmpty())
	assert.False(t, queue.IsFull())
	assert.True(t, queue.Push(4))

	assert.Equal(t, 2, queue.GetFront())
	assert.Equal(t, 4, queue.GetBack())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsFull())
}
