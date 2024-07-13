package main

import (
	"fmt"
	"sync"
	"time"
)

var ticket = 10
var mutex sync.Mutex

func saleTicket(name string) {
	for {
		mutex.Lock() // 上锁
		if ticket > 0 {
			time.Sleep(time.Millisecond * 500)
			fmt.Println(name, " total :", ticket)
			ticket--
		} else {
			mutex.Unlock() // 解锁
			fmt.Println("sale out!")
			break
		}
		mutex.Unlock() // 解锁
	}
}

func main() {

	// 临界资源修改的问题
	// goroutine 中对临界资源处理不当 ， 会导致数据不一致的问题

	a := 1

	go func() {
		a = 4
		fmt.Println("func - a =", a)
	}()

	a = 6
	fmt.Println("main - a - before =", a)
	time.Sleep(time.Second * 2)
	fmt.Println("main - a =", a)

	fmt.Println("------------")

	// 售票问题 - 临界资源处理不当

	// 解决办法 : 上锁 + 解锁 (高并发不推荐使用这种方法) 使得在某个时间段内，只有一个 goroutine 来访问这个共享数据
	//     1- 先创建一个锁 :  var mutex sync.Mutex
	//     2- 上锁 : mutex.Lock()
	//     3- 解锁 : mutex.Unlock()

	go saleTicket("door - 1")
	go saleTicket("door - 2")
	go saleTicket("door - 3")
	go saleTicket("door - 4")

	time.Sleep(time.Second * 10)

}
