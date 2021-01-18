package main

import (
	"fmt"
	"sync"
)

// go语言中的channel

var b chan int
var wg sync.WaitGroup

func main() {
	wg.Add(1)
	fmt.Println(b)
	b = make(chan int)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("从b中取值：", x)
	}()
	b <- 10
	fmt.Println("从b中存值：", 10)
	b = make(chan int, 16)
	fmt.Println(b)
	wg.Wait()
}
