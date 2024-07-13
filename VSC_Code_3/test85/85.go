package main

import (
	"fmt"
	"time"
)

func main() {

	// Select : 类似于 switch 语句
	//    1- Select 中的每个case 语句都必须是通道的操作
	//    2- Select 若多个case 语句都可以执行，那么会随机选择一个可以执行的 case 语句 ， 其余的 case 不会执行
	//    3- 如果没有 case 语句可以执行
	//         1- 如果有 default 语句 ：则执行 default 语句
	//         2- 如果没有 default 语句 ， Select 会阻塞， 直到有某个 case 可以执行

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(time.Second * 1)
		ch2 <- 200
	}()

	select {
	case num1 := <-ch1:
		fmt.Println(num1)
	case num2 := <-ch2:
		fmt.Println(num2)
		//default:
		//	fmt.Println("no run!")
	}

}
