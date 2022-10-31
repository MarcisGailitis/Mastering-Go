package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command line arguments")
		return
	}

	argument := os.Args[1]

	switch argument {
	case "0":
		fmt.Println("Zero")
	case "1":
		fmt.Println("One")
	case "2", "3", "4":
		fmt.Println("Two, Three or Four")
		fallthrough
	default:
		fmt.Println("Value:", argument)
	}

	value, err := strconv.Atoi(argument)
	if err != nil {
		fmt.Println("Cannot convert to int:", argument)
		return
	}

	switch {
	case value == 0: 
		fmt.Println("Zero")
	case value > 0:
		fmt.Println("Positive Int")
	case value < 0:
		fmt.Println("Negative Int")
	default:
		fmt.Println("This should not happen:", value)
	}

}
