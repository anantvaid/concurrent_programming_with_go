package main

import (
	"fmt"
	"os"
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
		fmt.Printf("%s FOUND in %s\n", search, filename)
	} else {
		fmt.Printf("%s NOT FOUND in %s\n", search, filename)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Usage: go run grepfiles.go <search_string> <file1> <file2> ...")
		return
	}

	search := args[0]
	files := args[1:]

	for _, file := range files {
		go grepfile(file, search)
	}

	time.Sleep(2 * time.Second)
}	