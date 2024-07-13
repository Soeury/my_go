package main

import "fmt"

func test() (int, int) {
	return 18, 20
}

func main() {

	// 匿名变量 (写完之后记得保存)

	a, _ := test() // 丢弃返回的第二个数   此时 a = 18
	_, b := test() // b = "chen"

	fmt.Println("a , b 's content is : ")
	fmt.Println(a, b)

	// 变量交换 :  a , b = b , a

	var i int = 18
	var k int = 20

	fmt.Println("Before swap :")
	fmt.Printf("i : %d , place : %p\n", i, &i)
	fmt.Printf("k : %d , place : %p\n", k, &k)

	i, k = k, i

	fmt.Println("After swap :")
	fmt.Printf("i : %d , place : %p\n", i, &i)
	fmt.Printf("k : %d , place : %p\n", k, &k)

}
