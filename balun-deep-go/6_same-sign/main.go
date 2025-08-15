package main

import "fmt"

func main() {
	x, y := 29, 30
	fmt.Printf("x = %d, y = %d, same sign: %v\n", x, y, FastSameSign(x, y))

	x, y = 29, -30
	fmt.Printf("x = %d, y = %d, same sign: %t\n", x, y, FastSameSign(x, y))
}

func SlowSameSign(x, y int) bool {
	return x > 0 && y > 0 || x < 0 && y < 0
}

func FastSameSign(x, y int) bool {
	return x^y >= 0
}
