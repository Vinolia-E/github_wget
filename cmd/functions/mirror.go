package functions

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"golang.org/x/net/html"
)

func MirrorWebsite(urlInput string) {
	domainName := "Example"
	response, err := http.Get(urlInput)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}

	defer response.Body.Close()

	doc, err := html.Parse(response.Body)
	if err != nil {
		fmt.Println("Failed to parse HTML:", err)
		return
	}

	parsedUrl, _ := url.Parse(urlInput)
	domainName = parsedUrl.Hostname()

	// Extracting links
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

	for _, resource := range resourses {
		if strings.HasPrefix(resource, "https://") && strings.Contains(resource, "https://") {
			// fmt.Println("Extracted links / Found resources: \n")
			// fmt.Println("Domain name: ", resource[len("https://"):], "\n")
			domainName = resource[len("https://"):]
		}
	}
	// Make a directory for the domain using the domain name
	err = os.MkdirAll(domainName, 0o755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	// Save the HTML file
	htmlFilePath := path.Join(domainName, "index.html")
	outFile, err := os.Create(htmlFilePath)
	if err != nil {
		fmt.Println("Failed to save HTML file:", err)
		return
	}

	defer response.Body.Close()
	io.Copy(outFile, response.Body)
}
