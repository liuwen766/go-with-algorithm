package main

import "fmt"

func main() {

	fmt.Println("方法method是作用于特定类型变量的函数，称该类型变量为接收者（Receiver）")

	p1 := NewPerson("小王子", 25)
	p1.Dream()
	fmt.Println("指针类型的接收者")
	fmt.Println(p1.age) // 25
	p1.SetAge(30)
	fmt.Println(p1.age) // 30

	fmt.Println("值类型的接收者")
	p2 := NewPerson("小王子", 25)
	fmt.Println(p2.age) // 25
	p2.SetAge2(30)      // (*p2).SetAge2(30)
	fmt.Println(p2.age) // 25

	fmt.Println("什么时候应该使用指针类型接收者?")
	fmt.Println("1.需要修改接收者中的值\n2.接收者是拷贝代价比较大的大对象\n3.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。")

	fmt.Println("匿名字段的说法并不代表没有字段名，而是默认会采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。")
}

//Person 结构体
type Person struct {
	name string
	age  int8
}

//NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//Dream Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

// SetAge 设置p的年龄
// 使用指针接收者
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

// SetAge2 设置p的年龄
// 使用值接收者
func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}
