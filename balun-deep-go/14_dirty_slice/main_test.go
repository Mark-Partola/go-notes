package main

import (
	"strings"
	"testing"
	"unsafe"
)

// make with cap set all the elements to 0
// but if we will rewrite them, we can skip the redundant work
func makeDirty(size int) []byte {
	var sb strings.Builder
	sb.Grow(size)

	pointer := unsafe.StringData(sb.String())
	return unsafe.Slice(pointer, size)
}

var Result []byte

func BenchmarkMake(b *testing.B) {
	for b.Loop() {
		Result = make([]byte, 0, 10<<20)
	}
}

func BenchmarkMakeDirty(b *testing.B) {
	for b.Loop() {
		Result = makeDirty(10 << 20)
	}
}
