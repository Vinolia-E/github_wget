package functions

import (
	"fmt"
	"os"
)

func BackgroundDownload() *os.File {
	fmt.Println("Output will be written to 'wget-log'.")
	file, err := os.OpenFile("wget-log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		os.Exit(0)
	}
	// defer file.Close()
	os.Stdout = file
	os.Stderr = file
	return file
}
