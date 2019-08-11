package main

import (
	"fmt"
	"strings"
)

/**
 * strings 来操作字符串
 */
func main() {
	var (
		url  string
		path string
	)
	// 输入，两个参数输入时中间加一个空格
	// fmt.Scanf("%s%s", &url, &path)
	url = "google.com/user/login"
	path = "c:/user/alan"

	url = urlProcess(url)
	path = pathProcess(path)

	fmt.Println(url)
	fmt.Println(path)
}

// 判断 url 是否以 http:// 开头，如果不是则添加 http://
func urlProcess(url string) string {
	result := strings.HasPrefix(url, "http://")
	if !result {
		url = fmt.Sprintf("http://%s", url)
	}

	return url
}

// 判断 path 是否以 / 结尾，如果不是，则加上 /
func pathProcess(path string) string {
	result := strings.HasSuffix(path, "/")
	if !result {
		path = fmt.Sprintf("%s/", path)
	}
	return path
}
