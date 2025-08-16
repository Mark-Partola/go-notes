package main

import "testing"

type Point struct {
	X, Y int
}

// because of level of indirection it will be slower
func BenchmarkWithPointers(b *testing.B) {
	points := [...]*Point{
		{0, 0}, {1, 1}, {2, 2}, {3, 3},
		{4, 4}, {5, 5}, {6, 6}, {7, 7},
	}

	for b.Loop() {
		for _, point := range points {
			point.X += 1
		}
	}
}

func BenchmarkWithIndices(b *testing.B) {
	points := [...]Point{
		{0, 0}, {1, 1}, {2, 2}, {3, 3},
		{4, 4}, {5, 5}, {6, 6}, {7, 7},
	}

	for b.Loop() {
		for idx := range points {
			points[idx].X += 1
		}
	}
}
