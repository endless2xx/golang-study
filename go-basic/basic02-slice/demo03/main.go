package main

import "fmt"

func main() {
	// 创建字符串切片，长度和容量都是 5
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	// 基于上面的切片，创建一个新的切片
	// 切片的长度为 3-2=1
	// 切片的容量为 4-2=2
	// 切片的起始位置是上一个切片索引为 2 的元素「Plum」开始
	slice := source[2:3:6]
	fmt.Println(slice)

	slice1 := make([]int, 5, 10)
	fmt.Println(cap(slice1))
	slice1 = append(slice1, 10)
	fmt.Println(cap(slice1))
	slice1 = append(slice1, 20)
	fmt.Println(cap(slice1))
	slice1 = append(slice1, 30)
	fmt.Println(cap(slice1))
	slice1 = append(slice1, 40)
	fmt.Println(cap(slice1))
	slice1 = append(slice1, 50)
	fmt.Println(cap(slice1))
	slice1 = append(slice1, 60)
	fmt.Println(cap(slice1))
	slice1 = append(slice1, 70)
	fmt.Println(cap(slice1))
	fmt.Println(slice1)
}
