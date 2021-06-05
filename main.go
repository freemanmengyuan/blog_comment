package main


import (
	"fmt"
	"github.com/gorhill/cronexpr"
)

/**
 * 临时测试 使用
 */
func main() {
	var (
		expr *cronexpr.Expression
		err error
	)
	// 那一分钟(0-59)，那小时(0-23)，那天（1-31），那月(1-12)，星期几(0-6)

	// 每一分钟执行一次
	if expr, err = cronexpr.Parse("* * * * *"); err != nil {
		fmt.Println(err)
		return
	}

	// 每隔5分钟执行一次
	if expr, err = cronexpr.Parse("*/5 * * * *"); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("import success")
	expr = expr
}