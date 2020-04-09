package main

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	// 通常のスライス
	s = s[1:4]
	fmt.Println(s)

	// 下限値を省略した記述
	// 下限値の既定値は０なので[0:2]と等価
	s = s[:2]
	fmt.Println(s)

	// 上限値を省略した記述
	// 上限値の既定値はスライスの長さ
	s = s[1:]
	fmt.Println(s)
}
