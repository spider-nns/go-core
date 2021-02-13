package main

import (
	"fmt"
	"sync"
	"time"
)
//需要维护 channel，
func main() {
	var exitChan = make(chan struct{})
	wg2.Add(1)
	go worker2(exitChan)
	time.Sleep(time.Second * 3)
	exitChan <- struct{}{}
	close(exitChan)
	wg2.Wait()
	fmt.Println("over")
}

var wg2 sync.WaitGroup

func worker2(exitChan chan struct{}) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-exitChan:
			break LOOP
		default:

		}
	}
	wg2.Done()
}
