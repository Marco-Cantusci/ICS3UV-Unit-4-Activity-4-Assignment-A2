/**
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-12-09
 * @fileoverview Calculator
 */

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	// declare variables
	var choice string
	var num1String string
	var num1Number float64
	var num2String string
	var num2Number float64
	var result float64

	reader := bufio.NewReader(os.Stdin)

	// selecting mathematical operation
	fmt.Println(
		"Which operation would you like to perform today? (Select by typing the letter in front of the operation.)",
		"A. Add",
		"B. Subtract",
		"C. Multiply",
		"D. Divide",
		"E. Absolute value",
		"F. Round",
		"G. Raise to an exponent",
		"H. Square root",
	)

	// input
	// choosing operation
	fmt.Print("What operation do you want to choose: ")
	choice, _ = reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	choice = strings.ToUpper(choice)

	// addition
	if choice == "A" {
		fmt.Print("Enter first number: ")
		num1String, _ = reader.ReadString('\n')
		num1String = strings.TrimSpace(num1String)
		num1Number, _ = strconv.ParseFloat(num1String, 64)

		fmt.Print("Enter second number: ")
		num2String, _ = reader.ReadString('\n')
		num2String = strings.TrimSpace(num2String)
		num2Number, _ = strconv.ParseFloat(num2String, 64)

		result = num1Number + num2Number
		fmt.Printf("%f + %f = %f\n", num1Number, num2Number, result)
	} else if choice == "B" { // subtraction
		fmt.Print("Enter first number: ")
		num1String, _ = reader.ReadString('\n')
		num1String = strings.TrimSpace(num1String)
		num1Number, _ = strconv.ParseFloat(num1String, 64)

		fmt.Print("Enter second number: ")
		num2String, _ = reader.ReadString('\n')
		num2String = strings.TrimSpace(num2String)
		num2Number, _ = strconv.ParseFloat(num2String, 64)

		result = num1Number - num2Number
		fmt.Printf("%f - %f = %f\n", num1Number, num2Number, result)
	} else if choice == "C" { // multiplication
		fmt.Print("Enter first number: ")
		num1String, _ = reader.ReadString('\n')
		num1String = strings.TrimSpace(num1String)
		num1Number, _ = strconv.ParseFloat(num1String, 64)

		fmt.Print("Enter second number: ")
		num2String, _ = reader.ReadString('\n')
		num2String = strings.TrimSpace(num2String)
		num2Number, _ = strconv.ParseFloat(num2String, 64)

		result = num1Number * num2Number
		fmt.Printf("%f * %f = %f\n", num1Number, num2Number, result)
	} else if choice == "D" { // division
		fmt.Print("Enter dividend: ")
		num1String, _ = reader.ReadString('\n')
		num1String = strings.TrimSpace(num1String)
		num1Number, _ = strconv.ParseFloat(num1String, 64)

		fmt.Print("Enter divisor: ")
		num2String, _ = reader.ReadString('\n')
		num2String = strings.TrimSpace(num2String)
		num2Number, _ = strconv.ParseFloat(num2String, 64)

		result = num1Number / num2Number
		fmt.Printf("%f / %f = %f\n", num1Number, num2Number, result)
	} else if choice == "E" { // Absolute value
		fmt.Print("Enter number: ")
		num1String, _ = reader.ReadString('\n')
		num1String = strings.TrimSpace(num1String)
		num1Number, _ = strconv.ParseFloat(num1String, 64)

		result = math.Abs(num1Number)
		fmt.Printf("The absolute value of %f = %f\n", num1Number, result)
	} else if choice == "F" { // Round
		fmt.Print("Enter number: ")
		num1String, _ = reader.ReadString('\n')
		num1String = strings.TrimSpace(num1String)
		num1Number, _ = strconv.ParseFloat(num1String, 64)

		result = math.Round(num1Number)
		fmt.Printf("Rounded value of %f = %f\n", num1Number, result)
	} else if choice == "G" { // exponent
		fmt.Print("Enter base number: ")
		num1String, _ = reader.ReadString('\n')
		num1String = strings.TrimSpace(num1String)
		num1Number, _ = strconv.ParseFloat(num1String, 64)

		fmt.Print("Enter exponent: ")
		num2String, _ = reader.ReadString('\n')
		num2String = strings.TrimSpace(num2String)
		num2Number, _ = strconv.ParseFloat(num2String, 64)

		result = math.Pow(num1Number, num2Number)
		fmt.Printf("%f ^ %f = %f\n", num1Number, num2Number, result)
	} else if choice == "H" { // square root
		fmt.Print("Enter number: ")
		num1String, _ = reader.ReadString('\n')
		num1String = strings.TrimSpace(num1String)
		num1Number, _ = strconv.ParseFloat(num1String, 64)

		result = math.Sqrt(num1Number)
		fmt.Printf("Square root of %f = %f\n", num1Number, result)
	} else {
		fmt.Println("Invalid. Please choose between A-H.")
	}

	fmt.Println("\nDone.")
}
