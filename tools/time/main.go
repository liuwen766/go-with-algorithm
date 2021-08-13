package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//now := time.Now()
	//fmt.Println("当前时间：", now)
	//
	//year := now.Year()     //年
	//month := now.Month()   //月
	//day := now.Day()       //日
	//hour := now.Hour()     //小时
	//minute := now.Minute() //分钟
	//second := now.Second() //秒
	//fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	//
	//timestamp1 := now.Unix()     //时间戳
	//timestamp2 := now.UnixNano() //纳秒时间戳
	//fmt.Println("时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总秒数")
	//fmt.Printf("current timestamp1:%v\n", timestamp1)
	//fmt.Printf("current timestamp2:%v\n", timestamp2)
	//
	//timeObj := time.Unix(timestamp1, 0) //将时间戳转为时间格式
	//fmt.Println("将时间戳转为时间格式:", timeObj)
	//
	//fmt.Println("时间操作:", "Add", "Sub", "Equal", "Before", "After")
	//
	//fmt.Println("定时器")
	////ticker := time.Tick(time.Second)
	////for i := range ticker {
	////	fmt.Printf("现在是第:%s秒\n",i) //每秒都会执行的任务
	////}
	//
	//fmt.Println("time格式化!")
	//// 24小时制
	//fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	//// 12小时制
	//fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	//fmt.Println(now.Format("2006/01/02 15:04"))
	//fmt.Println(now.Format("15:04 2006/01/02"))
	//fmt.Println(now.Format("2006/01/02"))
	//
	//fmt.Println("----------------------------")
	//fmt.Println(now)
	//// 加载时区
	//loc, err := time.LoadLocation("Asia/Shanghai")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//// 按照指定时区和指定格式解析字符串时间
	//timeObj, err = time.ParseInLocation("2006/01/02 15:04:05", "2021/08/02 21:55:20", loc)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(timeObj)
	//fmt.Println(timeObj.Sub(now))

	now := time.Now()
	successSent := func() (bool, error) {
		return false, nil
	}
	fmt.Println(time.Since(now))
	for i := 0; i < 10; i++ {
		fmt.Println(1 + rand.Intn(5))
	}
	rand.Int63n(6)
	if err := Retry(2*time.Second, 3, successSent); err != nil {
		fmt.Println("failed")
	} else {
		fmt.Println("success")
	}
	fmt.Println(time.Since(now))
}

//超时重试机制
type ConditionFunc func() (bool, error)
type RetryError struct {
	n int
}

func (e *RetryError) Error() string {
	return fmt.Sprintf("still failing after %d retries", e.n)
}
func Retry(interval time.Duration, maxRetries int, f ConditionFunc) error {
	if maxRetries <= 0 {
		return fmt.Errorf("maxRetries (%d) should be > 0", maxRetries)
	}
	tick := time.NewTicker(interval)
	defer tick.Stop()
	for i := 0; ; i++ {
		ok, err := f()
		if err != nil {
			return err
		}
		if ok {
			return nil
		}
		if i == maxRetries {
			break
		}
		<-tick.C
	}
	return &RetryError{maxRetries}
}
