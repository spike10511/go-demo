package main

import (
	"fmt"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	a, b := 24, 36
	fmt.Printf("最大公约数: %d\n最小公倍数: %d\n", gcd(a, b), lcm(a, b))
}

//思路分析：

//    使用辗转相除法计算最大公约数（GCD）。
//    最小公倍数（LCM）可以通过公式 LCM(a, b) = |a * b| / GCD(a, b) 计算
//结构思路：

//    定义两个函数 gcd 和 lcm。
//    在 main 函数中调用并打印结果。
