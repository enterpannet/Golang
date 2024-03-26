package main

import "fmt"

var score int

func main() {
	fmt.Println("กรุณาใส่คะแนน")
	fmt.Scanf("%d", &score)
	if score >= 80 {
		fmt.Println("คุณได้เกรด A")
	} else if score >= 70 {
		fmt.Println("คุณได้เกรด B")
	} else if score >= 60 {
		fmt.Println("คุณได้เกรด C")
	} else if score >= 50 {
		fmt.Println("คุณได้เกรด D")
	} else {
		fmt.Println("คุณได้เกรด F")
	}
}
