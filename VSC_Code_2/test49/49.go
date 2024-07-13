package main

import "fmt"

// 定义一个接口
type USB interface {
	output()
	input()
}

// 定义一个对象
type Mouse struct {
	name   string
	weight int
}

// 接口的实现
func (mouse Mouse) output() {
	fmt.Println("mouse output")
}

func (mouse Mouse) input() {
	fmt.Println("mouse input")
}

// 接口的调用
func test(u USB) {
	u.output()
	u.input()
}

type Key struct {
	name   string
	weight int
}

func (key Key) output() {
	fmt.Println("keyboard output")
}

func (key Key) input() {
	fmt.Println("keyboard input")
}

func main() {

	/*
		接口实现
		1. 定义一个接口

		    type name interface {
				methoud_1()
				methoud_2()
				methoud_3()
				......
			}
	*/

	m := Mouse{"luoji", 20}
	test(m)

	k := Key{"leishe", 100}
	test(k)

	var usb USB
	usb = k

	fmt.Println(usb)
	// fmt.Println(usb.name)   usb为USB类型 ， 将一个 Key 对象赋给 usb , usb 不能调用

	var str string = m.name
	fmt.Println("m.name = ", str)

}
