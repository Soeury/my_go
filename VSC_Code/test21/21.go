package main

import "fmt"

func f1() {
	fmt.Println("我是f1函数")
}

func main() {

	// 匿名函数就是没有名字的函数
	f1()
	f2 := f1
	f2()

	// 匿名函数，函数体后加一个()表示执行，通常只能执行一次
	func() {
		fmt.Println("我是匿名函数")
	}()

	// 也可以将匿名函数赋值，之后单独调用
	f3 := func() {
		fmt.Println("我是匿名函数赋值的f3函数")
	}

	f3()

	// 带参数的匿名函数
	func(a int, b int) {
		fmt.Println("带参匿名函数结果:", a+b)
	}(1, 2)

	//带返回值的匿名函数
	ret := func(m int, n int) int {
		return m + n
	}(3, 4)

	fmt.Println("带返回值的匿名函数ret的值是:", ret)

}
