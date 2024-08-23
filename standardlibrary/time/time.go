package main

import (
	"fmt"
	"time"
)

func main() {
	timeStampDemo()
}

func timeDemo() {
	now := time.Now() // 获取当前时间
	fmt.Printf("current time is %v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, month, day, hour, minute, second)
}

func timeStampDemo() {
	now := time.Now()
	timestamp := now.Unix()  // 秒级时间戳
	milli := now.UnixMilli() // 毫秒级时间戳
	micro := now.UnixMicro() // 微秒级时间戳
	nano := now.UnixNano()   // 纳秒级时间戳
	fmt.Println(timestamp, milli, micro, nano)
}
