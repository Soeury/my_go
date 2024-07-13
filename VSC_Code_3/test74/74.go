package main

import (
	"bufio" // 缓冲区提高读写效率
	"fmt"
	"os"
)

func main() {

	// bufio 包 通过缓冲区来提供效率   读入数据 -> 缓冲区 -> 写出数据

	path := "D:\\M_GO\\GO_work\\VSC_Code_3\\test74\\a.txt"
	file, _ := os.OpenFile(path, os.O_RDWR|os.O_APPEND, os.ModePerm)
	defer file.Close()

	// bufio 读入数据   reader := bufio.NewReader(file) , 之后将 newReader 当成正常文件使用即可
	reader := bufio.NewReader(file)
	buf := make([]byte, 1024)
	num1, _ := reader.Read(buf)
	fmt.Println(num1)
	fmt.Println(string(buf[:num1]))

	// bufio 键盘读入    inputData , error :=  reader.ReadString(读入的截至字符)
	// 假设截至字符是 'a' , 输入 contant , 那么 inputData = conta
	// 假设截至字符是 '\n' , 输入任何数据 ， 按下回车之前的所有数据都属于 inputData
	inputReader := bufio.NewReader(os.Stdin)
	readString, _ := inputReader.ReadString('\n')
	fmt.Println("键盘读入的数据是:", readString)

	// bufio 写入  两种 :    -1 Write(slice) 字节     -2 WriteString(字符串)
	writer := bufio.NewWriter(file)
	num2, _ := writer.WriteString(readString)
	fmt.Println(num2)

	writer.Flush() // 刷新缓冲区 ， 再缓冲区未满的情况下将数据写入   writer.Flush()

}
