package main

import (
	"fmt"
	"strings"
)

/**
 * @desc: 如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
 * 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
 * @data: 2022.2.7 14:54
 */
func main() {
	n := 8
	queens := solveNQueens(n)
	fmt.Println("N=", n, "时,N皇后问题解有:", len(queens), "种")
	fmt.Println(queens)
}

var res [][]string

func solveNQueens(n int) [][]string {
	res = make([][]string, 0)
	//初始化棋盘
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	//从第一行开始放皇后
	backtrack(board, 0)
	return res
}

/**
 * @desc: 回溯算法解n皇后问题
 * @data: 2022.2.7 15:06
 */
func backtrack(board [][]string, row int) {
	length := len(board)
	//终止条件→皇后已经放到最后一行
	if length == row {
		temp := make([]string, length)
		for i := 0; i < len(board); i++ {
			temp[i] = strings.Join(board[i], "")
		}
		res = append(res, temp)
	}

	for col := 0; col < length; col++ {
		if !isValid(board, row, col) {
			continue
		}
		board[row][col] = "Q"
		//放下一行的皇后
		backtrack(board, row+1)
		board[row][col] = "."
	}
}

/**
 * @desc: 不能同行  不能同列  不能同斜线
 * @data: 2022.2.7 14:59
 */
func isValid(board [][]string, row, col int) (res bool) {
	n := len(board)
	//不能同列
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	//不能同行
	for i := 0; i < n; i++ {
		if board[row][i] == "Q" {
			return false
		}
	}
	//不能同斜线,左上方
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	//不能同斜线,右上方
	for i, j := row, col; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}
