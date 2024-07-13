package main

import "fmt"

func main() {

	// String
	// len(str) 打印长度
	// str[i] 获取每个字符

	str := "hello,world"

	fmt.Println("str的长度是 : ", len(str))

	fmt.Print(str[0]) // str[0] 是单个字符 ， 所以打印出来的是对应的ASCII编码
	fmt.Printf("\n%c\n", str[0])

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
	}

	fmt.Println()

	// for range 循环 i --- 下标   v --- 字符

	for i, v := range str {
		fmt.Print(i)
		fmt.Printf("%c ", v)
	}

	// string字符串是不能修改的
	// str[2] = 'A'

}
