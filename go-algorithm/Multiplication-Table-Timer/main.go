package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}

	elapsed := time.Since(start)
	fmt.Printf("执行时间: %s\n", elapsed)
}

//思路分析：
//       使用嵌套循环来打印乘法表。
//       利用 time 包来测量程序运行时间。
//结构思路：
//       main 函数作为程序入口。
//       嵌套循环用于输出乘法表。
//       通过 time.Now() 和 time.Since() 计算执行时间。
