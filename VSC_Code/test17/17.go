package main

import "fmt"

func updata(arr2 [4]int) {

	fmt.Println("arr2原始数据:", arr2)
	arr2[0] = 9
	fmt.Println("arr2修改后的数据:", arr2)
}

func updata2(s2 []int) {
	fmt.Println("修改前的数据:", s2)
	s2[0] = 8
	fmt.Println("修改后的数据:", s2)
}

func main() {

	// 值传递 ：  传递的是副本 ， 修改副本数据 ， 不会改动原始数据
	// 对象:  int  float64  bool  string  array  struct

	// 数组 :     名称 := [个数]数据类型{val1 , val2 ...}
	arr := [4]int{1, 2, 3, 4}
	fmt.Println("调用前的数据:", arr)

	updata(arr)
	fmt.Println("调用后的数据:", arr)

	// 引用传递 ： 传递的是地址，导致多个变量指向同一空间
	// 对象:  slice  map  channel

	// 切片 ： 可以扩容的数组
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println("调用前的数据", s1)

	updata2(s1)
	fmt.Println("调用后的数据:", s1)
}
