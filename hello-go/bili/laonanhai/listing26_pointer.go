package main

import "fmt"

func main() {

	// 1. &取变量地址
	n1 := 18
	p1 := &n1
	fmt.Println(p1)
	fmt.Printf("%T \n", p1)
	// 2. *取地址里的值
	m1 := *p1
	fmt.Println(m1)
	fmt.Printf("%T \n", m1)

	var a *int
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b["沙河娜扎"] = 100
	fmt.Println(b)

}
