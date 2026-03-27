// Author @raymondproguy
package main

import (
	"fmt"
	"strconv"
)

// The program entry point
func main() {
	fmt.Println("Welcome to the ProGuy basic CLI calculator")
	fmt.Print("\n")
	for {
		var input1 string
		fmt.Println("Enter any first number: ")
		fmt.Scan(&input1)

		num1, err := strconv.Atoi(input1)
		if err != nil {
			fmt.Println("Error: enter a valid number")
			continue
		}

		var input2 string
		fmt.Println("Enter any second number: ")
		fmt.Scan(&input2)

		num2, err := strconv.Atoi(input2)
		if err != nil {
			fmt.Println("Error: enter a valid number")
			continue
		}

		var choice string
		fmt.Println("Choose any operator: add, sub, mul, div, help, quit")
		fmt.Print("> ")
		fmt.Scan(&choice)

		switch choice {
		case "add":
			fmt.Println("Anwser: ")
			fmt.Print("\n")
			fmt.Println(add(num1, num2))
		case "sub":
			fmt.Println("Anwser: ")
			fmt.Print("\n")
			fmt.Println(sub(num1, num2))
		case "mul":
			fmt.Println("Anwser: ")
			fmt.Print("\n")
			fmt.Println(mul(num1, num2))
		case "div":
			if num2 == 0 {
				fmt.Println("Error Non diveable by zero (0)")
				fmt.Print("\n")
				continue
			}
			fmt.Println("Anwser: ")
			fmt.Print("\n")
			fmt.Println(div(num1, num2))
		case "help":
			fmt.Println("Commands......")
			fmt.Println("add  -->  Adds two numbers")
			fmt.Println("sub  -->  Substract two numbers")
			fmt.Println("mul  -->  Multiply two numbers")
			fmt.Println("div  -->  Divids two numbers")
			fmt.Println("help -->  All the commands")
			fmt.Println("quit -->  Quits the program")


		case "quit":
			fmt.Println("Thanks for using the ProGuy's basic calculator program")
			return
			default :
			fmt.Println("Invalid input or an error pls run the help command, ")
		}
	}
}


// Basic functions that runs the maths logic.
func add(f, s int) int {
	return f + s
}

func sub(f, s int) int {
	return f - s
}

func mul(f, s int) int {
	return f * s
}

func div(f, s int) int {
	return f / s
}