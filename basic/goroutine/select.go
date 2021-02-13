package main

import "fmt"

func main() {

	fmt.Println("main over")
	select {}
	//deadlock

}
func selectOne() {
	ch := make(chan int, 1)
	go func(chan int) {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}(ch)
	for i := 0; i < 10; i++ {
		println(<-ch)
	}
}

func selectTwo() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			{
				fmt.Printf("case: %d\n", x)
			}
		case ch <- i:
		default:
			{
				fmt.Println(<-ch)
			}
		}
	}
}

//for循环
func selectThree(ch1, ch2 chan int) <-chan int {
	out := make(chan int, 3)
	go func() {
		defer close(out)
		for {
			select {
			//ch1 为nil或被关闭 !ok
			case v1, ok := <-ch1:
				if !ok {
					ch1 = nil
					continue
				}
				out <- v1
			case v2, ok := <-ch2:
				if !ok {
					ch2 = nil
					continue
				}
				out <- v2
			}

			if ch1 == nil && ch2 == nil {
				break
			}
		}
	}()
	return out
}

//select 多路监听通道
//当监听的通道没有状态是可读或可写的，select 是阻塞的
//没有case 的select 是阻塞的,可用于阻塞main函数
//只要监听的通道中有一个状态是可读或可写的，select 就不会阻塞，而是进入处理就绪通道的分之流程
//如果监听的通道有多个可读或可写的状态，则select 随机读取一个处理

//扇入，扇出
//扇入就是和，当生产者速度很慢时，需要使用扇入技术聚合多个生产者满足消费者，eg: 耗时的加解密服务
//扇出就是分，当消费者速度很慢时，需要使用扇出技术，eg: web服务器并发请求处理
