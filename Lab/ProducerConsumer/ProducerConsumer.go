package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numThreads = 10
	size       = 5
	numLoops   = 2
)

// Producer creates items and sends them to the channel
func producer(id int, theChan chan int, numLoops int, wg *sync.WaitGroup) {
	for i := 0; i < numLoops; i++ {
		time.Sleep(time.Second) // Simulate some work by sleeping for 1 second
		item := id + 10 + i     // Unique ID for each item produced by each thread
		fmt.Println("Producer ", id, "produced", item)
		theChan <- item // Send the produced item to the channel
	}
	wg.Done()
}

// Consumer gets items from the channel and processes them
func consumer(id int, theChan chan int, wg *sync.WaitGroup) {
	for item := range theChan { // Continue to consume items until the channel is closed
		fmt.Println("Consumer ", id, "consumed", item)
		time.Sleep(time.Second) // Simulate some work by sleeping for 1 second
	}
	wg.Done()
}

func main() {
	//used 2 wait groups to avoid deadlock
	var producerWG sync.WaitGroup
	var consumerWG sync.WaitGroup

	// Create a buffered channel to hold items
	theChan := make(chan int, size)

	// Start producer goroutines
	producerWG.Add(numThreads / 2) // Add the number of producer goroutines to the WaitGroup
	for i := 0; i < numThreads/2; i++ {
		go producer(i, theChan, numLoops, &producerWG) // Start each producer goroutine
	}

	// Start consumer goroutines
	consumerWG.Add(numThreads / 2) // Add the number of consumer goroutines to the WaitGroup
	for i := 0; i < numThreads/2; i++ {
		go consumer(i, theChan, &consumerWG) // Start each consumer goroutine
	}

	// Wait for all consumers to finish
	consumerWG.Wait()
	close(theChan)

	fmt.Println("All producers and consumers have finished.")
}
