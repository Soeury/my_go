package main

import (
	"fmt"
	"os"
)

func main() {

	// 向文件中写入数据 ： -打开文件，建立连接    -写入数据    -关闭文件

	// 1. 打开文件   file , err := os.OpenFile(路径 ， 属性 ， 权限)
	file, err1 := os.OpenFile("D:\\M_GO\\GO_work\\VSC_Code_3\\test70\\fruit.txt", os.O_RDONLY|os.O_WRONLY|os.O_APPEND, os.ModePerm)

	if err1 != nil {
		fmt.Println(err1)
		return
	}

	defer file.Close()

	// 2.1   file.Write([]byte slice) (n int , err error) 写入数据(字节写入)
	// 注意 ， 在 openfile 里面设置了 os.O_APPEND 才可以将我们写入的数据拼接到原有文件后 ， 否则是直接将原内容清空后再写入新数据
	bs := []byte{65, 66, 67, 68, 69}

	num, err2 := file.Write(bs)
	fmt.Println(num)
	fmt.Println(err2)
	fmt.Println("-----")

	// 2.2   file.WriteString(string_data) (n int , err error) 写入数据(字符串写入)

	num3, err3 := file.WriteString("grape and peach")
	fmt.Println(num3)
	fmt.Println(err3)
	fmt.Println("-----")

}
