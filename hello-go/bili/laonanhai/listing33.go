package main

import (
	"fmt"
	"reflect"
	"unicode"
)

// 判断一个字符串中汉字的数量
func main() {

	s1 := "hello你好"
	CountHan(s1)

	Palindromic("lol")

	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

}

func CountHan(str string) int {
	var han int
	for _, c := range str {
		if unicode.Is(unicode.Han, c) {
			han++
		}
	}
	fmt.Println(han)
	return han
}

func Palindromic(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if string(str[i]) != string(str[len(str)-1-i]) {
			fmt.Println(reflect.TypeOf(str[i]))
			fmt.Println(string(str[i]))
			fmt.Println(string(str[len(str)-1-i]))
			fmt.Println("不是回文")
			return false
		}
	}
	fmt.Println("是回文")
	return true
}

func f1() int {
	x := 5

	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return x
}

func f3() (y int) {
	x := 5

	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func f5() (x int) {
	defer func() {
		x++
	}()
	return 5
}
