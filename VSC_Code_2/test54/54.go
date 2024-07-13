package main

import "fmt"

/*
    注意全局作用域和函数内部两种不同的命名方式

	1-  全局作用域 :  type  new_type_name  type_name
	2-  函数内部 :    type  new_type_name  =  type_name

*/

type m_int int

func main() {

	//  type 别名

	//  1. 在全局作用域下 type 一个新的数据类型  type  new_type_name  type_name
	//  (此时相当于双胞胎 ， 双胞胎是两个人)
	//  此时 这两个数据类型 属于两种数据类型 ，要进行相关的操作必须转换成一种类型

	var a m_int = 10
	var b int = 20

	//  c := a + b  注意：这里会报错，是不被允许的操作 ， 因为 a 和 b 属于两种数据类型，必须强制转换成一种数据类型才可以相加

	c := b + int(a)
	fmt.Println("a + b = ", c)

	//  2. 在函数体内部 type 一个新的数据类型  type  new_type_name = type_name
	//  (此时相当于给一个人起了两个名字 ， 这两个名字指向的是同一个人)
	//  此时 这里个数据类型本质属于一种数据类型 ， 进行相关的操作时不需要进行强制转换

	type m_float64 = float64

	var x m_float64 = 3.1
	var y float64 = 2.1

	z := x + y // 这里不会报错 ， 是被允许的操作
	fmt.Println("x + y = ", z)

}
