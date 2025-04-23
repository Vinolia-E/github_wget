package functions

import (
	"fmt"
	"os"
	"strings"
)

func ReadInputFile(filename string) []string {

	var results []string

	fileByte, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading the Input file: ", err)
		os.Exit(0)
	}

	inputString := string(fileByte)
	splited := strings.Split(inputString, "\n")
	// Remove empty lines
	for _, lines := range splited {
		if lines == "" {
			continue
		} else {
			results = append(results, lines)
		}
	}

	return results
}
