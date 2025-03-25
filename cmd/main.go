package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
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

	if flags.URL == "" {
		fmt.Println("Error: No URL provided")
		return
	}

	// // Handle background modeg
	if flags.Background {
		fmt.Println("Output will be written to 'wget-log'.")
		file, err := os.OpenFile("wget-log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening log file:", err)
			return
		}
		defer file.Close()
		os.Stdout = file
		os.Stderr = file
	}

	// Print start time
	startTime := time.Now()
	fmt.Printf("start at %s\n", startTime.Format("2006-01-02 15:04:05"))
	fmt.Print("sending request, awaiting response... ")

	// Send GET request
	// response, err := http.Get(url)
	response, err := http.Get(flags.URL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	// Print response status
	fmt.Printf("status %s\n", response.Status)
	if response.StatusCode != http.StatusOK {
		fmt.Println("Download failed: Received status", response.Status)
		return
	}


	filename := path.Base(flags.URL)
	if flags.OutputFile != "" {
		filename = flags.OutputFile
	}

	// Determine file path
	filepath := filename
	if flags.Path != "" {
		fmt.Println("file path 1 : ", flags.Path)
		err := os.MkdirAll(flags.Path, 0777)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
		filepath = path.Join(flags.Path, filename)
		fmt.Println("file path 7 : ", flags.Path)
	}

	
	// Create the output file
	outFile, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outFile.Close()
	
	// Copy data from response to file
	size, err := io.Copy(outFile, response.Body)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}
	
	// // Convert size to appropriate unit (bytes, KB, MB, GB)
	fileSizeStr := functions.FormatSize(size)
	
	fmt.Printf("content size: %s\n", fileSizeStr)
	fmt.Println("saving file to:", filepath)
	fmt.Printf("Downloaded [%s]\n", flags.URL)

	// Print finish time
	endTime := time.Now()
	fmt.Printf("finished at %s\n", endTime.Format("2006-01-02 15:04:05"))
}

