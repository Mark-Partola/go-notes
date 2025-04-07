package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	group := New()

	group.Do(func() error {
		return fmt.Errorf("something happened")
	})

	time.Sleep(time.Second)

	for i := range 10 {
		group.Do(func() error {
			time.Sleep(time.Duration(i) * time.Second)
			fmt.Println(i)
			return nil
		})
	}

	if err := group.Wait(); err != nil {
		fmt.Println(err)
	}
}

type errGroup struct {
	wg   sync.WaitGroup
	done chan struct{}
	err  error
	once sync.Once
}

func New() *errGroup {
	return &errGroup{
		done: make(chan struct{}),
	}
}

func (group *errGroup) Do(task func() error) {
	select {
	case <-group.done:
		return
	default:
	}

	group.wg.Add(1)
	go func() {
		defer group.wg.Done()

		select {
		case <-group.done:
			return
		default:
			err := task()
			if err != nil {
				group.once.Do(func() {
					group.err = err
					close(group.done)
				})
			}
		}
	}()
}

func (group *errGroup) Wait() error {
	group.wg.Wait()
	return group.err
}
