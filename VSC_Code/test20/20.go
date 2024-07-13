package main

import "fmt"

func test(a int, b int) {
	fmt.Println(a, b)
}

func main() {

	// func()本质是一种数据类型 ,   比如  var name func(int ,int)
	// test() 是函数调用
	// test 一个特殊的指针变量，指向函数体的内存地址

	fmt.Printf("%T", test) // 这里说明函数test的类型是func(int , int)
	fmt.Println()

	var test2 func(int, int) // 一个函数类型的变量
	test2 = test
	fmt.Println(test)
	fmt.Println(test2) // 两个函数指向同一块地址 (有点像引用传递)

	test2(1, 2)

}
