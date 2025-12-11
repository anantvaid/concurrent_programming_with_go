package main

import "fmt"

func writeFile(fileName string) {
	fmt.Printf("Opening the file %s\n", fileName)
	defer fmt.Println("File closed")

	fmt.Println("Processing and writing to the file")
	fmt.Println("Disk Error")
	panic("Disk full error")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic Caught", r)
		}
	}()
	writeFile("sample.txt")
}
