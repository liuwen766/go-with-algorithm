package main

import (
	"fmt"
	"time"
)

func goRoutineA(a <-chan int) {
	val := <-a
	fmt.Println("goRoutineA received the data", val)
}
func goRoutineB(b chan int) {
	val := <-b
	fmt.Println("goRoutineB received the data", val)
}
func main() {

	for i := 0; i < 10; i++ {
		//源码目录  runtime/chan.go
		ch := make(chan int, 5)
		//ch := make(chan int)  //发送方和接收方要同步就绪，只有在两者都 ready 的情况下，数据才能在两者间传输
		go goRoutineA(ch)
		go goRoutineB(ch)
		ch <- 3
		time.Sleep(time.Second * 1)
	}

}
