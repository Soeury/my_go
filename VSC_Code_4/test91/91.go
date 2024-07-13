package main

import "fmt"

// 数据类型 int64的别名
type wow int64

// 数据类型前加上~表示支持这个数据类型的所有别名 , 比如下面的 int64 有一个别名 wow , 加了~表示支持这个 wow 别名
type MyType interface {
	int8 | ~int64 | float32 | float64 | byte
}

func GetMaxNum[T MyType](a T, b T) T {
	if a > b {
		return a
	}

	return b
}

func main() {

	// 自定义泛型约束 - 当约束的类型过多的时候 采用直接引用接口的方式，需要使用的数据类型全部在接口中

	var a wow = 10
	var b wow = 90

	ret := GetMaxNum[wow](a, b)
	fmt.Println(ret)
}
