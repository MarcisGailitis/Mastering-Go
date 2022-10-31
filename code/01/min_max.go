package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Please provide a command line arguments")
	}

	fmt.Println("You entered:", args[1:])

	var min, max int
	var non_int []string

	for inx, value := range args[1:] {
		i, err := strconv.Atoi(value)
		if err != nil {
			non_int = append(non_int, value)
		}

		if inx == 1 {
			min, max = i, i
			continue
		}

		if i < min {
			min = i
		} else if i > max {
			max = i
		}

	}

	fmt.Println("min:", min, "max:", max)
	fmt.Println("Non-int input:", non_int)

}
