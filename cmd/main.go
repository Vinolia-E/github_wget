package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
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

	// if flags.URL == nil {
	if flags.URL == nil && flags.InputFile == "" {
		fmt.Println("Error: No URL provided")
		return
	}

	if flags.InputFile != "" {
		inputs := functions.ReadInputFile(flags.InputFile)
		fmt.Println(inputs)

		// flags.URL = url
		flags.URL = append(flags.URL, inputs...)
	}

	// Handle background mode
	if flags.Background {
		// functions.BackgroundMode()

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

	// Print start time
	startTime := time.Now()
	fmt.Printf("start at %s\n", startTime.Format("2006-01-02 15:04:05"))
	fmt.Print("sending request, awaiting response... ")

	// StartDownloads(args, flags.URL, flags.Path)
	// Send GET request

	for _, url := range flags.URL {
		fmt.Println("THIS IS THE URL", url)
		response, err := http.Get(url) //(flags.URL)
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

		filename := path.Base(url) //(flags.URL)
		if flags.OutputFile != "" {
			filename = flags.OutputFile
		}

		// Determine file path
		filepath := filename
		if flags.Path != "" {
			// fmt.Println("file path 1 : ", flags.Path)
			if strings.HasPrefix(flags.Path, "~/") {
				homeDir, err := os.UserHomeDir()
				if err != nil {
					fmt.Println("Error getting home directory:", err)
					return
				}
				filepath = path.Join(homeDir, flags.Path[2:], filename)
				// fmt.Println("file path is : ", filepath)
				// return
			} else {

				err := os.MkdirAll(flags.Path, 0o777)
				if err != nil {
					fmt.Println("Error creating directory:", err)
					return
				}
				filepath = path.Join(flags.Path, filename)
			}
			// fmt.Println("file path is : ", filepath)
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
		fmt.Printf("Downloaded %s\n", flags.URL)

		// Print finish time
		endTime := time.Now()
		fmt.Printf("finished at %s\n", endTime.Format("2006-01-02 15:04:05"))
	}
}

func StartDownloads(args, urls []string, outputDir string) {

	flags := &functions.ArgValues{}
	flags.ParseFlags(args)
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			if flags.IsMirror {
				MirrorWebsite(u, outputDir)
			} else {
				DownloadFile(u, "", outputDir)
			}
		}(url)
	}
	wg.Wait()
}

func MirrorWebsite(baseURL, outputDir string) {
	fmt.Println("Mirroring website:", baseURL)
	htmlFilePath := DownloadFile(baseURL, "index.html", outputDir)
	if htmlFilePath == "" {
		fmt.Println("Failed to download main page.")
		return
	}

	resources := ExtractResources(htmlFilePath)
	var wg sync.WaitGroup
	for _, resource := range resources {
		wg.Add(1)
		go func(res string) {
			defer wg.Done()
			DownloadFile(res, "", outputDir)
		}(resource)
	}
	wg.Wait()
	fmt.Println("Website mirrored successfully.")
}

func ExtractResources(filePath string) []string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading HTML file:", err)
		return nil
	}
	var resources []string
	re := regexp.MustCompile(`(src|href)="([^"]+\.(css|js|png|jpg|jpeg|gif|svg))"`)
	for _, match := range re.FindAllStringSubmatch(string(data), -1) {
		resources = append(resources, match[2])
	}
	return resources
}

func DownloadFile(url, outputFile, outputDir string) string {
	fmt.Println("Downloading:", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Download failed: Received status", resp.Status)
		return ""
	}

	filename := outputFile
	if filename == "" {
		filename = path.Base(url)
	}

	if outputDir != "" {
		os.MkdirAll(outputDir, os.ModePerm)
		filename = path.Join(outputDir, filename)
	}

	outFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return ""
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return ""
	}

	fmt.Printf("Downloaded %s\n", url)
	return filename
}

// var (
// 	rateLimit  int64 = 0 // Default: No rate limit
// 	mirrorMode bool  = false
// )

// func main() {
// 	args := os.Args[1:]
// 	if len(args) < 1 {
// 		fmt.Println("Usage: go run main.go [FLAGS] <URL>")
// 		return
// 	}

// 	outputDir, inputFile, background, urls := parseArguments(args)

// 	if inputFile != "" {
// 		fileUrls, err := readURLsFromFile(inputFile)
// 		if err != nil {
// 			fmt.Println("Error reading file:", err)
// 			return
// 		}
// 		urls = append(urls, fileUrls...)
// 	}

// 	if len(urls) == 0 {
// 		fmt.Println("Error: No URL provided")
// 		return
// 	}

// 	if background {
// 		enableBackgroundMode()
// 	}

// 	startDownloads(urls, outputDir)
// }
