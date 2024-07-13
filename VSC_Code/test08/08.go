package main

import "fmt"

func main() {

	// 算数运算符 (打印结果) +  -  *  /  %  ++  --
	a := 10
	b := 3

	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	c := 8
	c++

	d := 6
	d--

	fmt.Println(c)
	fmt.Println(d)

	// 关系运算符 (打印true 和 false)  ==  !=  >  <  >=  <=

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)

	// 逻辑运算符   &&  ||  !(取反)

	m := 2
	n := 3

	// 不能输出
	if m == 2 && n == 2 {
		fmt.Println("m == 2 && n == 2")
	}

	// 能输出
	if m == 2 || n == 2 {
		fmt.Println("m == 2 || n == 2")
	}

	// 能输出
	if !(m == 2 && n == 2) {
		fmt.Println("!(m == 2 || n == 2)")
	}

}
