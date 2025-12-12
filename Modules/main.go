package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Printf("Generated ID: %s\n", id.String())
	// Let's run go mod tidy to ensure dependencies are managed
}
