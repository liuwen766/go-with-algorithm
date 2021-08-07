package main

import "fmt"

/**
* @desc: 1)使用稀疏数组，来保留类似前面的二维数组（棋盘、地图）等等
         2)把稀疏数组存盘，并且可以从新恢复原来的二维数组数
* @data: 2021.8.7 13:58
*/
func main() {
	//step1.创建数组
	var arr [11][11]int //棋盘
	arr[2][3] = 1       //黑子
	arr[3][4] = 2       //白子
	//step2.查看数组
	fmt.Println("原始数组：")
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], "  ")
		}
		fmt.Println()
	}
	//step3.转出稀疏数组
	//思路：（1）遍历 chessman,如果我们发现有一个元素的值不为0,创建一个node结构体
	//（2）将其放入到对应的切片
	var sparseArr []ValNode
	//存储数组规模
	sparseArr = append(sparseArr, ValNode{
		row: 11,
		col: 11,
		val: 0,
	})
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] != 0 {
				sparseArr = append(sparseArr, ValNode{
					row: i,
					col: j,
					val: arr[i][j],
				})
			}
		}
	}
	//step4.查看稀疏数组
	fmt.Println("稀疏数组")
	fmt.Println(sparseArr)
	//( todo 可以将稀疏数组进行存储)
	//step5.恢复原始数组
	//恢复数组规模
	var restoreArr [11][11]int
	for i, val := range sparseArr {
		if i != 0 {
			restoreArr[val.row][val.col] = val.val
		}
	}
	//step6.查看原始数组
	fmt.Println("原始数组")
	for _, row := range restoreArr {
		for _, col := range row {
			//fmt.Printf("%d\t", col)
			fmt.Print(col, "  ")
		}
		fmt.Println()
	}
	//fmt.Println(restoreArr)
}

type ValNode struct {
	row int
	col int
	val int
}
