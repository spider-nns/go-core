package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	getBaidu()
}

func basicHttp() {
	resp0, err := http.Get("www.baidu.com")
	if err != nil {
		//handler error
	}
	//程序在使用完response后必须关闭
	defer resp0.Body.Close()
	resp1, err := http.Post("www.baidu.com", "image/jpeg", nil)
	defer resp1.Body.Close()
	resp2, err := http.PostForm("www.baidu.com", url.Values{"key": {"value"}, "usernamee": {"password"}})
	defer resp2.Body.Close()

}
func getBaidu() {
	get, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Printf("get failed, err:", err)
		return
	}
	defer get.Body.Close()
	bodyAll, err := ioutil.ReadAll(get.Body)
	if err != nil {
		fmt.Printf("read from get.Body failed, err:%v\n", err)
		return
	}
	fmt.Print(string(bodyAll))
}

func getWithParam() {
	apiUrl := "http://127.0.0.1:9090/get"
	params := url.Values{}
	params.Set("name", "spider")
	params.Set("msg", "2020.2.12，write code with goland")
	uri, err := url.ParseRequestURI(apiUrl)
	fmt.Printf("request uri is:%s\n", uri)
	if err != nil {
		fmt.Printf("parse URL requestUrl failed, maybe it`s not a useful uri， err:%v\n", err)
	}
	uri.RawQuery = params.Encode()
	fmt.Printf("request param is:%s\n", uri)
	get, err := http.Get(uri.String())
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	defer get.Body.Close()
	all, err := ioutil.ReadAll(get.Body)
	if err != nil {
		fmt.Printf("get reso failed, err:%v\n", err)
		return
	}
	println(string(all))

}
func getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	param := r.URL.Query()
	fmt.Println(param.Get("name"))
	fmt.Println(param.Get("msg"))
	answer := `{"code":"200","msg":""}`
	w.Write([]byte(answer))
}

func postHttp() {
	url := "http://127.0.0.1:9091/post"
	//表单
	//contentType := "application/x-www-form-urlencoded"
	//data := "name-spider&age=24"
	//json
	contentType := "application/json"
	data := `{"name":"spider"}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed ,err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	respAll, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed , err:%v\n", err)
		return
	}
	fmt.Printf(string(respAll))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	//form
	r.ParseForm()
	fmt.Println(r.PostForm)
	fmt.Println(r.PostForm.Get("name"))
	//json
	all, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request.body failed,err:%v\n", err)
		return
	}
	fmt.Println(string(all))
	answer := `{"code":200,"msg":""}`
	w.Write([]byte(answer))

}

//自定义client
func customerClient() {
	//管理头域、重定向策略、...
	client := http.Client{CheckRedirect: nil}
	client.Get("")
	request, _ := http.NewRequest("GET", "", nil)
	request.Header.Add("IF-NONE-Match", `W/"wyzzy"`)
	resp, _ := client.Do(request)
	println(resp)
}

//自定义 transport
func customerTransport() {
	//管理 代理、tls配置、keep-alive、压缩、...
	//client transport 类型都可以安全的被多个 goroutine 同时使用,一次建立、尽量复用
	tr := &http.Transport{TLSClientConfig: &tls.Config{RootCAs: nil},
		DisableCompression: true}
	client := &http.Client{Transport: tr}
	get, _ := client.Get("")
	fmt.Println(get)
}
