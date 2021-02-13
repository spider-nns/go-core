package main

import (
	"fmt"
	"sync"
)

func main() {

	//waitGroup 计数器不能为空
	//waitGroup 不是引用类型，需要传地址
	var wg4 sync.WaitGroup
	wg4.Add(10)
	for i := 0; i < 10; i++ {
		go g(i, &wg4)
	}
	wg4.Wait()
	fmt.Println("over")
}
func g(i int, wg *sync.WaitGroup){
	fmt.Println(i)
	wg.Done()
}
