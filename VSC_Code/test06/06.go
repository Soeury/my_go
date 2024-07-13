package main

import "fmt"

func main() {

	// 字符串类型

	var str string = "hello,vsCode"
	fmt.Println(str)

	var c1 = 'A' // 字符的数据类型是 int
	var c2 = "A" // 字符串的数据类型是 string

	fmt.Printf("c1 type : %T , c1 num : %d\n", c1, c1)
	fmt.Printf("c2 type : %T , c2 num : %s\n", c2, c2) //  只有Printf可以格式化数据...

	c3 := '中'
	fmt.Printf("c3 type : %T , c3 num : %d\n", c3, c3)

	// 字符串拼接 (+)

	s1 := "apple and "
	s2 := "grape"
	fmt.Println(s1 + s2)

	// 转义字符
	fmt.Printf("my name is \"vsCode\"\n") // 转义双引号  \"
	fmt.Println("a1:18\ta2:20")           // 制表符 \t

}
