package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
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

// strconv 包实现了基本数据类型和其他字符串表示的转换
func main() {
	//atoi 字符串 -> int 类型
	//itoa int -> 字符串

	s1 := "2021"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		log.Fatalf("string:%s convert int fail\n", s1)
	} else {
		log.Printf("type:%T value:%#v\n", i1, i1)
	}

	i2 := 2021
	itoa := strconv.Itoa(i2)
	log.Println(itoa)

	//parse 系列
	log.Printf(
		"parseBool 接收 1，0，t，f，T，F，true，false，True，False，TRUE，FALSE")
	parseBool, err := strconv.ParseBool("1")
	fmt.Println(parseBool)

	//parseInt 接收正负号 parseUnit 不接受正负号，用于无符号整型
	//base 进制 2-36 base 为0 前导判断
	//bitSize 结果必须能无溢出赋值的整数类型
	parseInt, err := strconv.ParseInt("-2021", 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parseInt) //127

	//parseFloat bitSize 32->float32
	float, err := strconv.ParseFloat("2021.212", 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(float)

	//format系列,将给定的类型数据格式化成string类型数据的功能
	//FormatBool、FormatInt、FormatUnit
	fmt.Printf("FormatBool:%s\n", strconv.FormatBool(true))
	fmt.Printf("FormatInt:%s\n", strconv.FormatInt(2021, 36))
	fmt.Printf("FormatUnit:%s\n", strconv.FormatUint(2021, 36))
	//bitSize f 来源类型，32，64，
	//fmt 表示格式
	//'f' -ddd.dddd
	//'b' -ddddp+_ ddd 指数为二进制
	//'e' -d.dddde+-dd 十进制指数
	//'E' -d.ddddE+-dd 十进制指数
	//'g' 指数很大时用'e'格式
	//'G' 指数很大时用'E'格式
	//prec 控制精度，排除指数部分,对 'f' 'e' 'E' 表示小数点后的数字个数,对 'g' 'G' 控制总的数字个数
	//prec 为 -1 代表使用最少数量的、但又必须的数字来表示f
	fmt.Printf("FormatFloat:%s\n", strconv.FormatFloat(2021.2, 'f', -1, 32))

	//
	//isPrint 是否时可打印的 unicode.isPrint r 必须时 字母(广义),数字、标点、符号、ASCII空格
	fmt.Printf("isPrint 是否可打印：%t\n", strconv.IsPrint('a'))
	fmt.Printf("isPrint 是否可打印：%t\n", unicode.IsPrint(','))
	//CanBackquote 	返回一个字符串是否可以不被修改的表示为一个单行的、没有空格和tab之外控制字符的反引号字符串
	fmt.Printf("CanBackquote 字符串是否可以不被修改的表示为一个单行、没有空格和tab之外的控制字符串的反引号字符串:%t\n", strconv.CanBackquote("a"))

}
