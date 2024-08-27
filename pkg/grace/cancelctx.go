package grace

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func CancelHandler(ctx context.Context, index int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Goroutine %d started\n", index)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Goroutine %d received shutdown\n", index)
			time.Sleep(2 * time.Second)
			return
		default:
			fmt.Printf("Goroutine %d working...\n", index)
			time.Sleep(1 * time.Second)
		}
	}
}

func CancelMain() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go CancelHandler(ctx, i, &wg)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	<-signalChan
	fmt.Println("\nReceived shutdown signal, canceling context...")

	cancel()

	wg.Wait()
	fmt.Println("All goroutines stopped, exiting...")
}
