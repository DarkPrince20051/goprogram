package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Calculator")
	fmt.Println("-----------------")

	for {
		fmt.Print("Enter first number: ")
		num1Str, _ := reader.ReadString('\n')
		num1Str = strings.TrimSpace(num1Str)
		num1, err := strconv.ParseFloat(num1Str, 64)
		if err != nil {
			fmt.Println("Invalid number, please try again.")
			continue
		}

		fmt.Print("Enter second number: ")
		num2Str, _ := reader.ReadString('\n')
		num2Str = strings.TrimSpace(num2Str)
		num2, err := strconv.ParseFloat(num2Str, 64)
		if err != nil {
			fmt.Println("Invalid number, please try again.")
			continue
		}

		fmt.Print("Do you want to calculate all operations at once? (yes/no): ")
		calculateAllStr, _ := reader.ReadString('\n')
		calculateAll := strings.TrimSpace(strings.ToLower(calculateAllStr))

		if calculateAll == "yes" {
			printAllResults(num1, num2)
		} else {
			fmt.Print("Enter operator (+, -, *, /): ")
			operator, _ := reader.ReadString('\n')
			operator = strings.TrimSpace(operator)

			result, err := calculate(num1, num2, operator)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("Result: %.2f\n", result)
		}

		fmt.Print("Do you want to perform another calculation? (yes/no): ")
		continueCalc, _ := reader.ReadString('\n')
		continueCalc = strings.TrimSpace(strings.ToLower(continueCalc))
		if continueCalc != "yes" {
			break
		}
	}
}

func calculate(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return math.Abs(num1 + num2), nil
	case "-":
		return math.Abs(num1 - num2), nil
	case "*":
		return math.Abs(num1 * num2), nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("cannot divide by zero")
		}
		if num2 > num1 {
			return math.Abs(num2 / num1), nil
		}
		return math.Abs(num1 / num2), nil
	default:
		return 0, errors.New("invalid operator")
	}
}

func printAllResults(num1, num2 float64) {
	addResult, _ := calculate(num1, num2, "+")
	subResult, _ := calculate(num1, num2, "-")
	mulResult, _ := calculate(num1, num2, "*")
	divResult, err := calculate(num1, num2, "/")
	if err != nil {
		fmt.Println("Division Error:", err)
		divResult = 0
	}

	fmt.Printf("Addition (absolute value): %.2f\n", addResult)
	fmt.Printf("Subtraction (absolute value): %.2f\n", subResult)
	fmt.Printf("Multiplication (absolute value): %.2f\n", mulResult)
	fmt.Printf("Division (absolute value): %.2f\n", divResult)
}
