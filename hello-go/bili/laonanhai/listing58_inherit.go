package main

import "fmt"

// 结构体模拟继承

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s会动！", a.name)
}

// dog继承animal
type dog struct {
	feet uint8
	animal
}

func (d dog) wang() {
	fmt.Printf("%s在叫：汪汪汪！", d.name)
}

func main() {
	d1 := dog{
		feet:   4,
		animal: animal{name: "张三"},
	}

	fmt.Println(d1)
	d1.wang()
	d1.move()

}
