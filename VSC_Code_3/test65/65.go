package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 随机数 math/rand 包

	/*
		    注意 :  go 语言下 在并发下 ， 设置随机数种子 rand.Seed() 和 rand.Intn() or rand.Int() 不能共存
			所以，之后使用随机数种子，请写以下代码:

			    seed := rand.Now().Unix()
				src := rand.NewSource(seed)
				rnd := rand.New(src)

				rnd为我们之后需要使用的随机数生成器

	*/

	// 1.  rand.Seed(num)  +  rand.Int()  生成一个非常大的随机数，只有改变num才会改变生成的数字
	//rand.Seed(101)
	num1 := rand.Int()
	fmt.Println(num1)

	// 2. rand.Seed(num)  +  rand.Intn(num1) + num2  生成一个 [num2 , num1 + num2 - 1] 之间的数字
	//rand.Seed(109)
	num2 := rand.Intn(10) + 20 // 表示生成[20 , 29] 之间的数字
	fmt.Println(num2)

	// 3. 使用时间戳来更新 rand.Seed(time.Now().Unix())
	//rand.Seed(time.Now().Unix())

	seed := time.Now().Unix()
	src := rand.NewSource(seed)
	rnd := rand.New(src)

	for i := 0; i < 10; i++ {
		num := rnd.Intn(51) + 50 // 生成[50 , 100] 之间的数字
		fmt.Println("num =", num)
	}
}
