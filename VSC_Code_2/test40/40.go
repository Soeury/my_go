package main

import "fmt"

func return_ptr() *[4]int {

	arr := [4]int{1, 2, 3, 4}
	return &arr
}

// 传入的是地址，这个函数修改了这个地址上的值，所以a的值会被修改
func get_ptr(ptr *int) {

	fmt.Println("ptr = ", ptr)
	fmt.Println("*ptr = ", *ptr)
	*ptr = 100
}

func main() {

	// 指针与函数
	// 1. 指针作为函数返回值

	fmt.Println("--------------------------")
	ptr := return_ptr()
	fmt.Printf("ptr的类型: %T\n", ptr)
	fmt.Println("ptr的地址: ", &ptr)
	fmt.Printf("ptr->地址:  %p\n", ptr)
	fmt.Println("ptr->值: ", *ptr)
	fmt.Println((*ptr)[0])
	fmt.Println(ptr[0])

	// 2. 指针作为函数参数

	fmt.Println("--------------------------")
	a := 10
	fmt.Println("a = ", a)
	get_ptr(&a)
	fmt.Println("调用get_ptr之后 a = ", a)

}
