package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand/v2"
)

func think(index int){
	var X time.Duration
	X=time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second)//wait random time amount
	fmt.Println("Phil: ",index,"was thinking")
}

func eat(index int){
	var X time.Duration
	X=time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second)//wait random time amount
	fmt.Println("Phil: ",index,"was eating")
}


func getForks(index int, forks map[int]chan bool){
	forks[index] <- true
	forks[(index+1)%5] <-true
}

func putForks(index int, forks map[int]chan bool){
	<-forks[index]
	<-forks[(index+1)%5]
}

func doPhilStuff(index int, wg *sync.WaitGroup, forks map[int] chan bool){
	for {
		think(index)
		getForks(index,forks)
		eat(index)
		putForks(index,forks)
	}
	wg.Done()
}


func main() {
	var wg sync.WaitGroup
	philCount :=5
	wg.Add(philCount)

	forks:= make(map[int]chan bool)
	for k := range philCount{
		forks[k]=make(chan bool, 1)
	}//set up forks
	for N := range philCount {
		go doPhilStuff(N, &wg, forks)
	}//start philosophers
	wg.Wait() //wait here until everyone (10 go routines) is done

}//main
