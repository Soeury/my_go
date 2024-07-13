package main

import "fmt"

func main() {

	// channel 初识
	// 一个时间只有一个 goroutine 能够访问 channel 中的数据 ， 使用 channel 在不同的 goroutine 中传递内存数据
	// 通道接收数据后 ， 必须向外发送数据 ， 并且只有在发送数据之后才能继续接收 ，否则不能接收新的数据
	// 通道发送数据后 ， 必须接收到了数据才能继续发送 ， 否则不能发送新的数据
	// 创建通道后 ， 必须要使用 ， 否则出现死锁
	// 通道是 goroutine 之间的连接 ， 通道发送和接收数据必须在不同的 goroutine 上

	// name := make(chan type)

	// 1- 向外发送数据   data := <- name
	// 2- 通道接收数据   name <- data

	ch := make(chan bool) // 一个bool类型的数据

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println(i)
		}

		ch <- true // 循环结束后，向通道中写入数据，表示要结束了
	}()

	data := <-ch // 读取通道中的数据
	fmt.Println("data = ", data)

	// 下面的通道创建了之后只有数据读入 ， 没有在其他的 goroutine 中发送出数据 ， 所以会报错
	//ch2 := make(chan int)
	//ch2 <- 6

}
