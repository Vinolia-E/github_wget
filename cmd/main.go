package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid number of arguements\nUsage: go run . <URL>")
		return
	}

	// flag := ArgValues{}
	args := os.Args[1:]
	url := ""
	filename := ""
	startTime := time.Now()
	for _, arg := range args {
		if strings.HasPrefix(arg, "https://") || strings.HasPrefix(arg, "http://") {
			argsplit := strings.Split(arg, "/")
			filename = argsplit[len(argsplit)-1]
			url = arg
		}
	}
	// fmt.Println(url)
	// fmt.Println(filename)

	fmt.Println("Start at:", startTime.Format("2006-01-02 15:04:05"))
	fmt.Println("Sending request to:", url)
	err := DownloadFile(filename, url)

	if err != nil {
		fmt.Println("Error: ", err)
	}else {
		fmt.Println("Download complete!")
	}

	// flag.ParseFlags(args)

	// fmt.Println(args)
	fmt.Println("End!")
}

func DownloadFile(filename, url string) error {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return DownloadHTTPFile(filename, url)
	}
	return fmt.Errorf("unsupported protocol: %s", url)
}

func DownloadHTTPFile(filename, url string) error {
	download, err := http.Get(url)
	if err != nil {
		return err
	}

	defer download.Body.Close()

	if download.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", download.Status)
	} else {
		fmt.Println("status",http.StatusOK, "OK")
	}

	output, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, download.Body)
	return err
}
