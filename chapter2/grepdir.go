package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func grepfile(filename, search string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", filename, err)
		return
	}

	content := string(data)

	if strings.Contains(content, search) {
		fmt.Printf("MATCH: %s found in %s\n", search, filename)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Usage: go run grepdir.go <search_string> <directory_path>")
		return
	}

	search := args[0]
	dir := args[1]

	// Read directory entries
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error reading directory %s: %v\n", dir, err)
		return
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			filepath := filepath.Join(dir, entry.Name())
			go grepfile(filepath, search)
		}
	}

	// Temporary wait for all goroutines
	time.Sleep(2 * time.Second)
}