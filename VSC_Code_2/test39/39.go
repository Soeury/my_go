package main

import "fmt"

func main() {

	// 指针与数组
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Printf("arr1 的地址: %p\n ", &arr1)

	var p *[4]int = &arr1 // 指向arr1数组的指针p

	fmt.Println("p指针的地址: ", &p)
	fmt.Printf("p->地址: %p\n ", p)
	fmt.Println("p->地址->值: ", *p)
	fmt.Println()

	var num int = (*p)[0]
	fmt.Println("(*p)[0] = ", num)

	// 简化写法: 指针指向谁，就代表谁  可以用 (*p)[index] 的方式 ， 也可以用 p[index] 的方式
	var num1 int = p[0]
	fmt.Println("p[0] = ", num1)

	p[0] = 66
	fmt.Println("p[0] = ", p[0])
	fmt.Println(arr1)

	// 指针数组
	a := 5
	b := 6
	c := 7
	d := 8

	var arr2 = [4]*int{&a, &b, &c, &d}
	fmt.Println(arr2)

	fmt.Println()
	fmt.Println("before modyfing *arr2[0] : ")
	fmt.Println("*arr2[0] = ", *arr2[0], " a = ", a)
	*arr2[0] = 99

	fmt.Println()
	fmt.Println("after modyfing *arr2[0] = 99 : ")
	fmt.Println("*arr2[0] = ", *arr2[0], " a = ", a)

	fmt.Println()
	fmt.Println("after modifing 'a' = 55 :")
	a = 55
	fmt.Println("*arr2[0] = ", *arr2[0], " a = ", a)
	fmt.Println()

}
