package main

import "fmt"

/*
    go 语言结构体嵌套
	1. 模拟继承性
        type A struct {
			name1  type
			name2  type
		}

		type B struct {
			A               // 匿名字段表示继承 ， 这样B就可以直接访问A中的变量，即 B.name1 = _ , 不需要 B.A.name1 = _
		}

		2. 模拟聚合性
		type C struct {
			name3  type
			name4  type
		}

		type D struct {
			name5  C       // 此时不是继承关系，D不能直接访问C中的变量，必须  D.C.name3 = _
		}

*/

// 定义一个 "父类" 结构体
type Person struct {
	name string
	age  int
}

// 定义一个 "子类" 结构体
type Student struct {
	Person        // 匿名字段表示继承
	school string // Student 结构体自身拥有的变量
}

func main() {

	// go 语言使用结构体嵌套实现继承关系

	p1 := Person{"chen", 18}
	fmt.Println(p1)

	s1 := Student{Person{"zhao", 14}, "School_A"}
	fmt.Println(s1)

	fmt.Println("s1.Person.name = ", s1.Person.name)
	fmt.Println("s1.name = ", s1.name)

	// 提升字段 :
	// 对于结构体 Student 来说， Person 属于匿名字段 ， Person 中的 name 和 age 就叫做提升字段
	// 提升字段可以通过变量名字直接访问，不需要使用中间的 Person 结构体

	var s2 Student
	s2.name = "sun"
	s2.age = 56
	s2.school = "School_B"

	fmt.Println("s2 = ", s2)

}
