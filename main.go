package main

import (
	"fmt"
	"os"
)

func main() {
	logfile, err := InitLogfile("test.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer logfile.Close()
	line, err := logfile.Getline(0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(line))
}
