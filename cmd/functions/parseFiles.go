package functions

import (
	"fmt"
	"strings"
)

func (input *ArgValues)ParseFlags(args []string) {
	for _, arg := range args {
		if arg == "-B" {
			input.Background = true
			fmt.Println("Output downloaded to the background and redirected to a log file")
		} else if strings.HasPrefix(arg, "-O=") {
			input.OutputFile = arg[len("-O="):]
			// fmt.Println("-0 flag passed")
		} else if strings.HasPrefix(arg, "https:/") || strings.HasPrefix(arg, "http:/") {
			input.HasFlag = true
			input.URL = arg
			// fmt.Println("valid url")
		} else if strings.HasPrefix(arg, "-p=") {
			input.Path = arg[len("-p="):]
			// fmt.Println("Path to save file")
		} else if strings.HasPrefix(arg, "--rate-limit=") {
			input.RateLimit = arg[len("--rate-limit="):]
			fmt.Printf("download speed limit is : %v", input.RateLimit)
		} else if arg == "--mirror" {
			input.IsMirror = true
			fmt.Println("Instance of mirror")
		} else if strings.HasPrefix(arg, "-i="){
			input.InputFile = arg[len("-i="):]
			fmt.Println("Downlowd different files from a file")
		}
		// fmt.Println("Hello, complete the perse function")
	}
}