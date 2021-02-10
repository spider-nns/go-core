package main

import (
	"fmt"
	"log"
	"time"
)

var originTime = "2006-01-02 15:04:05"
var originTimePM = "2006-01-02 15:04:05 PM"

func main() {

	timeDemo()
	timeUnixStamp()
	timeComplex()
	timeParse()
	tickDemo()
}

func timeDemo() {
	now := time.Now()
	fmt.Printf("current time：%v\n\n", now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

}

func timeUnixStamp() {
	now := time.Now()
	timeStamp1 := now.Unix()
	timeStamp2 := now.UnixNano()
	fmt.Printf("current timeStamp1:%v\n", timeStamp1)
	fmt.Printf("current timeStamp2:%v\n", timeStamp2)
	//convert unixStamp to date
	timeObj := time.Unix(timeStamp1, 0)
	fmt.Println(timeObj)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

}

func durationDemo() {
	const (
		NanoSecond  time.Duration = 1 // 纳秒为单位 表示时间间隔，可表示的最大长度段大约 290 年
		Microsecond               = 1000 * NanoSecond
		Millisecond               = 1000 * Microsecond
		Second                    = 1000 * Millisecond
		Minute                    = 60 * Second
		Hour                      = 60 * Minute
	)
}
func timeComplex() {
	now := time.Now()
	//add
	later := now.Add(time.Hour * 12)
	log.Printf("formatTime: %s, %d", later.Format(originTime), later.Unix())
	laterPm := now.Add(time.Hour * 13)
	log.Printf("formatTime: %s, %d", laterPm.Format(originTimePM), laterPm.Unix())
	//sub
	subDuration := later.Sub(now)
	reverseDuration := later.Add(-12 * time.Hour)
	log.Printf("time sub:%d Hour", subDuration/time.Second)
	log.Printf("subDuration is:%f", subDuration.Hours())
	log.Printf("reverseDuration is:%s", reverseDuration.Format(originTime))
	//equal
	location, _ := time.LoadLocation("Asia/Shanghai")
	timeObj, _ := time.ParseInLocation(originTime, now.String(), location)
	equal0 := timeObj.Equal(now)
	// == 比较时区和地点
	if timeObj == now {
		println("timeObj == now")
	}
	//equal 考虑时区
	log.Printf("time equal with location:%v\n", equal0)
	//before
	if timeObj.Before(now) {
		println("timeObj is before now ")
	}
	//after
	if now.After(timeObj) {
		println("now is after timeObj ")
	}

}

func timeParse() {
	now := time.Now()
	fmt.Println(now)
	//加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	//按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation(originTime, "2021-02-11 00:50:51", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Printf("sub with location:%f\n", now.Sub(timeObj).Minutes())
}

func tickDemo(){
	ticker := time.Tick(time.Second)
	for i:= range ticker{
		//每秒执行
		fmt.Println(i)
	}
}
