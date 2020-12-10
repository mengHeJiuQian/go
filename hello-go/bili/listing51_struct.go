package main

import (
	"fmt"
	"unsafe"
)

// 定一个结构体
type person struct {
	name, gender string
}

// 传的是参数的拷贝
func f1(x person) {
	x.gender = "女"
}

// 传的是内存地址
func f2(x *person) {
	(*x).gender = "女"
}

type part1 struct {
	a bool
	b int32
	c int8
	d int64
	e byte
}
type part2 struct {
	a bool
	e byte
	c int8
	b int32
	d int64
}

func main() {
	var p person
	p.name = "刘洋"
	p.gender = "男"
	f1(p)
	fmt.Println(p.gender)
	f2(&p)
	fmt.Println(p.gender)

	var p2 = new(person)
	fmt.Println(&p2)
	fmt.Printf("%p \n", p2)
	fmt.Printf("%T \n", p2)
	fmt.Printf("%p \n", &p2)
	fmt.Printf("%T \n", &p2)

	// 结构体初始化方式一,key-value
	var p3 = person{
		name:   "刘洋",
		gender: "男",
	}
	fmt.Printf("%#v \n", p3)

	// 结构体初始化方式二,值列表
	var p4 = person{
		"男",
		"刘洋",
	}
	fmt.Printf("%#v \n", p4)
	fmt.Printf("%T \n", p4)

	p11 := part1{}
	p22 := part2{}
	fmt.Printf("part1 size: %d\n", unsafe.Sizeof(p11))
	fmt.Printf("part2 size: %d\n", unsafe.Sizeof(p22))
}
