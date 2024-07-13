package main

import (
	"fmt"
	"runtime"
)

func main() {

	// goroutine 的调度与终止  -  runtime包(获取系统信息 ， 操控 goroutine)

	// 1. runtime 包的使用

	fmt.Println("获取cpu数量:", runtime.NumCPU())
	fmt.Println("获取goroot目录:", runtime.GOROOT())
	fmt.Println("获取操作系统:", runtime.GOOS)

	// go 调度

	go func() {
		for i := 1; i <= 500; i++ {
			fmt.Println("goroutine--", i)
		}
	}()

	// runtime.Gosched() 当前函数让出时间片给其他的 goroutine
	//  不一定能成功 , 但是大概率能让其他的goroutine先执行，然后再执行当前函数

	for i := 1; i <= 500; i++ {
		runtime.Gosched()
		fmt.Println("main---------------", i)
	}

}
