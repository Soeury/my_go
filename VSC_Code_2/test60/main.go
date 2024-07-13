package main

/*
    import :

	    1. import   "path"   ---    默认导入包
		2. import R "path"   ---    包的别名 ： 通过写明访问该包下的函数
		3. import . "path"   ---    简便使用 ： 直接调用该包下的函数(有可能出现名称冲突问题 ， 不推荐)
		4. import _ "path"   ---    匿名导包 ： 仅让该包执行 init() 函数

*/

import (
	"fmt"

	_ "M_GO/GO_work/VSC_Code_2/test60/test1"
)

func init() {
	fmt.Println("main - init - 1")
}

func main() {

	/*
		init() 函数 ：

			1. 匿名导入包时，执行该包中的 init() 函数 ， 该函数先于 main() 函数自动执行
			2. init() 函数没有参数 ， 没有返回值 ， 不能被调用
			3. 可以有多个 init() 函数 ， 按照包的依赖关系顺序执行(即从最深处导入的包中的 init() 开始执行)
				一个包中同一个文件中的 init() 函数按顺序执行 ，同一个包中多个文件按文件首字母排序顺序执行
			4. 执行顺序:
		        -1 初始化导入的包
				-2 对包中声明的变量进行计算和分配初始值
				-3 执行 init() 函数

	*/

	fmt.Println("main - ")

}
