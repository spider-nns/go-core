package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(time.Microsecond * 50)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	select {
	case <-time.After(time.Second * 1):
		fmt.Println("overslept")
		//过期退出
	case <-ctx.Done():
		fmt.Println(ctx.Err())

	}
}

//返回父上下文副本，将 deadline 调整为不迟于 d
//如果父上下文的 deadline 已经早于 d，则 WithDeadline(parent, d) 在语义上等同于父上下文
//当截止日过期时，当调用返回的 cancel 函数时，或者当父上下文的 Done 通道关闭时，返回上下文的 Done 通道将被关闭
//以最先发生的情况为准
