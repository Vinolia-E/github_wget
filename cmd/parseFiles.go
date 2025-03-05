package main

import (
	"fmt"
	"strings"
)

type ArgValues struct {
	URL []string
	Flag string
	OutputFile string
	Path string
	Background bool
	InputFile string
	IsMirror bool
	RateLimit string
}

func (value *ArgValues) ParseFlags(args []string) {
	// value := ""
	for _, arg := range args {
		if arg == "-B" {
			value.Background = true
			// fmt.Println("Output will be written to "wget-log"")
		} else if strings.HasPrefix(arg, "-O=") {
			value.OutputFile = arg[len("-O="):]
			// fmt.Println("-0 flag passed")
		} else if strings.HasPrefix(arg, "https:/") || strings.HasPrefix(arg, "http:/") {
			value.URL = append(value.URL, arg)
			fmt.Println("valid url")
		} else if strings.HasPrefix(arg, "-p=") {
			value.Path = arg[len("-p="):]
			fmt.Println("Path to save file")
		} else if strings.HasPrefix(arg, "--rate-limit=") {
			value.RateLimit = arg[len("--rate-limit="):]
			fmt.Println("download speed limit")
		} else if arg == "--mirror" {
			value.IsMirror = true
			// fmt.Println("Instance of mirror")
		} else if strings.HasPrefix(arg, "-i="){
			value.InputFile = arg[len("-i="):]
			fmt.Println("Downlowd different files from a file")
		} else {
			fmt.Println("undefined arguements yet")
		}
		fmt.Println(value)
	}
}