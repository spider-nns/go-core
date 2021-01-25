package main

import (
	"runtime"
)

func main() {
	//获取当前 GOMAXPROCS 值， n >1 设置值，否则查询当前值
	println("GOMAXPROCS=", runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(10)
	println("GOMAXPROCS=", runtime.GOMAXPROCS(0))

	go func() {
		defer func() {
			println("anonymous func goroutine over，exec Goexit")
		}()
		println("---> before Goexit")
		runtime.Goexit()
		println("after Goexit")
	}()

	//time.Sleep(3 * time.Second)
	runtime.Gosched()
}

//goroutine 特性
//1.go 的执行是非阻塞的，不会等待
//2.go 后面的函数返回值会被忽略
//3.调度器不能保证多个 goroutine 的执行次序
//4.没有父子 goroutine 概念，所有 goroutine 都是被平等调度的
//5.go 程序在执行时户会单独为main函数创建一个 goroutine，遇到其他 go 关键字时再去创建其他的 goroutine
//6.go 没有暴露 goroutine id 给用户，所以不能在一个 goroutine 里面显示的操作另一个 goroutine，不过 runtime 包里面提供了一些函数访问和设置 goroutine 的相关信息

//GOMAXPROCS 设置或查询可以并发执行的 goroutine 数目
//Goexit 结束当前 goroutine 运行
//Gosched 放弃当前调度执行的机会，将当前 goroutine 放到队列中等待下次被调度
