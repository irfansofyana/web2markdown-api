package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// ConversionResult struct
type ConversionResult struct {
	URL   string `json:"url,omitempty"`
	Title string `json:"title,omitempty"`
	Data  string `json:"data,omitempty"`
	Error error  `json:"error,omitempty"`
}

func main() {
	var client = &http.Client{}
	var cvtResult ConversionResult

	endpoint := "https://web2markdown-api.vercel.app/api/convert"
	body := []byte(`{
		"url": "https://kirandev.com/http-post-golang"
	}`)
	request, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if err = json.NewDecoder(response.Body).Decode(&cvtResult); err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(fmt.Sprintf("%s.md", cvtResult.Title))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(cvtResult.Data)
	if err != nil {
		log.Fatal(err)
	}
}
