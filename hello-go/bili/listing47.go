package main

import "fmt"

func f(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

// 多层台阶，一次可以上一层，也可以上两层，求台阶n一共有多少走法
func steps(floors int) int {
	if floors <= 1 {
		return 1
	}
	if floors == 2 {
		return 2
	}
	return steps(floors-1) + steps(floors-2)
}

func main() {
	ret := f(20)
	fmt.Println(ret)

	ret2 := steps(10)
	fmt.Println(ret2)
}
