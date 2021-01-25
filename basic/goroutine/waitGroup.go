package main

import (
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var urls = []string{
	"http://www.baidu.com",
	"http://www.163.com",
	"http://www.126.com",
}

func main() {
	for _, url := range urls {
		//每一个url 启动一个 goroutine，同时给 wg 加 1
		wg.Add(1)
		//Launch a goroutine to fetch the url
		go func(url string) {
			//defer wg.Add(-1)
			defer wg.Done()
			//http call
			resp, err := http.Get(url)
			if err == nil {
				println(resp.Status)
			}
		}(url)
	}
	wg.Wait()
}
