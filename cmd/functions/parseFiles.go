package functions

import (
	"fmt"
	"strings"
)

func ParseFlags(args []string) {
	value := ""
	for _, arg := range args {
		if arg == "-B" {
			fmt.Println("-B, output downloaded to the background and redirected to a log file")
		} else if strings.HasPrefix(arg, "-O=") {
			value = arg[len("-O="):]
			fmt.Println("-0 flag passed")
		} else if strings.HasPrefix(arg, "https:/") || strings.HasPrefix(arg, "http:/") {
			fmt.Println("valid url")
		} else if strings.HasPrefix(arg, "-p=") {
			value = arg[len("-p="):]
			fmt.Println("Path to save file")
		} else if strings.HasPrefix(arg, "--rate-limit=") {
			value = arg[len("--rate-limit="):]
			fmt.Println("download speed limit")
		} else if arg == "--mirror" {
			fmt.Println("Instance of mirror")
		} else if strings.HasPrefix(arg, "-i="){
			value = arg[len("-i="):]
			fmt.Println("Downlowd different files from a file")
		}
		//  else {
		// 	fmt.Println("undefined arguements yet")
		// }
		fmt.Println(value)
	}
}