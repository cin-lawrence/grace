package grace

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func sendRequest(ctx context.Context) error {
	req, err := http.NewRequestWithContext(
		ctx, "GET", "http://mockserver:8123", nil,
	)
	if err != nil {
		log.Printf("can not create request: %v", err)
		return err
	}

	req.Header.Set("Contenxt-Type", "application/json")
	client := &http.Client{Timeout: 0}
	log.Println("Sending request to mockserver...")
	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.Canceled {
			log.Printf("can not send request: %v", err)
			return ctx.Err()
		}
		log.Printf("another error when sending request: %v", err)
		return err
	}
	log.Println("After received response")
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error reading response body: %v", err)
	}
	log.Printf("Received response: %s", body)
	return nil
}

func httpStop(ctx context.Context) error {
	for {
		if err := sendRequest(ctx); err != nil {
			log.Printf("httpStop: received error: %v", err)
			return err
		}
		log.Println("httpStop: Sucessful, sleep for 3s")
		time.Sleep(3 * time.Second)
	}
}

func httpFork(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	if err := httpStop(ctx); err != nil {
		if errors.Is(err, context.Canceled) {
			log.Println("HttpFork: received graceful shutdown")
		} else {
			log.Println("HttpFork: Unexpected error")

		}
		return
	}
}

func HttpStopMain() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go httpFork(ctx, &wg)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	select {
	case sig := <-signalChan:
		log.Printf("\nReceived signal: %s, shutting down...\n", sig)
		cancel()
	}

	log.Println("Cancaled children goroutine, waiting...")
	wg.Wait()
	log.Println("Shutdown gracefully")

}
