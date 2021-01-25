package main

func main() {

	//invoke typeMethod
	//普通方法调用
	t := &T{}
	t.set(5)
	println("普通方法调用，t:", t.get())
	var t1 = &T{}
	t1.set(5)
	println("普通方法调用，t:", t1.get())
	//方法值==带有闭包的函数变量 x.M
	var mt = &T{}
	s := mt.set
	s(5)
	println("方法值调用", mt.get())
	//方法表达式调用
	e := T{}
	(*T).set(&e, 1)
	get := T.get
	println("方法表达式调用", get(e))
	//方法表达式
	var at = &T{}
	(*T).set(at, 123)
	f := (*T).set
	f(at, 666)
	println(T.get(*at))
	//方法集校验
	//字面量直接显式调用，编译器不会进行方法集的自动转换，会严格校验方法集
	(*Data)(&struct{}{}).testPointer()
	(*Data)(&struct{}{}).testValue()
	(Data)(struct{}{}).testValue()
	//
	//(Data)(struct{}{}).testPointer()
}

//类型方法本质上是函数
//typeInstanceName.methodName(paramList)
//调用方式，方法集，方法变量，方法表达式
//struct 可以嵌入任何类型字段，可以嵌套自身的指针类型字段

//方法值
//变量 x 的静态类型是 T,M 是类型 T 的一个方法，x.M 被称为方法值
//x.M 是一个函数类型变量，可以赋值给其他变量
//方法值接受者被隐式绑定到方法值的闭包环境上
//类似 java 普通对象引用调用方法

//方法表达式
//类型方法调用转换为函数调用，需要显示指定接受者
//类似 java 内部类
//方法表达式，编译器不会做自动转换

//方法集
//1.命名类型 自定义类型 如果底层类型是内置的自定义类型，可以继承操作（计算+赋值）
//2.接受者为值类型和指针类型的的方法的集合称为方法集
//3.编译器在编译期间能够识别出调用关系，自动转化 类似java 装箱、拆箱
//4.字面量直接显式调用，编译器不会进行方法集的自动转换，会严格校验方法集

type T struct {
	a int
}

func (t *T) set(a int) {
	t.a = a
}

func (t T) get() int {
	return t.a
}

//方法集合校验
type Data struct{}

func (Data) testValue() {}

func (*Data) testPointer() {}

//Data 方法集是 testValue
//*Data 方法集是 testPointer testValue
