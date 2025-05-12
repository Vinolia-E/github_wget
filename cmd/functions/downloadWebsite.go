package functions

import (
	"fmt"
	"net/http"
	"os"
	"path"
)

func MirrorWeb(url string) {
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

foldername := path.Base(url)
er := os.MkdirAll(foldername, 0o755)
if er != nil {
	fmt.Println("Error creating directory:", err)
	return
}
	fmt.Println("folder name:", foldername)
	return
}

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"path"
// 	"strings"
// 	"time"
// )

// func DownloadFile(url string, flags *ArgValues) {
// 	response, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer response.Body.Close()

// 	// Print response status
// 	fmt.Printf("status %s\n", response.Status)
// 	if response.StatusCode != http.StatusOK {
// 		fmt.Println("Download failed: Received status", response.Status)
// 		return
// 	}
// 	// if flags.IsMirror {
// 	// 	// fmt.Println("Skipping download for mirror mode")
// 	// 	MirrorWebsite(url)
// 	// 	// MirrorWeb(url)
// 	// 	fmt.Println("Mirror mode  implementation trial. Debug in download.go")
// 	// 	return
// 	// }

// 	filename := path.Base(url)
// 	if flags.OutputFile != "" {
// 		filename = flags.OutputFile
// 	}

// 	// Determine file path
// 	filepath := filename
// 	if flags.Path != "" {
// 		if strings.HasPrefix(flags.Path, "~/") {
// 			homeDir, err := os.UserHomeDir()
// 			if err != nil {
// 				fmt.Println("Error getting home directory:", err)
// 				return
// 			}
// 			filepath = path.Join(homeDir, flags.Path[2:], filename)
// 		} else {
// 			err := os.MkdirAll(flags.Path, 0o777)
// 			if err != nil {
// 				fmt.Println("Error creating directory:", err)
// 				return
// 			}
// 			filepath = path.Join(flags.Path, filename)
// 		}
// 	}

// 	// Create the output file
// 	outFile, err := os.Create(filepath)
// 	if err != nil {
// 		fmt.Println("Error creating file:", err)
// 		return
// 	}
// 	defer outFile.Close()

// 	// Copy data from response to file
// 	size, err := io.Copy(outFile, response.Body)
// 	if err != nil {
// 		fmt.Println("Error saving file:", err)
// 		return
// 	}

// 	// Convert size to appropriate unit
// 	fileSizeStr := FormatSize(size)

// 	fmt.Printf("content size: %s\n", fileSizeStr)
// 	fmt.Println("saving file to:", filepath)
// 	fmt.Printf("Downloaded %s\n", url)

// 	// Print finish time
// 	endTime := time.Now()
// 	fmt.Printf("finished at %s\n", endTime.Format("2006-01-02 15:04:05"))
// }

// func MirrorWebsite(urlInput string) {
// 	response, err := http.Get(urlInput)
// 	if err != nil {
// 		fmt.Println("Error fetching URL:", err)
// 		return
// 	}

// 	defer response.Body.Close()

// 	doc, err := html.Parse(response.Body)
// 	if err != nil {
// 		fmt.Println("Failed to parse HTML:", err)
// 		return
// 	}

// 	parsedUrl, _ := url.Parse(urlInput)
// 	domainName := parsedUrl.Hostname()

// 	// Extracting links
// 	resourses := []string{}
// 	var walker func(*html.Node)

// 	walker = func(n *html.Node) {
// 		if n.Type == html.ElementNode {
// 			for _, attr := range n.Attr {
// 				if (n.Data == "a" || n.Data == "link") && attr.Key == "href" {
// 					resourses = append(resourses, attr.Val)
// 				}
// 			}
// 		}

// 		for c := n.FirstChild; c != nil; c = c.NextSibling {
// 			walker(c)
// 		}
// 	}
// 	walker(doc)

// 	for _, resource := range resourses {
// 		if strings.HasPrefix(resource, "https://") && strings.Contains(resource, "https://") {
// 			domainSplit := strings.Split(resource[len("https://"):], "/")
// 			domainName = domainSplit[0]
// 		}
// 	}
// 	// Make a directory for the domain using the domain name
// 	err = os.MkdirAll(domainName, 0o755)
// 	if err != nil {
// 		fmt.Println("Error creating directory:", err)
// 		return
// 	}
// 	// Save the HTML file
// 	htmlFilePath := path.Join(domainName, "index.html")
// 	outFile, err := os.Create(htmlFilePath)
// 	// outFile, err := os.Create(domainName)
// 	if err != nil {
// 		fmt.Println("Failed to save HTML file:", err)
// 		return
// 	}

// 	defer response.Body.Close()
// 	io.Copy(outFile, response.Body)
// }
