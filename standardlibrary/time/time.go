package main

import (
	"fmt"
	"time"
)

func main() {
	timeZoneDemo()
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

func timeZoneDemo() {
	//secondsEastOfUTC := int((8 * time.Hour).Seconds())

}
