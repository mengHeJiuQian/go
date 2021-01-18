package main

import (
	"fmt"
	"sync"
)

// mutex互斥锁

var wg sync.WaitGroup
var lock sync.Mutex
var x = 0

func add() {
	for i := 0; i < 10000; i++ {
		lock.Lock()
		x++
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
