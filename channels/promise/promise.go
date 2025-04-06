package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	p := Promise(func(resolve func(), _ func()) {
		time.Sleep(2 * time.Second)
		resolve()
	})

	p.Then(func() {
		fmt.Println("fulfilled 1")
		time.Sleep(time.Second)
	}).Then(func() {
		fmt.Println("fulfilled 2")
		time.Sleep(time.Second)
	}).Then(func() {
		fmt.Println("fulfilled 3")
		wg.Done()
	})

	p.Catch(func() {
		defer wg.Done()
		fmt.Println("rejected")
	})

	wg.Wait()
}

type promise struct {
	fulfilled chan struct{}
	rejected  chan struct{}
}

func Promise(create func(resolve func(), reject func())) *promise {
	promise := &promise{
		fulfilled: make(chan struct{}),
		rejected:  make(chan struct{}),
	}

	finalize := func() {
		close(promise.fulfilled)
		close(promise.rejected)
	}

	go func() {
		create(
			func() {
				promise.fulfilled <- struct{}{}
				finalize()
			},
			func() {
				promise.rejected <- struct{}{}
				finalize()
			},
		)
	}()

	return promise
}

func (p *promise) Then(fn func()) *promise {
	cond := sync.NewCond(&sync.Mutex{})

	go func() {
		_, ok := <-p.fulfilled
		if ok {
			fn()
			cond.Signal()
		}
	}()

	return Promise(func(resolve func(), _ func()) {
		cond.L.Lock()
		cond.Wait()
		cond.L.Unlock()
		resolve()
	})
}

func (p *promise) Catch(fn func()) *promise {
	cond := sync.NewCond(&sync.Mutex{})

	go func() {
		_, ok := <-p.rejected
		if ok {
			fn()
			cond.Signal()
		}
	}()

	return Promise(func(resolve func(), _ func()) {
		cond.L.Lock()
		cond.Wait()
		cond.L.Unlock()
		resolve()
	})
}
