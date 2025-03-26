package functions

import (
	"fmt"
	"os"
	"strings"
)

func ReadInputFile(filename string) []string {

	fileByte, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading the Input file: ", err)
		os.Exit(0)
	}

	inputString := string(fileByte)
	results := strings.Split(inputString, "\n")

	return results
}
