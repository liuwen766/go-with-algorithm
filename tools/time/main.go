package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("当前时间：", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Println("时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总秒数")
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)

	timeObj := time.Unix(timestamp1, 0) //将时间戳转为时间格式
	fmt.Println("将时间戳转为时间格式:", timeObj)

	fmt.Println("时间操作:", "Add", "Sub", "Equal", "Before", "After")

	fmt.Println("定时器")
	//ticker := time.Tick(time.Second)
	//for i := range ticker {
	//	fmt.Printf("现在是第:%s秒\n",i) //每秒都会执行的任务
	//}

	fmt.Println("time格式化!")
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))

	fmt.Println("----------------------------")
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err = time.ParseInLocation("2006/01/02 15:04:05", "2021/08/02 21:55:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))

}
