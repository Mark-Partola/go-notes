package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func SetupGracefulShutdown() <-chan os.Signal {
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	return ch
}

func worker(shutdownCh <-chan os.Signal) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-shutdownCh:
			fmt.Println("graceful shutdown..")
			time.Sleep(time.Second)
			return
		case <-ticker.C:
			fmt.Println("work")
		}
	}

}

func main() {
	wg := sync.WaitGroup{}
	shutdownCh := SetupGracefulShutdown()

	wg.Add(1)
	go func() {
		worker(shutdownCh)
		wg.Done()
	}()

	wg.Wait()
}
