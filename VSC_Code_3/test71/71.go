package main

import (
	"fmt"
	"io" //   io.EOF    io.Copy    ...
	"os"
)

func m_copy(source string, des string, bufSize int) {

	// 原文件
	sourceFile, err1 := os.Open(source)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer sourceFile.Close()

	// 目的文件
	desFile, err2 := os.OpenFile(des, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	defer desFile.Close()

	// for 循环实现  原文件 -> 缓冲区 -> 目的文件

	var buf = make([]byte, bufSize) // 可以不需要设置容量 ？ ？ ？

	for {
		// 读取
		n1, err3 := sourceFile.Read(buf)

		if err3 == io.EOF || n1 == 0 {
			fmt.Println("copy over!")
			break
		}

		// 写出
		_, err4 := desFile.Write(buf[:n1])

		if err4 != nil {
			fmt.Println("found err:", err4)
			return
		}
	}
}

func io_copy(source string, des string) {

	// 原文件
	sourceFile, err5 := os.Open(source)
	if err5 != nil {
		fmt.Println(err5)
		return
	}
	defer sourceFile.Close()

	// 目标文件
	desFile, err6 := os.OpenFile(des, os.O_WRONLY|os.O_CREATE, 0777) // 这里 0777 和 os.ModePerm 可以互换 ， 本质是一样的
	if err6 != nil {
		fmt.Println(err6)
		return
	}
	defer desFile.Close()

	// io :  size , err := io.Copy(目标文件 ， 原文件)
	written, err7 := io.Copy(desFile, sourceFile)
	if err7 != nil {
		fmt.Println("found err:", err7)
		return
	}

	fmt.Println("File's Size:", written)
}

func os_copy(source string, des string) {

	fileBuf, err8 := os.ReadFile(source) // 不建议使用，因为 os.ReadFile(原文件) 是一次性读取到缓冲区 ， 缓冲区可能没有那么大的空间

	if err8 != nil {
		fmt.Println("found err:", err8)
	}

	fmt.Println("copy over!")
	os.WriteFile(des, fileBuf, os.ModePerm)

}

func main() {

	// 文件复制

	// 1. 手动实现
	source := "D:\\M_GO\\GO_work\\VSC_Code_3\\test71\\photo\\rabbit.jpg"
	des := "D:\\M_GO\\GO_work\\VSC_Code_3\\test71\\copy_1.jpg"
	m_copy(source, des, 1024)

	// 2. os.Copy()
	source = "D:\\M_GO\\GO_work\\VSC_Code_3\\test71\\photo\\rabbit.jpg"
	des = "D:\\M_GO\\GO_work\\VSC_Code_3\\test71\\copy_2.jpg"
	io_copy(source, des)

	// 3. os.ReadFile()   +    os.WriteFile()
	source = "D:\\M_GO\\GO_work\\VSC_Code_3\\test71\\photo\\rabbit.jpg"
	des = "D:\\M_GO\\GO_work\\VSC_Code_3\\test71\\copy_3.jpg"
	os_copy(source, des)

}
