package main

import "fmt"

func main() {

	// 数组   name := [size]type{val , val...}

	nums := [4]int{1, 2, 3, 4}

	l := len(nums)
	c := cap(nums)

	fmt.Println("数组长度:", l)
	fmt.Println("数组容量:", c)

	fmt.Println(nums)

	fmt.Println(nums[0])
	fmt.Println(nums[1])
	fmt.Println(nums[2])
	fmt.Println(nums[3])

	nums[0] = 9
	fmt.Println(nums[0])

	fmt.Println("-----")

	// 数组初始化

	// 1. 常规定义 var name = [size]type{val , val ...}
	var arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)

	// 2. 快速定义  name := [size]type{val , val...}
	arr2 := [5]int{6, 7, 8, 9, 10}
	fmt.Println(arr2)

	// 3.不确定数据个数：[...]
	arr3 := [...]int{3, 6, 9, 12}
	fmt.Println(arr3)
	fmt.Println("arr3数据个数是:", len(arr3))

	// 4.下标定义数组，未被定义的直接采用默认值输出
	arr4 := [5]int{1: 10, 4: 40}
	fmt.Println(arr4)

	// 数组遍历
	// 1. for
	// 2. for range   返回两个值 (index , val)

	arr5 := [5]int{2, 5, 1, 6, 8}

	for i := 0; i < len(arr5); i++ {
		fmt.Println(arr5[i])
	}

	fmt.Println("---------")

	for i := 0; i < len(arr5); i++ {
		fmt.Printf("%d ", arr5[i])
	}

	fmt.Println()
	fmt.Println("---------")

	for index, val := range arr5 {
		fmt.Println(index, val)
	}

}
