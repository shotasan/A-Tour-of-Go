package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// ifで評価する前に変数宣言が可能
	if v := math.Pow(x, n); v < lim {
		return v
	}
	// スコープを外れるのでvは使えない
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
