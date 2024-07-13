package main

import (
	"errors"
	"fmt"
)

// 返回错误类型
func gauge_age(age int) error {

	if age < 0 {
		return errors.New("年龄输入错误")
	}

	fmt.Println(age)
	return nil
}

func main() {

	// 创建自己的错误 error

	/*

	   1.   name := errors.New(content)
	        此时 ， name 变量的内容就是 content

	   2.   name := fmt.Errorf(content)
	        格式化定义错误 ， 此时 ， name 变量的内容就是 content

	*/

	errorinfo := errors.New("我是一个错误")
	fmt.Println(errorinfo)
	fmt.Printf("%T\n", errorinfo) // *errors.errorString

	err1 := gauge_age(-1)

	if err1 != nil {
		fmt.Println(err1)
	}

	err2 := fmt.Errorf("我是一个错误%d", 404)
	fmt.Println(err2)
	fmt.Printf("%T\n", err2)

}
