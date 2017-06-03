package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	count := make(chan int)
	wg.Add(2)
	fmt.Println("Start goroutines...")
	go printCounts("A", count)
	go printCounts("B", count)
	fmt.Println("Channel begin")
	count <- 1
	fmt.Println("Waiting to finish...")
	wg.Wait()
	fmt.Println("\nTerminating Program...")
}

func printCounts(label string, count chan int) {
	defer wg.Done()
	for {
		val, ok := <-count
		if ok == false {
			fmt.Println("Channel was closed")
			return
		}
		fmt.Printf("Count: %d received from %s \n", val, label)
		if val == 10 {
			fmt.Printf("Channel closed from %s \n", label)
			// Close the Channel
			close(count)
			return
		}
		val++
		// Send count back to the other goroutine
		count <- val
	}
}
