package main

/**
* @desc: 1)使用稀疏数组，来保留类似前面的二维数组（棋盘、地图等等
         2)把稀疏数组存盘，并且可以从新恢复原来的二维数组数
* @data: 2021.8.7 13:58
*/
func main() {
	//step1.创建数组

	//step2.查看数组

	//step3.转出稀疏数组
	//思路：（1）遍历 chessman,如果我们发现有一个元素的值不为0,创建一个node结构体
	//（2）将其放入到对应的切片即可

	//step4.查看稀疏数组

	//step5.恢复原始数组

	//step6.查看原始数组

}

type ValNode struct {
	row int
	col int
	val int
}
