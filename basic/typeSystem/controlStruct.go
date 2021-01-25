package main

func main() {
	ifTest()
	switchTest()
	forTest()
	labelTest()
}

func labelTest() {
	println("------ label start -----")
label1:
	for true {
		for i := 1; i < 15; i++ {
			println("1th for loop")
			if i == 5 {
				println(i)
				break
			}

		}
		for i := 1; i < 15; i++ {
			println("2th for loop")
			if i == 10 {
				println(i)
				break label1
			}
		}
	}
}

func forTest() {
	println("------for start ------")
	// for init;condition;post{}
	for i := 1; i <= 5; i++ {
		print(i)
	}
	//for  {
	//	println("cpu work")
	//}
	//for 1==1 {
	//	println(1)
	//}1
	println()
}

func switchTest() {
	println("------  switch start ------")
	//switch 表达式可以省略 多重if else
	switch {
	//顺序没要求
	default:
		{
			println("switch default out")
		}
	case true:
		{
			println("true")
		}
		//强制执行下一个 case
		fallthrough
	case false:
		{
			println("false")
		}
	}
	//another switch
	//default 可以不存在
	switch l := len("switch condition"); l {

	case 2:
		println("switch condition  is 2")
	case 3, 4, 5, 6, 7, 8, 9, 10:
		return
	case 11, 12, 13, 14, 15:
		println("switch condition in (11,12,13,14,15,16)")
	case 16:
		println("condition is 16")
		fallthrough
	default:
		println("switch over")
	}
}

func ifTest() {
	//if
	if condition := "if"; len(condition) > 5 {
		println("condition length greater than 5")
	} else {
		return
		println("condition length less than 5")
	}
}
