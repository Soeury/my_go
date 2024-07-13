package main

import (
	"fmt"
	"strings"
)

func main() {

	// strings 包常用函数

	str := "yuanFang,niZenmekan?"
	fmt.Println("str =", str)
	fmt.Println("----------")

	// Contains( string1 , "string2") 是否包含指定字符or字符串 ， 返回 Bool 值 ， (精确查找，即 string1 中是否完全包含 string2)
	fmt.Println(strings.Contains(str, "n")) // 注意 ， 双引号引起来代表字符串，这里必须使用双引号
	fmt.Println(strings.Contains(str, "ni"))
	fmt.Println(strings.Contains(str, "kn"))
	fmt.Println("----------")

	// ContainsAny( string1 , "string2") 是否包含指定字符or字符串的任意一个部分 ，返回bool值 (模糊查找 ， 即string1中是否包含string2的某一部分)
	fmt.Println(strings.ContainsAny(str, "kn")) // str中是否包含  'k'   or   'n'   or   'kn'
	fmt.Println(strings.ContainsAny(str, "bc"))
	fmt.Println("----------")

	// Count(string1 , string2) 统计指定内容 在字符串中出现的次数 (即返回sting2 在 string1 中出现的次数)
	num := strings.Count(str, "n")
	fmt.Println("num =", num)
	num2 := strings.Count(str, "kn")
	fmt.Println("num2 =", num2)
	fmt.Println("----------")

	// HasPrefix(string1 , string2) 某字符串是否以指定字符串开头 , 返回 bool 值 , (即 string1 是否以 string2 开头)
	fmt.Println(strings.HasPrefix(str, "yuan"))
	fmt.Println(strings.HasPrefix(str, "http"))
	fmt.Println("----------")

	// HasSuffix(string1 , string2) 某字符串是否以指定字符串结尾 ， 返回 bool 值 (即 string1 是否以 string2 结尾)
	fmt.Println(strings.HasSuffix(str, ".mp4"))
	fmt.Println(strings.HasSuffix(str, "kan?"))
	fmt.Println("----------")

	// Index(string1 , string2) 寻找指定字符串第一次出现的位置 ， 找到了返回下标 ， 没找到返回 -1
	index1 := strings.Index(str, "an")
	fmt.Println("index1 =", index1)
	index2 := strings.Index(str, "kn")
	fmt.Println("index2 =", index2)
	fmt.Println("----------")

	// IndexAny(string1 , string2) 寻找指定字符串中的任意字符第一次出现的位置 ， 找到了返回下标，未找到返回 -1
	index3 := strings.IndexAny(str, "kn") // 找 'k'  or  'n'  or  'kn'  中最早出现的字符串的位置
	fmt.Println("index3 =", index3)
	index4 := strings.Index(str, "bc")
	fmt.Println("index4 =", index4)
	fmt.Println("----------")

	// LastIndex(string1 , string2) 寻找指定字符串最后一次出现的位置 ， 找到了返回下标，找不到返回 -1
	index5 := strings.LastIndex(str, "kan")
	fmt.Println("index5 =", index5)
	index6 := strings.LastIndex(str, "kn")
	fmt.Println("index6 =", index6)
	fmt.Println("----------")

	// LastIndexAny(string1 , string2) 寻找指定字符串的任意字符在原字符串中最后出现的位置 ， 找到了返回下标，为找到返回 -1
	index7 := strings.LastIndexAny(str, "kn") // 即找  'k'  or  'n'  or  'kn' 中最后出现的字符串的位置
	fmt.Println("index7 =", index7)
	index8 := strings.Index(str, "bc")
	fmt.Println("index8 =", index8)
	fmt.Println("----------")

	// Join(slice , string2) 用指定符号进行字符串拼接 ( 两个参数:  Join(切片 , 需要拼接的字符串) )
	s1 := []string{"a", "b", "c", "d"}
	str1 := strings.Join(s1, "-")
	fmt.Println("str1 =", str1)

	str2 := strings.Join(s1, " ")
	fmt.Println("str2 =", str2)
	fmt.Println("----------")

	// Split(string1 , string2) 按照指定符号进行分割 ( 两个参数：  Split(需要分割的字符串 ， 分割字符串))
	str3 := strings.Split(str1, "-")
	fmt.Println("str3 =", str3)

	str4 := strings.Split(str2, "-")
	fmt.Println("str4 =", str4)
	fmt.Println("----------")

	// Repeat(string1 , int) 重复拼接自己  (即 将string1 重复拼接 int 次 ， 返回拼接后的字符串 )

	str5 := strings.Repeat("A", 6)
	fmt.Println("str5 =", str5)
	fmt.Println("----------")

	// Replace(string1 , be_replaced_string , replaced_string , int) 替换原有字符串中的某一个写字符串 , 饭hi替换后的字符串
	// Replace(原始字符串 ， 被替换的字符 ， 替换的目的字符 ， 替换的个数) ， int = -1 代表全部替换
	str6 := strings.Replace(str, "n", "*", -1)
	fmt.Println("str6 =", str6)
	fmt.Println("----------")

	// ToUpper(string1) 将字母全部替换成大写
	str7 := strings.ToUpper(str)
	fmt.Println("str7 =", str7)

	str8 := strings.ToUpper(str6)
	fmt.Println("str8 =", str8)
	fmt.Println("----------")

	// ToLowe(string1 将字母全部替换成小写)
	str9 := strings.ToLower(str)
	fmt.Println("str9 =", str9)
	fmt.Println("----------")

}
