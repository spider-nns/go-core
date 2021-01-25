package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	//go fmt.Println(<-ch)
	go func() {
		fmt.Println(<- ch)
	}()
	ch <- 5
	time.Sleep(1 * time.Second)
	//deadlock!
}
//go 语句后面的函数调用，参数会先求值先于调用
//无缓冲的通道必须有接收才能发送
//使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为同步通道
