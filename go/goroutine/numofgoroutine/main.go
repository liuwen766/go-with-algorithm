package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	pool := New(10)
	now := time.Now()
	println(runtime.NumGoroutine())
	for i := 0; i < 3000; i++ {
		pool.Add(1)
		go func() {
			time.Sleep(time.Second)
			println("liu", runtime.NumGoroutine())
			pool.Done()
		}()
	}
	pool.Wait()
	since := time.Since(now)
	fmt.Println("总共用时：", since)
	println(runtime.NumGoroutine())
}

type pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

func New(size int) *pool {
	if size <= 0 {
		size = 1
	}
	return &pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

func (p *pool) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	for i := 0; i > delta; i-- {
		<-p.queue
	}
	p.wg.Add(delta)
}

func (p *pool) Done() {
	<-p.queue
	p.wg.Done()
}

func (p *pool) Wait() {
	p.wg.Wait()
}
