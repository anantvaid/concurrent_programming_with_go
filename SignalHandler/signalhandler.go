package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			fmt.Println("Working... (Processing requests)")
			time.Sleep(1 * time.Second)
		}
	}()

	fmt.Println("App Started. Press Ctrl+C to test graceful shutdown...")

	sig := <-sigs

	fmt.Println()
	fmt.Println("--- Signal Received! ---")
	fmt.Printf("Signal: %v\n", sig)

	fmt.Println("Cleaning up connections...")
	time.Sleep(2 * time.Second) // Simulate finishing active requests
	fmt.Println("Cleanup done. Exiting now.")

	os.Exit(0)
}
