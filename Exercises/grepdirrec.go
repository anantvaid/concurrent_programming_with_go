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

func searchPath(path, search string) {
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Printf("Error accessing %s: %v\n", path, err)
		return
	}

	// If it's a FILE → search inside the file.
	if !fi.IsDir() {
		grepfile(path, search)
		return
	}

	// If it's a DIRECTORY → recursively search inside.
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Error reading directory %s: %v\n", path, err)
		return
	}

	for _, entry := range entries {
		// Build full path
		entryPath := filepath.Join(path, entry.Name())

		// Spawn new goroutine for each file/directory
		go searchPath(entryPath, search)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Usage: go run grepdirrec.go <search_string> <directory_path>")
		return
	}

	search := args[0]
	startDir := args[1]

	// Start recursive search from root directory
	go searchPath(startDir, search)

	// Temporary waiting for goroutines
	time.Sleep(2 * time.Second)
}