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
	fmt.Println("当前时间：", now)
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

/**
 * 调度多个任务
 */
// 代表一个任务
type CronJob struct {
	expr *cronexpr.Expression
	nextTime time.Time
}
func TestCronexprMoreJob(t *testing.T) {
	var (
		cronJob *CronJob
		expr *cronexpr.Expression
		now time.Time
		scheduleTable map[string]*CronJob
	)
	// 当前时间
	now = time.Now()
	// 分配空间
	scheduleTable = make(map[string]*CronJob)

	// 定义两个cronjob
	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr: expr,
		nextTime: expr.Next(now),
	}
	// 将任务注册到调度表
	scheduleTable["job1"] = cronJob

	expr = cronexpr.MustParse("*/10 * * * * * *")
	cronJob = &CronJob{
		expr: expr,
		nextTime: expr.Next(now),
	}
	// 将任务注册到调度表
	scheduleTable["job2"] = cronJob

	// 启动一个协程进行调度
	go func() {
		var (
			jobName string
			cronJob *CronJob
			now time.Time
		)
		// 不停的检查调度表
		for {
			now = time.Now()
			for jobName, cronJob = range scheduleTable {
				// 判断是否过期
				if cronJob.nextTime.Before(now) || cronJob.nextTime.Equal(now) {
					// 启动协程执行任务
					go func(jobName string) {
						fmt.Println("执行调度", jobName)
					}(jobName)
					// 计算下次调度的时间
					cronJob.nextTime = cronJob.expr.Next(now)
					fmt.Println("计算出下次调度的时间是：", cronJob.nextTime, jobName)
				}
			}
			// 暂停100ms避免cpu占满
			// time.Sleep(100*time.Millisecond) // 睡眠100毫秒
			select {
			case <- time.NewTimer(100 * time.Millisecond).C:
			}
		}
	}()
	time.Sleep(100 * time.Second)
}

func TestSimple(t *testing.T) {
	arr := [5]int{1,2,3,4,5}
	// 启动一个协程进行测试
	go func() {
		for {
			for k, v := range arr {
				fmt.Println(k, v)
			}
			time.Sleep(100*time.Millisecond) // 睡眠100毫秒
		}
	}()

	time.Sleep(100 * time.Second)
}

func TestSelect(t *testing.T) {
	fmt.Println("hello world")

	select {
	case <- time.NewTimer(1000 * time.Millisecond).C:
	}
}