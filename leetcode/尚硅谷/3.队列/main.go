package main

import "fmt"

func main() {
	fmt.Println("队列是一个有序列表，可以用数组或是链表来实现。\n遵循先入先出的原则。即：先存入队列的数据，要先取出。后存入的要后取出")
	fmt.Println("使用数组实现队列的思路\n1.创建一个数组 arrary,是作为队列的一个字段\n2. front初始化为-1\n3.real,表示队列尾部，初始化为-1\n4.完成队列的基本查找\nAddqueue//加入数据到队列\nGetqueue/从队列取出数据\nShowqueue/显示队列")
}

type Queue struct {
}

/**
* @desc: 将尾指针往后移：rear+1, front==rear【空】
         若尾指引rear小于等于队列的最大下标 Maxsize-1,则将数据存入rear所指的数组元素中，
         否则无法存入数据。rear== Maxsize-1[队列满
* @data: 2021.8.7 14:57
*/
func (this *Queue) AddQueue(val int) (err error) {

	return err
}
