package main

import "fmt"

func main() {

	// 深拷贝与浅拷贝
	// 深拷贝 ： 拷贝的是数据本身, 向像值类型的数据都是深拷贝
	// 浅拷贝 ： 拷贝的是数据的地址，导致多个变量指向同一块内存，引用类型的数据，默认都是浅拷贝

	s1 := []int{1, 2, 3, 4}
	fmt.Println(s1)

	s2 := make([]int, 0, 10)

	// for循环拼接数据实现切片深拷贝
	for i := 0; i < len(s1); i++ {
		s2 = append(s2, s1[i])
	}

	fmt.Println("before:")
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println("after:")
	s2[0] = 99
	fmt.Println(s1)
	fmt.Println(s2)

	// copy函数实现深拷贝
	// copy(slice2 , slice1) 意思是将 slice1 复制给 slice2 , 容量不足时只复制一部分，其余保持不变

	fmt.Println("----------")
	s3 := []int{6, 7, 8}
	fmt.Println(s1)
	fmt.Println(s3)

	/*

		fmt.Println("copy(s1 , s3) : ")
		copy(s1, s3)
		fmt.Println(s1)   // 6 7 8 4
		fmt.Println(s3)   // 6 7 8

	*/

	fmt.Println("copy(s3 , s1) : ")
	copy(s3, s1)
	fmt.Println(s1) // 1 2 3 4
	fmt.Println(s3) // 1 2 3

}
