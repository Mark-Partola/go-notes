package main

import (
	"slices"
	"testing"
)

func TestTransformer(t *testing.T) {
	ch := generate(10)
	expect := []int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81}
	result := collect(transform(ch, func(v int) int {
		return v * v
	}))

	if !slices.Equal(result, expect) {
		t.Errorf("expected %v, got %v", expect, result)
	}
}

func TestTransformerEmpty(t *testing.T) {
	ch := generate(0)
	expect := []int{}
	result := collect(transform(ch, func(v int) int {
		return v * v
	}))

	if !slices.Equal(result, expect) {
		t.Errorf("expected %v, got %v", expect, result)
	}
}

func generate(count int) <-chan int {
	ch := make(chan int, count)

	go func() {
		defer close(ch)
		for i := range count {
			ch <- i
		}
	}()

	return ch
}

func collect(ch chan int) []int {
	res := make([]int, 0)
	for v := range ch {
		res = append(res, v)
	}
	return res
}
