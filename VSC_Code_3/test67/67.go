package main

import (
	"fmt"
	"os"
)

func main() {

	// go 语言 io流
	// i : input   读入
	// o : output  读出

	// file 类是在 os包中的 ， 里面有 FileInfo 接口，封装了一系列文件相关的方法

	// (os: 绝对路径原来是这么写的......从盘符开始 ， 用斜右下双斜杠)
	m_fileInfo, err := os.Stat("D:\\M_GO\\GO_work\\VSC_Code_3\\test67\\psg.txt")

	if err != nil {
		fmt.Println(err) //
		return           // 不要忘记了 return
	}

	fmt.Println(m_fileInfo.Name())    // 文件名字
	fmt.Println(m_fileInfo.Mode())    // 文件读写属性:  -  ---  ---  --- (这里 - rwx rwx rwx 表示这是一个文件，可读 可写 可执行)
	fmt.Println(m_fileInfo.Size())    // 文件大小
	fmt.Println(m_fileInfo.IsDir())   // 是否是一个目录
	fmt.Println(m_fileInfo.ModTime()) // 修改时间
	fmt.Println(m_fileInfo.Sys())     // 更多信息

}
