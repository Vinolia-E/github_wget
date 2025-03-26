package functions

import (
	"fmt"
	"os"
	"strings"
)

func (input *ArgValues) ParseFlags(args []string) {
	for _, arg := range args {
		if arg == "-B" {
			input.Background = true
		} else if strings.HasPrefix(arg, "-O=") {
			input.OutputFile = arg[len("-O="):]
			// fmt.Println("-0 flag passed")
		} else if strings.HasPrefix(arg, "https://") || strings.HasPrefix(arg, "http://") || strings.HasPrefix(arg, "ftp://") {
			input.HasFlag = true
			// input.URL = arg
			input.URL = append(input.URL, arg) // arg
			// fmt.Println("valid url")
		} else if strings.HasPrefix(arg, "-P=") {
			input.Path = arg[len("-P="):]
			// input.OutputFile = args[i+1]
			fmt.Println("Path to save file")
		} else if strings.HasPrefix(arg, "--rate-limit=") {
			if strings.HasSuffix(arg, "k") || strings.HasSuffix(arg, "m") || strings.HasSuffix(arg, "K") || strings.HasSuffix(arg, "M") {
				input.RateLimit = arg[len("--rate-limit="):]
				fmt.Printf("download speed limit is : %v\n", input.RateLimit)
			} else {
				fmt.Println("Invalid rate limit, \nRate limit has sufix 'm' or 'k'. ")
				os.Exit(0)
			}
		} else if arg == "--mirror" {
			input.IsMirror = true
			fmt.Println("Instance of mirror")
		} else if strings.HasPrefix(arg, "-i=") {
			input.InputFile = arg[len("-i="):]
			fmt.Println("Download different files from a file")
		}
	}
}
