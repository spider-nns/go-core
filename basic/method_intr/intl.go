package main

func main() {

}

//接口
//接口是一组方法签名的集合，一个具体类型实现接口不行哟啊在语法上显示声明
//只要具体类型的方法集是接口方法集的超集，就代表该类型实现了接口，编译器在编译时会进行方法集的校验

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}
type ReadWriter interface {
	Reader
	Writer
}

//
//type ReadWriter interface {
//	Reader
//	Writer(p []byte) (n int, err error)
//}
//type ReadWriter interface {
//	Reader(p []byte) (n int, err error)
//	Write(p []byte) (n int, err error)
//}

//接口绑定具体类型的实例的过程称为接口初始化
//1.实例赋值接口
//如果具体类型实例的方法集是某个接口的方法集的超集，称该类型实现了接口
//可以将该具体类型的实例直接复制给接口类型的变量,编译器会进行静态的类型检查
//接口被初始化后，调用接口的方法就相当于调用接口绑定的具体类型方法
//2.接口变量赋值接口变量
//已经初始化的接口变量a直接赋值给另一种接口变量b，要求b的方法集是a的方法集的子集
//编译器会在编译时进行方法集的静态检查，此时接口b绑定的具体实例是接口变量a绑定的具体实例的副本
