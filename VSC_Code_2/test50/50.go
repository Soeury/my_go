package main

import "fmt"

type Animal interface {
	eat()
	sleep()
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func (dog Dog) eat() {
	fmt.Println(dog.name, "is eating")
}

func (dog Dog) sleep() {
	fmt.Println(dog.name, "is sleeping")
}

func (cat Cat) eat() {
	fmt.Println(cat.name, "is eating")
}

func (cat Cat) sleep() {
	fmt.Println(cat.name, "is sleeping")
}

func m_test(a Animal) {
	fmt.Println("test")
	a.eat()
	a.sleep()
}

func main() {

	// 接口实现多态
	// Animal 中有 cat 和 dog ， cat 既是 cat ， 也是 animal ， 这就是多态，即一个事物多种形态

	dog := Dog{"John"}
	dog.eat()
	dog.sleep()
	m_test(dog) // dog 既是 Dog , 也是 Animal , dog 属于接口的实现类对象 ， 可以传入 test 函数

	fmt.Println("-----------------")

	cat := Cat{"Alice"}
	cat.eat()
	cat.sleep()
	m_test(cat) // cat 既是 Cat , 也是 Animal , cat 属于接口的实现类对象 ， 可以传入 test 函数

	fmt.Println("-----------------")

	var animal Animal
	animal = dog // 此时 animal 相当于 dog , 可以使用 Animal 中的方法 , 但是animal属于 Animal 类型， 所以不能使用 Dog 结构体中的变量
	fmt.Println(animal)
	animal.eat()
	animal.sleep()
	m_test(animal)
	//  其中 animal.name 是不允许的

	fmt.Println("-----------------")

	animal = cat
	fmt.Println(animal)
	animal.eat()
	animal.sleep()

	fmt.Println("-----------------")
}
