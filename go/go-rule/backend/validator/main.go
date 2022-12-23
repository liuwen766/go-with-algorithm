package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

//3.1.1.1【必须】按类型进行数据校验
//所有外部输入的参数，应使用validator进行白名单校验，校验内容包括但不限于数据长度、数据范围、数据类型与格式，
//校验不通过的应当拒绝。
func main() {
	validate = validator.New()
	validateVariable()
}

var validate *validator.Validate

func validateVariable() {
	email := "abc@qq.com"
	err := validate.Var(email, "required,email")
	if err != nil {
		fmt.Println(err)
		return
		//停止执行
	}
	// 验证通过，继续执行
	//...
}
