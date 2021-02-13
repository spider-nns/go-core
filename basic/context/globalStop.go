package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg1.Add(1)
	go worker1()
	time.Sleep(time.Second * 3)
	exit = true
	wg1.Wait()
	fmt.Println("over")

}
//全局变量
//1.跨包调用时不容易统一
//2.如果worker中在启动 goroutine,就不太好控制
var wg1 sync.WaitGroup
var exit bool

func worker1(){
	for  {
		fmt.Println("worker")
		time.Sleep(time.Second)
		if exit {
			break
		}
	}
	wg1.Done()
}

