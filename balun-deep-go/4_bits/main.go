package main

import (
	"fmt"
)

func main() {
	// inversion for signed and unsigned values
	var x uint8 = 3
	var y int8 = 3
	fmt.Println(^x, ^y)

	// check is square of 2
	for i := range 9 {
		fmt.Println(i, IsSquare2Short(i))
	}
}

func IsSquare2(n int) bool {
	var result = false
	for n > 0 {
		if n&1 == 1 {
			if result {
				return false
			}
			result = true
		}
		n >>= 1
	}
	return result
}

/**
 * 010000 n
 * &
 * 001111 n-1
 * =
 * 000000 0
 */
func IsSquare2Short(n int) bool {
	return n&(n-1) == 0
}
