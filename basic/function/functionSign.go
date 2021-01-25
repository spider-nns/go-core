package main

import (
	"fmt"
	"math"
)

func main() {
	printFuncSign(5)
	do(calculation, 3, 4)
	//匿名函数
	sum := func(a, b int) int {
		return a + b
	}
	fmt.Printf("匿名函数赋值给变量作为参数被调用返回值为函数，out:%d\n", doInput(sum, 1, 2)(1, 2))

}

//函数作为参数
func doInput(f func(int, int) int, a, b int) func(a, b int) int {
	return f
}
func printFuncSign(i int) {
	fmt.Printf("printFuncSign 函数签名为:%T\n", printFuncSign)
}

//type 定义函数类型
type op func(int, int) float64

func calculation(a, b int) (c float64) {
	c = math.Sqrt(float64(a*a + b*b))
	fmt.Printf("c is %v\n", c)
	return
}

//函数参数为函数类型
func do(f op, a, b int) float64 {
	return f(a, b)
}

//1.函数类型又叫函数签名，就是函数定义首行去掉函数名和参数名、{
//2.两个函数类型相同的条件是 拥有相同的参数列表，返回值列表(列表元素的次序，个数和类型都相同)
//3.type 可以定义函数类型

//匿名函数
//可以看作函数字面量，所有直接使用函数类型变量的地方都可以由匿名函数代替
//匿名函数可以直接赋值给函数变量,可以作为实参，也可以作为返回值,还可以直接被调用
