package main

import "fmt"

func main() {

	// map & slice 案例
	// 1. 每个map 分别存储一个人的信息， name sex  address age
	// 2. 将每个map 放入一个切片中
	// 3. 将切片按顺序打印输出

	user1 := make(map[string]string)
	user1["name"] = "A"
	user1["age"] = "18"
	user1["addr"] = "China"
	user1["sex"] = "male"

	user2 := map[string]string{"name": "B", "age": "16", "addr": "American", "sex": "female"}
	user3 := map[string]string{"name": "C", "age": "20", "addr": "African", "sex": "male"}

	s := make([]map[string]string, 0, 5)
	s = append(s, user1, user2, user3)

	fmt.Println(s)

	for _, v := range s {
		fmt.Printf("v[name]:%s\n", v["name"])
		fmt.Printf("v[age]:%s\n", v["age"])
		fmt.Printf("v[addr]:%s\n", v["addr"])
		fmt.Printf("v[sex]:%s\n", v["sex"])
		fmt.Println("----------------")
	}

}
