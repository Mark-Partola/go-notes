package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arr := [...]int{1, 2, 3, 4, 5}

	pointer := unsafe.Pointer(&arr)
	elementSize := unsafe.Sizeof(int(0))

	first := *(*int)(unsafe.Add(pointer, 0*elementSize))
	second := *(*int)(unsafe.Add(pointer, 1*elementSize))
	dangerous := *(*int)(unsafe.Add(pointer, 10*elementSize))

	fmt.Println(first, second, dangerous)
}
