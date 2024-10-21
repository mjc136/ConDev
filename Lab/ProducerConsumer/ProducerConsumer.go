package main

import (
	"sync"
)

func Consumer() {

}
func Producer() {

}

func doStuff() {

}

func main() {
	totalRoutines := 10
	arrived := 0
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	var theLock sync.Mutex
	for i := 0; i < totalRoutines; i++ {
		go doStuff()
	}
}
