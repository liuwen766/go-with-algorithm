package main

import (
	"fmt"
)

func main() {

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
