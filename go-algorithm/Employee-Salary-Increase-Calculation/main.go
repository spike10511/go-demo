package main

import (
	"fmt"
)

func calculateRaise(salary float64, performanceRating int) float64 {
	var raisePercentage float64
	switch performanceRating {
	case 1:
		raisePercentage = 0.05 // 5%
	case 2:
		raisePercentage = 0.10 // 10%
	case 3:
		raisePercentage = 0.15 // 15%
	case 4:
		raisePercentage = 0.20 // 20%
	default:
		raisePercentage = 0.0
	}
	return salary * raisePercentage
}

func main() {
	salary := 50000.0
	performanceRating := 3
	raise := calculateRaise(salary, performanceRating)
	fmt.Printf("员工的涨薪金额: %.2f\n", raise)
}

//思路分析：

//    根据员工的绩效评分进行涨薪计算，可以设置不同的评分区间和对应的涨薪百分比。
//结构思路：

//    定义 calculateRaise 函数，根据绩效评分计算涨薪金额。
//    在 main 函数中调用并输出结果。
