package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// wg is used to wait for the program to finish the goroutines spawned.
var wg sync.WaitGroup

func main() {
	wg.Add(2)

	fmt.Println("Staring Goroutines")
	// Start a goroutine with the label A
	go printCounts("A")
	// Start a goroutine with the label B
	go printCounts("B")
	// waiting for the goroutines to finish
	fmt.Println("Waiting to Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
	fmt.Println(runtime.Version())
}

func printCounts(label string) {
	defer wg.Done()
	for count := 1; count <= 10; count++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Count %d from %s\n", count, label)
	}
}
