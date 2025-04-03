package main

type Semaphore interface {
	Acquire()
	Release()
}
