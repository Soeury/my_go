package main

import "fmt"

func main() {

	// 位运算符   &   |   ^   >>   <<

	a := 60 // 00111100
	b := 13 // 00001101

	c := a & b
	fmt.Printf("c num_10 : %d  ,  c num_2 : %b\n", c, c)

	c = a | b
	fmt.Printf("c num_10 : %d  ,  c num_2 : %b\n", c, c)

	c = a ^ b
	fmt.Printf("c num_10 : %d  ,  c num_2 : %b\n", c, c)

	a = 60
	c = a << 2
	fmt.Printf("c num_10 : %d ,  c num_2 : %b\n", c, c)

	a = 60
	c = a >> 2
	fmt.Printf("c num_10 : %d  ,  c num_2 : %b\n", c, c)

	// 赋值运算符   =  -=  +=  *=  %=  /=  &=  |=  ^=  ...

	d := int64(21)
	var e int64

	e = d
	e += d
	fmt.Println(e)

	e = d
	e -= d
	fmt.Println(e)

	//...
}
