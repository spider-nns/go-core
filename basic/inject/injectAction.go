package main

import (
	"fmt"
	"github.com/codegangsta/inject"
)

type S1 interface{}
type S2 interface{}

func Format(name string, company S1, level S2, age int) {
	fmt.Printf("name=%s,company=%s,level=%s,age=%d!\n", name, company, level, age)
}

type Staff struct {
	Name    string `inject`
	Company S1     `inject`
	Level   S2     `inject`
	Age     int    `inject`
}

func main() {
	inj := inject.New()
	inj.Map("tom")
	inj.MapTo("tencent", (*S1)(nil))
	inj.MapTo("T4", (*S2)(nil))
	inj.Map(24)
	//函数反转调用
	inj.Invoke(Format)

	s := Staff{}

	jet := inject.New()
	jet.Map("tom")
	jet.MapTo("tencent", (*S1)(nil))
	jet.MapTo("T4", (*S2)(nil))
	jet.Map(23)
	jet.Apply(&s)
	fmt.Printf("s=%v\n", s)

}
