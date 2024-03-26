package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(prompt string) float64 {
	fmt.Printf("%s", prompt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		panic(fmt.Sprintf("%v must only contain a number", prompt))
	}
	return value
}

func add(value1, value2 float64) float64 {
	return value1 + value2
}

func subtract(value1, value2 float64) float64 {
	return value1 - value2
}

func multiply(value1, value2 float64) float64 {
	return value1 * value2
}

func divide(value1, value2 float64) float64 {
	return value1 / value2
}

func getOperator() string {
	fmt.Print("Operator is ( + - * /):")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func main() {
	value1 := getInput("value 1 =")
	value2 := getInput("value 2 =")
	var result float64
	switch operator := getOperator(); operator {
	case "+":
		result = add(value1, value2)
	case "-":
		result = subtract(value1, value2)
	case "*":
		result = multiply(value1, value2)
	case "/":
		result = divide(value1, value2)
	default:
		panic("wrong operator")
	}
	fmt.Printf("Result = %v", result)
}