package main

import "fmt"

/**
问题描述:

请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。

给定一个string，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。
*/

func main() {
	s := "abcdefghijklm"

	fmt.Print(reverse(s))
}

func reverse(s string) string {
	i, j := 0, len(s)-1
	str := []rune(s)

	for i < j {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
	return string(str)
}
