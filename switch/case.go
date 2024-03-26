package main

import "fmt"

func inputC() {
	var input int
	fmt.Scanf("%d", &input)
	fmt.Println(input)
	switch input {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("default")
	}
}
func color() {
	var input string
	fmt.Scanf("%s", &input)
	switch input {
	case "blue":
		fmt.Println("#0000FF")
	case "green":
		fmt.Println("#008000")
	case "pink":
		fmt.Println("#FFC0CB")
	case "yellow":
		fmt.Println("#FFFF00")
	default:
		fmt.Println("No Color")

	}
}
func main() {
	color()
}
