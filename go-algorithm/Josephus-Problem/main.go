package main

import (
	"fmt"
)

func josephus(n int, k int) int {
	if n == 1 {
		return 0
	}
	return (josephus(n-1, k) + k) % n
}

func main() {
	n := 7                         // 总人数
	k := 3                         // 每次删除的第 k 个
	survivor := josephus(n, k) + 1 // +1 使得结果从 1 开始
	fmt.Printf("最后幸存者的位置: %d\n", survivor)
}

//思路分析：

//    经典的约瑟夫环问题可以通过链表或循环数组来实现。
//    采用循环列表来模拟。
//结构思路：

//    使用递归实现约瑟夫环算法。
//    在 main 函数中设置参数并输出最后的幸存者位置。
