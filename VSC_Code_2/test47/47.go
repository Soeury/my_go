package main

import "fmt"

// 方法初识
/*
	    1. 方法 : 类似于函数，但与函数有所不同，方法只能由指定的接收者来调用，其他人不可调用，只要接收者不同，方法命名可以相同

	    func (receiver type) methoud_name() {
			content
		}

		创建一个 type 类型的 receiver ， 然后 receiver.methoud_name() 代表使用这个方法

		2. 函数 ： 可以任意被调用，命名不能冲突

		func name() {
			content
		}

		3. 疑问 :
            1- 只能有一个接收者吗?
			2- 接收者只能是 struct 吗?
*/

type Dog struct {
	name string
	age  int
}

func (dog Dog) eat() {
	fmt.Println("Dog , eating")
}

type Cat struct {
	name string
	age  int
}

func (cat Cat) eat() {
	fmt.Println("Cat , eating")
}

func main() {

	dog := Dog{"John", 2}
	dog.eat()

	cat22 := Cat{"Amy", 1}
	cat22.eat()

}
