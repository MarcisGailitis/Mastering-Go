package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Printf("Type: %T, content: ", args)
	fmt.Println(args)
	for inx, arg := range args {
		fmt.Println(inx, arg)
	}
}