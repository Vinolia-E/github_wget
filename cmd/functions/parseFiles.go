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
		} else if strings.HasPrefix(arg, "https://") || strings.HasPrefix(arg, "http://") || strings.HasPrefix(arg, "ftp://") {
			input.HasFlag = true
			input.URL = append(input.URL, arg) // arg
		} else if strings.HasPrefix(arg, "-P=") {
			input.Path = arg[len("-P="):]
			// fmt.Println("Path to save file")
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
		} else if strings.HasPrefix(arg, "-i=") {
			input.InputFile = arg[len("-i="):]
		}
	}
}
