package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	PrettyPrint(Convert("192.168.0.1"))
	PrettyPrint(Convert("192.288.0.1"))
	PrettyPrint(Convert("127.0.0.1"))
	PrettyPrint(Convert("192.288.0.1.9"))
	PrettyPrint(Convert("255.255.255.0"))
}

func PrettyPrint(address uint32, err error) {
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%032b\n", address)
}

var ErrInvalidIPv4Address = errors.New("invalid IPv4 address")

func Convert(address string) (uint32, error) {
	octetsCount := 4
	segments := strings.Split(address, ".")

	if len(segments) != octetsCount {
		return 0, ErrInvalidIPv4Address
	}

	var result uint32
	for idx := range octetsCount {
		octet, err := strconv.Atoi(segments[idx])
		if err != nil {
			return 0, ErrInvalidIPv4Address
		}

		if octet < 0 || octet > 255 {
			return 0, ErrInvalidIPv4Address
		}

		offset := (octetsCount - idx - 1) * 8
		result |= uint32(octet) << offset
	}

	return result, nil
}
