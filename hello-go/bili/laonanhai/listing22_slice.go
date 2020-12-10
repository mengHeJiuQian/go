package main

import "fmt"

func main() {
	// 定义一个存放int类型元素的切片
	var s1 []int
	var s2 []string
	fmt.Println(s1, s2)

	// 切片初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"xx", "xx", "xx"}
	fmt.Println(s1, s2)

	fmt.Printf("len(s1)=%d, cap(s1)=%d \n", len(s1), cap(s1))

	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4]

	fmt.Println(s3)

}
