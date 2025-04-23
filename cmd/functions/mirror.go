package functions

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"

	"golang.org/x/net/html"
)

func MirrorWebsite(urlInput string, baseDir string) {
	resp, err := http.Get(urlInput)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}

	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Failed to parse HTML:", err)
		return
	}

	parsedUrl, _ := url.Parse(urlInput)
	domain := parsedUrl.Hostname()

	// Create directory named after the domain

	dormainFolder := path.Join(baseDir, domain)
	err = os.MkdirAll(dormainFolder, 0o755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	// Save the HTML file
	htmlFilePath := path.Join(dormainFolder, "index.html")
	outFile, err := os.Create(htmlFilePath)
	if err != nil {
		fmt.Println("Failed to save HTML file:", err)
		return
	}

	defer resp.Body.Close()
	io.Copy(outFile, resp.Body)


//Extracting links
	resourses := []string{}
	var walker func(*html.Node)

	walker = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, attr := range n.Attr {
				if (n.Data == "a" || n.Data == "link") && attr.Key == "href" {
					resourses = append(resourses, attr.Val)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walker(c)
		}
	}
	walker(doc)

	fmt.Println("Extracted links / Found resources:")
	for _, resource := range resourses {
		fmt.Println("-", resource)
	}
}
