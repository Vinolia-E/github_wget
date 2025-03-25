package functions

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(filename, url string) error {
	// Make http request:
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	// Check for succesful response:
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", response.Status)
	}

	// Create the output file:

	output, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer output.Close()

	// Write the response body to file

	_, err = io.Copy(output, response.Body)
	return err
}
