package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	doCall(ctx)
}

type respData struct {
	resp *http.Response
	err  error
}

func doCall(ctx context.Context) {
	transport := http.Transport{
		//请求频繁可定义全局的 client 对象并启用长链接
		//请求不频繁使用短链接
		DisableKeepAlives: true}
	client := http.Client{
		Transport: &transport,
	}
	respChan := make(chan *respData, 1)

	req, err := http.NewRequest("GET", "http://127.0.0.1:8088/", nil)
	if err != nil {
		fmt.Printf("new request failed ,err :%v\n", err)
		return
	}
	//超时的 ctx 创建新的 client
	req = req.WithContext(ctx)
	var wg9 sync.WaitGroup
	wg9.Add(1)
	defer wg9.Wait()
	go func() {
		resp, err := client.Do(req)
		fmt.Printf("client.do resp:%v，err:%v\n", resp, err)
		rd := &respData{
			resp: resp,
			err:  err,
		}
		respChan <- rd
		wg9.Done()
	}()
	select {
	case <-ctx.Done():
		fmt.Println("call api timeout")
	case result := <-respChan:
		fmt.Println("call server api success")
		if result.err != nil {
			fmt.Printf("call server api failed:%v\n", result.err)
			return
		}
		defer result.resp.Body.Close()
		data, _ := ioutil.ReadAll(result.resp.Body)
		fmt.Printf("resp:%v\n", string(data))
	}

}
