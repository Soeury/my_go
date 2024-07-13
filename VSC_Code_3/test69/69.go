package main

import (
	"fmt"
	"os"
)

func main() {

	// 读取文件内容

	// 0. fileinfo , err := os.stat(path)  通过 FileInfo.Mode() 找到文件的读写属性(rwx)
	fileInfo, _ := os.Stat("D:\\M_GO\\GO_work\\VSC_Code_3\\test69\\book.txt")
	fmt.Println(fileInfo.Mode())

	// 1. 打开文件，建立连接 :  file , err := os.Open(path)
	file1, err1 := os.Open("D:\\M_GO\\GO_work\\VSC_Code_3\\test69\\book.txt")

	if err1 != nil {
		fmt.Println(err1)
		// return
	}

	fmt.Println(file1.Name())

	// 2. 关闭连接 : 延迟最后执行  :  defer file_name.Close()
	defer file1.Close()

	// 3. 读取文件内容 :  num , err := file.Read([]byte类型的切片) , 返回读取到的 字符数 和 错误
	// 每次读取到了字符时 err = nil , 未读取到字符时 ， 表示到达文件尾 ， err = io.EOF
	// 每次读取的字符数是根据切片的大小来决定的

	var bs = make([]byte, 3, 1024)
	num, m_err := file1.Read(bs) // 将读到的数据放入切片中
	fmt.Println(num)
	fmt.Println(string(bs))
	fmt.Println(m_err)
	fmt.Println("-----")

	num, m_err = file1.Read(bs)
	fmt.Println(num)
	fmt.Println(string(bs))
	fmt.Println(m_err)
	fmt.Println("-----")

	num, m_err = file1.Read(bs)
	fmt.Println(num)
	fmt.Println(string(bs))
	fmt.Println(m_err)
	fmt.Println("-----")

	num, m_err = file1.Read(bs)
	fmt.Println(num)
	fmt.Println(string(bs))
	fmt.Println(m_err)
	fmt.Println("-----")

	// 3. 设置文件打开方式 ， 建立连接
	// (不知道文件的读写属性就使用这种打开方式) : file , err := os.OpenFile(path , 属性|属性 , os.ModePerm)

	//file2, err2 := os.OpenFile("D:\\M_GO\\GO_work\\VSC_Code_3\\test69\\book.txt", os.O_RDONLY|os.O_WRONLY, os.ModePerm)

	//if err2 != nil {
	//	fmt.Println(err2)
	//	// return
	//}

	//fmt.Println(file2.Name())
}
