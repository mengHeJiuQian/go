package main

import "fmt"

func main() {
	// 1. 计算数组所有元素之和
	arr1 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range arr1 {
		sum += v
	}
	fmt.Println("sum = ", sum)

	// 2. 计算数组中两个元素之和等于8的元素下标
	for i := 0; i < len(arr1)-1; i++ {
		for j := i + 1; j < len(arr1); j++ {
			if arr1[i]+arr1[j] == 8 {
				fmt.Printf("(%d, %d)\n", arr1[i], arr1[j])
			}
		}
	}
}
