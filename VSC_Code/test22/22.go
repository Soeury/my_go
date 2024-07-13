package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

// 高阶函数 可以接收一个函数作为参数(被接收的函数称为回调函数，也可以是匿名函数)
func oper(a int, b int, f func(int, int) int) int {
	ret := f(a, b)
	return ret
}

func main() {

	// 回调函数 - 高阶函数
	ret1 := add(1, 2)
	fmt.Println(ret1)

	ret2 := oper(3, 4, add)
	fmt.Println(ret2)

	ret3 := oper(5, 3, sub)
	fmt.Println(ret3)

	ret4 := oper(12, 0, func(a int, b int) int {

		if b == 0 {
			return -1
		}

		return a / b
	})

	fmt.Println(ret4)

}
