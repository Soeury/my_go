package main

import (
	"fmt"

	"M_GO/GO_work/VSC_Code_2/test45/pojo"
)

func main() {

	// 结构体跨包调用 (总算把问题解决了.........解决文档参考：CSDN：后端收藏夹："两种愉快的方式使用go.mod引用自己开发的包")

	/*
	   1. 在顶层文件夹下建立一个 go.mod 文件并且执行 go mod tidy , 这个 go.mod 文件中 module + 最顶层的文件夹名称
	   之后，这是唯一的一个 go.mod 文件 ， 此后不再在子文件夹或者文件下建立 go.mod 文件 ，
	   要导入任何自定义的包的路径必须从这个最顶层文件夹开始，具体例子如上:

	   2. 给文件内的各个包都建立一个 go.mod 文件 ， 其中 ， 被调用的包只需  ( moudle + 包名称   go 1.22.2 ) 即可
	   主动调用其他包的包需要另外进行操作 ， 即加入 require 和 replace

	   require 被调用包名 v0.0.0
	   replace 被调用包名 => 被调用包相对于主动调用包的 '相对路径'

	*/

	var user1 pojo.User
	user1.Name = "chen"
	user1.Age = 18
	user1.Sex = "male"

	fmt.Println(user1)

}
