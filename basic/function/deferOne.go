package main

import "fmt"

func main() {
	defer call()
}
func  call(){
	defer func() {fmt.Println("打印前")}()
	defer func() {fmt.Println("打印中")}()
	defer func() {fmt.Println("打印后")}()
	panic("触发异常")
	//最后执行panic
}
