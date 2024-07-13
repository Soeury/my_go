package main

import "fmt"

// 有两个参数的函数 - 有一个返回值的函数
func add(a int, b int) int {
	return a + b
}

// 有一个参数的函数 - 没有返回值的函数
func m_print(s string) {
	fmt.Println(s)
}

// 没有参数的函数 - 有返回值的函数
func m_print2() int {
	return 20
}

// 有两个参数的函数 - 有多个返回值的函数
func m_swap(a int, b int) (int, int) {
	return b, a
}

// 返回更大值函数
func get_max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {

	// 函数声明  -  参数是形式参数，简称形参
	// func 函数名称(参数 数据类型 ...)(返回值类型) {实现}

	// 函数调用  -  参数是实际参数，简称实参
	// 函数名称(参数)

	ret1 := add(1, 2)
	fmt.Println(ret1)

	m_print("hello")

	ret2 := m_print2()
	fmt.Println(ret2)

	ret3, ret4 := m_swap(25, 36) // 交换其实可以直接   a , b = b , a   就行
	fmt.Println(ret3, ret4)

	ret5 := get_max(12, 34)
	fmt.Println("the max is :", ret5)
}
