package functions

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func MirrorWeb(url string) {
	// filepath := ""
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error getting the URL:", err)
		return
	}

	defer response.Body.Close()

	/**/
	docinfo, err := html.Parse(response.Body)
	if err != nil {
		fmt.Println("Failed to parse HTML:", err)
		return
	}

	// Extracting links
	resourses := []string{}
	count := 1
	var walker func(*html.Node)

	walker = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if (n.Data == "a" || n.Data == "link" || n.Data == "img") && (attr.Key == "href" || attr.Key == "src") {
					resourse := n.Data + " : " + attr.Key + " : " + attr.Val
					exist := false

					for _, rs := range resourses {
						if rs == resourse {
							exist = true
							break
						}
					}

					if !exist {
						resourses = append(resourses, resourse)
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walker(c)
		}
	}
	walker(docinfo)

	for _, resource := range resourses {
		fmt.Println("Resource from html:", count, " :", resource)
		count++
	}
	/**/

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
	filepath := path.Join(foldername, filename)
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

	/*

	 */

	fmt.Printf("content size: %s\n", fileSizeStr)
	fmt.Println("saving file to:", filepath)
	fmt.Printf("Downloaded %s\n", url)
	// Print finish time
	endTime := time.Now()
	fmt.Printf("finished at %s\n", endTime.Format("2006-01-02 15:04:05"))
}

func DownloadFile2(url string, flags *ArgValues) {
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
	// if flags.OutputFile != "" {
	// 	filename = flags.OutputFile
	// }

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
			// filepath = path.Join(flags.Path, filename)
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
