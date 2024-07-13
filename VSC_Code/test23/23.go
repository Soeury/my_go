package main

import "fmt"

func add() func() int {

	i := 0

	f := func() int {
		i += 1
		return i
	}

	return f // f 的类型是func() int
}

func main() {

	// 闭包结构
	// 在一个外层函数中，有一个内层函数，内层函数中，会操作一个外层函数的局部变量
	// 并且外层函数的返回值就是内层函数
	// 这样的结构称为闭包结构

	r1 := add()

	// fmt.Println(r1) // r1 实际上是一个函数，打印的是地址

	fmt.Println(r1())
	fmt.Println(r1())
	fmt.Println(r1())
	fmt.Println(r1())
	fmt.Println(r1())

	r2 := add()

	fmt.Println(r2())
	fmt.Println(r2())
	fmt.Println(r2())
	fmt.Println(r1())
}
