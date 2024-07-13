package main

import (
	"fmt"
	"os"
)

// 文件操作相关 :  os包  :   Stat  Mkdir  MkdirAll  Remove  RemoveAll  Creat ...

func main() {

	// 创建目录(文件夹)与文件

	// 1. 创建文件夹:  err := os.Mkdir(路径\\文件名 ， os.ModePerm) ,  从最后两个 \\ 开始往前的所有路径必须提前存在

	err := os.Mkdir("D:\\M_GO\\GO_work\\VSC_Code_3\\test68\\m_file", os.ModePerm)

	if err != nil {
		fmt.Println(err)
		// return
	}
	fmt.Println("m_file 创建成功!")

	// 2. 创建层级文件夹:  err := os.MkdirAll (路径\\文件名1\\文件名2 ， os.ModePerm) 这里 c 前的路径不完全存在，所以不能用 Mkdir

	err1 := os.MkdirAll("D:\\M_GO\\GO_work\\VSC_Code_3\\test68\\m_file2\\a\\b\\c", os.ModePerm)

	if err1 != nil {
		fmt.Println(err)
		// return
	}
	fmt.Println("m_file2/a/b/c 创建成功!")

	// 3. 删除文件夹 ： err := os.Remove(路径) ， 这个文件夹必须是空的

	err3 := os.Remove("D:\\M_GO\\GO_work\\VSC_Code_3\\test68\\m_file")

	if err3 != nil {
		fmt.Println(err3)
		// return
	}
	fmt.Println("m_file 删除成功!")

	// 4. 删除非空文件夹(慎用) : err := os.RemoveAll(路径)

	err4 := os.RemoveAll("D:\\M_GO\\GO_work\\VSC_Code_3\\test68\\m_file2")

	if err4 != nil {
		fmt.Println(err4)
		// return
	}
	fmt.Println("m_file2/a/b/c 删除成功!")

	// 5. 创建文件 :   file , _ err := os.Creat(路径)   无论文件里放了什么 ， 每次creat之后会将一个新的文件覆盖在原文件上
	// f, err5 := os.Create("D:\\M_GO\\GO_work\\VSC_Code_3\\test68\\my.txt")
	// fmt.Println(f.Name())
	// fmt.Println(err5)
	// fmt.Println("my.txt 文件创建成功!")

	os.Remove("D:\\M_GO\\GO_work\\VSC_Code_3\\test68\\my.txt")
	fmt.Println("my.txt 已移除!")

}
