package main

import "fmt"

// 自定义类型和类型别名
type myInt int
type yourInt = int

func main() {
	var n1 myInt
	n1 = 1
	fmt.Println(n1)
	fmt.Printf("%T \n", n1)

	var n2 yourInt
	n2 = 1
	fmt.Println(n2)
	fmt.Printf("%T \n", n2)
}
