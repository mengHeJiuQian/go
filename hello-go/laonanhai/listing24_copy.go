package main

import "fmt"

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1
	var a3 = make([]int, 3, 3)
	copy(a3, a1)

	fmt.Println(a1, a2, a3)

	x1 := [...]int{1, 3, 5}
	s1 := x1[:]
	fmt.Println(s1, len(s1), cap(s1))

	s1 = append(s1[:1], s1[2:]...)
	fmt.Println(s1, len(s1), cap(s1))

	fmt.Println(x1)

	// 切片的练习题
	var a = make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
