package main

import "fmt"

func main() {

	// 多维数组
	// 1.for遍历
	// 2.for range遍历

	arr := [5][4]int{
		{1, 2, 3, 4},
		{2, 5, 4, 3},
		{3, 4, 5, 6},
		{4, 5, 6, 7},
		{5, 6, 7, 8},
	}

	fmt.Println(arr[0][0])

	fmt.Println("--------")

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}

	fmt.Println("--------")

	for index, val := range arr {
		fmt.Printf("%d- ", index)

		for _, val2 := range val {
			fmt.Printf("%d ", val2)
		}

		fmt.Println()
	}

}
