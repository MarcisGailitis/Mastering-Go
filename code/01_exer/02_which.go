package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// allows to search for multiple executable files, separated by a space

func main() {
	args := os.Args
	if len(args) == 1 {
		return
	}

	// read all arguments and store them in a slice
	files := args[1:]
	
	pathSlice := filepath.SplitList(os.Getenv("PATH"))

	// iteraring over argument slice
	for _, file := range files {
		
		for _, directory := range pathSlice {
			fullPath := filepath.Join(directory, file)
			fileInfo, err := os.Stat(fullPath)

			if err == nil {
				mode := fileInfo.Mode()
				if mode.IsRegular() {
					if mode&0111 != 0 {
						fmt.Println(fullPath)
					}
				}
			}
		}
	}
}
