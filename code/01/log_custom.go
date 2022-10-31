package main

import (
	"os"
	"log"
	"fmt"
	"path/filepath"
)

var LOGFILE_ABSOLUTE = filepath.Join(os.TempDir(), "log_custom.go")

func main() {
	
	logfile, err := os.OpenFile(LOGFILE_ABSOLUTE, os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644) // create, write or append only

	if err != nil {
		fmt.Println(err)
		return
	}

	defer logfile.Close() //the defer stmt tells Go to execute the statement just begore the current function returns
	iLog := log.New(logfile, "log_custom.go ", log.LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Mastering Go 3rd edition")

}