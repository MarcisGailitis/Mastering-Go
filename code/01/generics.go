package main

import "fmt"

func Print[T any](s []T) {
	for _, value := range s {
		fmt.Print(value, " ")
	}
	fmt.Println()
}

func main() {
	string_slice := []string{"one", "two", "three"}
	nr_slice := []int{1,2,3}
	Print(string_slice)
	Print(nr_slice)
}
