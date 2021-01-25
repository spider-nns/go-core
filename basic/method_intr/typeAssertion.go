package main

import (
	"fmt"
)

type Inter interface {
	Ping()
	Pang()
}
type Anter interface {
	Inter
	String()
}
type St struct {
	Name string
}

func (St) Ping() {
	println("ping")
}
func (*St) Pang() {
	println("pang")
}
func main() {
	st := &St{"andes"}
	var i interface{} = st
	if o, ok := i.(Inter); ok {
		o.Ping()
		o.Pang()
	}
	// i 没有实现接口 Inter 会 panic
	//anter := i.(Anter)
	//anter.String()
	if p, ok := i.(Anter); ok {
		fmt.Println("dont exec")
		p.String()
		println("p.string")
	}
	if s, ok := i.(*St); ok {
		fmt.Printf("%s", s.Name)
	}
}

//接口类型断言 i.(TypeName)
//i 必须是具体类型变量(non-interface type xxx on left)
//TypeName 可以是接口类型，也可以是具体类型名

//接口断言的两层语义
//如果TypeName 是一个具体类型名，则类型断言用于判断接口变量i绑定的实例类型是否就是具体类型 TypeName
//如果TypeName 是一个接口类型名，则类型断言用于判断接口变量i绑定的实例类型是否同时实现了TypeName接口

//空接口有两个字段，一个是实例字段，另一个是指向绑定实例的指针，只有两个都为 nil 时，空接口才为 nil

