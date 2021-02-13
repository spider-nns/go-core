package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//WithValue 函数将请求作用域的数据与 Context 对象建立关系
//提供的键必须是可比较的，并且不应该是string类型或者其他内置类型，避免使用上下文在包之间发生冲突
//WithValue 定义 自己的类型，为了避免在分配给interface{}时进行分配，上下文键通常具有具体类型 struct{}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*50)
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"),"2021.2.13")
	wg8.Add(1)
	go worker8(ctx)
	time.Sleep(time.Second * 5)
	cancel()//通知子 goroutine 结束
	wg8.Wait()
	fmt.Println("over")
}

type TraceCode string

var wg8 sync.WaitGroup

func worker8(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	value := ctx.Value(key)
LOOP:
	for {
		fmt.Printf("worker8, trace code :%s\n", value)
		time.Sleep(time.Microsecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg8.Done()
}
