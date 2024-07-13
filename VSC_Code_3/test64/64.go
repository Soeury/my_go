package main

import (
	"fmt"
	"time"
)

func get_time() {

	now := time.Now()
	fmt.Println(now)

	year := now.Year()
	month := now.Month()
	dat := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, dat, hour, minute, second)
}

func main() {

	//  时间 - 时间戳

	// 1. 自定义函数获取时间 (比较麻烦 ， 所以有了 2. )
	get_time()
	fmt.Println("-------------------")

	// 2. Format 方式将当前时间格式化为字符串 :    当前时间.Format(格式按照 2006-01-02 15:04:05 来参照)

	now := time.Now()
	m_time1 := now.Format("2006-01-02 15:04:05") // 这是24小时表示法
	fmt.Println(m_time1)
	m_time2 := now.Format("2006-01-02 03:04:05 PM") // 这是12小时表示法
	fmt.Println(m_time2)
	m_time3 := now.Format("2006/01/02 15:04:05")
	fmt.Println(m_time3)
	m_time4 := now.Format("15:04 2006/01/02")
	fmt.Println(m_time4)
	fmt.Println("-------------------")

	// 3. 将字符串格式化为当前时间，之后就可以用   时间.Month() 方式将时间信息全部取出
	// -1. loc , _ := time.LoadLocation("Asia/Shanghai")  返回 时区 和 error
	// -2. m_time , _ := time.ParseInLocation("2006/01/02 15:04:05" , string_time , loc)   , 返回string_time对应的时间和error

	loc, _ := time.LoadLocation("Asia/Shanghai")
	m_time, _ := time.ParseInLocation("2006/01/02 15:04:05", "2005/10/15 14:28:42", loc)
	fmt.Println(m_time)

	month := m_time.Month()
	day := m_time.Day()

	fmt.Println("m_time's month =", month)
	fmt.Println("m_time's day =", day)
	fmt.Println("-------------------")

	// 4. 时间戳 ， 从 1970/1/1  08:00:00 到当前时间的总毫秒数，也被称为 Unix 时间戳

	//     1- 时间 -> 时间戳(毫秒数) :   num_ms := time.Now().Unix()
	timestamp := time.Now().Unix()
	fmt.Println(timestamp)

	//     2- 时间戳(毫秒数) -> 时间 :   time := time.Unix(num_ms , 0)
	time := time.Unix(timestamp, 0)
	fmt.Println(time)
}
