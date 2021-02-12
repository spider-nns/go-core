package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	//demo0()
	//readFile()
	//readFileLoop()
	//readFileUseBufioReadLine()
	readFullFile()
}
func demo0() {
	file, err := os.Open("./go.mod")
	if err != nil {
		fmt.Println("open file failed!,err:", err)
		return
	}
	defer file.Close()
}

func readFile() {
	file, err := os.Open("./go-core.log")
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer file.Close()
	var temp = make([]byte, 128)
	n, err := file.Read(temp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("read file failed,err:", err)
		return
	}
	fmt.Printf("读取了%d字节数据", n)
	fmt.Println(string(temp[:n]))

}
func readFileLoop() {
	file, _ := os.Open("./go-core.log")
	defer file.Close()
	var content []byte
	var tmp = make([]byte, 128)
	for {
		read, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed，err:", err)
			return
		}
		//拼接两个slice
		content = append(content, tmp[:read]...)

	}
	fmt.Println(string(content))
}

//按行读取
func readFileUseBufioReadLine() {
	file, err := os.Open("./basic/file/fileRead.go")
	if err != nil {
		fmt.Println("open file failed，err", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed,err", err)
			return
		}
		fmt.Println(line)
	}
}
//读取整个文件
func readFullFile(){
	content, err := ioutil.ReadFile("./basic/file/fileRead.go")
	if err != nil {
		fmt.Println("read file failed,err", err)
		return
	}
	fmt.Println(string(content))
}
