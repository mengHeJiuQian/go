package main

import "fmt"

// 匿名字段,person属性的名，默认是类型的名字
type person struct {
	string
	int
}

func main() {
	p1 := person{
		"刘洋",
		24,
	}
	fmt.Println(p1)
	fmt.Println(p1.string)
	fmt.Println(p1.int)

}
