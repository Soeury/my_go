package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {

	// seek 实现断点传续
	srcFile := "D:\\M_GO\\GO_work\\VSC_Code_3\\test71\\photo\\rabbit.jpg"
	desFile := "D:\\M_GO\\GO_work\\VSC_Code_3\\test73\\des.jpg"
	tempFile := "D:\\M_GO\\GO_work\\VSC_Code_3\\test73\\temp.txt"

	file1, _ := os.Open(srcFile)
	file2, _ := os.OpenFile(desFile, os.O_CREATE|os.O_RDWR, 0777)
	file3, _ := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, 0777)

	defer file1.Close()
	defer file2.Close()

	// 处理临时文件中的数据
	buf := make([]byte, 10, 1024)
	file3.Seek(0, io.SeekStart)
	n, _ := file3.Read(buf)
	countStr := string(buf[:n])
	fmt.Println("countStr =", countStr)

	count, _ := strconv.ParseInt(countStr, 10, 64)

	// seek 源文件和目的文件的位置
	file1.Seek(count, io.SeekStart)
	file2.Seek(count, io.SeekStart)

	// 传输数据
	bufData := make([]byte, 10, 1024)
	total := int(count)

	for {
		readnum, err1 := file1.Read(bufData)
		if err1 == io.EOF {
			fmt.Println("trans over!")
			file3.Close()
			os.Remove(tempFile) // 移除 temp 文件  os.Remove(path)
			break
		}

		writenum, _ := file2.Write(buf[:readnum])

		// 更新临时文件
		total = total + writenum
		file3.Seek(0, io.SeekStart)
		file3.WriteString(strconv.Itoa(total))
	}

}
