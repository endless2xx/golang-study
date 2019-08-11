package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (this *Student) init(name string, age int) {
	this.Name = name
	this.Age = age
}

// 结构体的方法
func main() {
	var stu Student
	// (&stu).init("Tom", 29) 等价
	stu.init("Tom", 29)

	fmt.Println(stu)
}
