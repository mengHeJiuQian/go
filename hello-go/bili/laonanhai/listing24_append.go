package main

import "fmt"

func main() {
	s1 := []string{"北京", "上海", "深圳"}
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d \n", s1, len(s1), cap(s1))

	s2 := append(s1, "广州", "杭州")
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d \n", s2, len(s2), cap(s2))

	s2[0] = "武汉"

	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d \n", s2, len(s2), cap(s2))

	// 使用append删除切片中的某个元素
	a1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s3 := a1[:]
	// 删除索引为1的元素
	s3 = append(a1[:1], a1[2:]...)
	fmt.Println(s3)
}
