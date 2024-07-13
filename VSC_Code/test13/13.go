package main

import "fmt"

func main() {

	// for 循环语句

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	j := 10

	for j >= 1 {
		fmt.Println(j)
		j--
	}

	// 5 * 5 矩阵

	for m := 1; m <= 5; m++ {
		for n := 1; n <= 5; n++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// 乘法表

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d \t", j, i, i*j)
		}
		fmt.Println()
	}

	// break continue 语句

	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Print(i)
	}

	fmt.Println()

	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Print(i)
	}

	fmt.Println()

}
