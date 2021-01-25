package main

import "runtime"

func main() {
	c := make(chan  struct{})
	c1 := make(chan  int,10)
	go func(i chan struct{},j chan int) {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
			c1 <- i
		}
		close(c1)
		println(sum)
		//写通道
		c <- struct{}{}

	}(c,c1)
	println("NumGoroutine=",runtime.NumGoroutine())

	println("NumGoroutine=",runtime.NumGoroutine())
	//此时 c1 通道已经关闭， 但是还是可以继续读取
	for v := range c1{
		println(v)
	}
	//读通道,通过通道进行同步等待
	<-c
}

//chan channel 有类型的管道，通过通信来共享内存 chan dataType
//make(chan dataType) 创建存放 dataType 的无缓冲的管道 len、cap 都是 0
//make(chan dataType,10) 创建一个有 10 个缓冲的通道，存放 dataType 类型元素， len 代表没有被读取的元素数，cap 代表整个通道的容量
//无缓冲的通道可以用于通信，也可以用于两个 goroutine 的同步
//有缓冲的通道主要用于通信

//goroutine 运行结束后，写到缓冲通道里面的数据不会消失，可以缓冲和适配两个 goroutine 处理速率不一致的情况
//缓冲通道和消息队列类似，有削峰和增大吞吐量的功能


//操作不同状态的chan 会引发三种行为
//panic
//1.向已经关闭的通道写数据
//2.重复关闭通道的通道
//阻塞
//1.向为初始化的通道写数据或读数据都会导致当前goroutine 的永久阻塞
//2.向缓冲区已满的通道写入数据会导致 goroutine 阻塞
//3.通道中没有数据，读取改通道会导致 goroutine 阻塞
//非阻塞
//1.读取已经关闭的通道不会引发阻塞，而是立即返回通道元素类型的零值，可以使用command,ok 语法判断通道是否已经关闭
//2.向有缓冲且没有满的通道读/写不会引发阻塞
