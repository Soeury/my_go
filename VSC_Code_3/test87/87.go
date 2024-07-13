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

func reflectInfo(input interface{}) {

	// 注意 ： 把这里的 getType 看成一个大的结构体 ， getValue 看成我们这个结构体的一个对象 ， 就不难理解了

	// 获取参数类型 :  传入的对象  ->  结构体(getType)
	getType := reflect.TypeOf(input)

	tName := getType.Name()
	tKind := getType.Kind()

	fmt.Println("get type is :", tName)
	fmt.Println("get kind is :", tKind)

	// 获取值 :  传入的对象  ->  镜像对象(getValue)  ->  for(num)通过镜像对象将传入对象中的值拿出来
	getValue := reflect.ValueOf(input)
	fmt.Println("getValue is :", getValue)
	fmt.Println("getType is :", getType)

	fmt.Println("------------------------------")

	// 获取结构体中的字段 ：
	/*
		1- 首先获取类型 ：  getType := reflect.TypeOf(传入的对象)
		2- 获取字段个数 ：  num := getType.NumFiled()
		3- 获取字段 :      filed := getType.Filed(index)
		4- 获取值 ：       value := getValue.Filed(index).Interface()
	*/

	for i := 0; i < getType.NumField(); i++ {
		filed := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Println("字段名称:", filed.Name, "  字段类型:", filed.Type, "  字段数值:", value)
	}

	fmt.Println("------------------------------")

	// 获取方法(同上)
	for i := 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Println("方法名称:", method.Name, "方法类型:", method.Type)
	}
}

func main() {

	// 反射获取变量信息
	u1 := User{"chen", 18, "male"}
	reflectInfo(u1)

}
