package main

import "fmt"

type Animal struct {
	name string
	age  int
}

func (animal Animal) eat() {
	fmt.Println(animal.name, "is eating")
}

type Cat struct {
	Animal
}

type Dog struct {
	Animal
}

// Cat 类重写的 Animal 类的方法 :   命名相同，方法的接收者不同
func (cat Cat) eat() {
	fmt.Println(cat.name, "is eating - cat")
}

// Cat 类自己的方法
func (cat Cat) miao() {
	fmt.Println(cat.name, "is miao-ing")
}

func main() {

	// 方法重写  子类重写父类的方法

	// Dog 结构体继承了 Animal 结构体 ， 所以 Dog 可以使用 父类的方法
	dog1 := Dog{Animal{"John", 2}}
	dog1.eat()

	a1 := Animal{"animal_Alice", 12}
	a1.eat()

	cat1 := Cat{Animal{"Amy", 1}}

	// 这里 Cat 重写了父类的 eat 方法 ， 所以这里调用的是重写后的方法
	cat1.eat()

	// 这里调用的是 Cat 自身的方法
	cat1.miao()

}
