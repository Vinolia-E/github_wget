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

// Function downloading all the filles from the website when flag --mirror is used.

func MirrorWeb(url string, flags *ArgValues) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error getting the URL:", err)
		return
	}

	defer response.Body.Close()

	docinfo, err := html.Parse(response.Body)
	if err != nil {
		fmt.Println("Failed to parse HTML:", err)
		return
	}

	// Extracting links
	resourses := []string{}

	var walker func(*html.Node)

	walker = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if (n.Data == "a" || n.Data == "link" || n.Data == "img") && (attr.Key == "href" || attr.Key == "src") {
					resourse := n.Data + " " + attr.Key + " " + attr.Val
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

	filepaths := []string{}
	subfiles := ""

	walker(docinfo)

	// Print response status
	fmt.Printf("status %s\n", response.Status)
	if response.StatusCode != http.StatusOK {
		fmt.Println("Download failed: Received status", response.Status)
		return
	}

	foldername := path.Base(url)
	filename := "index.html"

	er := os.MkdirAll(foldername, 0o755)
	if er != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	// fmt.Println("folder name:", foldername)
	filepath := path.Join(foldername, filename)

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

	for _, resource := range resourses {
		if !strings.HasSuffix(resource, ".css") && !strings.HasPrefix(resource, "img") {
			continue
		} else {
			fmt.Println("Resource: ", resource)

			getLink := strings.Split(resource, " ")
			resp, err := http.Get(getLink[2])
			if err != nil {
				fmt.Println("Error getting the URL:", err)
				return
			}

			defer resp.Body.Close()

			subfiles = path.Base(resource)
			filepaths = append(filepaths, subfiles)

			filepath = path.Join(foldername, subfiles)
			output, err := os.Create(filepath)
			if err != nil {
				fmt.Println("Error creating file:", err)
				return
			}
			defer output.Close()

			// Copy data from response to file
			data, err := io.Copy(output, resp.Body)
			if err != nil {
				fmt.Println("Error saving file:", err)
				return
			}
			fmt.Println("subfiles: ", subfiles)
			// fmt.Println("Data: ", data)
			fileSizeStr := FormatSize(data)
			fmt.Printf("content size: %s\n", fileSizeStr)
			fmt.Println("saving file to:", filepath)
		}
	}
	DownloadHtmlFile(url, flags)

	fileSizeStr := FormatSize(data)
	fmt.Printf("content size: %s\n", fileSizeStr)
	fmt.Println("saving file to:", filepath)
	fmt.Printf("Downloaded %s\n", url)
	// Print finish time
	endTime := time.Now()
	fmt.Printf("finished at %s\n", endTime.Format("2006-01-02 15:04:05"))
}
