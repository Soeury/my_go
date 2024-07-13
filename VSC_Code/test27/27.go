package main

import "fmt"

func main() {

	// 切片 slice : 可以看成是可变长度的数组
	// 数组：
	arr := [4]int{1, 4, 6, 9}
	fmt.Println(arr)

	// 切片：
	var s []int
	fmt.Printf("%T %T\n", s, arr)

	if s == nil {
		fmt.Println("切片s为空, 默认值是nil")
	}

	s2 := []int{1, 4, 2, 5}
	fmt.Println(s2)
	fmt.Println(s2[3])

	// make创建切片  name := make([]type , length , capacity)
	// 最后是根据length来决定切片中数据的个数
	fmt.Println()

	s3 := make([]int, 5, 10)
	fmt.Println(s3)
	fmt.Println("s3.len = ", len(s3))
	fmt.Println("s3.cap = ", cap(s3))

	s3[0] = 100
	fmt.Print(s3)

	// 切片扩容 :
	// 1. slice1 = append (slice1 , val1 , val2)
	// 2. slice1 = append (slice1 , slice2...)    其中slice2...代表将切片s2中的数据全部取出，放进切片slice1中

	s1 := make([]int, 0, 5)
	s1 = append(s1, 1, 2)

	fmt.Println()
	fmt.Println()
	fmt.Println(s1)
	s1 = append(s1, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(s1)

	s4 := make([]int, 3, 5)

	for i := 0; i < 3; i++ {
		s4[i] = 1
	}

	s1 = append(s1, s4...)
	fmt.Println(s1)

}
