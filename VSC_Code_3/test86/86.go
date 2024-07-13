package main

import (
	"fmt"
	"time"
)

func main() {

	// timer定时器 ： timer 包下的一些有关 channel 的用法

	// 1. time.Tick(time.second) 与 channel 无关
	//ticker := time.Tick(time.Second)

	//for t := range ticker {
	//	fmt.Println(t)
	//}

	// timer -> 一次性的时间触发   ticker -> 持续性的触发事件
	// 2.  timer := time.NewTimer(time.second)
	//     stop := timer.Stop()
	timer := time.NewTimer(time.Second * 2)
	fmt.Println(time.Now())
	timeChan := timer.C
	fmt.Println(<-timeChan)

	fmt.Println("-----------------------")

	timer2 := time.NewTimer(time.Second * 5)

	go func() {
		<-timer2.C // 这里会阻塞 5s
		fmt.Println("goroutine over!")
	}()

	time.Sleep(time.Second * 2) // 休眠 2s
	stop := timer2.Stop()       // 这里的 stop 会让前面的阻塞停止 ，继续向后执行 ， 返回 bool 类型

	if stop {
		fmt.Println("timer2 stop!")
	} else {
		fmt.Println("timer2 not stop!")
	}

	fmt.Println("-----------------------")
	// 3. timerChan := time.After(time.second * n)  这里直接返回 channel 不像上面那样  timer.C 把通道拿出来

	timerChan := time.After(time.Second * 3)
	fmt.Println(time.Now())
	fmt.Println(<-timerChan)

	fmt.Println("-----------------------")
	// 4.  time.AfterFunc(time.second * n , func_name)  没有返回值 ， 这里表示在指定时间后触发函数执行

	time.AfterFunc(time.Second*2, get_sum)
	time.Sleep(time.Second * 3)

}

func get_sum() {
	fmt.Println("i am get_sum")
}
