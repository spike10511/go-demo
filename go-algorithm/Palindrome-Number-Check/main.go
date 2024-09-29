package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func main() {
	num := 121
	fmt.Printf("%d 是回文数: %v\n", num, isPalindrome(num))
}

//思路分析：

//    将整数转换为字符串，然后判断其是否为回文。
//结构思路：

//    定义 isPalindrome 函数。
//    使用双指针技巧判断字符是否对称。
