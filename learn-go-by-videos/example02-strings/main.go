package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abc world abc"
	// 字符串替换，将字符串 str 中的 abc 替换成 www 替换一次
	// n = -1 表示替换所有
	// n = 0 表示一个也不替换
	result := strings.Replace(str, "abc", "www", -1)
	fmt.Println(result)

	// Count 方法返回字符串 str 中包含 子字符 "abc" 的次数
	count := strings.Count(str, "abc")
	fmt.Println(count)

	// Repeat 将字符串 str 重复 n 次后返回
	resultAfterRepeat := strings.Repeat(str, 3)
	fmt.Println(resultAfterRepeat)

	// 将字符串转为大写
	result = strings.ToUpper(str)
	fmt.Println(result)

	// 将字符串转为小写
	result = strings.ToLower(str)
	fmt.Println(result)
}
