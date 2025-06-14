package functions

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

// DownloadHtmlFile downloads the HTML file from the specified URL and saves it to a file.
// It takes the URL to download and a pointer to ArgValues containing command-line flags and options.
func DownloadHtmlFile(url string, args *ArgValues) {
	response, err := http.Get(url)
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
	filename := path.Base(url)

	// Determine file path
	filepath := filename
	if args.Path == "" {
		args.Path = filename
		if strings.HasPrefix(args.Path, "~/") {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				fmt.Println("Error getting home directory:", err)
				return
			}
			filepath = path.Join(homeDir, args.Path[2:], filename)
		} else {
			err := os.MkdirAll(args.Path, 0o777)
			if err != nil {
				fmt.Println("Error creating directory:", err)
				return
			}
			filepath = path.Join(args.Path, "index.html")

		}
	} else {
		fmt.Println("File path should not be applied when mirroring a website")
		return
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

	// Convert size to appropriate unit
	fileSizeStr := FormatSize(size)

	fmt.Printf("content size: %s [%s]\n", size, fileSizeStr)
	fmt.Println("saving file to:", filepath)
	fmt.Printf("Downloaded %s\n", url)

	// Print finish time
	endTime := time.Now()
	fmt.Printf("finished at %s\n", endTime.Format("2006-01-02 15:04:05"))
}
