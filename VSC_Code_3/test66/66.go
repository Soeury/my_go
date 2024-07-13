package main

import (
	"fmt"
	"time"
)

func main() {

	// 定时器与时间判断

	// 1. time.Tick(time._) 实现定时器
	ticker := time.Tick(time.Second)

	for t := range ticker {
		fmt.Println(t)
	}

	fmt.Println("---------------------")

	// 2. for 循环 + sleep 实现定时器

	for i := 0; i < 10; i++ {
		fmt.Println(time.Now())
		time.Sleep(time.Second)
	}

	fmt.Println("---------------------")

	// 3. 时间判断

	now := time.Now()
	fmt.Println(now)

	later := now.Add(time.Hour)
	fmt.Println(later)

	sub_time := now.Sub(later)
	fmt.Println(sub_time)

	is_before := now.Before(later)
	fmt.Println(is_before)

	is_after := now.After(later)
	fmt.Println(is_after)

	is_equal := now.Equal(later)
	fmt.Println(is_equal)

}
