package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
	Sex  string
}

func (user User) Say(m string) {
	fmt.Println("User say :", m)
}

func (user User) PrintInfo() {
	fmt.Println("user name:", user.Name, " | user age:", user.Age, " | user sex:", user.Sex)
}

func miao() {
	fmt.Println("miao miao")
}

func fruit(p string) {
	fmt.Println("i love", p)
}

func add(n int, m int) int {
	return n + m
}

func main() {

	// 反射调用方法
	// 1- value := reflect.ValueOf(对象)
	// 2- value.MethodByName("方法名字").Call(参数)

	p := User{"chen", 18, "God"}
	value := reflect.ValueOf(p)

	// 方法没有参数就写 .Call(nil)
	value.MethodByName("PrintInfo").Call(nil)

	// 方法有参数就 .Call(slice)
	// slice := make([]reflect.Value , 参数个数)
	// slice[0] = reflect.ValueOf(参数)
	s1 := make([]reflect.Value, 1)
	s1[0] = reflect.ValueOf("iloveyou")
	value.MethodByName("Say").Call(s1)

	fmt.Println("-----------------")

	// 反射调用函数
	// 1- name := reflect.ValueOf(函数名)
	// 2- name.Call(参数)
	// 3- 参数没有就填 nil , 有就填入切片 ， 有返回值需要有切片来接收数据

	// 1- 无参写 nil
	v1 := reflect.ValueOf(miao)
	v1.Call(nil)

	// 2- 有参用切片
	v2 := reflect.ValueOf(fruit)
	s2 := make([]reflect.Value, 1)
	s2[0] = reflect.ValueOf("apple")
	v2.Call(s2)

	// 3- 有返回值用切片接收
	v3 := reflect.ValueOf(add)
	s3 := make([]reflect.Value, 2)
	s3[0] = reflect.ValueOf(1)
	s3[1] = reflect.ValueOf(2)

	ret := v3.Call(s3)
	fmt.Println(ret[0].Interface()) // ！ ！ ！

}
