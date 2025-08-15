package main

import "fmt"

func main() {
	fmt.Println(SearchStrict(0b00011111))
	fmt.Println(SearchStrict(0b00010101))
	fmt.Println(SearchStrict(0b1010110))
}

func SearchStrict(pattern int) []string {
	restaurants := []struct {
		title  string
		bitmap int
	}{{
		title:  "A",
		bitmap: 0b00010101,
	}, {
		title:  "B",
		bitmap: 0b1010110,
	}, {
		title:  "C",
		bitmap: 0b01110000,
	}, {
		title:  "D",
		bitmap: 0b10111100,
	}, {
		title:  "E",
		bitmap: 0b1010110,
	}}

	var res []string
	for _, r := range restaurants {
		// xor for strict search
		if r.bitmap^pattern == 0 {
			res = append(res, r.title)
		}
	}
	return res
}
