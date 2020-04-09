package main

import (
	"fmt"
)

//Vertex : hoge
type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  //構造体のフィールドの初期化
	v2 = Vertex{X: 1}  //Xだけを初期化。Yは暗黙的に0になる
	v3 = Vertex{}      //初期化しない。XとYは０になる
	p  = &Vertex{1, 2} //構造体のポインタを取得
)

func main() {
	fmt.Println(p)
	fmt.Println(*p) //ポインタを通して値を取得
	fmt.Println(v1, p, v2, v3)
}
