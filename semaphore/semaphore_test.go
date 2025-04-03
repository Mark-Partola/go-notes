package main

import (
	"sync"
	"testing"
)

func BenchmarkSemaphoreCondAtomics(b *testing.B) {
	semaphore := NewCondAtomicsSemaphore(3)
	for b.Loop() {
		run(semaphore)
	}
}

func BenchmarkSemaphoreCondMutex(b *testing.B) {
	semaphore := NewCondMutexSemaphore(3)
	for b.Loop() {
		run(semaphore)
	}
}

func BenchmarkSemaphoreChannel(b *testing.B) {
	semaphore := NewChannelSemaphore(3)
	for b.Loop() {
		run(semaphore)
	}
}

func run(s Semaphore) {
	workers := 10
	wg := sync.WaitGroup{}

	wg.Add(workers)
	for range workers {
		go func() {
			s.Acquire()
			s.Release()
			wg.Done()
		}()
	}

	wg.Wait()
}
