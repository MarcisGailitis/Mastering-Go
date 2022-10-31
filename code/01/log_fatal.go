package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("FATAL: Hello world!")
	} else {
		log.Panic("PANIC: Hello world!")
	}
}