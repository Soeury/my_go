package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func writeOnly(ch chan<- int) {
	ch <- 100
	ch <- 90
	ch <- 80
	ch <- 70
	ch <- 60
	close(ch) //  for range 必须要等到 close(channel) 才会结束读取
	wg.Done()
}

func readOnly(ch <-chan int) {

	for v := range ch {
		fmt.Println(v)
	}

	wg.Done()
}

func main() {

	// 定向通道 只能写入数据或者读出数据
	// 单项通道一般使用在 函数的 参数 和 返回值 中
	// 1- 只写入 : ch := make(chan <- type)
	// 2- 只读出 : ch := make(<- chan type)

	// 1- 只写入通道  ch := make(chan <- type)
	//ch1 := make(chan<- int)
	//ch1 <- 200
	// data1 := <-ch1 // cannot receive from send-only channel ch1

	// 2- 只读出通道  ch := make(<- chan type)
	// ch2 := make(<-chan int)
	// ch2 <- 200 // cannot send to receive-only channel ch2

	ch3 := make(chan int, 5)

	wg.Add(2)
	go writeOnly(ch3)
	go readOnly(ch3)

	wg.Wait()

}
