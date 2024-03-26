package main

import (
	"fmt"
)

func main() {
	// i := 0
	// for i < 10 {
	// 	fmt.Println("i=", i)
	// 	i = i + 1
	// }
	// for j := 0; j < count; j++ {
	// 	fmt.Println("j=", j)
	// }
	var input string
	for {
		// fmt.Println("สวัสดีๆๆๆๆ")
		fmt.Scanf("%s", &input)
		fmt.Println("input = ", input)
		if input == "exit" {
			break
		}
	}
}
