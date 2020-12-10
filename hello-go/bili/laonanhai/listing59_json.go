package main

import (
	"encoding/json"
	"fmt"
)

// 结构体转成json字符串

type person struct {
	Name   string
	Age    int
	salary int `json:"salary"`
}

func main() {
	p1 := person{
		Name:   "张三",
		Age:    9000,
		salary: 11111,
	}
	bytes, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v", err)
		return
	}
	fmt.Printf("%#v \n", string(bytes))

}
