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
	flags := &functions.ArgValues{}
	flags.ParseFlags(args)

	if (flags.Reject || flags.Exclude || flags.ConvertLinks )&& !flags.IsMirror {
		fmt.Println("Error: --reject, --exclude, and --convert-links can only be used with --mirror flag")
		return
	}

	if !flags.Background{
		// Print start time on the console if not in background mode
		startTime := time.Now()
		fmt.Printf("start at %s\n", startTime.Format("2006-01-02 15:04:05"))
		fmt.Print("sending request, awaiting response... ")
	}


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
		output := functions.BackgroundDownload()
		startTime := time.Now()
		fmt.Printf("start at %s\n", startTime.Format("2006-01-02 15:04:05"))
		fmt.Print("sending request, awaiting response... ")
		if output != nil {
			defer output.Close()
		}
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
}


