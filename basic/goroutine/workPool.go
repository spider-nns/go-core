package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 100)
	defer close(jobs)
	results := make(chan int, 100)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= 5; j++ {
		fmt.Printf("---> job-produce: %d\n", j)
		jobs <- j
		time.Sleep(2 * time.Second)
	}

	for a := 1; a <= 5; a++ {
		<-results
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("	worker:%d start job %d\n", id, j)
		time.Sleep(3 * time.Second)
		fmt.Printf("	worker:%d end job:%d\n", id, j)
		results <- j * 2
	}
}
