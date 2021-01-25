package main

import (
	"fmt"
)

func main() {
	p := new(person)
	fmt.Println(p)
	p2 := of("nns", 23)
	fmt.Println(p2)

}

//类型方法，实现面向对象
//自定义类型继承底层类型的操作集合(map)不会继承旧类型方法
//struct 匿名字段 只给出字段类型，类型必须为命名类型或命名类型的指针
//一个结构体里面不能同时存在某个类型和其指针的匿名字段
//如果嵌入的字段来自其他包，需要加上包名，并且是可以导出的类型

//类型方法
//func (t TypeName) methodName(paramList)(returnList){ methodBody}
//1.可以为命名类型增加方法(除了接口)，非命名类型不能自定义方法
//2.为类型增加方法限制，方法定义必须和类型定义在一个包中
//3.不能为内置的预声明类型增加方法，并且不在一个包中
//4.方法的命名空间的可见行和变量一样
//5.使用 type 定义的自定义类型不能调用原有类型的方法，但是底层类型支持的运算可以被新类型继承
type person struct {
	name string
	age  int
}

//推荐使用构造函数，当结构发生变化时，构造函数可以屏蔽细节
func of(name string, age int) person {
	return person{name: name, age: age}
}
