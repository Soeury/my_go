package main

import (
	"fmt"
	"strconv"
)

func main() {

	// strconv包常用函数

	s1 := "true" // 字符串除了 "1"  "true" 之外解析成 bool 类型都是 flase
	s2 := "100"
	s3 := "25"

	// 1. ParseBool(string_data)  , return bool_data , error : 将字符串类型解析成 bool 类型
	b1, _ := strconv.ParseBool(s1)
	fmt.Printf("%T\n", b1)
	fmt.Println(b1)

	fmt.Println("-----")

	// 2. FormatBool(bool_data)  , return string_data : 将 bool 类型还原为字符串类型
	str1 := strconv.FormatBool(b1)
	fmt.Printf("%T\n", str1)
	fmt.Println(str1)

	fmt.Println("-----")

	// 3. ParseInt(string_data , base_int , bit_size)  , return int_data , error : 将字符串解析成 int 类型
	i1, _ := strconv.ParseInt(s2, 10, 64) // (字符串 ， 进制 ， int的类型-一般都写 64)
	fmt.Printf("%T\n", i1)
	fmt.Println(i1)

	fmt.Println("-----")

	// 4. FormatInt(int_data , base_int )  ,  return string_data : 将 int 类型还原为字符串类型
	str2 := strconv.FormatInt(i1, 10)
	fmt.Printf("%T\n", str2)
	fmt.Println(str2)

	fmt.Println("-----")

	// 5. Atoi(string_data)  , return int_data , error : string -> int 快速转换
	i2, _ := strconv.Atoi(s3)
	fmt.Printf("%T\n", i2)
	fmt.Println(i2)

	fmt.Println("-----")

	// 6. Itoa(int_data) , return string_data :  int -> string 快速转换

	str3 := strconv.Itoa(i2)
	fmt.Printf("%T\n", str3)
	fmt.Println(str3)

}
