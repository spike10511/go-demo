package main

import (
	"fmt"
)

func isArmstrong(n int) bool {
	a := n / 100
	b := (n / 10) % 10
	c := n % 10
	return n == (a*a*a + b*b*b + c*c*c)
}

func main() {
	fmt.Println("水仙花数:")
	for i := 100; i < 1000; i++ {
		if isArmstrong(i) {
			fmt.Println(i)
		}
	}
}

//思路分析：

//    水仙花数是指一个三位数，其各位数字的立方和等于它本身。
//    循环遍历 100 到 999 的数字，进行判断。
//结构思路：

//    使用 isArmstrong 函数判断水仙花数。
//    在 main 函数中进行循环并打印结果。
