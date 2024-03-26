package main

import "fmt"

func zeroi(ivalue int) {
	ivalue = 0

}
func zeropointer(ipointer *int){
	*ipointer=0
}
func main() {
	i := 1
	fmt.Println("i =", i)
	zeroi(i)
	fmt.Println("I in func =", i)
	zeropointer(&i)
	fmt.Println("i in zero pointer =",i)
	fmt.Println("i address",&i)
}
