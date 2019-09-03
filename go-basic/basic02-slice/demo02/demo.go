package main

import "fmt"

func main() {
	slice1 := []int{10, 20, 30, 40, 50}
	// 修改切片中的值
	slice1[2] = 300
	fmt.Println(slice1)

	newSlice := slice1[1:3]
	newSlice[1] = 35
	fmt.Println(slice1)

	fmt.Println(newSlice)
	newSlice = append(newSlice, 60)
	fmt.Println(newSlice)
	fmt.Println(slice1)
}
