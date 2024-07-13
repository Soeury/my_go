package main

import "fmt"

func main() {

	// go 指针初识
	// & 取地址符号 ->   &val
	// * 取值符号   ->   *ptr

	var a int = 10
	var b = &a

	fmt.Println("a的值是: ", a)
	fmt.Println("a的地址是: ", &a)

	fmt.Println("b的值是: ", b)
	fmt.Println("b指向的内存空间上的值 *b 是: ", *b)

	*b = 66
	fmt.Println("修改 *b 后 a 的值是: ", a)

}
