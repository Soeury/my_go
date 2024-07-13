package main

import "fmt"

type I interface {
	// 一个名为 I 的空接口
}

func assert_string(i interface{}) {
	s := i.(string)
	fmt.Println(s)
}

func assert_int(i interface{}) {
	v, ok := i.(int)

	if ok {
		fmt.Println("i is type 'int' ")
		fmt.Println(v)
	} else {
		fmt.Println("i is not type 'int' ")
	}
}

// 若断言类型同时实现了 switch 中的多个类型，则取第一个 case , 所以空接口类型放在最后面进行 case
func test(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int type")
	case string:
		fmt.Println("string type")
	case float64:
		fmt.Println("float type")
	case bool:
		fmt.Println("bool type")
	case nil:
		fmt.Println("nil type")
	case I:
		fmt.Println("I type")
	default:
		fmt.Println("unknow type!")
	}
}

func main() {

	// 接口断言 实现类型转换 将接口类型的对象转化为对应的数据类型
	// 假如当前接口类型的对象是16 ，准备参与数学运算，则必须转换为 int 类型才能参加运算 ， 接口断言实现了这个类型转换的过程

	/*
		    1.  t := i.(type)
		        变量名称 := 接口对象名称.(数据类型)
		        若 i 的实际类型不是括号中的数据类型，无操作则输出 panic:

			2.  v , ok := i.(type)     ->     ok-idiom 形式
			    ok == true  i 的实际类型是括号中的数据类型
				ok != true  i 的实际类型不是括号中的数据类型

			3.  switch i.(type) {
			case 1:
				    ...
			case 2:
				    ...
			case 3:
				    ...
			default:
				    ...
			}
	*/

	a := "hello"
	assert_string(a)

	/*
		b := 18
		assert_string(b)
	*/

	c := 25
	assert_int(c)

	d := "world"
	assert_int(d)

	fmt.Println("-------------------")

	e := 66
	f := "chen"
	g := 3.14
	h := true
	var n I

	test(e)
	test(f)
	test(g)
	test(h)
	test(n)

}
