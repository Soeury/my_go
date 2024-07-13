package main

import "fmt"

/*
	recover :

		1. recover 的作用是捕获panic , 回复程序正常运行
		2. recover 必须配合 defer 一起使用
		3. recover 不需要传入参数，但是有返回值 ， 返回值就是 panic 的内容
		4. recover语句一般放在匿名函数中 ， 后面加()直接执行
		5. defer msg := recover()
*/

func test(i int) {

	// recover语句放在匿名函数中进行!!!
	// 简而言之 :  recover = defer + 匿名函数 + 返回值
	defer func() {
		msg := recover()

		if msg != nil {
			fmt.Println("msg = ", msg, "--- 程序恢复正常")
		}
	}()

	defer fmt.Println("test --- 1")
	defer fmt.Println("test --- 2")
	fmt.Println("test --- 3")

	if i == 1 {
		panic("find panic!") // panic抛出异常
	}

	fmt.Println("test --- 4")
	defer fmt.Println("test --- 5")
}

func main() {

	// panic 用来抛出异常 ， 当程序中出现了 panic语句，会终止其后要执行的代码 ， 若panic前执行的代码中有defer语句
	// 则将所有的 defer 语句按照逆序执行完毕之后程序抛出 panic 异常 ， panic后的代码都不会执行

	defer fmt.Println("main --- 1")
	defer fmt.Println("main --- 2")
	fmt.Println("main --- 3")

	test(1)

	fmt.Println("main --- 4")
	defer fmt.Println("main --- 5")
	defer fmt.Println("main --- 6")

}
