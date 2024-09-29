package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapST := make(map[rune]rune)
	mapTS := make(map[rune]rune)

	for i, c := range s {
		if mapped, ok := mapST[c]; ok && mapped != rune(t[i]) {
			return false
		}
		if mapped, ok := mapTS[rune(t[i])]; ok && mapped != c {
			return false
		}
		mapST[c] = rune(t[i])
		mapTS[rune(t[i])] = c
	}
	return true
}

func main() {
	s, t := "egg", "add"
	fmt.Printf("'%s' 和 '%s' 是同构数: %v\n", s, t, isIsomorphic(s, t))
}

//思路分析：

//    同构数是指两个数字的每位数字可以通过一个一一映射关系转化。
//    可以使用 map 来存储数字之间的映射关系。

//结构思路：

//    使用两个 map 来存储字符映射关系。
//    逐字符检查映射是否一致。
