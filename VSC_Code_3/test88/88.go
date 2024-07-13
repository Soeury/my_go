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

func main() {

	// 反射修改对象的值 ， 操作的对象只能是指针
	// 步骤:
	// 1-  value := reflect.ValueOf(&对象)
	// 2-  newValue := value.Elem()
	// 3-  判断: newValue.CanSet() ?= true
	// 4-  newValue.Set_Type(修改后的值)

	var num float64 = 3.14
	fmt.Println("before num :", num)

	value := reflect.ValueOf(&num)
	newValue := value.Elem()

	fmt.Println(newValue.Kind())
	fmt.Println(newValue.Type())
	fmt.Println(newValue.CanSet())

	if newValue.CanSet() {
		newValue.SetFloat(6.28)
	}

	fmt.Println("after num :", num)
	fmt.Println("-----------------")

	user := User{"chen", 18, "male"}
	fmt.Println("before :", user)

	value2 := reflect.ValueOf(&user)
	newValue2 := value2.Elem()

	if newValue2.CanSet() {
		newValue2.FieldByName("Name").SetString("zhao")
		newValue2.FieldByName("Age").SetInt(22)
		newValue2.FieldByName("Sex").SetString("female")
	}

	fmt.Println("after :", user)

}
