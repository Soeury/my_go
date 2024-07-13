package main

import "fmt"

func listen_select(ch chan<- int, quit <-chan bool) {

	x := 1
	y := 1

	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit: // 一开始堵塞 ， 只有当quit通道中写入数据后才能读出数据(这里就是等后面将10个数据输出后来做终止用的)
			fmt.Println("flag =", flag)
			return
		}
	}
}

func main() {

	// select 实现斐波那契数列

	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			num := <-ch // 一开始堵塞，只有select中将数据写入通道后才能读取，这里只读取10个数据
			fmt.Println(num)
		}
		quit <- true
	}()

	listen_select(ch, quit)
}
