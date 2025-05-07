package functions

import (
	"fmt"
	"os"
	"strings"
)

func (input *ArgValues) ParseFlags(args []string) {
	mirrorMode := false // Flag to track if mirror is used.
	track := false      // Flag to track if a source file is provided (-i=)

	for _, arg := range args {
		if arg == "-B" {
			input.Background = true
		} else if strings.HasPrefix(arg, "-O=") {
			input.OutputFile = arg[len("-O="):]
		} else if strings.HasPrefix(arg, "https://") || strings.HasPrefix(arg, "http://") || strings.HasPrefix(arg, "ftp://") {
			input.HasFlag = true
			input.URL = append(input.URL, arg)
		} else if strings.HasPrefix(arg, "-P=") {
			input.Path = arg[len("-P="):]
		} else if strings.HasPrefix(arg, "--rate-limit=") {
			if strings.HasSuffix(arg, "k") || strings.HasSuffix(arg, "m") || strings.HasSuffix(arg, "K") || strings.HasSuffix(arg, "M") {
				input.RateLimit = arg[len("--rate-limit="):]
				fmt.Printf("download speed limit is : %v\n", input.RateLimit)
			} else {
				fmt.Println("\n\nInvalid rate limit, Rate limit has sufix 'm' or 'k'. \n")
				os.Exit(0)
				// return
			}
		} else if strings.HasPrefix(arg, "--mirror") {
			input.IsMirror = true
			mirrorMode = true
		} else if strings.HasPrefix(arg, "--convert-links") {
			if !mirrorMode {
				fmt.Println("error: --convert-links can only be used with --mirror")
				return
			}
			input.ConvertLinks = true
		} else if strings.HasPrefix(arg, "-R=") || strings.HasPrefix(arg, "--reject=") {
			if !mirrorMode {
				fmt.Println("error: --reject can only be used with --mirror")
				return
			}

			if strings.HasPrefix(arg, "-R=") {
				input.RejectFlag = arg[len("-R="):]
			} else {
				input.RejectFlag = arg[len("--reject="):]
			}
		} else if strings.HasPrefix(arg, "-X=") || strings.HasPrefix(arg, "--exclude=") {
			if !mirrorMode {
				fmt.Println("error: --exclude can only be used with --mirror")
				return
			}
			if strings.HasPrefix(arg, "-X=") {
				input.ExcludeFlag = arg[len("-X="):]
			} else {
				input.ExcludeFlag = arg[len("--exclude="):]
			}
		} else if strings.HasPrefix(arg, "-i=") {
			track = true
			input.InputFile = arg[len("-i="):]
		} else {
			fmt.Printf("\nInvalid argument: %s\n", arg)
			return
		}
	}

	// Validate background mode restrictions
	if input.Background {
		if input.InputFile != "" || input.Path != "" {
			fmt.Errorf("-B flag should not be used with -i or -P flags")
			return
		}
	}
	// Ensure --mirror is not combined with incompatible flags
	if input.IsMirror {
		if input.OutputFile != "" || input.Path != "" || input.RateLimit != "" ||
			input.InputFile != "" || input.Background {
			fmt.Errorf("error: --mirror can only be used with --convert-links, --reject, --exclude, and a URL. No other flags are allowed")
			return
		}
	} else {
		// Ensure --convert-links, --reject, and --exclude are only used with --mirror
		if input.ConvertLinks || input.RejectFlag != "" || input.ExcludeFlag != "" {
			fmt.Errorf("error: --convert-links, --reject, and --exclude can only be used with --mirror")
			return
		}
	}
	// Ensure a URL or source file is provided for valid execution
	if input.URL == nil && !track {
		fmt.Errorf("error: URL not provided")
		return
	}

	return
}

// isValidAttribute checks if an HTML tag attribute is valid for processing
func isValidAttribute(tagName, attrKey string) bool {
	return (tagName == "a" && attrKey == "href") ||
		(tagName == "img" && attrKey == "src") ||
		(tagName == "script" && attrKey == "src") ||
		(tagName == "link" && attrKey == "href")
}
