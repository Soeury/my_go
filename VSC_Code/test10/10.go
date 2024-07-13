package main

// Scanf 输入问题

import "fmt"

func main() {

	var a int64
	var b float64
	var c int64
	var d int64
	var e string

	// Scanln输入两个数据 ，没有要求一定要输入两个数据，输入一个数据之后直接回车 另一个数据会按默认值输出

	fmt.Println("输入两个数： 1 , 整数  2 , 小数 ：")
	fmt.Scanln(&a, &b)

	fmt.Printf("a 的值是 %d\n", a)
	fmt.Printf("b 的值是 %f\n", b)

	// Scan 要求输入几个数据，一定要输入几个数据，非则程序会一直等待

	fmt.Println("输入两个整数 : ")
	fmt.Scan(&c, &d)
	fmt.Printf("c 的值是 %d\n", c)
	fmt.Printf("d 的值是 %d\n", d)

	// Scanf 除了输入相关的  (& + 变量)  外 , 还需要在前面说明 (% + 类型)
	// ?  ?  ?

	fmt.Println("请输入一个字符串: ")
	fmt.Scanf("%s \n", &e)
	fmt.Printf("e 的值是 %s", e)

}
