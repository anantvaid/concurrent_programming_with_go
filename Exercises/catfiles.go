package main

import (
	"fmt"
	"os"
	"time"
)

func catfile(filename string) {
	// Read entire file
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", filename, err)
		return
	}

	// Print file contents
	fmt.Printf("=== %s ===\n", filename)
	fmt.Println(string(data))
}

func main() {
	args := os.Args[1:]

	for _, arg := range args {
		go catfile(arg)
	}

	time.Sleep(2 * time.Second)
}