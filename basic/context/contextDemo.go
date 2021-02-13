package main

import (
	"fmt"
	"sync"
	"time"
)

//推荐以参数的方式显示传递 Context
//Context 作为第一个参数
//给一个函数方法传递 Context，不要传递nil,不知道传递神，就是用 context.TODO（）
//Context 的value相关方法应该传递请求域的必要数据，不应该用于传递可选参数
//Context 是线程安全的
func main() {
	wg.Add(1)
	go worker()
	wg.Wait()
	fmt.Println("over")
}
var wg sync.WaitGroup

func worker(){
	for  {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}
	wg.Done()
}