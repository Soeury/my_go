package main

import "fmt"

type error interface {
	Error() string
}

type M_error struct {
	msg  string
	code int
}

func (p M_error) Error() string {
	return fmt.Sprint("错误类型: ", p.msg, " 状态码: ", p.code)
}

// 这里返回的 temp 是 M_error 类型的对象 , 可以这么理解 : error 接口中有一个 Error 方法 , Error 方法的接收者是一个 M_error 对象
// 所以这里的 M_error 属于多态 ， 即 M_error 既是 M_Error 又是 error , 所以返回的 temp 的类型可以属于 error 类型
func test(i int) (int, error) {
	if i != 0 {
		temp := M_error{"非零数字", 404}

		return i, temp // 这里用了&符 , 后面的断言需要将括号内的类型更改为 *M_error , 但是 , 为什么不用&符也可以正常输出?
	}

	return i, nil
}

func main() {

	// 自定义错误类型 ： 通过自定义 Error() string 函数来实现自定义的错误
	var number int
	fmt.Scanln(&number)

	num, err := test(number)

	if err != nil {
		fmt.Println(err)

		m_err, ok := err.(M_error) // ok_idiom 的断言方式 , ok = true 代表 err 属于 M_error 类型

		if ok {
			fmt.Println("m_err.msg =", m_err.msg)
			fmt.Println("m_err.code =", m_err.code)
		}
	}

	fmt.Println("num =", num)
}
