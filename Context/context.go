// Learning the use of Context in Go
package main

import (
	"context"
	"fmt"
	"time"
)

func slowDatabaseQuery(ctx context.Context) {
	// Simulate a slow database query with a 5-second sleep
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Database query completed")
	case <-ctx.Done():
		fmt.Println("Database query canceled:", ctx.Err())
	}
}

func main() {
	// Create a context that will be canceled after 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	fmt.Println("Database Query Starts now")
	slowDatabaseQuery(ctx)

	fmt.Println("Main: Execution after return.")
}
