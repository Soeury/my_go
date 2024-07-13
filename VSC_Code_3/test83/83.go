package main

import (
	"fmt"
	"strconv"
	"time"
)

func get_string(ch chan string) {

	for i := 1; i <= 10; i++ {
		ch <- "string - " + strconv.Itoa(i)
		fmt.Println("wrtie data : ", "string - "+strconv.Itoa(i))
	}

	close(ch)
}

func main() {

	// 非缓冲通道 ：写入一个数据后堵塞 ， 这个数据读完之后才能继续写入数据  ch := make(chan type)
	// 缓冲通道 : 给这个通道设置一个容量 ， 写入数据后容量满了才堵塞 ， 读出数据后才能继续写入  ch := make(chan type , capacity)

	ch := make(chan int)
	fmt.Println("len(ch) :", len(ch), "cap(ch) :", cap(ch)) // 非缓冲通道默认 len = 0  , cap = 0

	ch2 := make(chan int, 5)
	fmt.Println("len(ch2) :", len(ch2), "cap(ch2) :", cap(ch2)) // 0  10

	ch2 <- 2
	ch2 <- 4
	fmt.Println("len(ch2) :", len(ch2), "cap(ch2) :", cap(ch2))
	ch2 <- 6
	ch2 <- 8
	ch2 <- 10
	fmt.Println("len(ch2) :", len(ch2), "cap(ch2) :", cap(ch2))
	// ch2 <- 12 // 通道容量满了之后不能再添加数据

	fmt.Println("---------------------")

	ch3 := make(chan string, 5)
	go get_string(ch3)

	// 取出数据 (等待一段时间将数据读入通道后再取出)
	for v := range ch3 {
		time.Sleep(time.Millisecond * 500)
		fmt.Println(v)
	}

}
