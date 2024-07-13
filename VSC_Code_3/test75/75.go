package main

import (
	"fmt"
	"os"
)

func listDir(filePath string, tabInt int) {

	tab := "|--"
	dir := filePath
	fileInfo, _ := os.ReadDir(dir)

	for i := 0; i < tabInt; i++ {
		tab = "    " + tab
	}

	for _, file := range fileInfo {
		filename := dir + "\\" + file.Name()
		fmt.Printf("%s%s\n", tab, filename)

		if file.IsDir() {
			listDir(filename, tabInt+1)
		}
	}
}

func main() {

	// 递归遍历文件夹
	listDir("D:\\M_GO\\GO_work\\VSC_Code_3", 0)

}
