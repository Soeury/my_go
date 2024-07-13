package main

import "fmt"

// 自定义泛型结构体 template
type MyStruct[T string | int] struct {
	Id   T
	Name string
}

// 自定义泛型接口
type MyInterface[T int | string] interface {
	Print(data T)
}

func main() {

	// 定义泛型类型

	//  type  自己写的  固定死的

	// 1. 自定义泛型切片
	type mSlice[T string | int] []T

	var s1 mSlice[int] = []int{25, 36}
	fmt.Println(s1)

	// 2. 自定义泛型 map
	type mMap[KEY int | string, VALUE float64 | bool | int] map[KEY]VALUE

	var m1 mMap[string, int] = map[string]int{
		"apple": 100,
		"grape": 99,
	}
	fmt.Println(m1)

	// 3. 自定义泛型通道
	//type Mychan[T int | string] chan T

	// 4. 注意 : 在前面定义了约束，最终还是要看后面的那个类型，后面的类型才是底层类型 ， 如像下面这个定义其实没有意义，(没有用到T)
	type miao[T int | string] int

	var a miao[int] = 90
	fmt.Printf("a type = %T\n", a)
	var b miao[string] = 90
	fmt.Printf("b type = %T\n", b)
	//var c miao[string] = "abc"

}
