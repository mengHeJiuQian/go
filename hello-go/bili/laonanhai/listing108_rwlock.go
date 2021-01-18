package main

import (
	"fmt"
	"sync"
	"time"
)

// 演示读写锁

var (
	x = 0
	//lock sync.Mutex
	lock sync.RWMutex
	gw   sync.WaitGroup
)

func read() {
	gw.Add(1)
	defer gw.Done()
	lock.Lock()
	//time.Sleep(time.Millisecond)
	lock.Unlock()
}

func write() {
	gw.Add(1)
	defer gw.Done()
	lock.Lock()
	x++
	//time.Sleep(time.Millisecond * 5)
	lock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		go read()
	}

	for i := 0; i < 1000; i++ {
		go write()
	}
	gw.Wait()
	fmt.Println(time.Now().Sub(start))
}
