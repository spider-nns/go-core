package main

import (
	"errors"
	"fmt"
)

func main() {

	r := fmt.Errorf("500:auth failed")
	if r != nil {
		fmt.Println(r)
	}
	err := errors.New("500:errors action")
	if err != nil {
		fmt.Println(err)
	}
}

//defer 语句放在 err 判断后面，避免出现 panic
//错误逐级向上传递过程，应该是不断完善的

//异常 发生未期待的未知行为
//错误 发生非期望的已知行为，已知定义好的错误类型

//类型安全的go 运行时不会出现编译器运行和运行时都无法捕获的错误,出现的异常都是可捕获的错误

//异常处理规则
//1.程序发生的错误导致程序不能容错继续执行，此时程序应该主动调用 panic 或 由运行时抛出 panic
//2.程序虽然发生错误，但是程序能够容错继续执行，此时应该使用错误返回值的方式处理错误，或者在可能发生运行时错误的非关键分之上使用 recover 捕获 panic

/**
 错误 errors
 untrapped error 未捕获错误-异常  ｜  可捕获错误 trapped error
								｜ compile errors、runtime errors、logic errors
*/
