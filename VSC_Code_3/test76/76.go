package main

import "fmt"

func hello() {

	for i := 1; i <= 1000; i++ {
		fmt.Println("hello---", i)
	}
}

func main() {

	// goroutine : 协程  -> 只在某个函数调用前面加上 go 即可
	// 1-  当遇到 goroutine 时 ， goroutine 的调用立即返回 ， 与函数不同 ， go 不等待 goroutine 执行结束
	// 2-  goroutine 后 main 函数中的语句与 goroutine 并发执行
	// 3-  main 函数是一个主 goroutine ， 当这个主 goroutine 执行结束后， 其他的 goroutine 就不会执行

	go hello() // 这里可能会出现 主 goroutine 执行完毕而其他的 goroutine 还未执行完毕的情况

	for i := 1; i <= 1000; i++ {
		fmt.Println("main----------------", i)
	}

}
