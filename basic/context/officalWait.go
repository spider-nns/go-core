package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)
//Go1.7 引入，简化对于处理单个请求的多个 goroutine 之间与请求域的数据、取消信号、截止时间等操作
//对服务器传入的请求应该创建上下文，而对服务器的传出调用应该接受上下文
//他们之间的Hanns户调用链必须传递上下文，或者使用 WithCancel、WithDeadline、withTimeout、WithValue 创建派生上下文
//当一个上下文被取消时，派生的所有上下文也被取消
func main() {
	//backGround TODO 是 Context 树结构的最顶层的 Context ,本质上都是 emptyCtx 结构体类型
	//是不可取消、没有设置截止时间、没有携带任何值的 Context
	ctx, cancelFunc := context.WithCancel(context.Background())
	wg5.Add(1)
	go worker5(ctx)
	time.Sleep(time.Second * 3)
	//通知子 goroutine 结束
	cancelFunc()
	wg5.Wait()
	fmt.Println("over")

}

var wg5 sync.WaitGroup

func worker5(ctx context.Context) {
	go worker6(ctx)
LOOP:
	for {
		fmt.Println("worker5")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
	wg5.Done()
}

func worker6(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker6")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}

}
