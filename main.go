package main

import (
	"fmt"
)

// deferはすぐに評価されるが呼び出し元の関数がreturnするまで実行されない
func main() {
	defer fmt.Println("wrold")

	fmt.Println("hello")
}
