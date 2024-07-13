package main

import (
	"fmt"
	"time"
)

// 写入数据到通道
func test(ch chan int) {

	for i := 1; i <= 10; i++ {
		time.Sleep(time.Millisecond * 300)
		ch <- i
	}

	close(ch)
}

func test2(ch2 chan int) {

	for i := 1; i <= 10; i++ {
		time.Sleep(time.Millisecond * 300)
		ch2 <- i
	}

	close(ch2)
}

func main() {

	// 关闭通道

	ch := make(chan int)
	go test(ch)

	// 1. 手动for循环拿出 channel 数据 ， 一般不使用 ， 通常使用 for range

	for {
		time.Sleep(time.Millisecond * 300)
		v, ok := <-ch

		if !ok {
			fmt.Println("channel close!")
			break
		}
		fmt.Println("data :", v)
	}

	// 2. for range 自动识别通道是否关闭(底层也是 ok-idiom 实现的)

	ch2 := make(chan int)
	go test2(ch2)

	for v := range ch2 {
		fmt.Println("data - 2 :", v)
	}

}
