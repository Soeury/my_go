package main

import "fmt"

func main() {

	// 通过数组创建切片  slice := arr_name[index_start(include) : index_end(the pre include)]
	// 这样创建的切片的底层数组就是当前数组, len 是切割的数据量, cap是从start到数组末尾的数据量

	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("arr数组:", arr)

	s1 := arr[:9]  // index_start  -  index_8
	s2 := arr[5:]  // index_5  -  index_end
	s3 := arr[2:6] // index_2  -  index_5
	s4 := arr[:]   // index_start  -  index_end

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	fmt.Printf("arr_address:%p\n", &arr)
	fmt.Printf("s1_address:%p\n", s1)
	fmt.Printf("s2_address:%p\n", s2)
	fmt.Printf("s3_address:%p\n", s3)
	fmt.Printf("s4_address:%p\n", s4)

	fmt.Printf("s1_len : %d , s1_cap : %d\n", len(s1), cap(s1))
	fmt.Printf("s2_len : %d , s2_cap : %d\n", len(s2), cap(s2))
	fmt.Printf("s3_len : %d , s3_cap : %d\n", len(s3), cap(s3))
	fmt.Printf("s4_len : %d , s4_cap : %d\n", len(s4), cap(s4))

	// 在没有超过数组arr最开始的容量时，切片和数组指向的是同一块内存
	// 上面打印的内存地址不一样是因为切片指向的是arr数组的一部分，实际上指向的地址还是同一个
	// 在超过切片超过容量后，切片会重新分配一块内存，此时切片和数组指向的就不是同一块内存了

	arr[0] = 38
	fmt.Println("arr数组:", arr)
	fmt.Println(s4)
	fmt.Println(s3)

	s4[1] = 25
	fmt.Println("arr数组:", arr)
	fmt.Println(s4)
	fmt.Println(s3)

	s3[2] = 66
	fmt.Println("arr数组:", arr)
	fmt.Println(s4)
	fmt.Println(s3)

	s1 = append(s1, 1, 1, 1, 1, 1, 1, 1, 1)
	fmt.Println("arr数组:", arr)
	fmt.Println(s1)

	fmt.Printf("arr_address:%p\n", &arr)
	fmt.Printf("s1_address:%p\n", s1)

	// 此时切片s1和数组arr已经不在同一块内存，无论改变arr还是s1，对另一方都不会有影响
	arr[0] = 64
	fmt.Println("arr数组:", arr)
	fmt.Println(s1)

	s1[0] = 99
	fmt.Println("arr数组:", arr)
	fmt.Println(s1)

}
