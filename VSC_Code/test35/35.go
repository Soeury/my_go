package main

import "fmt"

func main() {

	//  map遍历 ：
	// map 是无序的，每次遍历的结果都不一样，不能通过index来遍历，只能通过key来遍历
	// map 属于引用类型
	// len(map)获取长度 ， 表示map 当前所拥有的key的数量
	// key 可以是所有能够进行比较的类型，如 bool  float  int  string...

	// var map1 = map[string]int{"go": 100, "java": 90, "c": 80}

	map1 := make(map[string]int)
	map1["go"] = 100
	map1["java"] = 90
	map1["c"] = 80

	fmt.Println(map1)

	for k, v := range map1 {
		fmt.Println(k, v)
	}

	map2 := map1

	map2["go"] = 0
	fmt.Println(map2)
	fmt.Println("length of map2 (means the num of key that exists) = ", len(map2))

}
