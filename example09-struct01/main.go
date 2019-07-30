package main

import (
	"fmt"
	"time"
)

// 结构体中使用 匿名字段
type Car struct {
	name  string
	color string
}

type Train struct {
	Car
	int   // 同一种类型的匿名字段只能出现一次
	start time.Time
	color string
}

func main() {
	// 访问匿名字段
	var t Train
	t.name = "HuoCar" // t.Car.name = ""
	t.color = "绿皮"    // t.Car.color = "" 默认使用自己的字段
	t.int = 16

	fmt.Println(t)
}
