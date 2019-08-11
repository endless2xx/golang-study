package main

import "fmt"

type Student struct {
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Score float32 `json:"score"`
}

// 结构体
func main() {
	// 初始化方式一：
	var stu1 Student
	stu1.Age = 18
	stu1.Name = "Tom"
	stu1.Score = 99

	fmt.Println(stu1)
	fmt.Printf("Name: %p\n", &stu1.Name)
	fmt.Printf("Age: %p\n", &stu1.Age)
	fmt.Printf("Score: %p\n", &stu1.Score)

	// 初始化方式二： 等价于 var stu2 *Student = &Student{}
	var stu2 = &Student{
		Age:  120,
		Name: "Amy",
	}
	fmt.Println(stu2)

	// 初始化方式三
	var stu3 = Student{
		Name: "Tom",
	}
	fmt.Println(stu3)
}
