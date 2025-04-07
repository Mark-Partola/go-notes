package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	workers := 10
	wg := sync.WaitGroup{}
	wg.Add(workers)
	singleflight := New()

	for idx := range workers {
		go func() {
			defer wg.Done()
			result, err := singleflight.Do("request", func() (any, error) {
				fmt.Println("passed: ", idx)
				time.Sleep(time.Second)
				return idx, nil
			})

			fmt.Println(result, err)
		}()
	}

	wg.Wait()
}

type (
	call struct {
		done   chan struct{}
		result any
		err    error
	}
	singleflight struct {
		mx    sync.Mutex
		calls map[string]*call
	}
)

func New() *singleflight {
	return &singleflight{
		mx:    sync.Mutex{},
		calls: make(map[string]*call),
	}
}

func (s *singleflight) Do(key string, task func() (any, error)) (any, error) {
	s.mx.Lock()
	if call, found := s.calls[key]; found {
		s.mx.Unlock()
		return s.wait(call)
	}

	call := &call{
		done: make(chan struct{}),
	}
	s.calls[key] = call
	s.mx.Unlock()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				call.err = fmt.Errorf("panic happened: %v", r)
			}

			s.mx.Lock()
			close(call.done)
			delete(s.calls, key)
			s.mx.Unlock()
		}()

		call.result, call.err = task()
	}()

	return s.wait(call)
}

func (s *singleflight) wait(call *call) (any, error) {
	<-call.done
	return call.result, call.err
}
