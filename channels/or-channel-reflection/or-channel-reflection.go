package main

import (
	"fmt"
	"reflect"
)

func main() {
	ch := make(chan int)
	vch := reflect.ValueOf(ch)

	ok := vch.TrySend(reflect.ValueOf(10))
	fmt.Println(ok, vch.Len(), vch.Cap())

	arms := []reflect.SelectCase{
		{Dir: reflect.SelectDefault},
		{Dir: reflect.SelectRecv, Chan: vch},
	}

	idx, vRecv, ok := reflect.Select(arms)
	fmt.Println(idx, vRecv, ok)

	vRecv.Close()
}
