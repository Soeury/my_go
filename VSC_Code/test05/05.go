package main

import "fmt"

func main() {

	// bool 类型

	var temp bool
	fmt.Println(temp) //  bool 类型的数据不初始化 默认是false

	temp = true
	fmt.Println(temp)

	// int 类型 (默认是int64)
	// uint - 无符号，最小是0
	// int - 有符号，可以是负数
	// byte = uint8 (0 - 255)

	var num1 int64 = 289
	var num2 byte = 255
	fmt.Println(num1, num2)

	// float 类型 (默认是float64 , 默认保留6位小数)

	var num3 float64 = 3.141592
	fmt.Printf("%.4f", num3) // 只有Printf可以格式化数据(带%)      %.4f 意思是保留四位小数(会自动四舍五入)

}
