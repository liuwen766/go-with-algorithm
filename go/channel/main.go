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

		ch := make(chan int, 5)
		//ch := make(chan int)
		go goRoutineA(ch)
		go goRoutineB(ch)
		ch <- 3
		time.Sleep(time.Second * 1)
	}

}
