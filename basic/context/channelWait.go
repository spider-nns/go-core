package main

import "fmt"

func main() {
  channel := make(chan bool,10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			channel <- true
		}(i)
	}
	for i := 0; i < 10; i++ {
		<- channel
	}
}
