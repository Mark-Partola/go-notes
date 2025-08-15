package main

import (
	"fmt"
)

func main() {
	val := 100
	pointer := &val

	fmt.Println(*pointer)
	process(&pointer)
	fmt.Println(*pointer)
}

func process(tmp **int) {
	val := 200

	*tmp = &val
}
