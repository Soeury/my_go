package main

import (
	"fmt"
	"os"
)

/*  error 是一个内置的接口类型

type error interface {
	error() string
}

*/

func main() {

	// 错误类型   -   实际上是一个接口类型

	file, err := os.Open("aaa.txt")

	if err != nil {
		fmt.Println(err)

		// 断言取出错误信息
		val, ok := err.(*os.PathError)

		// 如果 ok == true ， 代表 err 确实是 *os.PathError 类型 ， 那么这时候就可以取出 val 中的值 ， 即  Os , Path , Err
		if ok {
			fmt.Println(val.Op)
			fmt.Println(val.Path)
			fmt.Println(val.Err)
		}

		return
	}

	fmt.Println(file.Name())
}
