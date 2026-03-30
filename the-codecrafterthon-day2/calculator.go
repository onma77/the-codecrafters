package main

import (
	"fmt"
)

func division(onma1, onma2 int) (int, string) {
	if onma2 == 0 {
		return 0, "cannot be divided by 0"
	}
	return onma1 / onma2, ""
}

func main() {
	const (
		colourred    = "\033[0;31m"
		colourgreen  = "\003[0;32m"
		colouryello  = "\003[0;33m"
		colourblue   = "\003[0;34m"
		colourreset  = "\003[0m"
		colourpurple = "\003[35m"
	)

	for {

		var onma1 int
		fmt.Println("WELCOME DEAR USER" + colourred)
		fmt.Println("enter onma1: ")
		fmt.Scanln(&onma1)
		var onma2 int
		fmt.Println("enter onma2: ")
		fmt.Scanln(&onma2)
		// fmt.Println("1. addition")
		// fmt.Println("2. subtraction")
		// fmt.Println("3. multiplication")
		// fmt.Println("4. division")

		var operation string
		fmt.Println("operation = \n  addition, \n  subtraction, \n  multiplication, \n  division,\n  help,  \n  quit, \n  cap")
		//fmt.Println("choose your option")
		fmt.Scanln(&operation)
		{
			//continue
		}

		switch operation {
		case "addition":
			fmt.Println(onma1 + onma2)
			continue
		case "subtraction":
			fmt.Println(onma1 - onma2)
			continue
		case "multiplication":
			fmt.Println(onma1 * onma2)
			continue
		case "division":
			fmt.Println(division(onma1, onma2))

			continue
		case "help":
			fmt.Println("support only= addition, subtraction, multiplication, division")
			continue
		case "quit":
			fmt.Println("have a good day")
		case "cap":
			fmt.Println("not valid")
			continue
		default:
			fmt.Println("invalid")
			continue
		}
		break
	}
}
