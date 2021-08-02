package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	fmt.Print("我是fmt.Print(),我不换行")
	fmt.Printf("我是fmt.Printf(),%s,我占位", "哈哈哈")
	fmt.Println("我是fmt.Println(),我换行")

	fmt.Println("---------------------------")

	s1 := fmt.Sprint("沙河小王子")
	name := "沙河小王子"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("沙河小王子")
	fmt.Printf("s1:%s,s2:%s,s3:%s", s1, s2, s3)

	fmt.Println("---------------------------")

	err := fmt.Errorf("这是一个错误")
	fmt.Println("自定义err:", err)

	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误:%w", e)
	fmt.Println("自定义err:", w)

	fmt.Println("---------------------------")

	fmt.Fprint(os.Stdout, "向标准输出写入内容,我不换行")
	file, err := os.OpenFile("./tools/fmt/test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件错误，err:", err)
	}
	fmt.Fprintf(file, "向文件中写入内容:%s", "hello world!")

	fmt.Println("---------------------------")

	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	o := struct{ name string }{"小王子"}
	fmt.Printf("%s:%v\n", "这是结构体", o)
	fmt.Printf("%s:%+v\n", "这是添加了字段名的结构体", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")

	fmt.Println("---------------------------")

}
