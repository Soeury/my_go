package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// seek 接口 - 可以看成是光标的位置 ：  file_name.Seek(offset int , io.[seekStart , seekCurrent , seekend])

	// 1. 建立连接
	file, err1 := os.OpenFile("D:\\M_GO\\GO_work\\VSC_Code_3\\test72\\book.txt", os.O_RDWR, os.ModePerm)

	if err1 != nil {
		fmt.Println(err1)
		return
	}

	defer file.Close()

	// 2. seek操作文件中的光标 : 注意 ： 取出的是光标右边的字符串 ， 取完之后光标向右移动到该取出的字符串的右侧 ，
	// 字符串长度取决于切片大小
	file.Seek(2, io.SeekStart) // (光标向右的偏移量 ， 相对的位置)
	buf := make([]byte, 2)
	file.Read(buf)
	fmt.Println(string(buf))

	file.Seek(2, io.SeekCurrent)
	file.Read(buf)
	fmt.Println(string(buf))

	file.Seek(0, io.SeekEnd) // seekEnd 的目的是为了让光标到达文件尾部 ， 从而向后拼接数据
	file.WriteString(" andforeverlove")

	file.Seek(0, io.SeekStart) // file.Seek(0, io.SeekStart) 之后再拼接数据会将原来的数据覆盖
	file.WriteString("updata the data!")

}
