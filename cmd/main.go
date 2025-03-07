package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Usage: go run main.go [FLAGS] <URL>")
		return
	}

	// Default values
	var outputFile, outputDir string
	background := false
	var url string

	// Manual argument parsing
	for i := 0; i < len(args); i++ {
		if args[i] == "-B" {
			background = true
		} else if strings.HasPrefix(args[i], "-O=") {
			outputFile = strings.TrimPrefix(args[i], "-O=")
		} else if strings.HasPrefix(args[i], "-P=") {
			outputDir = strings.TrimPrefix(args[i], "-P=")
		} else {
			url = args[i]
		}
	}

	if url == "" {
		fmt.Println("Error: No URL provided")
		return
	}

	// Handle background mode
	if background {
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
	fmt.Println("sending request, awaiting response...")

	// Send GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Print response status
	fmt.Printf("status %s\n", resp.Status)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Download failed: Received status", resp.Status)
		return
	}

	// Determine filename
	filename := path.Base(url)
	if outputFile != "" {
		filename = outputFile
	}

	// Determine file path
	filepath := filename
	if outputDir != "" {
		if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
		filepath = path.Join(outputDir, filename)
	}

	fmt.Println("saving file to:", filepath)

	// Create the output file
	outFile, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outFile.Close()

	// Copy data from response to file
	size, err := io.Copy(outFile, resp.Body)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}

	// Convert size to appropriate unit (bytes, KB, MB, GB)
	fileSizeStr := formatSize(size)

	fmt.Printf("Downloaded [%s]\n", url)
	fmt.Printf("content size: %s\n", fileSizeStr)

	// Print finish time
	endTime := time.Now()
	fmt.Printf("finished at %s\n", endTime.Format("2006-01-02 15:04:05"))
}

// Converts file size to KB, MB, or GB for readability
func formatSize(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%d bytes", size)
	} else if size < 1024*1024 {
		return fmt.Sprintf("%.2f KB", float64(size)/1024)
	} else if size < 1024*1024*1024 {
		return fmt.Sprintf("%.2f MB", float64(size)/(1024*1024))
	}
	return fmt.Sprintf("%.2f GB", float64(size)/(1024*1024*1024))
}
