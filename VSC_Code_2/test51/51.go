package main

import "fmt"

type OBJ interface {
	// 一个空接口 , 可以用名称来表示 , 也可以用 interface{} 来表示
}

type Dog struct {
	name  string
	color string
}

type Cat struct {
	name   string
	weight int
}

// test1 == test2
func test1(p OBJ) {
	fmt.Println(p)
}

func test2(p interface{}) {
	fmt.Println(p)
}

func main() {

	// 空接口 可以看成是一种是数据类型 存储任意类型的数据

	var a1 OBJ = Dog{"John", "red"}
	var a2 OBJ = Cat{"Alice", 18}
	var a3 OBJ = 100
	var a4 OBJ = "hello"
	var a5 OBJ = 3.14

	fmt.Println("------------------------")
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)

	fmt.Println("------------------------")
	test1(a1)
	test1(a2)
	test2(a3)
	fmt.Println("------------------------")

	// 空接口定义map
	var m1 = make(map[string]OBJ) // var m1 = make(map[string]interface{})

	m1["A"] = 100
	m1["B"] = 3.14
	m1["C"] = "world"
	m1["D"] = Dog{"Robot", "white"}

	fmt.Println(m1)

	// 空接口定义切片
	var s1 = make([]OBJ, 0, 10)
	s1 = append(s1, a1, a2, a3, a4, a5)
	s1 = append(s1, 66, "mine")

	fmt.Println(s1)

	fmt.Println("------------------------")

	// 以下是一些疑问 :

	var temp OBJ // 初始默认为 nil
	var m2 = make([]interface{}, 0, 15)

	if temp == nil {
		fmt.Println("temp == nil")
	} else {
		fmt.Println("temp != nil")
	}

	for i := 0; i < 10; i++ {
		fmt.Scan(&temp)
		m2 = append(m2, temp)
	}

	fmt.Println(m2)
}
