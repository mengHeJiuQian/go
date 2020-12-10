package main

import "fmt"

func main() {

	p1()
	p2()
	p3()

}

func p1() {
	fmt.Println("p1")
}

func p2() {
	fmt.Println("p2")
	//defer recover()
	defer func() {
		i := recover()
		fmt.Println(i)

	}()
	panic("系统出现异常")
}

func p3() {
	fmt.Println("p3")
}
