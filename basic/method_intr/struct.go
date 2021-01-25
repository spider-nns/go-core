package main

import (
	"fmt"
	"reflect"
)

func main() {

	checkType()
	stringExchangeSlice()
}

func defineType() {
	type T1 string
	var t1 T1
	t1 = "t1"
	fmt.Printf("t1 type: %v\n", reflect.TypeOf(t1).Kind())
	type T2 T1
	var t2 T2
	fmt.Printf("t1 type: %v\n", reflect.TypeOf(t2).Kind())
	type T3 []string
	var t3 T3
	fmt.Printf("t1 type: %v\n", reflect.TypeOf(t3).Kind())
	type T4 T3
	var t4 T4
	fmt.Printf("t1 type: %v\n", reflect.TypeOf(t4).Kind())
	type T5 []T1
	var t5 T5
	fmt.Printf("t1 type: %v\n", reflect.TypeOf(t5).Kind())
	type T6 T5
	var t6 T6
	fmt.Printf("t1 type: %v\n", reflect.TypeOf(t6).Kind())
}

//类型系统
//命名类型，未命名类型
//命名类型
//类型可以通过标识符表示 基本类型 + 用户自定义类型
//预声明的20 个基本类型，预声明类型

//未命名类型
//一个类型由 预声明类型、关键字、操作符组成 称作未命名类型/类型字面量
//复合类型都是未命名类型

//底层类型 underlying type
//1.预声明类型、类型字面量 底层类型是他们自身
//2.自定义类型 type newType oldType ，
//  newType 是逐层递归向下查找的,直到查找到的 ordType 是预声明类型或类型字面量为止

//type T1 string     string
//type T2 T1 		 string
//type T3 []string   []string
//type T4 T3		 []string
//type T5 []T1		 []T1
//type T6 T5		 []T1

//底层类型在类型赋值和类型强制转换使用

//类型相同
//1.两个命名类型相同的条件是两个类型声明的语句完全相同
//2.命名类型和未命名类型永远不相同
//3.两个未命名类型相同的条件是他们的类型声明字面量结构相同
//4.通过类型别名语句声明的两个类型相同

//类型直接赋值 var t1 T2 = t1
//如果类型 T1 变量 t1 可以赋值给类型 T2 变量 t2
//1. T1 T2 类型相同
//2.具有相同的底层类型，并且至少又一个是未命名类型
//3.T2 是接口类型， T1是基本类型，T1的方法集是 T2方法集的超集
//4.T1 和 T2 都是通道类型，拥有相同的元素类型,并且至少有一个是未命名类型
//5.t1 是预声明标识符，T2 是 pointer,function,slice,map,channel,interface 中的一个
//6.t1 是一个字面常量值，可以用来表示类型 T的值

//类型强制转换
// var a T = (T)(b)
//非常量类型的变量 x 可以强制转化并传递给类型 T
//1.x可以直接赋值给T类型变量
//2.x的类型和 T 具有相同的底层类型
//3.x的类型和 T 都是未命名的指针类型，并且指针指向的类型具有相同的底层类型
//4.x的类型和 T 都是整型，或者都是浮点数
//5.x的类型都是复数类型
//6.x是整数值或[]byte类型的值，T是string类型
//7.x是字符串，T是[]byte或者[]rune
//strconv string 数字
//unsafe 指针和 integer 直接转换

func stringExchangeSlice() {
	println("----- string exchange slice -----")
	s := "hello,世界"
	var a []byte
	a = ([]byte)(s)
	var b string
	b = string(a)
	var c []rune
	c = []rune(s)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}

type Map map[string]string
type iMap Map
type slice []int

func (m Map) Print() {
	for _, v := range m {
		fmt.Println(v)
	}
}

func (m iMap) Print() {
	for _, m := range m {
		fmt.Println(m)
	}
}

func (s slice) Print() {
	for _, v := range s {
		fmt.Println(v)
	}
}

func checkType() {
	mp := make(map[string]string, 10)
	mp["hi"] = "hi"
	var ma Map = mp
	ma.Print()
	//有相同的底层类型，但是都是命名类型，没有未命名类型
	//强制类型转换
	var im iMap = (iMap)(ma)
	im.Print()

	var i interface{ Print() } = ma
	i.Print()

	s1 := []int{1, 2, 3}
	var s2 slice
	s2 = s1
	s2.Print()

}
