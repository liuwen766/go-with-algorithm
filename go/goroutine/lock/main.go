package main

import (
	"fmt"
	"sync"
	"time"
)

var x int64
var wg sync.WaitGroup

var lock sync.Mutex

var rwLock sync.RWMutex

func add() {
	for i := 0; i < 500000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 解锁
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)

	rwLock.Lock()
	time.Sleep(10 * time.Millisecond)
	rwLock.Unlock()
	fmt.Println("读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来。")
	rwLock.RLock()
	time.Sleep(1 * time.Millisecond)
	rwLock.RUnlock()

	fmt.Println("Go语言中可以使用sync.WaitGroup来实现并发任务的同步")

	fmt.Println("Go语言中可以使用sync.Once确保某些操作在高并发的场景下只执行一次。例如只加载一次配置文件、只关闭一次通道等。")

	fmt.Println("Go语言中可以使用sync.Map来为map加锁来保证并发的安全性。避免：fatal error: concurrent map writes错误")
}
