package main

import (
	"fmt"
	"sync"
)

// go语言中的channel
// f1函数操作将一百个数放ch1中，f2函数中是从ch1中取数后计算平方值放到ch2

var gw sync.WaitGroup

func f1(ch1 chan int) {
	defer gw.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch2 chan int, ch1 chan int) {
	defer gw.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}

	close(ch2)
}
func main() {
	gw.Add(2)
	var ch1 chan int = make(chan int, 100)
	var ch2 chan int = make(chan int, 100)
	go f1(ch1)
	go f2(ch2, ch1)
	gw.Wait()
	for ret := range ch2 {
		fmt.Println(ret)
	}

}
