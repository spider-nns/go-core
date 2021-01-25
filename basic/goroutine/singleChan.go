package main

import "fmt"

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}

func counter(out chan<- int) {
	defer close(out)
	for i := 0; i < 100; i++ {
		out <- i
	}
}

func squarer(out chan<- int, in <-chan int) {
	defer close(out)
	for i := range in {
		out <- i * i
	}
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

//在函数传参及任何赋值操作中可以将双向通道转换为单向通道，但反过来是不可以的

//channel  nil 			非空  				空的					满了 						没满
//接收     阻塞		   接收值			    阻塞					接收值						接收值
//发送     阻塞		   发送值 			   发送值 				阻塞							发送值
//关闭	  panic  关闭成功，读完数据后返回零值   关闭成功，返回零值		关闭数据，读完数据后返回零值		关闭成功，读完数据后返回零值