package main

import "fmt"

type Person struct {
	name    string
	age     int
	sex     string
	address Address
}

type Address struct {
	city  string
	state string
}

func main() {

	// 结构体嵌套
	var p1 Person
	p1.name = "chen"
	p1.age = 18
	p1.sex = "male"
	p1.address = Address{
		city:  "ShangHai",
		state: "China",
	}

	fmt.Println(p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)
	fmt.Println(p1.sex)
	fmt.Println(p1.address)

	fmt.Println("p1.address.city = ", p1.address.city)
	fmt.Println("p1.address.state = ", p1.address.state)

}
