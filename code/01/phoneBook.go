package main

import (
	"fmt"
	"os"
	"path"
)

type Entry struct {
	Name    string
	Surname string
	Phone   string
}

var data = []Entry{}

func search(key string) *Entry {
	for index, value := range data {
		if value.Surname  == key {
			return &data[index]
		}
	}
	return nil
}

func list() {
	for _, value := range data {
		fmt.Println(value)
	}
}

func printHelp(filename string) {
	fmt.Printf("Usage: %s search|list <args>\n", filename)
}

func main() {
	arguments := os.Args
	filename := path.Base(arguments[0])

	data = append(data, Entry{"Name_1", "Surame_1", "Phone_1"})
	data = append(data, Entry{"Name_2", "Surame_2", "Phone_2"})
	data = append(data, Entry{"Name_3", "Surame_3", "Phone_3"})

	// no args provided, arguments[0] = filename
	if len(arguments) == 1 {
		printHelp(filename)
		return
	}

	// at least one arg
	switch arguments[1] {
	case "list":
		list()
	case "search":
		if len(arguments) != 3 {
			fmt.Printf("Usage: %s search Surname\n", filename)
		} else {
			result := search(arguments[2])
			if result == nil {
				fmt.Printf("Entry \"%v\" not found\n", arguments[2])
				return
			}
			fmt.Println(*result)
		}
	default:
		printHelp(filename)
	}
}
