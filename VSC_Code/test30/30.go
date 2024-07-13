package main

import "fmt"

func main() {

	// 值类型 : 数组   副本修改之后不影响母本
	arr := [4]int{1, 2, 3, 4}
	arr2 := arr

	fmt.Println(arr, arr2)

	arr2[0] = 66
	fmt.Println("arr2修改之后:")
	fmt.Println(arr, arr2)

	// 引用类型 : 切片   副本修改之后影响母本

	fmt.Println()
	s := []int{1, 2, 3, 4}
	s2 := s
	fmt.Println(s, s2)

	s2[0] = 99
	fmt.Println("s2修改之后:")
	fmt.Println(s, s2)

	fmt.Printf("s_address  : %p \ns2_address : %p", s, s2)

}
