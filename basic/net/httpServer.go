package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func main() {
	defaultServer()
}

//ListenAndServe 使用指定的监听地址和处理器启动一个 HTTP 服务端
//处理器参数通常是 nil，表示使用包变量 DefaultServeMux 作为处理器
//handle 和 handleFunc 函数可以向 DefaultServeMux 添加处理器
func defaultServer() {
	http.HandleFunc("/foo", fooHandler)
	http.HandleFunc("/bar", func(respWrite http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(respWrite, "hello, %q", html.EscapeString(request.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":9000", nil))
}

//管理服务端
func serverKeeper() {
	s := &http.Server{
		Addr:           ":9000",
		//Handler:        fooHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//...
	log.Fatal(s.ListenAndServe())
}

func fooHandler(write http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(write, "hello, 2021.2.12")
}
