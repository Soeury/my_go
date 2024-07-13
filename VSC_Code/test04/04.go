package main

import "fmt"

var temp int = 18 // 全局变量

func main() {

	fmt.Println(temp)
	var temp int = 20 // 局部变量 (优先考虑 , 可以被再次定义)
	fmt.Println(temp)

	const url string = "www.baidu.com" // 显示定义常量
	const url2 = "www.hqu.edu.cn"      // 隐式定义常量(自动类型推导)
	const a, b, c = 18, "chen", 3.14

	fmt.Println(url)
	fmt.Println(url2)
	fmt.Println(a, b, c)

	// iota 常量计数器  ->   可以看成是某个const组中的常量的编号(从0开始 ， 没有给定常数值的常量默认使用iota值，在每个const中会刷新第一个值)
	const (
		e = iota     // iota 0
		f            // iota 1
		g            // iota 2
		h = "string" // string  (iota 3)
		i            // string  (iota 4)
		j = 100      // 100  (iota 5)
		k            // 100  (iota 6)
		l = iota     // iota 7
		m = iota     // iota 8
	)

	fmt.Println(e, f, g, h, i, j, k, l, m)

	const (
		o = iota // iota 0
		p        // iota 1
		q        // iota 2
		r = iota // iota 3
	)

	fmt.Println(o, p, q, r)
}
