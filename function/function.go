package main

import "fmt"

func hello() {
	fmt.Println("สวัสดีๆๆๆๆ")

}
func plus(num1 int, num2 int) {
	result := num1 + num2
	fmt.Println(result)
}
func plusreturn(num1 int, num2 int) int {
	return num1 + num2
}
func plus3value(num1, num2, num3 int) int {
	return num1 + num2 + num3
}
func main() {
	hello()
	plus(2, 10)
	result := plusreturn(55, 100)
	fmt.Println(result)
	result2 := plus3value(12, 52, 62)
	fmt.Println(result2)
}
