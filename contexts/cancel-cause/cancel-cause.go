package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancelCause(context.Background())

	cancel(errors.New("something happened"))

	request(ctx)
}

func request(ctx context.Context) {
	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop()

	select {
	case <-timer.C:
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println(ctx.Err(), context.Cause(ctx))
	}
}
