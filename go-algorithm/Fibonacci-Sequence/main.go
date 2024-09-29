package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	n := 10
	fmt.Printf("第 %d 个斐波那契数: %d\n", n, fibonacci(n))
}

//思路分析：

//    斐波那契数列可以通过递归、动态规划或迭代来实现。
//    使用迭代方法更为高效。
//结构思路：

//    使用迭代方式计算斐波那契数。
//    在 main 函数中设置要计算的数并输出结果。
