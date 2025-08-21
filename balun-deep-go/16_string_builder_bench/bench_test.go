package main

import (
	"strings"
	"testing"
)

func BenchmarkSimpleConcatenation(b *testing.B) {
	str := "test"
	for b.Loop() {
		str += "test"
	}

	_ = str
}

func BenchmarkConcatenationWithStringBuilder(b *testing.B) {
	builder := strings.Builder{}
	builder.WriteString("test")
	for b.Loop() {
		builder.WriteString("test")
	}

	_ = builder.String()
}

func BenchmarkConcatenationWithStringBuilderOptimized(b *testing.B) {
	builder := strings.Builder{}
	builder.Grow(4 + b.N*4)
	builder.WriteString("test")
	for b.Loop() {
		builder.WriteString("test")
	}

	_ = builder.String()
}
