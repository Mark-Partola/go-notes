package main

import (
	"fmt"
	"time"
)

func main() {
	limiter := NewLimiter(10, time.Second)

	ticker := time.NewTicker(90 * time.Millisecond)
	for v := range ticker.C {
		if limiter.Allow() {
			fmt.Println(v)
		} else {
			fmt.Println("discarded")
		}
	}
}

type limiter struct {
	ch chan struct{}
}

func NewLimiter(n int, period time.Duration) *limiter {
	limiter := &limiter{
		ch: make(chan struct{}, n),
	}

	leakingInterval := time.Duration(period.Nanoseconds() / int64(n))
	go limiter.startLeaking(leakingInterval)
	return limiter
}

func (l *limiter) Allow() bool {
	select {
	case l.ch <- struct{}{}:
		return true
	default:
		return false
	}
}

func (l *limiter) startLeaking(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		<-l.ch
	}
}
