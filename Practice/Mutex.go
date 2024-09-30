package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var total int64

func adds(n int, theLock *sync.Mutex) bool {
	for i := 0; i < n; i++ {
		theLock.Lock()
		total += 1
		theLock.Unlock()
	}
	wg.Done() //let wait group know we have finished
	return true
}

func main() {
	var theLock sync.Mutex

	total = 0

	wg.Add(10)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		go adds(1000, &theLock)
	}
	wg.Wait()
	fmt.Println(total)
}
