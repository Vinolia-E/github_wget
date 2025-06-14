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

// Function that dowloads all files.

// DownloadFile downloads a file from the specified URL and saves it to disk.
// The output location and filename can be customized using the provided ArgValues.
// If the IsMirror flag is set in args, the function delegates to MirrorWeb for mirroring.
// The function handles creating necessary directories, expanding the home directory (~),
// and prints status messages about the download progress and result.
//
// Parameters:
//   - url: The URL of the file to download.
//   - args: Pointer to ArgValues struct containing options such as output path, filename, and mirror mode.
//
// The function prints errors and status information to standard output.
func DownloadFile(url string, args *ArgValues) {
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
	if args.IsMirror {
		MirrorWeb(url, args)
		return
	}

	filename := path.Base(url)
	if args.OutputFile != "" {
		filename = args.OutputFile
	}

	// Determine file path
	filepath := filename
	if args.Path != "" {
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
			filepath = path.Join(args.Path, filename)
		}
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

	fmt.Printf("content size: %d [%s]\n", size, fileSizeStr)
	fmt.Println("saving file to:", filepath)
	fmt.Printf("Downloaded %s\n", url)

	// Print finish time
	endTime := time.Now()
	fmt.Printf("finished at %s\n\n", endTime.Format("2006-01-02 15:04:05"))
}
