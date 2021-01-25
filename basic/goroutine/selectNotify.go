package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func main() {

	done := make(chan struct{})
	ch := GenerateIntA(done)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//发送通知，告诉生产者停止生产

	close(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//此时生产者已经退出
	println("NumGoroutine=", runtime.NumGoroutine())
}
func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	lable:
		for {
			select {
				case ch <- rand.Int():
				//读取已经关闭的通道不会引起阻塞，也不会导致panic,而是立即返回该通道存储类型的零值
				//关闭 select 监听伤的某个通道能使  select 立即感知到这种通知，然后进行相应的处理
				case <-done:
					break lable
			}
		}
		close(ch)
	}()
	return ch
}
