package main

import (
	json2 "encoding/json"
	"fmt"
)

type outOfService struct {
	ClockStart string `json:"clock_start"`
	ClockEnd   string `json:"clock_end"`
}

type night struct {
	//匿名字段实现结构体继承
	os    *outOfService `json:"os"`
	Ide   string        //默认使用字段名
	Music string        `json:"music"`
}
type sl struct {
	Clock string `json:"clock"`
}

func main() {
	json()
}

//Tag 大写字段可导出
func json() {
	outOfService := &outOfService{ClockStart: "19:00", ClockEnd: "02:00"}
	n := &night{os: outOfService, Ide: "goland", Music: "Taylor Swift"}
	fmt.Printf("outOfService from %s to %s, use %s like a spider\n", n.os.ClockStart, n.os.ClockEnd, n.Ide)
	data, err := json2.Marshal(n)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
	fmt.Printf("%#v\n", data)
	ng := &night{}
	err = json2.Unmarshal(data, ng)
	if err != nil {
		fmt.Println(err)
		return
	}
	i := new(sl)
	i.Clock="01:30"
	marshal, err := json2.Marshal(&i)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(marshal))
	fmt.Printf("%#v\n", marshal)
}
