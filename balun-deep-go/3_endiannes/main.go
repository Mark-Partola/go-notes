package main

import (
	"fmt"
	"unsafe"
)

func main() {
	if IsLittleEndian() {
		fmt.Println("little endian")
	} else {
		fmt.Println("big endian")
	}

	value := 0x12345678

	pointer := unsafe.Pointer(&value)

	fmt.Print("0x")
	for i := range 4 {
		b := *(*int8)(unsafe.Add(pointer, i))
		fmt.Printf("%x", b)
	}

	fmt.Println()
}

func IsLittleEndian() bool {
	value := 0x0001
	pointer := (*int8)(unsafe.Pointer(&value))
	return *pointer == 1
}

func IsBigEndian() bool {
	return !IsLittleEndian()
}
