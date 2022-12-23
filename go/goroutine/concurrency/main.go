package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
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

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("闭包：", i)
		}()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
	fmt.Println("main goroutine done!")

	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch:
			fmt.Println("打印出10以内的奇数:", x)
		case ch <- i:
		}
	}

	dataChan := make(chan int)
	go newgoroutine(dataChan)
	// Some application processing (but forgot to consume data from the channel (dataChan))
	return
}

func newgoroutine(dataChan chan int) {
	data := 6
	dataChan <- data
	return
}
