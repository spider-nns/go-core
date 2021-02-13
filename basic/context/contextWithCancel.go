package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	for n := range gen(ctx){
		fmt.Println(n)
		if n==5 {
			break
		}
	}
}

//WithCancel 返回带有新 Done 通道的父节点的副本
//当调用返回的 cancel 函数或当关闭父上下文的Done通道时，将关闭返回上下文的 Done 通道
func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
