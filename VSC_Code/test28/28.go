package main

import "fmt"

func main() {

	// 切片底层
	// 可以将切片看出一个指针，指向一块内存地址，地址上存放一个数组
	// 每次添入数据超过容量后，会重新指向一块更大的内存，将原来的数据copy过去，并将原来的容量翻倍

	// 切片相当于指针，指向一个内存块，不需要&打印地址，而数组需要

	s := []int{1, 2, 3}
	fmt.Println(s)
	fmt.Printf("len: %d  cap:%d\n", len(s), cap(s))
	fmt.Printf("%p", s)
	fmt.Println()
	fmt.Println()

	s = append(s, 4, 5, 6)
	fmt.Println(s)
	fmt.Printf("len: %d  cap:%d\n", len(s), cap(s))
	fmt.Printf("%p", s)
	fmt.Println()
	fmt.Println()

	s = append(s, 7, 8, 9)
	fmt.Println(s)
	fmt.Printf("len: %d  cap:%d\n", len(s), cap(s))
	fmt.Printf("%p", s)
	fmt.Println()
	fmt.Println()

	s = append(s, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19)
	fmt.Println(s)
	fmt.Printf("len: %d  cap:%d\n", len(s), cap(s))
	fmt.Printf("%p", s)
	fmt.Println()
	fmt.Println()

	for i := 20; i <= 47; i++ {
		s = append(s, i)
	}
	fmt.Println(s)
	fmt.Printf("len: %d  cap:%d\n", len(s), cap(s))
	fmt.Printf("%p", s)
	fmt.Println()
	fmt.Println()

}
