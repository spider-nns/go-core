package main

import "time"

func main() {
	c := make(chan int)
	go func(c chan int) {
		println("g")
		o :=<-c
		println(o)
	}(c)
	time.Sleep(1 * time.Second)
	println("m")
	c <- 1

	//fmt.Println(<-c)
}
