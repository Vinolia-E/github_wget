package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid number of arguements\nUsage: go run . <URL>")
		return
	}

	flag := ArgValues{}
	args := os.Args[1:]
	startTime := time.Now()

	fmt.Println("Start at:", startTime.Format("2006-01-02 15:04:05"))

	// functions.ParseFlags(args)
	flag.ParseFlags(args)


	fmt.Println(args)
	fmt.Println("End!")
}
