package main

import "fmt"

type MySlice[T int | string | float64] []T

// 1. 泛型方法 , 使用时直接 对象.方法名称(参数) 即可
func (s MySlice[T]) Sum() T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

// 2. 泛型函数
func Add[T int | string | float64](a T, b T) T {
	return a + b
}

func main() {

	// 泛型函数与方法

	// 泛型用途 ： -1 泛型类型  -2 泛型函数  -3 泛型方法

	var s1 MySlice[int] = []int{1, 2, 3, 4}
	fmt.Println(s1.Sum())

	var s2 MySlice[string] = []string{"miao", "wang"}
	fmt.Println(s2.Sum())

	var s3 MySlice[float64] = []float64{3.14, 6.28}
	fmt.Println(s3.Sum())

	fmt.Println("-------")

	fmt.Println(Add[int](1, 2))
	fmt.Println(Add(2, 4)) // 当然这里也可以自动类型推导
	fmt.Println(Add[string]("a", "b"))
	fmt.Println(Add[float64](3.1, 6.2))

}
