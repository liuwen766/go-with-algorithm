package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Monster struct {
	Name string  `json:"monsterName"`
	Age  int     `json:"monsterAge"`
	Sal  float64 `json:"monsterSal"`
	Sex  string  `json:"monsterSex"`
}

func main() {
	m := Monster{
		Name: "玉兔精",
		Age:  20,
		Sal:  888.99,
		Sex:  "female",
	}
	marshal, _ := json.Marshal(m)
	fmt.Println("json result:", string(marshal))

	fmt.Println("变量、 interface{}和 reflect.Value是可以相互转换的，这点在实际开发中，会经常使用到。")
	fmt.Println("变量<--->interface{}<--->reflect.Value")
	reflectTest(100)
}

type Stu struct {
}

func reflectTest(b interface{}) {
	refType := reflect.TypeOf(b)
	fmt.Println("refType:", refType)

	refValue := reflect.ValueOf(b)

	res := refValue.Int() + 2
	fmt.Println("res:", res)

	refInter := refValue.Interface()
	num := refInter.(int)
	fmt.Println("num:", num)

}
