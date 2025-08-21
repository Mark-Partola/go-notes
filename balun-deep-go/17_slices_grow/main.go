package main

import "fmt"

func main() {
	a := make([]int, 0)
	a = append(a, []int{1, 2, 3, 4, 5}...)
	fmt.Println(len(a), cap(a))

	b := make([]uint8, 0)
	b = append(b, []uint8{1, 2, 3, 4, 5}...)
	fmt.Println(len(b), cap(b))

	type A struct {
		_pad [14]byte
	}

	c := make([]A, 0)
	c = append(c, []A{{}, {}, {}, {}, {}}...)
	fmt.Println(len(c), cap(c))
}
