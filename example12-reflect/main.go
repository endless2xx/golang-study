package main

import (
	"fmt"
	"reflect"
)

func test(b interface{}) {
	// 通过反射获取变量的类型
	t := reflect.TypeOf(b)
	fmt.Printf("typeOf b is: %s \n", t)

	// 通过反射获取变量的值
	v := reflect.ValueOf(b)
	fmt.Printf("valueOf b is: %v \n", v)
	k := v.Kind()                       // 类别
	fmt.Printf("KindOf v is: %s \n", k) // output：struct

	iv := v.Interface()
	stu, ok := iv.(Student)
	if ok {
		fmt.Printf("%v %T \n", stu, stu)
	}
}

type Student struct {
	Name  string
	Age   int
	Score float32
}

// 反射：在运行时动态获取变量的相关信息
func main() {
	var a Student = Student{
		Name:  "Tom",
		Age:   19,
		Score: 99,
	}
	test(a)
}
