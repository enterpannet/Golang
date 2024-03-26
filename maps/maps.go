package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println("Product =", product)

	// add data
	product["Macbook"] = 40000
	product["Mac"] = 410000
	product["book"] = 402000
	product["Iphone"] = 140000
	fmt.Println("Product =", product)

	delete(product, "Iphone")
	fmt.Println("Product =", product)
	value1 := product["book"]
	fmt.Println("value1", value1)
	courseName := map[string]string{"101": "jave", "102": "PHP"}
	fmt.Println(courseName)

}
