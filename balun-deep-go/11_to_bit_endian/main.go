package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func main() {
	var value uint32 = 0x12345678

	fmt.Printf("%032b\n", value)
	fmt.Printf("%032b\n", ToBigEndian(value))
}

func ToBigEndian(number uint32) uint32 {
	return (number>>24)&0xFF | (number>>8)&0xFF00 | (number<<8)&0xFF0000 | (number<<24)&0xFF000000
}

func TestConversion(t *testing.T) {
	tests := map[string]struct {
		number uint32
		result uint32
	}{
		"test case #1": {
			number: 0x00000000,
			result: 0x00000000,
		},
		"test case #2": {
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		"test case #4": {
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		"test case #5": {
			number: 0x01020304,
			result: 0x04030201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToBigEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
