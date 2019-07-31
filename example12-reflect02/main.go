package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string
	Age   int
	Score float32
}

func TestStruct(a interface{}) {
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)

	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)
}

func (p *Student) initStruct(name string) {
	p.Name = name
}

func main() {
	var a Student = Student{
		Name:  "Tom",
		Age:   19,
		Score: 100,
	}
	TestStruct(a)
	fmt.Println(a)
}
