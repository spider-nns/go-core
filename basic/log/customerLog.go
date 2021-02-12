package main

import (
	"log"
	"os"
)

func init() {
}
func main() {
	logger := log.New(os.Stdout, "<New>", log.Llongfile|log.Ldate|log.Ltime)
	logger.Println("自定义的logger记录日志")
}
