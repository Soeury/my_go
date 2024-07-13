package main

import "fmt"

// 冒泡排序:
func bubble_sort(arr [6]int, s string) {
	for i := 1; i <= len(arr); i++ {
		flag := 0

		for j := 0; j < len(arr)-i; j++ {
			if s == "asc" {
				if arr[j] > arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					flag = 1
				}
			}

			if s == "desc" {
				if arr[j] < arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
					flag = 1
				}
			}
		}

		if flag == 1 {
			fmt.Println(arr)
		} else {
			break
		}
	}
}

func main() {

	// 数组 值类型 - 值传递

	arr1 := [4]int{2, 5, 7, 9}
	fmt.Println("arr1 : ", arr1)

	arr2 := arr1 // 改变了arr2 不影响arr1
	fmt.Println("arr2 : ", arr2)
	arr2[2] = 100
	fmt.Println("arr2 : ", arr2)

	fmt.Println()

	// 数组排序： 冒泡排序 ：

	arr := [6]int{4, 2, 6, 1, 7, 3}
	bubble_sort(arr, "asc")
	fmt.Println("-------------")
	bubble_sort(arr, "desc")

}
