package main

import (
	"fmt"
)

func sieveOfEratosthenes(n int) []int {
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for p := 2; p*p <= n; p++ {
		if isPrime[p] {
			for i := p * p; i <= n; i += p {
				isPrime[i] = false
			}
		}
	}
	primes := []int{}
	for p := 2; p <= n; p++ {
		if isPrime[p] {
			primes = append(primes, p)
		}
	}
	return primes
}

func main() {
	n := 50
	fmt.Printf("小于 %d 的质数: %v\n", n, sieveOfEratosthenes(n))
}

//思路分析：

//    使用埃拉托斯特尼筛法来高效找出质数。
//结构思路：

//    使用布尔数组记录质数状态。
//    在 main 函数中调用并输出结果。
