package main

import "fmt"

func main() {

	s := make([]int, 0, 5)
	s = append(s, 1, 2)
	fmt.Println(s)

	s2 := make([]int, 3, 5)
	s2 = append(s2, 1, 2)
	fmt.Println(s2)

	// 如果要通过 slice := make([]type , length , cap)方式新建一个切片，并且开始没有数据之后要加入数据，必须:
	// slice := make([]type , 0 , cap)

	s3 := make([]int, 0, 10)
	s3 = append(s3, 1, 2, 3, 4, 5)
	fmt.Println(s3)
}
