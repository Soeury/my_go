package main

import "fmt"

// 函数接收不定数量的参数，类型全部都是int
func add(nums ...int) {

	sum := 0
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
		sum += nums[i]
	}

	fmt.Println("sum:", sum)
}

func main() {

	// 可变参数:
	// func 函数名称(变量名称 ...数据类型) {}   ->   代表函数接收不定数量的参数，类型全部都是一种

	// 可变参数最多只有一个，并且可变参数必须写在所有参数的最后面
	// 错误示范：
	// func test(str ...string , nums ...int)

	add(1, 3, 5, 7, 9, 100)
}
