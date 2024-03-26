package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"java", "php"}
	fmt.Println(courseName)
	courseName = append(courseName, "Golang")
	fmt.Println(courseName)

	courseWed := courseName[2:3]
	fmt.Println(courseWed)
}
