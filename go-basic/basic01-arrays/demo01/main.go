package main

import "fmt"

// Currency 货币符号
type Currency int

// USD 美元
const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)

func main() {
	// 1：声明使用默认值初始化的数组
	array1 := [3]int{}
	fmt.Println("1: ", array1)
	// 2：使用字面量初始化数组
	// 字面量不够时剩下的元素初始化为零值
	array2 := [3]int{0, 1}
	fmt.Println("2: ", array2)
	// 3: 指定特定元素的值，其他元素默认零值
	array3 := [3]int{2: 10, 1: 100}
	fmt.Println("3: ", array3)
	// 4: 更具字面量计算数组长度
	array4 := [...]int{1, 3, 5, 7, 9, 11}
	fmt.Println("4: ", array4)
	array5 := [...]int{1: 100, 8: 200, 3, 4, 3, 2, 1}
	fmt.Println("5: ", array5)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(symbol)
	// 元素是指针类型
	array6 := [...]*int{new(int), new(int)}
	fmt.Println(array6)
	*array6[0] = 100
	*array6[1] = 200
	fmt.Println(*array6[1])

	// 数组的赋值操作
	array7 := [5]string{}
	array8 := [...]string{"Red", "Blue", "Green", "Yellow", "pink"}
	array7 = array8
	fmt.Println(array7)
	fmt.Println(array8)
	// 指针类型数组的赋值操作
	array9 := [...]*string{new(string), new(string), new(string)}
	*array9[0] = "Red"
	*array9[1] = "Blue"
	*array9[2] = "Green"
	// 指针类型数组的赋值操作
	array10 := [3]*string{}
	array10 = array9
	fmt.Println(array10)
}
