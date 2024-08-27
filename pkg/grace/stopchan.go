package grace

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func StopChanHandler(stopChan chan bool, index int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Goroutine %d started\n", index)
	for {
		select {
		case <-stopChan:
			fmt.Printf("Goroutine %d received shutdown\n", index)
			fmt.Println("Sleeping for 2 seconds")
			time.Sleep(2 * time.Second)
			return
		default:
			fmt.Printf("Goroutine %d working...\n", index)
			time.Sleep(1 * time.Second)
		}
	}
}

func StopChanMain() {
	stopChan := make (chan bool)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i ++ {
		wg.Add(1)
		go StopChanHandler(stopChan, i, &wg)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-signalChan:
		fmt.Printf("Received signal: %s, shutting down...\n", sig)
		close(stopChan)
	}

	fmt.Println("Before wait group...")
	wg.Wait()
	fmt.Println("All goroutine stopped, exiting...")
}
