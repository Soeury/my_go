package main

import (
	"encoding/json"
	"fmt"
)

// 二次编码
// 结构体后面的 `json:" 内容 "` 属于二次编码
//   -1  "-" 表示该字段在转换后的 Json 中不会呈现
//   -2  "小写的字段名"  表示该字段名称在 Json 中以小写的形式呈现
//   -3  ",数据类型" 表示叫原先的数据类型换成新的数据类型

// json.Marshal 和  json.MarshalIndent 和  json.Unmarshal 的应用

type Json struct {
	Name    string   `json:"-"`       // -1
	Content []string `json:"content"` // -2
	Time    int      `json:",string"` // -3
	Price   float64  `json:",string"`
	Fine    bool     `json:"fine"`
}

func main() {

	//   1- 通过结构体生成 json 文本
	p := Json{"Tenchen", []string{"C++", "Python", "Go", "Java"}, 60, 3.14, true}

	// 非格式化输出
	ret, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Found err:", err)
	}

	fmt.Println(string(ret))
	fmt.Println("-----------------------")

	// 格式化输出
	ret2, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		fmt.Println("Found err:", err)
	}

	fmt.Println(string(ret2)) // 返回的是切片 ， 转换成文本需要 string([]slice)
	fmt.Println("-----------------------")

	//  2- 通过 map 生成 Json 文本
	m := make(map[string]interface{}, 5) // 表示有5个键值对

	m["name"] = "Tenchen"
	m["content"] = []string{"C++", "Go", "Python", "Java"}
	m["price"] = 3.14
	m["time"] = 60
	m["fine"] = true

	// 非格式化输出
	ret3, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Found err:", err)
		return
	}

	fmt.Println(string(ret3))
	fmt.Println("-----------------------")

	// 格式化输出
	ret4, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		fmt.Println("Found err:", err)
		return
	}

	fmt.Println(string(ret4))
	fmt.Println("-----------------------")

	//  3- json 解析到结构体  这里要注意字段名称相匹配的问题！！！
	//  err := json.Unmarshal([]byte( JsonBuf ), &name)
	JsonBuf := `{"content":["C++","Python","Go","Java"],"Time":"60","Price":"3.14","fine":true}`

	var pre_ret1 Json
	err = json.Unmarshal([]byte(JsonBuf), &pre_ret1)
	if err != nil {
		fmt.Println("Found err:", err)
		return
	}

	fmt.Println(pre_ret1)
	fmt.Println("-----------------------")

	// Json 解析到指定字段的结构体中
	type Json2 struct {
		Content []string `json:"content"`
	}

	var pre_ret Json2
	err = json.Unmarshal([]byte(JsonBuf), &pre_ret) // 要用引用的方式把 存储的变量放进去
	if err != nil {
		fmt.Println("Found err:", err)
		return
	}

	fmt.Println(pre_ret)
	fmt.Println("-----------------------")

	//  4- json 解析到 map 切片中
	m = make(map[string]interface{}, 5)
	err = json.Unmarshal([]byte(JsonBuf), &m)
	if err != nil {
		fmt.Println("Found err:", err)
		return
	}

	fmt.Println(m)
	fmt.Println("-----------------------")

	// 将切片中的每个字段拿出来
	for k, v := range m {
		fmt.Printf("%v\t-------->  %v\n", k, v)

		if k == "fine" {
			fmt.Println(v)
		} else if k == "Time" {
			fmt.Println(v)
		} else if k == "content" {
			fmt.Println(v)
		} else if k == "map" {
			fmt.Println(v)
		} else if k == "Price" {
			fmt.Println(v)
		}

		fmt.Println()
	}

}
