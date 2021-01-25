package main

import (
	"fmt"
	"reflect"
)

func main() {

	println(pointTest())
	arrayTest()
	sliceTest()
	mapTest()
	structTest()
}

func structTest() {
	fmt.Println("------ struct test ------")
	//struct 类型可以任意，存储空间连续,按照字段生命时的顺序存放
	//形式有两种,一种是 struct 类型字面量，另一种是 type 声明的自定义 struct 类型
	//字面量
	s := struct {
		filedName string
	}{}
	s.filedName = "struct?"
	fmt.Printf("%v\n", s)
	//自定义类型
	type person struct {
		name string
		age  int
	}
	type student struct {
		number int
		data   *person
	}
	p := &person{
		name: "nns",
		age:  23,
	}
	stu := student{
		number: 1,
		data:   p,
	}
	fmt.Println(*p)
	fmt.Println(stu.number, *stu.data)
}

func mapTest() {
	fmt.Println("------ check map ------")
	//map[K]T k 任意类型 T 值类型
	//1.create
	ma := map[string]int{"a": 1, "b": 2}
	fmt.Println(ma["a"])
	fmt.Println(ma["A"])
	//2.make
	mp1 := make(map[int]string)
	mp2 := make(map[int]string, 10)
	mp1[1] = "golang"
	mp2[2] = "java"
	fmt.Println(mp1)
	fmt.Println(mp2)
	//3.update
	mp1[1] = "go"
	println(len(mp2))
	delete(mp2, 2)
	mp1[2] = "gogo"
	//3.range 不保证顺序 并发不保证安全
	for k, v := range mp1 {
		fmt.Println("key:=", k, ",value:=", v)
	}
	//不能直接修改map value 内元素的值，如果想修改，必须整体赋值

}

func sliceTest() {
	fmt.Println("----- check slice ----- ")
	//变长数组,引用类型，指针指向数组，{arrPoint,len,cap}
	//1.创建 由数组创建，内置函数创建make([]Type,cap)
	arr := [...]int{1, 2, 3, 4, 5}
	arr1 := [...]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3}
	fmt.Printf("arr type is:%v\n", reflect.TypeOf(arr))
	fmt.Printf("arr1 type is:%v\n", reflect.TypeOf(arr1))
	//同种类型和长度的数据类型相同
	fmt.Printf("arr1 type == arr type ?:%v\n", reflect.TypeOf(arr1) == reflect.TypeOf(arr))
	fmt.Printf("arr1 type == arr2 type ?:%v\n", reflect.TypeOf(arr1) == reflect.TypeOf(arr2))
	//数组创建
	s0 := arr[0:]
	fmt.Printf("%v\n", s0)
	//make
	s1 := make([]int, 10)
	fmt.Printf("%v\n", s1)
	//2.len(slice)，cap(slice)
	fmt.Println("len(slice) get len:", len(s1))
	fmt.Println("cap(slice) get cap:", cap(s1))
	//3.append 追加元素，copy() 复制长度最小的
	s1a := append(s1, 6)
	fmt.Println("s1a type is:", reflect.TypeOf(s1a))
	fmt.Printf("%v\n", s1a)
	dst := []int{1, 2, 3}[0:]
	fmt.Printf("dst is : %v\n", dst)
	fmt.Printf("%v type is: %v AND cap is: %v \n", dst, reflect.TypeOf(dst), cap(dst))
	i := copy(dst, s1a)
	fmt.Println(i)
	fmt.Printf("%v\n", dst)
}

func arrayTest() {
	fmt.Println("------- check array ------")
	//[n]elementType
	array := [...]float64{1.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Println(array)
	//指定总长度，通过索引值进行初始化
	a1 := [3]int{0: 1, 2: 3}
	a2 := [...]int{0: 1, 2: 3}
	fmt.Println(a1)
	fmt.Println(a2)
	//length,range
	fmt.Println("len(array_variable) get length of array ", len(a2))
	for v := range a2 {
		fmt.Println(v)
	}
}

func pointTest() *int {
	fmt.Println("------- check point ------")
	//声明 *T 支持多级指针 **T ,&指针变量 获取指针地址
	//1. *T 出现再 = 左边 声明指针 ，*T 出现在 = 右边表示取指针指向的值
	a := 11
	p := &a
	fmt.Println(*p)
	fmt.Println(p)
	//2. 通过 . 访问字段
	//3. 禁止指针运算
	//4. 函数中允许返回局部变量的地址，栈逃逸机制允许局部变量分配在堆上
	return p
}

/**
复合类型
	指针 * pointType，
	数组 [n] elementType，
	切片 [] elementType，
	字典 map [keyType] valueType，
	通道 chan valueType，
	结构 struct { fieldName fieldType}，
	接口 interface { method (inputParams)(returnParams)}

*/
