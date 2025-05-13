package functions

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

func MirrorWeb(url string) {
	// filepath := ""
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error getting the URL:", err)
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
	foldername := filename

	er := os.MkdirAll(foldername, 0o755)
	if er != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	fmt.Println("folder name:", foldername)
	filepath := path.Join(foldername, "index.html")
	// filepath := path.Join(foldername, filename)

	output, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer output.Close()

	// Copy data from response to file
	data, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}
	fileSizeStr := FormatSize(data)

	fmt.Printf("content size: %s\n", fileSizeStr)
	fmt.Println("saving file to:", filepath)
	fmt.Printf("Downloaded %s\n", url)
	// Print finish time
	endTime := time.Now()
	fmt.Printf("finished at %s\n", endTime.Format("2006-01-02 15:04:05"))
}
