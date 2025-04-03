package main

import (
	"fmt"
	"net/http"
)

type result struct {
	status  int
	address string
}

func main() {
	addresses := []string{"https://example.com", "https://example.org"}
	result := BatchProcess(addresses)
	fmt.Println(result)
}

func BatchProcess(addresses []string) result {
	// buffer для того, чтобы не заблокироваться на записи и не выйти из горутины
	// когда еще никто не слушает
	out := make(chan result, 1)

	for _, address := range addresses {
		go func() {
			status, err := process(address)
			if err != nil {
				return
			}

			select {
			case out <- result{status, address}:
			default:
				return
			}
		}()
	}

	return <-out
}

func process(address string) (int, error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		return 0, err
	}

	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	return res.StatusCode, nil
}
