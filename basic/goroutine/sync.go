package main

import (
	"fmt"
	"sync"
	"time"
)

//竞态问题
var x int64
var wg1 sync.WaitGroup
var lock sync.Mutex

func main() {
	wg1.Add(2)
	go add()
	go add()
	wg1.Wait()
	time.Sleep(time.Second)
	fmt.Printf("x is：%d\f", x)
}

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg1.Done()
}
