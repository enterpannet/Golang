package main

import "fmt"

var productName [2]string

func main() {
	productName[0]="macbook"
	productName[1]="iphone"
	price :=[2]float32{40000,20000}
	fmt.Println(productName)
	fmt.Println(price)

}