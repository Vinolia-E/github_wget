package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid number of arguements\nUsage: go run . <URL>")
		return
	}

	args := os.Args[1:]
	startTime := time.Now()
	url := ""
	fmt.Println("Start at:", startTime.Format("2006-01-02 15:04:05"))

	for _, arg := range args {
		if strings.Contains(arg, "https:") {
			url = arg
		}
	}

	fmt.Println(url)

}