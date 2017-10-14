package main

import (
        "calculator/calc"
	"fmt"
	"strconv"
	"os"
)
func main() {

        var operation string
	fmt.Println("Please select an operation: +, -, *, /")
        fmt.Scanln(&operation)

	var num1 string
	fmt.Println("Please input the first number")
	fmt.Scanln(&num1)

	var num2 string
	fmt.Println("Please input the second number")
	fmt.Scanln(&num2)

	firstNumber := stringToInt(num1)
	secondNumber := stringToInt(num2)
	firstFloat := stringToFloat64(num1)
	secondFloat := stringToFloat64(num2)
	fmt.Println(calc.Add(firstNumber, secondNumber))
	fmt.Println(calc.Subtract(firstNumber, secondNumber))
	fmt.Println(calc.Multiply(firstFloat, secondFloat))
	fmt.Println(calc.Divide(firstFloat, secondFloat))

}

func stringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}


func stringToFloat64(str string) float64 {
	i, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}
