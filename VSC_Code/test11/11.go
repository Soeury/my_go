package main

import "fmt"

func main() {

	// if语句

	var score int = 85

	if score >= 90 && score <= 100 {
		fmt.Println("A")
	}

	if score >= 80 && score < 90 {
		fmt.Println("B")
	}

	if score >= 70 && score < 80 {
		fmt.Println("C")
	}

	if score >= 60 && score < 70 {
		fmt.Println("D")
	}

	// 嵌套if 密码验证
	var a int
	var b int
	var num int = 202204

	fmt.Println("请输入密码：")
	fmt.Scanln(&a)

	if a == num {
		fmt.Println("请再次输入密码：")
		fmt.Scanln(&b)

		if b == num {
			fmt.Println("登录成功")
		} else {
			fmt.Println("密码错误，登录失败")
		}

	} else {
		fmt.Println("登陆失败")
	}

}
