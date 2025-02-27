package main

import (
	"fmt"
	"os"
	"time"
	"wget/cmd/functions"
)

type ArgValues struct {
	URL []string
	Flag string
	OutputFile string
	Path string
	Background bool
	InputFile string
	IsMirror bool
	RateLimit string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid number of arguements\nUsage: go run . <URL>")
		return
	}

	args := os.Args[1:]
	startTime := time.Now()
	// url := ""
	fmt.Println("Start at:", startTime.Format("2006-01-02 15:04:05"))

	functions.ParseFlags(args)


	fmt.Println("End!")
}


