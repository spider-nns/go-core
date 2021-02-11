package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		name    string
		age     int
		married bool
	)
	//scan 默认空格分割
	//fmt.Scan(&name, &age, &married)
	//Scanf 根据 format 参数指定的格式去读取由空白符号分割的值保存到传递给本函数的参数中
	//fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
	fmt.Scanln(&name, &age, &married)
	fmt.Printf("scan get name:%s age:%d married:%t\n", name, age, married)
	bufioDemo()

}
func bufioDemo() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("输入内容:")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Printf("%#v\n", text)
}
