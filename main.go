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
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// 配列の長さより要素数が少ない場合は０が入る。多い場合はコンパイルエラー
	// primes := [6]int{2, 3, 5, 7, 11} //=> [2 3 5 7 11 0]
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
