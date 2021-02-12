package main

import (
	"log"
	"os"
)

func init() {
	//设置输出位置，默认是标准错误输出
	file, err := os.OpenFile("go-core.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println("open current log file err:", err)
		return
	}
	//设置标准logger输出选项
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	//设置前缀,方便检索
	log.SetPrefix("[spider]")
	log.SetOutput(file)
}
func main() {
	log.Println("这是一条很普通的日志")
	v := "很普通的"
	log.Printf("这是一条%s\n", v)
	//fatal 系列函数会写入日志信息后调用 os.Exit(1)
	log.Fatalln("这是一条会触发fatal的日志")
	log.Panicln("这是一条会触发panic的日志")
}
