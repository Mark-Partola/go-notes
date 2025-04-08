package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	request(ctx)
}

func request(ctx context.Context) {
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	select {
	case <-timer.C:
		fmt.Println("timer")
	case <-ctx.Done():
		fmt.Println("cancelled")
	}
}
