package test1

import (
	"fmt"

	_ "M_GO/GO_work/VSC_Code_2/test60/test2"
)

func init() {

	fmt.Println("test1 - init - 1")
}

func init() {

	fmt.Println("test1 - init - 2")
}

func init() {

	fmt.Println("test1 - init - 3")
}

func add(m int, n int) int {
	return m + n
}
