package main

import (
	"fmt"
)

func main() {
	var value uint32 = 0x12345678

	fmt.Printf("%032b\n", value)
	fmt.Printf("%032b\n", ToBigEndian(value))
}

func ToBigEndian(number uint32) uint32 {
	return (number >> 24) | (number>>8)&0xFF00 | (number<<8)&0xFF0000 | (number<<24)&0xFF000000
}
