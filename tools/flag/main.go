package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("执行:flag.exe a b c d")
	fmt.Println("获取命令行参数")
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
	fmt.Println("是一个存储命令行参数的字符串切片，它的第一个元素是执行文件的名称")

	////定义命令行参数方式1
	//var name string
	//var age int
	//var married bool
	//var delay time.Duration
	//flag.StringVar(&name, "name", "张三", "姓名")
	//flag.IntVar(&age, "age", 18, "年龄")
	//flag.BoolVar(&married, "married", false, "婚否")
	//flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")
	//
	////解析命令行参数
	//flag.Parse()
	//fmt.Println(name, age, married, delay)
	////返回命令行参数后的其他参数
	//fmt.Println(flag.Args())
	////返回命令行参数后的其他参数个数
	//fmt.Println(flag.NArg())
	////返回使用的命令行参数个数
	//fmt.Println(flag.NFlag())

}
