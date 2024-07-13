package main

import "fmt"

func main() {

	// go数据类型转换 -> 显示转换    new val := type_new (old val)

	a := 5.0
	b := int64(a)

	fmt.Printf("a num : %.1f  a type : %T\n", a, a)
	fmt.Printf("b num : %d    b type : %T\n", b, b)

	// 报错  ->  整形不能转换成 bool 类型
	// c := bool(b)
	// fmt.Print("c num : %d  c type : %T\n", c, c)

	// 数据转换范围问题

	var d int32 = 300
	e := int64(d) // int32 -> int64  小范围 -> 大范围  √
	f := int8(d)  // int32 -> int8   大范围 -> 小范围  ×

	fmt.Printf("d num : %d  d type : %T\n", d, d)
	fmt.Printf("e num : %d  e type : %T\n", e, e)
	fmt.Printf("f num : %d   f type : %T\n", f, f) // 精度丢失

}
