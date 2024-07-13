package main

import "fmt"

func main() {

	// 集合map : 键值对
	// map[key]value

	// 1. 声明变量初始化map  此时只创建了map1对象这么名字，并没有创建一个map对象，此时 map1 == nil
	var map1 map[int]string

	// 2. make初始化map  这种初始化的方式用的多，简洁明了  make方式创建了一个map对象，此时 map2 != nil
	var map2 = make(map[string]string)
	map4 := make(map[string]int)

	// 3. 声明时创建变量
	var map3 = map[string]int{"go": 100, "java": 80, "python": 40, "c": 60, "r": 20, "php": 0}

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)
	fmt.Println(map4)

	if map1 == nil {
		fmt.Println("map1 == nil")
	} else {
		fmt.Println("map1 != nil")
	}

	if map2 == nil {
		fmt.Println("map2 == nil")
	} else {
		fmt.Println("map2 != nil")
	}

	map2["A"] = "apple"
	map2["B"] = "banana"
	map2["C"] = "cat"
	map2["D"] = "dog"
	map2["E"] = "element"
	fmt.Println(map2)

	map2["A"] = "angry" // 修改了某个 key 对应的 value 之后，会覆盖原来的 value
	fmt.Println(map2)

}
