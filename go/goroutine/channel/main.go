package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
	"time"
)

func main() {

	//demo1()

	demo2()

	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}

	ch := make(chan int)
	fmt.Println("对于无缓冲通道，必须先启用goroutine从通道接收值")
	go recv(ch) //
	ch <- 10
	fmt.Println("发送成功")

	ch11 := make(chan int)
	ch22 := make(chan int)
	go counter(ch11)
	go squarer(ch22, ch11)
	printer(ch22)

}

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

// demo1 通道误用导致的bug
func demo1() {
	// 禁止GC
	debug.SetGCPercent(-1)
	fmt.Println("程序开始前：", runtime.NumGoroutine())
	defer fmt.Println("程序结束后：", runtime.NumGoroutine())

	wg := sync.WaitGroup{}

	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	wg.Add(3)
	for j := 0; j < 3; j++ {
		go func() {
			for {
				//task := <-ch
				//// 这里假设对接收的数据执行某些操作
				//fmt.Println("获取值：", task)
				task, ok := <-ch
				if ok {
					// 这里假设对接收的数据执行某些操作
					fmt.Println("获取值：", task)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

// demo2 通道误用导致的bug
func demo2() {
	// 禁止GC
	debug.SetGCPercent(-1)
	fmt.Println("程序开始前：", runtime.NumGoroutine())
	defer fmt.Println("程序结束后：", runtime.NumGoroutine())

	ch := make(chan string)
	go func() {
		// 这里假设执行一些耗时的操作
		time.Sleep(3 * time.Second)
		ch <- "job result"
	}()

	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(time.Second): // 较小的超时时间
		return
	}
}
