package main

import (
	"runtime"
	"time"
)

func main() {

	//并行是任何时刻、任一粒度的时间内都具备同时执行的能力，最简单的并行就是多机，多台机器同时处理
	//并发是在单位时间内同时进行，规定时间内多个请求都得到执行和处理，强调的是给外界的感觉
	//实际上内部是分时操作的，并发重点在避免阻塞，使程序不会因为一个阻塞而停止处理，单核分时操作系统就是一种并发设计

	//并行具有瞬时性，并行在于执行
	//并发具有过程性，并发在于结构

	//go func(){}
	//go 关键字后面必须跟一个函数


	//goroutine 特性
	//1.go 的执行是非阻塞的，不会等待
	//2.go 后面的函数返回值会被忽略
	//3.调度器不能保证多个 goroutine 的执行次序
	//4.没有父子 goroutine 概念，所有 goroutine 都是被平等调度的
	//5.go 程序在执行时户会单独为main函数创建一个 goroutine，遇到其他 go 关键字时再去创建其他的 goroutine
	//6.go 没有暴露 goroutine id 给用户，所以不能在一个 goroutine 里面显示的操作另一个 goroutine，不过 runtime 包里面提供了一些函数访问和设置 goroutine 的相关信息


	go func() {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
		println(sum)
		time.Sleep(1 * time.Second)
	}()

	println("NumCPU=",runtime.NumCPU())
	println("NumCgoCall=",runtime.NumCgoCall())
	println("NumGoroutine=",runtime.NumGoroutine())
	time.Sleep(5 * time.Second)
}
