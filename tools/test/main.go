package main

import "fmt"

func main() {
	fmt.Println("go test命令会遍历所有的*_test.go文件中符合上述命名规则的函数，" +
		"然后生成一个临时的main包用于调用相应的测试函数，" +
		"然后构建并运行、报告测试结果，最后清理测试中生成的临时文件。")
	fmt.Println("go test\ngo test -v\ngo test -cover")

}
