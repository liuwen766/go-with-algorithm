package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Println("Hello Goroutine!", i)
}
func main() {

	// 禁止GC
	debug.SetGCPercent(-1)
	fmt.Println("程序开始前：", runtime.NumGoroutine())
	defer fmt.Println("程序结束后：", runtime.NumGoroutine())

	fmt.Println("Go语言中一次创建十万左右的goroutine也是可以的。")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
	fmt.Println("main goroutine done!")

	fmt.Println("可以通过runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数")

	//runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}
