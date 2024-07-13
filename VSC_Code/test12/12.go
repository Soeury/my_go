package main

import "fmt"

func main() {

	// switch 语句  ->  适合判断一个数据是否符合某一个特定的值
	// 如果是判断区间最好用if语句判断

	// switch val {case1 : ...    case2 : ...    ...}

	var score int = 90

	switch score {
	case 90:
		fmt.Println("A")
	case 80:
		fmt.Println("B")
	case 70:
		fmt.Println("C")
	case 60:
		fmt.Println("D")
	default:
		fmt.Println("不及格")
	}

	// switch 语句没有值 默认采用 bool = true 的形式
	// 即以下代码会打印case true 的语句

	switch {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	default:
		fmt.Println("not found")
	}

	// switch  fallthrough穿透  ->  直接执行下一个case

	var key int = 10

	switch key {
	case 10:
		fmt.Println("key = 10")
		fallthrough
	case 20:

		if key == 10 {
			break
		}

		fmt.Println("key = 20")
	case 30:
		fmt.Println("key = 30")
	default:
		fmt.Println("not found")
	}
}
