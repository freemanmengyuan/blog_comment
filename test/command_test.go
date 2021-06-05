package test

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"testing"
	"time"
)

// 执行shell命令
func TestCommand(t *testing.T) {
	var (
		cmd *exec.Cmd
		err error
		out bytes.Buffer
	)
	// 生成cmd
	cmd = exec.Command("/bin/bash", "-c", "echo 1;echo 2; ")
	// 捕获输出
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	fmt.Printf("执行输出: %s \n", out.String())
}

/**
 * 执行shell命令，并捕获输出
 */
func TestCommandPrint(t *testing.T) {
	var (
		cmd    *exec.Cmd
		err    error
		output []byte
	)

	// 生成cmd
	cmd = exec.Command("/bin/bash", "-c", "echo 1;echo 2; ls -la")
	// 执行命令，捕获子进程的输出（pipe）
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}

	// 打印子进程的输出
	fmt.Printf("执行输出: %s \n", output)
}

/**
 * 执行一个cmd 让他执行2秒 sleep 2; echo hello;
 * 1秒的时候，我们杀死cmd
 */
type result struct {
	err    error
	output []byte
}

func TestCommandPrintFuncCancel(t *testing.T) {
	var (
		cmd        *exec.Cmd
		ctx        context.Context
		cancelFunc context.CancelFunc
		resultChan chan *result
		res        *result
	)
	/**
	 * 原理:
	 * context chan byte
	 * cancelFunc close (chan byte)
	 */
	resultChan = make(chan *result, 1000)
	// 继承系统的上下文对象, 返回上下文和取消函数
	ctx, cancelFunc = context.WithCancel(context.TODO())
	go func() {
		var (
			err    error
			output []byte
		)
		// 生成cmd
		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 2;echo hello;")
		// select(case <-ctx.Done();)
		// 执行命令，捕获子进程的输出（pipe）
		output, err = cmd.CombinedOutput()
		// 把任务的输出结果传给main协程
		resultChan <- &result{
			err:    err,
			output: output,
		}
	}()
	// 继续执行
	time.Sleep(5 * time.Second)
	// 清空chan byte,中断子协程运行
	cancelFunc()

	// 在main协程中等待子协程的退出，并打印执行结果
	res = <-resultChan
	// 打印子进程的输出
	fmt.Printf("执行输出: %s \n", res.output)
}
