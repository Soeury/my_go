package main

import (
	"fmt"
	"os"
)

func main() {

	// 什么是错误
	// 错误: 可能出现问题的地方出现了问题
	// 异常: 不该出现问题的地方出现了问题

	file, error := os.Open("aaa.txt")

	if error != nil {
		fmt.Println("Found error:")
		fmt.Println(error)
		return
	}

	fmt.Println(file.Name()) // 获取文件名称  file.Name()
}
