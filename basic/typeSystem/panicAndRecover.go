package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer recover inner")
		except()
	}()
	do()
}

// panic 和 defer 一起使用，只有在defer后面函数体内被直接调用才能捕获panic终止异常,否则返回nil,异常继续向外传递,嵌套两层也会捕获失败
// init 函数中引发的panic 只能在init函数中捕获
// 函数不能捕获内部新启动的 goroutine 所抛出的panic ?
// 连续多个panic 只能出现在延迟调用里面，但是只有最后一次panic能被捕获

func do() {
	panic("main body panic")
	defer func() {
		println("anonymous first")
		panic("panic first")
	}()
	defer func() {
		println("anonymous second")
		panic("panic second")
	}()
}
func except() {
	recover()
}
