package main

import "fmt"

type User struct {
	name string
	age  int
	sex  string
}

func main() {

	// 结构体属于值类型的数据 , 我们通常都会定义一个指针类型的结构体变量

	var user1 User
	user1.name = "chen"
	user1.age = 18
	user1.sex = "male"

	fmt.Println("user1 = ", user1)
	fmt.Println("-----------------------")

	user2 := user1
	user2.name = "suan"
	fmt.Println("user1 = ", user1)
	fmt.Println("user2 = ", user2)

	// 要换成引用类型可以使用指针  此时 user3 是一个指向 user1 类型为 *User 的指针

	fmt.Println("-----------------------")
	var user3 *User = &user1

	user3.age = 25
	fmt.Println("user1 = ", user1)
	fmt.Println("user3 = ", *user3)

	// 也可以直接用 new(struct_name) 创建一个指针 ， 类型是 *struct_name
	// 用 new 创建的变量都是指针

	user4 := new(User)

	(*user4).name = "qian" // = user4.name
	(*user4).age = 60      // = uesr4.age
	(*user4).sex = "male"  // = user4.sex

	fmt.Println("-----------------------")
	fmt.Println("user4 = ", *user4)
	fmt.Printf("user4's type = %T\n", user4)

}
