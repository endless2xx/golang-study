package main

import (
	"fmt"
	"time"

	"github.com/gorhill/cronexpr"
)

func main() {
	var (
		expr *cronexpr.Expression
		err  error
	)
	if expr, err = cronexpr.Parse("*/5 * * * *"); err != nil {
		fmt.Println(err)
		return
	}
	// 当前时间
	now := time.Now()
	fmt.Println(now)
	// 下次调度时间
	nextTime := expr.Next(now)
	fmt.Println(nextTime)
}
