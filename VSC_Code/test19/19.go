package main

import "fmt"

// 求和函数
func get_sum(i int) int {

	if i == 1 {
		return 1
	}

	return get_sum(i-1) + i
}

func add(a int) {
	fmt.Println("函数中的a:", a)
}

func main() {

	// 递归函数调用自身

	sum := get_sum(4)
	fmt.Println("sum: ", sum)

	// defer 延迟函数 : 当函数执行到最后是，这些defer语句会被延迟到最后按照逆序执行

	fmt.Println("1")
	defer fmt.Println("2") // 最后执行
	fmt.Println("3")
	defer fmt.Println("4") // 倒数第二执行
	fmt.Println(5)
	defer fmt.Println("6") // 倒数第三执行

	fmt.Println("---")

	// add(a) 语句已经执行过了，准备打印输出，但是被defer延迟打印输出 ， 所以打印输出的是10，不是11
	a := 10
	fmt.Println("a += 1 之前的值:", a)
	defer add(a) // 倒数第四执行
	a += 1
	fmt.Println("a += 1 之后的值:", a)

}
