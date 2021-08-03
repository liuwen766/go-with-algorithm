package main

import (
	"fmt"
	"unsafe"
)

type NewInt int
type MyInt = int

type person struct {
	name string
	city string
	age  int8
}

func main() {

	fmt.Println("自定义类型，如：type MyInt int")
	fmt.Println("类型别名，如：type rune = int32")
	var a NewInt
	var b MyInt
	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int

	fmt.Println("方法一：对结构体变量进行实例化，使用 . 访问结构体的字段")
	var p1 person
	p1.name = "沙河娜扎"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)  //p1={沙河娜扎 北京 18}
	fmt.Printf("p1=%+v\n", p1) //p1={name:沙河娜扎 city:北京 age:18}
	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"沙河娜扎", city:"北京", age:18}

	fmt.Println("方法二：new关键字对结构体进行实例化，得到的是结构体的地址，也可以使用 . 访问结构体的字段")
	var p2 = new(person)
	fmt.Printf("%T\n", p2)     //*main.person
	fmt.Printf("%p\n", p2)     //0xc0000803c0
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"", city:"", age:0}
	p2.name = "七米"
	p2.age = 30
	p2.city = "成都"
	fmt.Printf("p2=%#v\n", p2) //p2=&main.person{name:"七米", city:"成都", age:30}

	fmt.Println("方法三：取结构体的地址实例化（使用&相当于new实例化）")
	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("%p\n", p3)     //0xc00008042c
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "七米"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"七米", city:"成都", age:30}

	fmt.Println("空结构体不占内存")
	var v struct{}
	fmt.Println(unsafe.Sizeof(v)) // 0

	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
		for k, v := range m {
			fmt.Println(k, "=>", v.name)
		}
		fmt.Println("---------")
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

	fmt.Println("构造函数需要自己实现，struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以可以返回结构体指针类型。")
	p9 := newPerson("张三", "沙河", 90)
	fmt.Printf("%#v\n", p9) //&main.person{name:"张三", city:"沙河", age:90}

}

type student struct {
	name string
	age  int
}

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
