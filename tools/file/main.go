package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	//读
	fileRead()
	bufioRead()
	ioutilRead()

	//写入
	fileWrite()
	bufioWrite()
	ioutilWrite()
}

func ioutilWrite() {
	str := "hello 沙河1"
	err := ioutil.WriteFile("./tools/file/write.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}

func bufioWrite() {
	file, err := os.OpenFile("./tools/file/write.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello沙河\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件
}

func fileWrite() {
	file, err := os.OpenFile("./tools/file/write.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "hello 沙河!"
	file.Write([]byte(str))        //写入字节切片数据
	file.WriteString("hello 小王子!") //直接写入字符串数据
}

func ioutilRead() {
	content, err := ioutil.ReadFile("./tools/file/main.go")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}

func fileRead() {
	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("./tools/file/main.go")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	// 关闭文件
	defer file.Close()

	//循环读文件
	var content []byte
	var tmp = make([]byte, 12)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}

func bufioRead() {

	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("./tools/file/main.go")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}

}
