package main

import "fmt"

/**
 * @desc: 给你一个字符串S、一个字符串T,请在字符串S里面找出：包含T所有字母的最小子串。
 * @data: 2022.1.26 23:12
 */
func main() {
	s := "ASDASDWQNDMKDXCDFSD"
	t := "ADS"
	res := minWindow(s, t)
	fmt.Println(res)
}

/**
 *  在 S(source) 中找到包含 T(target) 中全部字母的一个子串，且这个子串一定是所有可能子串中最短的
	1、我们在字符串 S 中使用双指针中的左右指针技巧，初始化 left = right = 0，把索引左闭右开区间 [left, right) 称为一个「窗口」。
	2、我们先不断地增加 right 指针扩大窗口 [left, right)，直到窗口中的字符串符合要求（包含了 T 中的所有字符）。
	3、此时，我们停止增加 right，转而不断增加 left 指针缩小窗口 [left, right)，直到窗口中的字符串不再符合要求（不包含 T 中的所有字符了）。同时，每次增加 left，我们都要更新一轮结果。
	4、重复第 2 和第 3 步，直到 right 到达字符串 S 的尽头。
    第 2 步相当于在寻找一个「可行解」，然后第 3 步在优化这个「可行解」，最终找到最优解
*/
func minWindow(s, t string) string {
	/**明确四个问题
	1、当移动 right 扩大窗口，即加入字符时，应该更新哪些数据？
	2、什么条件下，窗口应该暂停扩大，开始移动 left 缩小窗口？
	3、当移动 left 缩小窗口，即移出字符时，应该更新哪些数据？
	4、我们要的结果应该在扩大窗口时还是缩小窗口时进行更新？
	*/

	return ""
}
