package main

import (
	"fmt"
)

// Pi : コメントを記述しないとlintの警告が表示される
const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
