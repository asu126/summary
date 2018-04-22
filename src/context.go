package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	c2 := make(chan string)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
				if 2 == 2 {
					c2 <- "Wait repo sync end, repo sync successful ..."
				}
			}
		}
	}(ctx)
	// time.Sleep(10 * time.Second)

	select {
	case res := <-c2:
		cancel()
		fmt.Println(res)
	case <-time.After(10 * time.Second):
		fmt.Println("Wait repo sync timeout ...")
		fmt.Println("可以了，通知监控停止")
		cancel()
		return
	}
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	fmt.Println("下载中...")
	time.Sleep(5 * time.Second)
}
