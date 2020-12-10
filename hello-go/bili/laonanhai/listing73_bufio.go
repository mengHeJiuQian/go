package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	fmt.Println("请输入：")

	// 输入“a b c”,只能读到“a”,遇到空格就结束
	fmt.Scanln(&s)
	fmt.Printf("您输入的内容是：%s", s)

	var s2 string
	reader := bufio.NewReader(os.Stdin)
	s2, _ = reader.ReadString('\n')
	fmt.Printf("您输入的内容是：%s", s2)

}
