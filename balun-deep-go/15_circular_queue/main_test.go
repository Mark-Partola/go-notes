package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue[int](queueSize)

	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsFull())

	v, _ := queue.GetFront()
	assert.Equal(t, 0, v)
	v, _ = queue.GetBack()
	assert.Equal(t, 0, v)
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.False(t, queue.IsEmpty())
	assert.True(t, queue.IsFull())

	v, _ = queue.GetFront()
	assert.Equal(t, 1, v)
	v, _ = queue.GetBack()
	assert.Equal(t, 3, v)

	assert.True(t, queue.Pop())
	assert.False(t, queue.IsEmpty())
	assert.False(t, queue.IsFull())
	assert.True(t, queue.Push(4))

	v, _ = queue.GetFront()
	assert.Equal(t, 2, v)
	v, _ = queue.GetBack()
	assert.Equal(t, 4, v)

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.IsEmpty())
	assert.False(t, queue.IsFull())
}
