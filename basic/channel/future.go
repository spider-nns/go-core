package main

import (
	"fmt"
	"time"
)

//查询结构体
type query struct {
	sql    chan string
	result chan string
}

func execQuery(q query) {
	go func() {
		//获取输入
		sql := <-q.sql
		//访问数据库
		q.result <- "result from " + sql
	}()
}
func main() {
	//初始化
	q := query{make(chan string, 1), make(chan string, 1)}
	go execQuery(q)
	q.sql <- "select * from table"
	//发送参数
	//执行Query,无需准备参数
	//go execQuery(q)
	time.Sleep(1 * time.Second)
	fmt.Println(<-q.result)
}
