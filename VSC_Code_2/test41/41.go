package main

import "fmt"

//  结构体定义
//  struct_name 区分大小写(权限),首字母小写表示私有，大写表示公开

//   type struct_name struct {
//   val_name type
//   val_name type
//   }

type User struct {
	name string
	age  int
	sex  string
}

func main() {

	// 结构体创建方式 1.  var  name  struct_name

	var user1 User
	user1.name = "chen"
	user1.age = 18
	user1.sex = "female"

	fmt.Println(user1)

	// 结构体创建方式 2.  name :=  struct_name{}
	user2 := User{}
	user2.name = "zhao"
	user2.age = 20
	user2.sex = "male"

	fmt.Println(user2)

	// 结构体创建方式 3. name := struct_name{ val_name : _  ,  val_name : _  , ... , }

	user3 := User{
		name: "qian",
		age:  25,
		sex:  "male",
	}

	fmt.Println(user3)

	// 结构体创建方式 4. name := struct_name{ _  ,  _  ,  _  ,}   需要按顺序直接赋值

	user4 := User{"sun", 16, "female"}
	fmt.Println(user4)

	// 结构体获取数据  使用 . 运算符 ：  name.val_name
	fmt.Println("user1.name = ", user1.name)
	fmt.Println("user2.age  = ", user2.age)
	fmt.Println("user3.sex  = ", user3.sex)
}
