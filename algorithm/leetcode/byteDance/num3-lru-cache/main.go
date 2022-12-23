package main

import (
	"container/list"
	"fmt"
)

func main() {

	//capacity := 10
	//key := 2
	//value := 5
	//
	//obj := Constructor(capacity)
	//obj.Put(key, value)
	//val := obj.Get(key)
	//
	//fmt.Println(val)

	l := list.List{}
	l.Init()
	l.PushFront("1")
	l.PushFront("2")
	fmt.Println("l.Len()", l.Len())
	fmt.Println("l.Back().Value", l.Back().Value)
	fmt.Println("l.Front().Value", l.Front().Value)
}

type LRUCache struct {
	Cap  int
	Keys map[int]*list.Element
	List *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{}
}

func (this *LRUCache) Get(key int) int {
	return key
}

func (this *LRUCache) Put(key int, value int) {

}
