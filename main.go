package main

import (
	"fmt"
)

//Vertex : hoge
type Vertex struct {
	X int
	Y int
}

func main() {
	// 要素数６の配列を宣言
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// sはスライス
	// 変数[low:high]でlowからhighの一つ前を含むスライスを作る
	var s []int = primes[1:4]
	var m []int = primes[0:1]
	fmt.Println(s)
	fmt.Println(m)
}
