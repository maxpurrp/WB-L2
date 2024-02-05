package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func myWget(url string) error {

	// create an HTTP client
	var client http.Client

	// create or open the file to save the downloaded content
	file, err := os.Create("site.html")
	if err != nil {
		return err
	}
	defer file.Close()

	// create an HTTP request with a specified User-Agent and other headers
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Edge/97.0.1072.71 Yandex/22.1.2.155 Safari/537.36")
	request.Header.Set("Accept", "text/html, application/xhtml+xml, application/xml;q=0.8, image/webp, */*;q=0.9")

	// perform the HTTP request
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// check if the HTTP status is OK (200)
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("error when receiving the file. Status: %s", response.Status)
	}

	// copy the response body to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// get command-line arguments
	args := os.Args
	fmt.Println(args)

	// check if there are enough arguments
	if len(args) < 2 {
		fmt.Println("Usage: program <url>")
		return
	}

	// extract URL and depth from command-line arguments
	url := args[1]

	// call myWget function with the specified URL
	err := myWget(url)
	if err != nil {
		fmt.Printf("error: %s \n", err)
	}
	fmt.Println("Successfully")
}
