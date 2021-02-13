package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//writeAnsWriteString()
	//bufioWrite()
	//ioutilAction()
	//fileCopy()
}

//os.OpenFile 函数能够以指定模式打开文件
//os.o_WRONLY 只读
//os.o_CREATE 创建文件
//os.o_RDONLY 只读
//os.o_RDWR   读写
//os.o_TRUNC  清空
//os.o_APPEND 追加

//文件权限 八进制数 r 读 04,w 写 02，x 执行 01

func writeAnsWriteString() {
	file, err := os.OpenFile("file.txt", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("open file failed,err", err)
		return
	}
	defer file.Close()
	str := "2021.2.12 spider write go\n"
	file.Write([]byte(str)) //写入字节切片数据
	file.WriteString(str)   //直接写入字符串数据

}
func bufioWrite() {
	file, err := os.OpenFile("bufiofile.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file fialed ,err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("go,go,go\n")
	}
	writer.Flush() //将缓存中的内容写入文件
}
func ioutilAction() {
	str := "2021 ,go"
	err := ioutil.WriteFile("ioutil.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed ,err:", err)
		return
	}
}

func fileCopy() {
	srcFile := "./file.txt"
	destFile := "./destFile.txt"
	src, err := os.Open(srcFile)
	if err != nil {
		fmt.Printf("open %s failed,err:%v.\n", srcFile, err)
		return
	}
	defer src.Close()
	//写、创建 方式打开 dest 文件
	dst, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, 0644)
	defer dst.Close()
	written, err := io.Copy(dst, src)
	if err != nil {
		fmt.Printf("io copy from %s to %s failed,err:%v\n", srcFile, destFile, err)
	}
	fmt.Println(written)
}
func catMock(){

}