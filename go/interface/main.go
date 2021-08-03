package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {

	var peo People = &Student{}
	//var peo People = Student{}
	think := "sb"
	fmt.Println(peo.Speak(think))

	fmt.Println("请牢记接口（interface）是一种类型。")
	cat := Cat{}
	fmt.Println(cat.say())
	Dog := Dog{}
	fmt.Println(Dog.say())
	var say Sayer
	say = cat
	s := say.say()
	fmt.Println(s)

	var x Mover
	var wangcai = dog{} // 旺财是dog类型
	x = wangcai         // x可以接收dog类型
	x.move()
	var fugui = &dog{} // 富贵是*dog类型
	x = fugui          // x可以接收*dog类型
	x.move()

	//var x1 Mover
	//var wangcai1 = dog{} // 旺财是dog类型
	//x1 = wangcai1         // x不可以接收dog类型
	//var fugui1 = &dog{}  // 富贵是*dog类型
	//x1 = fugui1           // x可以接收*dog类型

	justifyType(11)
	justifyType("11")
}

func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}

func (d *dog) move1() {
	fmt.Println("狗会动")
}

type Mover interface {
	move()
}

type dog struct{}

func (d dog) move() {
	fmt.Println("狗会动")
}

// 摘自gin框架routergroup.go
type IRouter interface{}

type RouterGroup struct{}

var _ IRouter = &RouterGroup{} // 确保RouterGroup实现了接口IRouter

type Cat struct {
}

type Dog struct {
}

// Sayer 接口
type Sayer interface {
	say() string
}

func (c Cat) say() string {
	return "喵喵"
}

func (Dog) say() string {
	return "旺旺"
}
