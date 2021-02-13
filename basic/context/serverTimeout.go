package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		panic(err)
	}
}
func indexHandler(resp http.ResponseWriter, request *http.Request) {
	number := rand.Intn(2)
	fmt.Println(number)
	if number == 0 {
		time.Sleep(time.Second * 3)
		fmt.Fprintf(resp, "slow response")
		return
	}
	fmt.Fprint(resp, "quick response")
}
