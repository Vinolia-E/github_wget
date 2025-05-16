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

// Function that downloads the html file.
func DownloadHtmlFile(url string, flags *ArgValues) {
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
	if flags.Path == "" {
		flags.Path = filename
		if strings.HasPrefix(flags.Path, "~/") {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				fmt.Println("Error getting home directory:", err)
				return
			}
			filepath = path.Join(homeDir, flags.Path[2:], filename)
		} else {
			err := os.MkdirAll(flags.Path, 0o777)
			if err != nil {
				fmt.Println("Error creating directory:", err)
				return
			}
			filepath = path.Join(flags.Path, "index.html")
			
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

	fmt.Printf("content size: %s\n", fileSizeStr)
	fmt.Println("saving file to:", filepath)
	fmt.Printf("Downloaded %s\n", url)

	// Print finish time
	endTime := time.Now()
	fmt.Printf("finished at %s\n", endTime.Format("2006-01-02 15:04:05"))
}
