package main

import "time"

func hello(i int) {
	//println("hello", i)
}

func main() {
	for i := 0; i < 1000000; i++ {
		go hello(i + 1)
	}
	println("main")
	time.Sleep(time.Second)
	println("over")
}
