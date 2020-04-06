package main

import (
	"fmt"
)

func main() {
	sum := 1
	// whileはforだけを使う
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
