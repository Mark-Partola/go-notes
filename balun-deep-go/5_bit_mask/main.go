package main

import "fmt"

const (
	OpenModeIn = 1 << iota
	OpenModeOut
	OpenModeAppend
	OpenModeBinary

	// presets
	OpenModeInOut = OpenModeIn | OpenModeOut
)

func main() {
	Open("filename.out", OpenModeIn)
	Open("filename.out", OpenModeOut)
	Open("filename.out", OpenModeBinary)
	Open("filename.out", OpenModeInOut)
}

func Open(filename string, mask int8) {
	if mask&OpenModeIn != 0 {
		fmt.Println("in")
	}
	if mask&OpenModeOut != 0 {
		fmt.Println("out")
	}
	if mask&OpenModeAppend != 0 {
		fmt.Println("append")
	}
	if mask&OpenModeBinary != 0 {
		fmt.Println("binary")
	}
}
