package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main(){
	//50 ms timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*50)
	wg6.Add(1)
	go worker7(ctx)
	time.Sleep(time.Second * 5)
	cancel()//通知子 goroutine 结束
	wg6.Wait()
	fmt.Println("over")
}

var wg6 sync.WaitGroup

func worker7(ctx context.Context){
LOOP:
	for  {
		fmt.Println("db connecting ...")
		time.Sleep(time.Microsecond * 10)
		select {
			case <- ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg6.Done()
}