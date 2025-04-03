package main

import (
	"sync"
	"sync/atomic"
)

type CondAtomicsSemaphore struct {
	available *atomic.Int32
	cond      *sync.Cond
}

func NewCondAtomicsSemaphore(n int32) Semaphore {
	available := &atomic.Int32{}
	available.Add(n)
	mx := &sync.Mutex{}
	return &CondAtomicsSemaphore{
		available: available,
		cond:      sync.NewCond(mx),
	}
}

func (s *CondAtomicsSemaphore) Acquire() {
	for s.available.Load() == 0 {
		s.cond.L.Lock()
		s.cond.Wait()
		s.cond.L.Unlock()
	}

	s.available.Add(-1)
}

func (s *CondAtomicsSemaphore) Release() {
	s.available.Add(1)
	s.cond.Signal()
}
