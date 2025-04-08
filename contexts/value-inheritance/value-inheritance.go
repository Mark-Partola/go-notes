package main

import (
	"context"
	"fmt"
	"time"
)

type TraceIDKey struct{}
type UserIDKey struct{}

func main() {
	ctx := context.WithValue(
		context.WithValue(
			context.Background(),
			TraceIDKey{},
			"c6e72394-60a8-4916-8e4d-fd685a66ec79",
		),
		UserIDKey{},
		"9a78410f-0ddf-48b0-8e5e-74c79be5948e",
	)

	request(ctx)
}

func request(ctx context.Context) {
	traceID, ok := ctx.Value(TraceIDKey{}).(string)
	if ok {
		fmt.Println(traceID)
	}

	userID, ok := ctx.Value(UserIDKey{}).(string)
	if ok {
		fmt.Println(userID)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	traceID, ok = ctx.Value(TraceIDKey{}).(string)
	if ok {
		fmt.Println(traceID)
	}
}
