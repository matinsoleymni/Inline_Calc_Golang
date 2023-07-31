package main

import "fmt"

func main() {

	fmt.Print("Enter Number: ")
	var one int
	fmt.Scan(&one)

	fmt.Print("Enter Number: ")
	var two int
	fmt.Scan(&two)

	fmt.Print("Enter operation: ")
	var op string
	fmt.Scan(&op)

	calc(one, two, op)

	fmt.Println("\n\n24 is here ∞")
	fmt.Scan(&op)
}

func calc(numberOne int, numberTwo int, operation string) {
	fmt.Printf("\nYou Say: %v %v %v \n", numberOne, operation, numberTwo)

	if operation == "/" {
		result := numberOne / numberTwo
		fmt.Printf("result is ~= %v", result)
	} else if operation == "+" {
		result := numberOne + numberTwo
		fmt.Printf("result is = %v", result)
	} else if operation == "-" {
		result := numberOne - numberTwo
		fmt.Printf("result is = %v", result)
	} else if operation == "*" {
		result := numberOne * numberTwo
		fmt.Printf("result is = %v", result)
	}
}
