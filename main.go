package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	p := &i         //&でポインタをしゅとくする
	fmt.Println(p)  //ポインタ
	fmt.Println(*p) //ポインタを通して値を取得
	*p = 21         //ポインタを通して値を更新
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(j)
}
