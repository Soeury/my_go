package main

import "fmt"

func main() {

	// 指针使用 - 嵌套
	var a int = 10
	var p *int = &a

	fmt.Println("a的值是: ", a)
	fmt.Println("a的地址是: ", &a)
	fmt.Println("p指针指向地址是: ", p)
	fmt.Println("p指针的地址是: ", &p)
	fmt.Println("p指针指向的值是: ", *p)

	fmt.Println("---------------------------------")
	*p = 20
	fmt.Println("a的值是: ", a)
	fmt.Println("p指针指向的值是: ", *p)

	fmt.Println("---------------------------------")
	a = 30
	fmt.Println("a的值是: ", a)
	fmt.Println("p指针指向的值是: ", *p)

	// ptr 指针是指向指针p的一个指针, 类型是 **int , p指针的类型是 *int
	var ptr **int = &p

	fmt.Println("---------------------------------")
	fmt.Println("ptr指针指向的地址是: ", ptr)
	fmt.Println("ptr指针的地址是: ", &ptr)
	fmt.Println("ptr指针指向的值是: ", *ptr)
	fmt.Println("p指针指向地址是: ", p)
	fmt.Println("*ptr指向的值是: ", **ptr)
}
