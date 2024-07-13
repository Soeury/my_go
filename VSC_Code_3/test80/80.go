package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1 - ", i)
	}

	wg.Done()
}

func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("test2 --- ", i)
	}

	wg.Done() // 每执行一次 wg.Done()  , goroutine 的数量就 -1
}

func main() {

	// WaitGroup 代替 time.Sleep( )

	wg.Add(2) // wg.Add(num) 添加 goroutine 的数量

	go test1()
	go test2()

	fmt.Println("main ing")
	wg.Wait() // 等待 所有的 goroutine 执行完毕后 立刻 程序继续向后执行
	fmt.Println("main end")

}
