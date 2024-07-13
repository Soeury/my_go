package main

import "fmt"

func main() {

	// map 容器的一些用法

	// 1. 创建map
	// var map1 map[int]string
	map1 := make(map[int]string)

	// 2. 通过键值给map赋值
	map1[1] = "apple"
	map1[2] = "banana"
	map1[3] = "comment"
	map1[5] = "digital"

	fmt.Println(map1)

	// 3. 获取map数据
	fmt.Println(map1[1])
	fmt.Println(map1[2])
	fmt.Println(map1[3]) //  根据key , 获取value
	fmt.Println(map1[4]) //  键值不存在，则打印默认值，这里string的默认值是""

	// 4. ok-idiom 方式判断该键值是否存在 , map[key]返回两个值   value , ok := map[key]
	// 其中 value 是当键值为key时的value值，ok 表示该键值是否存在，不存在则 ok = flase  , 存在则 ok = true

	value, ok := map1[5]

	if ok {
		fmt.Println("key 存在 , map_value = ", value)
	} else {
		fmt.Println("key 不存在")
	}

	// 5. 修改 map 数据
	map1[1] = "arounding"
	fmt.Println(map1)

	// 6. 删除数据 delete
	// delete(map_name , key)   意思是将 map_name 中键值为 key 的数据

	delete(map1, 1) // 删除 map1 中键值为 1 的数据
	fmt.Println(map1)

	// 7. map长度 :  len(map_name)
	fmt.Println("map1 length = ", len(map1))

	// 8. 添入数据 键值不存在就是填入，存在就是修改
	map1[10] = "january"
	fmt.Println(map1)

}
