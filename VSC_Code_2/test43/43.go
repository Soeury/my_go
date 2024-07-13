package main

import "fmt"

// 普通结构体
type Student struct {
	name string
	age  int
	sex  string
}

// 匿名字段的结构体(不建议使用)
type Teacher struct {
	string
	int
}

func main() {

	var s1 = Student{name: "chen", age: 18, sex: "male"}
	fmt.Println("s1 = ", s1)

	// 匿名结构体
	s2 := struct {
		name string
		age  int
		sex  string
	}{
		name: "zhao",
		age:  90,
		sex:  "female",
	}

	fmt.Println("s2 = ", s2)

	// 匿名字段
	t1 := Teacher{"sun", 66}
	fmt.Println("t1 = ", t1)

	// 匿名字段中的类型代表变量名称，直接用 . 拿出来数据即可
	fmt.Println("t1.string = ", t1.string)
	fmt.Println("t1.int = ", t1.int)

}
