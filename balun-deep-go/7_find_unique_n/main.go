package main

import "fmt"

func main() {
	ns := []int{1, 1, 2, 3, 2, 4, 5, 4, 5}
	res := FindUnique(ns)
	fmt.Println(res)
}

func FindUnique(ns []int) int {
	res := 0
	for _, v := range ns {
		res ^= v
	}
	return res
}
