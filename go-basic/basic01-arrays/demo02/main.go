package main

import "fmt"

// 多维数组
func main() {
	// 二维数组
	var array1 [4][2]int
	fmt.Println(array1)
	// 使用字面量初始化
	array2 := [4][2]int{{1, 2}, {2, 3}, {3, 4}, {5, 6}}
	fmt.Println(array2)
	// 初始化两个维度的指定元素的值
	array3 := [4][2]int{1: {0: 1, 2}}
	fmt.Println(array3)

	//////////////////////////////////////////
	// 访问数组的元素
	array4 := [2][2]int{}
	array4[0][0] = 10
	array4[0][1] = 20
	array4[1][0] = 30
	array4[1][1] = 40
	fmt.Println(array4)

	//////////////////////////////////////////
	// 赋值: 两个相同维度的数组
	array5 := [2][2]int{}
	array6 := [2][2]int{}
	array6[0][0] = 10
	array6[0][1] = 20
	array6[1][0] = 30
	array6[1][1] = 40
	array5 = array6
	fmt.Println(array5)
	// 独立的将某个维度的元素赋值给另一个元素，如下
	// 将二维数组 array6 中的第二个元素中的数组 {30, 40} 赋值给了一维数组 array7
	array7 := [2]int{}
	array7 = array6[1]
	fmt.Println(array7)
}
