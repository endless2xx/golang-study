package main

import "fmt"

//// 切片的创建和初始化
func main() {
	// 创建一个字符串切片，长度和容量都为 5
	// 如果只指定长度，那么容量和长度相等
	slice1 := make([]string, 5)
	fmt.Println(slice1)
	// 使用长度和容量声明切片，切片的容量就是底层数组的长度
	slice2 := make([]int, 3, 5)
	fmt.Println(slice2)
	// /////////////////////////////////////////////////// 通过字面量创建切片

	slice3 := []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Println(slice3)

	slice4 := []int{90: 90, 99: 99}
	fmt.Println(slice4)
}
