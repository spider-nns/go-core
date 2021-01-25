package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Printf("Hello,world. 你好，世界! \n")
	iotaTest()
	complexTest()
	stringTest()
	intOut()
}

func intOut() {
	// 十进制
	a := 10
	fmt.Printf("%d \n", a)  // 10
	fmt.Printf("%b \n", a)  // 1010  占位符%b表示二进制

	// 八进制  以0开头
	b := 077
	fmt.Printf("%o \n", b)  // 77

	// 十六进制  以0x开头
	c := 0xff
	fmt.Printf("%x \n", c)  // ff
	fmt.Printf("%X \n", c)  // FF

	//splitter with '_'
	si := 123_456
	fmt.Printf("int splitter with '_' -> %d",si)
}
func stringTest() {
	fmt.Println("-------- check string ---------")
	s := "hello,world"
	fmt.Println("s len is: ", len(s))
	fmt.Println("string append use + ->", "hello"+",world")
	r := s[0]
	fmt.Println("s[0]: ", r)
	//s[1]='a' 不能修改某个字节单元
	byt := []byte(s)
	run := []rune(s)
	fmt.Println("string convert []byte ", byt)
	fmt.Println("string convert []rune ", run)
	sub := byt[6:]
	fmt.Println("sub: ", sub)
	fmt.Println(reflect.TypeOf(sub))
	b := s[6:]
	fmt.Println("b: ", b)
	fmt.Println(reflect.TypeOf(b)) // string 切片操作返回还是 string
	fmt.Println("------range-----")
	//range
	for i := 0; i < len(byt); i++ {
		fmt.Println(byt[i])
	}
	for index, value := range run {
		fmt.Println(index, value)
	}
}

func complexTest() {
	fmt.Println("--------- check complex ---------")
	c0 := complex(3.1, 5)
	fmt.Println(real(c0))
	fmt.Println(imag(c0))
}
func iotaTest() {
	fmt.Println("--------- check iota -------")
	// const 组合使用 iota 会逐行累加
	const x = iota
	const y = iota
	fmt.Println(x)
	fmt.Println(y)
	const (
		c0         = 1 << iota
		c1         = 1 << iota
		c2         = 3         //iota =2
		c3 float64 = 1 << iota //iota=3
	)
	const (
		Iota0 = iota
		Iota1
		Iota2
	)
	fmt.Println(c0)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Printf("iota0: %v\n", Iota0)
	fmt.Printf("iota0: %v\n", Iota1)
	fmt.Printf("iota0: %v\n", Iota2)
}

/**
  内置数据类型标示符号
  数值
		整型 12 不同类型的整型必须进行强制类型转换
			byte[uint8] int int8 int16 int32 int64
            uint uint8 uint16 uint32 uint64 uintptr
		特殊类型
			uint 32位机器上 uint32，64位机器上 uint64
			int  32位机器上 int32，64位机器上 int64
			uintptr 无符号整型，用于存储指针

		浮点型 2
			float32 float64 自动类型推断为 float64
		复数 2
			complex64 complex12  complex64 = complex32 + complex32
			var c0 complex64 = 3.1 + 5i
			三个内置函数处理复数类型
				var c1 = complex(2.1,3) 构造
				a := real(c1) 返回实部
				v := imag(c1) 返回虚部
  字符和字符串型 2
		string rune[int32]
 		可以访问但不能修改字节单元，底层二元数据结构 一个指针指向字节数组起点，一个是长度
		可以转换为 slice，基于字符串创建的切片和原数组指向相同的底层字符数组,一样不能修改
		对字符串的切片操作返回的还是字符串
		字符串尾部不包含 NULL 字符
  错误类型 1
	error
  布尔型 1
	bool 不能和整型数据转换
*/

/**
内置函数 15
make new len cap append
copy delete panic recover close
complex real image Print println
*/

/**
常量标示符 4
true false
iota 初始值 0,一组同时声明,逐行累加
nil
*/
/**
空白标识符
 _
*/
