package main

import "fmt"

func main() {
	ns := []int{1, 2, 5, 3, 4}
	fmt.Println(HasDuplicateN(ns))
	ns = []int{1, 2, 5, 3, 5, 4}
	fmt.Println(HasDuplicateN(ns))
}

func HasDuplicateN(ns []int) (bool, int) {
	var lookup int8

	for _, v := range ns {
		if lookup&(1<<v) > 0 {
			return true, v
		}

		lookup |= 1 << v
	}

	return false, 0
}
