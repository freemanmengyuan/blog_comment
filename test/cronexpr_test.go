package test

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"testing"
	"time"
)

/**
 * 执行调度
 */
func TestCronexprDemo1(t *testing.T) {
	var (
		expr *cronexpr.Expression
		err error
		now time.Time
		nextTime time.Time
	)
	// 那一分钟(0-59)，那小时(0-23)，那天（1-31），那月(1-12)，星期几(0-6)

	// 每隔5分钟执行一次
	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}
	// 当前时间
	now = time.Now()
	// 下次调度的时间
	nextTime = expr.Next(now)

	// 等待定时器超时
	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println("被调度了：", nextTime)
	})

	time.Sleep(5*time.Second)
	// fmt.Println(now)
	// fmt.Println(nextTime)
}