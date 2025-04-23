package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"wget/cmd/functions"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [FLAGS] <URL>")
		return
	}
	args := os.Args[1:]

	// Print start time
	startTime := time.Now()
	fmt.Printf("start at %s\n", startTime.Format("2006-01-02 15:04:05"))
	fmt.Print("sending request, awaiting response... ")

	flags := &functions.ArgValues{}
	flags.ParseFlags(args)

	if flags.URL == nil && flags.InputFile == "" {
		fmt.Println("Error: No URL provided")
		return
	}

	if flags.InputFile != "" && !flags.IsMirror {
		inputs := functions.ReadInputFile(flags.InputFile)
		flags.URL = append(flags.URL, inputs...)
	}

	// Handle background mode
	if flags.Background {
		fmt.Println("Output will be written to 'wget-log'.")
		file, err := os.OpenFile("wget-log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			fmt.Println("Error opening log file:", err)
			return
		}
		defer file.Close()
		os.Stdout = file
		os.Stderr = file
	}

	// Use a WaitGroup for async downloads
	var wg sync.WaitGroup

	for _, url := range flags.URL {
		wg.Add(1)
		go func(url string) { // Run each download in a goroutine
			defer wg.Done()
			functions.DownloadFile(url, flags)
		}(url)
	}

	wg.Wait() // Wait for all downloads to finish

	if flags.IsMirror {
		fmt.Println("Mirroring website... to be done")
		return
	}
}
