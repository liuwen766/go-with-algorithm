package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		appName := strconv.FormatInt(int64(time.Now().Minute()), 10)
		fmt.Println(appName)
		time.Sleep(time.Second * 6)
		fmt.Println(time.Now().Format("2006-01-02 15:04"))
	}
}
