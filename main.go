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
	v := Vertex{1, 2}
	p := &v             //構造体のポインタの取得
	fmt.Println((*p).X) //ポインタを通して構造体の値にアクセス
	fmt.Println(p.X)    //省略しても書ける
	p.X = 1e9
	fmt.Println(v)
}
