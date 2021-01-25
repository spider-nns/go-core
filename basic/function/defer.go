package main

import "fmt"

func main() {
	defer func() { println("first") }()
	defer func() { println("second") }()
	println("function body")
	return
	//this defer registry fail，not execute
	fmt.Printf("defer deal param: %d\n", f())
}

//defer 关键字,可以注册多个延迟调用 栈，子弹夹
//调用以先进后出 filo的顺序在函数返回前被执行,常用于保证一些资源最终一定能够得到回收和释放

//1.defer 后面必须是函数或者方法的调用 expression in defer must be function call
//2.defer 函数的实参在注册时通过值拷贝传递过去
//3.defer 必须先注册后才能执行，如果 defer 位于 return 之后，defer 因为没有注册，不会执行
//4.主动调用 os.Exit(int) 退出进程时，defer 将不再被执行,即使 defer 已经提前注册

//defer 位置不当可能会导致 panic ，一般defer 语句放在错误检查语句之后
//defer 会推迟资源的释放,尽量不要放在循环中,不要对有名返回值参数进行操作

func f() int {
	a := 0
	defer func(i int) {
		println("defer i=", i)
	}(a)
	a++
	return a
}
