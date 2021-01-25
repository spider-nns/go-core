package main

import "fmt"

const ct int = 5

func main() {
	funcDefine(5)
	v := 5
	println(chvalue(v))
	fmt.Println("v stack out:", v)
	i := 55
	fmt.Println("i stack out:", chpointer(&i))

	fmt.Println("-----------")
	multiParam(1, 2, 3)
	sli0 := []int{1, 2, 3}
	sli1 := make([]int, 3)
	multiParam(sli0...)
	multiParam(sli1...)
	suma(1, 2, 3)
	//形参为不定参数的函数和形参为切片的函数类型不相同
	fmt.Printf("%T\n", suma)
	fmt.Printf("%T\n", sumb)
	f := suma(1, 2, 3, 4)
	println("F", f)
}

func suma(arr ...int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return
}
func sumb(arr []int) (sum int) {
	for i := range arr {
		sum += i
	}
	return
}

func multiParam(param ...int) (sum int) {
	for i, v := range param {
		fmt.Println("add ", i)
		sum += v
	}
	fmt.Println("multi param sum is :", sum)
	return
}

func chpointer(a *int) int {
	*a = *a + 1
	fmt.Println("change a out:", *a)
	return *a
}

func chvalue(v int) int {
	v += 1
	fmt.Println("change v out:", v)
	return v
}

func funcDefine(param int) (cos, sin int) {
	println(param)
	//命名的返回值可以覆盖,必须指定返回
	cos = 1
	sin = 2
	return cos, sin
	//没有覆盖可以简写
	return
}

//1.函数是一种类型，函数类型变量可以作为其他函数的参数或者返回值,也可以直接调用执行
//2.函数支持多值返回[默认错误类型作为最后一个返回值]，支持闭包，支持可变参数
//3.函数生命关键字 func 函数名 参数列表 返回列表 函数体
//4.函数名首字母的大小写决定该函数在其他包的可见性
//5.函数可以没有返回值 默认返回 0,多个相邻的相同类型参数可以使用简写模式
//6.函数支持有名的返回值(1.16 beta1 命名的返回值不能覆盖),参数名就相当于函数体内最外层的局部变量,会被初始化为类型零值，最后 return 可以不带参数名直接返回

// 不支持默认值参数，不支持函数重载，不支持命名函数的嵌套定义，支持嵌套你名函数

//不定参数 param ...type
//1.所有不定参数类型必须相同
//2.不定参数必须是函数的最后一个参数
//3.不定参数在函数体内相当于切片，对切片的操作同样适合对不定参数的操作
//4.不定参数可以接受切片，slice... ,不能接受数组
//5.形参为不定参数的函数和形参为切片的函数类型不相同
