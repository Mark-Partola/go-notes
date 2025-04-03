package main

import (
	"sync"
)

type CondMutexSemaphore struct {
	available int32
	cond      *sync.Cond
}

func NewCondMutexSemaphore(n int32) Semaphore {
	mx := &sync.Mutex{}
	return &CondMutexSemaphore{
		available: n,
		cond:      sync.NewCond(mx),
	}
}

func (s *CondMutexSemaphore) Acquire() {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()

	for s.available == 0 {
		s.cond.Wait()
	}

	s.available--
}

func (s *CondMutexSemaphore) Release() {
	s.cond.L.Lock()
	defer s.cond.L.Unlock()
	s.available++
	s.cond.Signal()
}
