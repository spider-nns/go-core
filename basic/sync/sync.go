package main

import (
	"fmt"
	"sync"
)

type appType int32
const(
	ONE appType=iota
	TWO
)

type Counter struct {
	sync.Mutex
	Count int
}
func main() {
	var c Counter
	c.Lock()
	defer c.Unlock()
	c.Count++
	foo(c)
}
func foo(c Counter){
	//锁传递
	c.Lock()
	defer c.Unlock()
	fmt.Println(TWO)
}
