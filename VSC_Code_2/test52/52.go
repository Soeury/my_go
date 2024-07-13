package main

import "fmt"

type A interface {
	test1()
}

type B interface {
	test2()
}

// 在接口嵌套中，嵌套接口默认继承了被嵌套接口的所有方法 ， 这里 C 接口继承了 A B 接口的所有方法
type C interface {
	A
	B
	test3()
}

type Dog struct {
	// 暂时让它为空
}

func (dog Dog) test1() {
	fmt.Println("test1")
}

func (dog Dog) test2() {
	fmt.Println("test2")
}

func (dog Dog) test3() {
	fmt.Println("test3")
}

func main() {

	// 接口嵌套实现接口继承  ->  接口可以继承 ， 甚至可以多继承
	dog := Dog{}

	dog.test1()
	dog.test2()
	dog.test3()

	var pa A = dog
	pa.test1()
	// 结构体对象赋值给哪个接口 ， 只能使用该接口下的方法 ， 这里 pa 只能使用 test1 方法 ， 因为 A 接口只有 test1 方法

	var pb B = dog
	pb.test2()
	// 这里 pb 只能使用 test2 方法 ， 因为 B 接口只有 test2 方法

	var pc C = dog
	pc.test1()
	pc.test2()
	pc.test3()
	// 在接口嵌套中，嵌套接口默认继承了被嵌套接口的所有方法 ， 这里 C 接口继承了 A B 接口的所有方法
	// 所以 pc 可以使用 A B C 中所用的方法

}
