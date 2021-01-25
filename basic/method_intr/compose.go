package main

import "fmt"

func main() {
	compose()
	nesting()
	methodSet()
}

func methodSet() {
	println("---- methodSet ----")
	x := X{a: 1}
	y := Y{X: x}
	println(y.get())
	y.set(2)
	print(y.get)
	(*Y).set(&y, 3)
	Y.set(y, 3)

}

//组合
//type 新定义的新类型不会继承原有类型的方法
//命名结构类型[特殊]，可以嵌套其他的命名类型字段，外层的结构类型可以调用嵌入字段类型的方法（显式，隐式）
//内嵌字段唯一，不需要全路径访问
func compose() {
	x := X{a: 1}
	y := Y{
		X: x,
		b: 2,
	}
	z := Z{Y: y, c: 3}
	//z.a  z.Y.z  z.Y.X.a  等价
	println(z.a, z.Y.a, z.Y.X.a)
	z = Z{}
	z.a = 2
	println(z.a, z.Y.a, z.Y.X.a)
}

//多层嵌套，最好使用完全路径访问，初始化，如果出现嵌套字段同名，代表不同的字段
//内嵌字段的方法调用，有相同的方法，编译器优先从外向内逐层查找方法，外层方法会覆盖内层方法
func nesting() {
	x := X{a: 1}
	y := Y{X: x, b: 2}
	z := Z{Y: y, c: 3}
	z.Print()
	z.X.Print()
	z.Y.X.Print()
}
func (x X) Print() {
	fmt.Printf("In X,a=%d\n", x.a)
}
func (y Y) Print() {
	fmt.Printf("In Y,a=%d\n", y.a)
}
func (z Z) Print() {
	fmt.Printf("In Z,a=%d\n", z.a)
	z.Y.Print()
	z.Y.X.Print()
}

//组合的方法集
//1.如果类型 s 包含匿名字段 T,则 s 的方法集包含 T 的方法集
//2.如果类型 s 包含匿名字段 *T，则 s 的方法集包含 T,*T 的方法集合
//3.不管类型 s 中嵌入的匿名字段是 T 还是 *T,*S 方法集总是包含 T 和 *T 方法集

type X struct {
	a int
}
type Y struct {
	X
	b int
}
type Z struct {
	Y
	c int
}
type G struct {
	*X
}

func (x X) get() int {
	return x.a
}

func (x *X) set(i int) {
	x.a = i
}
