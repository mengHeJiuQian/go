package main

import (
	"fmt"
	"sync"
)

var gw sync.WaitGroup

func a() {
	defer gw.Done()
	for i := 0; i < 10; i++ {
		fmt.Print("A:", i)
		fmt.Print("	   ")
	}
}

func b() {
	defer gw.Done()
	for i := 0; i < 10; i++ {
		fmt.Print("B:", i)
		fmt.Print("	   ")
	}
}

func main() {
	// runtime.GOMAXPROCS(1)
	gw.Add(2)
	go a()
	go b()
	gw.Wait()
}
