package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	group, ctx := errgroup.WithContext(ctx)

	for range 10 {
		group.Go(func() error {
			duration := time.Duration(rand.Intn(10)) * time.Second
			timer := time.NewTimer(duration)
			defer timer.Stop()

			select {
			case <-timer.C:
				fmt.Println("timeout")
				return errors.New("timeout error")
			case <-ctx.Done():
				fmt.Println("cancelled")
				return nil
			}
		})
	}

	if err := group.Wait(); err != nil {
		fmt.Println()
		fmt.Println(err.Error())
	}
}
