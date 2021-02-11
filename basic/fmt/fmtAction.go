package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	format()
}
func format() {
	print("print 直接输出内容\n")
	fmt.Printf("printf 支持格式化输出字符串%s\n", "sdfsf")
	println("println函数会在输出内容结尾添加一个换行符")
	fmt.Fprintln(os.Stdout, "Fprint 会将内容输出到一个 io.writer 接口类型的变量中，通常往文件中写入内容")
	file, err := os.OpenFile("/Users/spider/Desktop/a.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错", err)
		return
	}
	name := "郑州小王子"
	fmt.Fprintf(file, "往文件中写入信息:%s\n", name)
	print(fmt.Sprintln("sprint系列会把传入的数据生成并返回一个字符串"))
	println(fmt.Sprint("sprint系列会把传入的数据生成并返回一个字符串"))
	println(fmt.Sprintf("content:%s", "sprint系列会把传入的数据生成并返回一个字符串"))
	err = fmt.Errorf("这是一个错误")
	println(err)
	err = errors.New("原始错误e")
	w := fmt.Errorf("Wrap 了一个错误:%w", err)
	println(w)
	//通用占位符 %printf
	/**
	       %v 值的默认格式表示
		   % + v ，结构体会添加字段名
	   	   %#v 值的go语法表示
	       %T 打印值的类型
		   %% %
	# 添加前导
	 布尔值 %t

	*/

	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	fmt.Printf("%t\n", true)
	o := struct{ name string }{"小王子"}
	fmt.Printf("%+v\n", o)
	fmt.Printf("%v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")

	fmt.Printf("%%b 二进制:%b\n", 8)
	fmt.Printf("%%d 十进制:%d\n", 8)
	fmt.Printf("%%o 八进制:%o\n", 8)
	fmt.Printf("%%x 十六进制a-f:%x\n", 65)
	fmt.Printf("%%x 十六进制A-F:%X\n", 65)
	fmt.Printf("%%c unicode:%c\n", 'a')
	fmt.Printf("%%q 单引号括起来go语法字符字面量，必要时会:%q\n", 1)

	println("----- float and complex format -----")
	//float and complex
	fmt.Printf("%%b 无小数部分、二进制指数的科学计数法:%b\n", 123.45)
	fmt.Printf("%%e 科学计数法:%b\n", 123.45)
	fmt.Printf("%%E 科学计数法:%E\n", 123.45)
	fmt.Printf("%%f 有小数但无指数部分:%f\n", 123.45)
	fmt.Printf("%%F 有小数但无指数部分:%F\n", 123.45)
	//g G 精度为所有数字的总和 ⚠️
	fmt.Printf("%%g 根据实际情况采用%%e或%%f格式:%.4g\n", 123.45)
	fmt.Printf("%%G 根据实际情况采用%%E或%%F格式:%.4G\n", 123.45)

	println("----- string and byte")
	//string and byte
	var str = "2020 大年三十写代码"
	fmt.Printf("%%s 直接输出字符串或者 []byte:%s\n", str)
	fmt.Printf("%%q 该值对应的双引号括起来的go语法字符串字面值，必要时户籍采用安全的转义表示:%q\n", str)
	fmt.Printf("%%x 每个字节用两字符十六进制表示 a-f:%x\n", str)
	fmt.Printf("%%X 每个字节用两字符十六进制数表示 A-F:%X\n", str)

	println("----- pointer -----")
	fmt.Printf("%%p 指针表示为十六进制 并加上前导的 0x:%p\n", &str)
	fmt.Printf("%%p 指针表示为十六进制 并加上前导的 0x:%#p\n", &str)

	println("----- 宽度标识符 -----")
	//宽度通过一个紧跟在百分号后面的十进制数指定，如果未指定宽度，表示值时除必须之外不做填充
	var f = 123456.78
	fmt.Printf("%%f    默认宽度、精度 :%f\n", f)
	fmt.Printf("%%9f   宽度9、默认精度:%f\n", f)
	fmt.Printf("%%.2f  默认宽度、精度2:%f\n", f)
	fmt.Printf("%%9.2f 宽度9、精度2  :%f\n", f)
	fmt.Printf("%%9.f  宽度9、默认精度:%f\n", f)

	//
	println("----- -----")
	fmt.Printf("%%+ 输出数值的正负号:%+b\n", -1)
	println("origin pointer address: ", &str)
	fmt.Printf("%%- 左对其，右侧填充:%-X\n", &str)
	fmt.Printf("%%+ 总是输出数值的正负号，对%%+q 会输出全是 ASCII 字符的输出%+q\n", 65)
	fmt.Printf("%%+ 总是输出数值的正负号，对%%+q 会输出全是 ASCII 字符的输出%q\n", 97)
	fmt.Printf("%%'' 正数前加空格,负数前加负号，对%%+q 会输出全是 ASCII 字符的输出%+X\n", -97)
	fmt.Printf("%%0 使用0填充前导，%0b\n", -1)

}

