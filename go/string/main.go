package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	s := "主备（32C256G，ESSD PL2 云盘）独享型"

	fmt.Println(s[strings.Index(s, "（")+3 : strings.Index(s, "C")])
	fmt.Println(s[strings.Index(s, "C")+1 : strings.Index(s, "G")])

	//cpu, err := strconv.Atoi(flavorName[strings.Index(flavorName, "（")+3 : strings.Index(flavorName, "C")])
	//if err != nil {
	//	fmt.Println("can't convert cpu to int")
	//}
	//mem, err := strconv.Atoi(flavorName[strings.Index(flavorName, "C")+3 : strings.Index(flavorName, "G")])
	//if err != nil {
	//	fmt.Println("can't convert mem to int")
	//}

	strs := []string{"主备（16C128G，ESSD PL1 云盘）独享型", " MySQL主备(2C8G,本地SSD盘)独享型",
		" 主备（16C128G，ESSD PL1 云盘）独享型", "MySQL独享包月套餐(4核8GB主备版SSD云盘)",
		"MySQL主备(16C32G,本地SSD盘)独享型", " MySQL主备(8C32G,本地SSD盘)通用型",
		" MySQL通用按小时套餐(4核16GB单机版SSD云盘)", "MySQL主备(24C192G,本地SSD盘)独享型",
		" MySQL主备(2C8G,本地SSD盘)独享型", " MySQL主备(2C4G,SSD云盘)独享型",
		"MySQL通用按小时套餐(2核4GB主备版SSD云盘)", " MySQL通用按小时套餐(2核4GB单机版SSD云盘)"}
	reg1 := regexp.MustCompile("[1-9]*[C核][1-9]*[G]")
	reg2 := regexp.MustCompile("[1-9]*")
	for i := range strs {
		//fmt.Println(reg1.FindString(strs[i])) // [2 2]
		//fmt.Println(reg2.FindAllString(reg1.FindString(strs[i]), 2))
		allString := reg2.FindAllString(reg1.FindString(strs[i]), 2)
		cpu, err := strconv.Atoi(allString[0])
		if err != nil {
			fmt.Println("can't convert cpu to int")
		}
		mem, err := strconv.Atoi(allString[1])
		if err != nil {
			fmt.Println("can't convert mem to int")
		}
		fmt.Println(cpu, mem)
	}
}
