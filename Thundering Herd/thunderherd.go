package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for id := range ch {
		fmt.Printf("Worker %d starting\n", id)
		time.Sleep(3 * time.Second)
		fmt.Printf("Worker %d done\n", id)
	}
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 5
	numTasks := 10
	ch := make(chan int, numTasks)

	fmt.Printf("Starting %d workers to process %d tasks\n", numWorkers, numTasks)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	for j := 0; j < numTasks; j++ {
		ch <- j
	}
	close(ch)

	wg.Wait()
	fmt.Println("All workers completed")

}
