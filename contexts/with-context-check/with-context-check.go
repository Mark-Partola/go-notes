package main

import "context"

func main() {
	ctx := context.Background()
	WithContextCheck(ctx, func() {
		println("action")
	})
}

func WithContextCheck(ctx context.Context, action func()) {
	if action == nil || ctx.Err() != nil {
		return
	}

	action()
}
