package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// allows to return multiple paths not only the 1st one

func main() {
	args := os.Args
	if len(args) == 1 {
		return
	}

	file := args[1]

	pathSlice := filepath.SplitList(os.Getenv("PATH"))

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
