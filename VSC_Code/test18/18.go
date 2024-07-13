package main

import "fmt"

var num int = 36

func test1() {
	num := 72
	fmt.Println(num)
}

func test2() {
	num := 108
	fmt.Println(num)
}

func main() {

	//变量的作用域 : 可以在不同作用域下初始化同一个变量一次 ， 变量的使用采用就近原则
	temp := 20

	// if 语句中可以直接初始化变量

	if b := 1; b < 10 {
		temp := 40
		fmt.Println(temp)
		fmt.Println(b)
	}

	fmt.Println(temp)

	fmt.Println("---")

	fmt.Println(num)
	test1()
	test2()

}
