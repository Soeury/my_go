package main

import "fmt"

// 写完了记得先保存...
// 变量创建了一定要使用

// loading workplace 问题：
// 再所在的go文件下新建文件名(lan-control)
// 然后再终端中输入： go mod init lan-control   (回车)
// 之后会出现一个 go.mod 文件 就没有问题了

func main() {

	// 第一种创建变量的方式 ：  var + 名称 + 类型 = 值
	var name string = "chen"
	var age int = 18
	fmt.Println(name, age)

	// 创建多个变量  var( 名称 + 类型    名称 + 类型   ... )
	var (
		word   string
		number int
	)

	word = "hello"
	number = 12
	fmt.Println(word, number)

	// 第二种创建变量的方式 ： 名称 := 值 (表示声明并初始化)   只有要修改此值的话不能再用 := 的方式  只能用 =
	m_name := "zhao"
	m_age := 18
	m_addr := "address"
	fmt.Println(m_name, m_age)
	fmt.Println(m_addr)

	m_name = "ming"
	fmt.Println(m_name)

	fmt.Printf("number: %d , 内存地址: %p", number, &number)

	// Print() 直接输出内容
	// Printf() 支持格式化输出内容, 如 %d  %p  %T  %s ...
	// Println() 会再输出内容后加一个换行符
}
